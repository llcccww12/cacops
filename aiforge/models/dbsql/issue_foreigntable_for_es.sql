delete from public.issue_es;
DROP FOREIGN TABLE public.issue_es;
DROP TRIGGER  IF EXISTS es_insert_issue on public.issue;
DROP FUNCTION  public.insert_issue_data;
DROP TRIGGER IF EXISTS es_udpate_issue_comment on public.comment;
DROP FUNCTION udpate_issue_comment;
DROP TRIGGER  IF EXISTS es_update_issue on public.issue;
DROP FUNCTION public.update_issue;
DROP TRIGGER IF EXISTS es_delete_issue on public.issue;
DROP FUNCTION public.delete_issue;


CREATE FOREIGN TABLE public.issue_es
(
    id bigint NOT NULL,
    repo_id bigint,
    index bigint,
    poster_id bigint,
    original_author character varying(255),
    original_author_id bigint,
    name character varying(255) ,
    content text,
    comment text,
    milestone_id bigint,
    priority integer,
    is_closed boolean,
    is_pull boolean,
    pr_id bigint,
    num_comments integer,
    ref character varying(255),
    deadline_unix bigint,
    created_unix bigint,
    updated_unix bigint,
    closed_unix bigint,
    is_locked boolean NOT NULL,
    amount bigint,
    is_transformed boolean NOT NULL
)SERVER multicorn_es
OPTIONS
    (
        host '192.168.207.94',
        port '9200',
        index 'issue-es-index',
        rowid_column 'id',
        default_sort '_id'
    )
;

delete from public.issue_es;
INSERT INTO public.issue_es(
	     id, 
         repo_id, 
         index, 
         poster_id, 
         original_author, 
         original_author_id, 
         name, 
         content, 
         milestone_id, 
         priority, 
         is_closed, 
         is_pull, 
         num_comments, 
         ref, 
         deadline_unix, 
         created_unix, 
         updated_unix, 
         closed_unix, 
         is_locked, 
         amount, 
         is_transformed,comment,pr_id)
	     SELECT 
         b.id, 
         b.repo_id, 
         b.index, 
         b.poster_id, 
         b.original_author, 
         b.original_author_id, 
         b.name, 
         b.content, 
         b.milestone_id, 
         b.priority, 
         b.is_closed, 
         b.is_pull, 
         b.num_comments, 
         b.ref, 
         b.deadline_unix, 
         b.created_unix, 
         b.updated_unix, 
         b.closed_unix, 
         b.is_locked, 
         b.amount, 
         b.is_transformed,
		 (select array_to_string(array_agg(content order by created_unix desc),',') from public.comment a where a.issue_id=b.id),
         (select id from public.pull_request d where  b.id=d.issue_id and b.is_pull=true)
         FROM public.issue b,public.repository c where b.repo_id=c.id and c.is_private=false;


CREATE OR REPLACE FUNCTION public.insert_issue_data() RETURNS trigger AS 
$def$
    DECLARE
        privateValue boolean=false;
    BEGIN
        select into privateValue is_private from public.repository where id=NEW.repo_id;
        if not privateValue then
            INSERT INTO public.issue_es(
            id, 
            repo_id, 
            index, 
            poster_id, 
            original_author, 
            original_author_id, 
            name, 
            content, 
            milestone_id, 
            priority, 
            is_closed, 
            is_pull, 
            num_comments, 
            ref, 
            deadline_unix, 
            created_unix, 
            updated_unix, 
            closed_unix, 
            is_locked, 
            amount, 
            is_transformed)
            VALUES (
            NEW.id, 
            NEW.repo_id, 
            NEW.index, 
            NEW.poster_id, 
            NEW.original_author, 
            NEW.original_author_id, 
            NEW.name, 
            NEW.content, 
            NEW.milestone_id, 
            NEW.priority, 
            NEW.is_closed, 
            NEW.is_pull, 
            NEW.num_comments, 
            NEW.ref, 
            NEW.deadline_unix, 
            NEW.created_unix, 
            NEW.updated_unix, 
            NEW.closed_unix, 
            NEW.is_locked, 
            NEW.amount, 
            NEW.is_transformed
            );
        end if;
        RETURN NEW;
    END;
$def$ 
LANGUAGE plpgsql;

DROP TRIGGER  IF EXISTS es_insert_issue on public.issue;

CREATE TRIGGER es_insert_issue
    AFTER INSERT ON public.issue
    FOR EACH ROW EXECUTE PROCEDURE insert_issue_data();

ALTER TABLE public.issue ENABLE ALWAYS TRIGGER es_insert_issue;

CREATE OR REPLACE FUNCTION public.udpate_issue_comment() RETURNS trigger AS 
$def$
    BEGIN
       if (TG_OP = 'DELETE') then
	       update public.issue_es  SET comment=(select array_to_string(array_agg(content order by created_unix desc),',') from public.comment where issue_id=OLD.issue_id) where id=OLD.issue_id;
	   	elsif (TG_OP = 'UPDATE') then
	       update public.issue_es  SET comment=(select array_to_string(array_agg(content order by created_unix desc),',') from public.comment where issue_id=NEW.issue_id) where id=NEW.issue_id;
       end if;
	   
       return null;
    END;
$def$ 
LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS es_udpate_issue_comment on public.comment;
CREATE TRIGGER es_udpate_issue_comment
    AFTER DELETE OR UPDATE ON public.comment
    FOR EACH ROW EXECUTE PROCEDURE udpate_issue_comment();

ALTER TABLE public.comment ENABLE ALWAYS TRIGGER es_udpate_issue_comment;


CREATE OR REPLACE FUNCTION public.update_issue() RETURNS trigger AS 
$def$
    declare
    BEGIN
	   UPDATE public.issue_es  
       SET content=NEW.content, 
       name=NEW.name,
       is_closed=NEW.is_closed,
       num_comments=NEW.num_comments,
       updated_unix=NEW.updated_unix,
       comment=(select array_to_string(array_agg(content order by created_unix desc),',') from public.comment where issue_id=NEW.id)
       where id=NEW.id;
       return new;
    END
$def$ 
LANGUAGE plpgsql;

DROP TRIGGER  IF EXISTS es_update_issue on public.issue;

CREATE TRIGGER es_update_issue
    AFTER UPDATE  ON public.issue
    FOR EACH ROW EXECUTE PROCEDURE update_issue();

ALTER TABLE public.issue ENABLE ALWAYS TRIGGER es_update_issue;

CREATE OR REPLACE FUNCTION public.delete_issue() RETURNS trigger AS 
$def$
    declare
    BEGIN
	   DELETE FROM public.issue_es  where id=OLD.id;
       return new;
    END
$def$ 
LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS es_delete_issue on public.issue;
CREATE TRIGGER es_delete_issue
    AFTER DELETE ON public.issue
    FOR EACH ROW EXECUTE PROCEDURE delete_issue();

ALTER TABLE public.issue ENABLE ALWAYS TRIGGER es_delete_issue;