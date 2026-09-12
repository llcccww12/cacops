-- 要处理项目从私有变为公有，并且从公有变成私有的情况
DELETE FROM public.repository_es;
DROP FOREIGN table if exists  public.repository_es;
DROP TRIGGER  IF EXISTS es_insert_repository on public.repository;
DROP FUNCTION public.insert_repository_data;
DROP TRIGGER  IF EXISTS es_update_repository on public.repository;
DROP FUNCTION public.update_repository;

DROP TRIGGER IF EXISTS es_delete_repository on public.repository;
DROP FUNCTION public.delete_repository;

DROP TRIGGER IF EXISTS es_udpate_repository_lang on public.language_stat;
DROP FUNCTION public.udpate_repository_lang;


CREATE FOREIGN TABLE public.repository_es (
    id bigint NOT NULL,
    owner_id bigint,
    owner_name character varying(255),
    lower_name character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    website character varying(2048),
    original_service_type integer,
    original_url character varying(2048),
    default_branch character varying(255),
    num_watches integer,
    num_stars integer,
    num_forks integer,
    num_issues integer,
    num_closed_issues integer,
    num_pulls integer,
    num_closed_pulls integer,
    num_milestones integer DEFAULT 0 NOT NULL,
    num_closed_milestones integer DEFAULT 0 NOT NULL,
    is_private boolean,
    is_empty boolean,
    is_archived boolean,
    is_mirror boolean,
    status integer DEFAULT 0 NOT NULL,
    is_fork boolean DEFAULT false NOT NULL,
    fork_id bigint,
    is_template boolean DEFAULT false NOT NULL,
    template_id bigint,
    size bigint DEFAULT 0 NOT NULL,
    is_fsck_enabled boolean DEFAULT true NOT NULL,
    close_issues_via_commit_in_any_branch boolean DEFAULT false NOT NULL,
    topics text,
    avatar character varying(64),
    created_unix bigint,
    updated_unix bigint,
    contract_address character varying(255),
    block_chain_status integer DEFAULT 0 NOT NULL,
    balance character varying(255) DEFAULT '0'::character varying NOT NULL,
    clone_cnt bigint DEFAULT 0 NOT NULL,
    license character varying(100),
    download_cnt bigint DEFAULT 0 NOT NULL,
    num_commit bigint DEFAULT 0 NOT NULL,
    git_clone_cnt bigint DEFAULT 0 NOT NULL,
    creator_id bigint NOT NULL DEFAULT 0,
    repo_type integer NOT NULL DEFAULT 0,
	lang character varying(2048),
    alias character varying(255),
    lower_alias character varying(255)
) SERVER multicorn_es
OPTIONS
    (
        host '192.168.207.94',
        port '9200',
        index 'repository-es-index',
        rowid_column 'id',
        default_sort '_id'
    )
;
delete from  public.repository_es;
    INSERT INTO public.repository_es (id,
                                        owner_id,
                                        owner_name,
                                        lower_name,
                                        name,
                                        description,
                                        website,
                                        original_service_type,
                                        original_url,
                                        default_branch,
                                        num_watches,
                                        num_stars,
                                        num_forks,
                                        num_issues,
                                        num_closed_issues,
                                        num_pulls,
                                        num_closed_pulls,
                                        num_milestones,
                                        num_closed_milestones,
                                        is_private,
                                        is_empty,
                                        is_archived,
                                        is_mirror,
                                        status,
                                        is_fork,
                                        fork_id,
                                        is_template,
                                        template_id,
                                        size,
                                        is_fsck_enabled,
                                        close_issues_via_commit_in_any_branch,
                                        topics,
                                        avatar,
                                        created_unix,
                                        updated_unix,
                                        contract_address,
                                        block_chain_status,
                                        balance,
                                        clone_cnt,
                                        num_commit,
                                        git_clone_cnt,
                                        creator_id,
                                        repo_type,
                                        lang,
                                        alias,
                                        lower_alias
                                        ) 
                                        SELECT
                                        id,
                                        owner_id,
                                        owner_name,
                                        lower_name,
                                        name,
                                        description,
                                        website,
                                        original_service_type,
                                        original_url,
                                        default_branch,
                                        num_watches,
                                        num_stars,
                                        num_forks,
                                        num_issues,
                                        num_closed_issues,
                                        num_pulls,
                                        num_closed_pulls,
                                        num_milestones,
                                        num_closed_milestones,
                                        is_private,
                                        is_empty,
                                        is_archived,
                                        is_mirror,
                                        status,
                                        is_fork,
                                        fork_id,
                                        is_template,
                                        template_id,
                                        size,
                                        is_fsck_enabled,
                                        close_issues_via_commit_in_any_branch,
                                        topics,
                                        avatar,
                                        created_unix,
                                        updated_unix,
                                        contract_address,
                                        block_chain_status,
                                        balance,
                                        clone_cnt,
                                        num_commit,
                                        git_clone_cnt,
                                        creator_id,
                                        repo_type,
                                        (select  array_to_string(array_agg(language order by percentage desc),',')  from public.language_stat a where a.repo_id=b.id),
                                        alias,
                                        lower_alias
                                        FROM  public.repository b where b.is_private=false;

DROP TRIGGER  IF EXISTS es_insert_repository on public.repository;

CREATE OR REPLACE FUNCTION public.insert_repository_data() RETURNS trigger AS 
$def$
    BEGIN
        if not NEW.is_private then
                INSERT INTO public.repository_es (id,
                                    owner_id,
                                    owner_name,
                                    lower_name,
                                    name,
                                    description,
                                    website,
                                    original_service_type,
                                    original_url,
                                    default_branch,
                                    num_watches,
                                    num_stars,
                                    num_forks,
                                    num_issues,
                                    num_closed_issues,
                                    num_pulls,
                                    num_closed_pulls,
                                    num_milestones,
                                    num_closed_milestones,
                                    is_private,
                                    is_empty,
                                    is_archived,
                                    is_mirror,
                                    status,
                                    is_fork,
                                    fork_id,
                                    is_template,
                                    template_id,
                                    size,
                                    is_fsck_enabled,
                                    close_issues_via_commit_in_any_branch,
                                    topics,
                                    avatar,
                                    created_unix,
                                    updated_unix,
                                    contract_address,
                                    block_chain_status,
                                    balance,
                                    clone_cnt,
                                    num_commit,
                                    git_clone_cnt,
                                    creator_id,
                                    repo_type,
                                    alias,
                                    lower_alias) VALUES
                                    (NEW.id,
                                    NEW.owner_id,
                                    NEW.owner_name,
                                    NEW.lower_name,
                                    NEW.name,
                                    NEW.description,
                                    NEW.website,
                                    NEW.original_service_type,
                                    NEW.original_url,
                                    NEW.default_branch,
                                    NEW.num_watches,
                                    NEW.num_stars,
                                    NEW.num_forks,
                                    NEW.num_issues,
                                    NEW.num_closed_issues,
                                    NEW.num_pulls,
                                    NEW.num_closed_pulls,
                                    NEW.num_milestones,
                                    NEW.num_closed_milestones,
                                    NEW.is_private,
                                    NEW.is_empty,
                                    NEW.is_archived,
                                    NEW.is_mirror,
                                    NEW.status,
                                    NEW.is_fork,
                                    NEW.fork_id,
                                    NEW.is_template,
                                    NEW.template_id,
                                    NEW.size,
                                    NEW.is_fsck_enabled,
                                    NEW.close_issues_via_commit_in_any_branch,
                                    NEW.topics,
                                    NEW.avatar,
                                    NEW.created_unix,
                                    NEW.updated_unix,
                                    NEW.contract_address,
                                    NEW.block_chain_status,
                                    NEW.balance,
                                    NEW.clone_cnt,
                                    NEW.num_commit,
                                    NEW.git_clone_cnt,
                                    NEW.creator_id,
                                    NEW.repo_type,
                                    NEW.alias,
                                    NEW.lower_alias);
        end if;
        RETURN NEW;
    END;
$def$ 
LANGUAGE plpgsql;


CREATE TRIGGER es_insert_repository
    AFTER INSERT ON public.repository
    FOR EACH ROW EXECUTE PROCEDURE insert_repository_data();

ALTER TABLE public.repository ENABLE ALWAYS TRIGGER es_insert_repository;

DROP TRIGGER  IF EXISTS es_update_repository on public.repository;

CREATE OR REPLACE FUNCTION public.update_repository() RETURNS trigger AS 
$def$
    BEGIN
      if OLD.is_private != NEW.is_private then
        if OLD.is_private and not NEW.is_private then
            --insert
            INSERT INTO public.repository_es (id,
                                    owner_id,
                                    owner_name,
                                    lower_name,
                                    name,
                                    description,
                                    website,
                                    original_service_type,
                                    original_url,
                                    default_branch,
                                    num_watches,
                                    num_stars,
                                    num_forks,
                                    num_issues,
                                    num_closed_issues,
                                    num_pulls,
                                    num_closed_pulls,
                                    num_milestones,
                                    num_closed_milestones,
                                    is_private,
                                    is_empty,
                                    is_archived,
                                    is_mirror,
                                    status,
                                    is_fork,
                                    fork_id,
                                    is_template,
                                    template_id,
                                    size,
                                    is_fsck_enabled,
                                    close_issues_via_commit_in_any_branch,
                                    topics,
                                    avatar,
                                    created_unix,
                                    updated_unix,
                                    contract_address,
                                    block_chain_status,
                                    balance,
                                    clone_cnt,
                                    num_commit,
                                    git_clone_cnt,
                                    creator_id,
                                    repo_type,
                                    lang, 
                                    alias,
                                    lower_alias) 
                                    SELECT
                                    id,
                                    owner_id,
                                    owner_name,
                                    lower_name,
                                    name,
                                    description,
                                    website,
                                    original_service_type,
                                    original_url,
                                    default_branch,
                                    num_watches,
                                    num_stars,
                                    num_forks,
                                    num_issues,
                                    num_closed_issues,
                                    num_pulls,
                                    num_closed_pulls,
                                    num_milestones,
                                    num_closed_milestones,
                                    is_private,
                                    is_empty,
                                    is_archived,
                                    is_mirror,
                                    status,
                                    is_fork,
                                    fork_id,
                                    is_template,
                                    template_id,
                                    size,
                                    is_fsck_enabled,
                                    close_issues_via_commit_in_any_branch,
                                    topics,
                                    avatar,
                                    created_unix,
                                    updated_unix,
                                    contract_address,
                                    block_chain_status,
                                    balance,
                                    clone_cnt,
                                    num_commit,
                                    git_clone_cnt,
                                    creator_id,
                                    repo_type,
                                    (select  array_to_string(array_agg(language order by percentage desc),',')  from public.language_stat a where a.repo_id=b.id),
                                    alias,
                                    lower_alias 
                                    FROM  public.repository b where b.id=NEW.id;
             INSERT INTO public.dataset_es(
                id, 
                title, 
                status, 
                category, 
                description, 
                download_times, 
                license, task, 
                release_id, 
                user_id, 
                repo_id, 
                created_unix, 
                updated_unix,file_name)
                SELECT 
                b.id, 
                b.title, 
                b.status, 
                b.category, 
                b.description, 
                b.download_times, 
                b.license, 
                b.task, 
                b.release_id, 
                b.user_id, 
                b.repo_id, 
                b.created_unix, 
                b.updated_unix,(select array_to_string(array_agg(name order by created_unix desc),',') from public.attachment a where a.dataset_id=b.id and a.is_private=false)
                FROM public.dataset b where b.repo_id=NEW.id;

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
                    (select id from public.pull_request d where d.issue_id=b.id)
                    FROM public.issue b where b.repo_id=NEW.id;
        
        end if;

        if not OLD.is_private and NEW.is_private then
            delete from public.issue_es where repo_id=NEW.id;
            -- delete from public.dataset_es where repo_id=NEW.id;
            delete from public.repository_es where id=NEW.id;
        end if;

      end if;

      if not NEW.is_private then
            raise notice 'update repo,the updated_unix is %',NEW.updated_unix;
            update public.repository_es SET description=NEW.description,
            name=NEW.name,
            lower_name=NEW.lower_name,
            owner_name=NEW.owner_name,
            website=NEW.website,
            updated_unix=NEW.updated_unix,
            num_watches=NEW.num_watches,
            num_stars=NEW.num_stars,
            num_forks=NEW.num_forks,
            topics=NEW.topics,
            alias = NEW.alias,
            lower_alias = NEW.lower_alias,
            avatar=NEW.avatar
            where id=NEW.id;
       end if;
       return new;
    END
$def$ 
LANGUAGE plpgsql;

CREATE TRIGGER es_update_repository
    AFTER UPDATE ON public.repository
    FOR EACH ROW EXECUTE PROCEDURE update_repository();

ALTER TABLE public.repository ENABLE ALWAYS TRIGGER es_update_repository;


DROP TRIGGER IF EXISTS es_delete_repository on public.repository;

CREATE OR REPLACE FUNCTION public.delete_repository() RETURNS trigger AS 
$def$
    declare
    BEGIN
       delete from public.issue_es where repo_id=OLD.id;
       delete from public.dataset_es where repo_id=OLD.id;
	   DELETE FROM public.repository_es  where id=OLD.id;
       return new;
    END
$def$ 
LANGUAGE plpgsql;


CREATE TRIGGER es_delete_repository
    AFTER DELETE ON public.repository
    FOR EACH ROW EXECUTE PROCEDURE delete_repository();

ALTER TABLE public.repository ENABLE ALWAYS TRIGGER es_delete_repository;



DROP TRIGGER IF EXISTS es_udpate_repository_lang on public.language_stat;

CREATE OR REPLACE FUNCTION public.udpate_repository_lang() RETURNS trigger AS 
$def$
    DECLARE
        privateValue bigint;
    BEGIN
       if (TG_OP = 'UPDATE') then
           select into privateValue updated_unix from public.repository where id=NEW.repo_id;
	       update public.repository_es  SET updated_unix=privateValue,lang=(select  array_to_string(array_agg(language order by percentage desc),',') from public.language_stat where repo_id=NEW.repo_id) where id=NEW.repo_id;
	   elsif (TG_OP = 'INSERT') then
           select into privateValue updated_unix from public.repository where id=NEW.repo_id;
	       update public.repository_es  SET updated_unix=privateValue,lang=(select  array_to_string(array_agg(language order by percentage desc),',')  from public.language_stat where repo_id=NEW.repo_id) where id=NEW.repo_id;
       elsif (TG_OP = 'DELETE') then
           if exists(select 1 from public.repository where id=OLD.repo_id) then
	           update public.repository_es  SET lang=(select  array_to_string(array_agg(language order by percentage desc),',')  from public.language_stat where repo_id=OLD.repo_id) where id=OLD.repo_id;
	       end if;
       end if;
       return NEW;
    END;
$def$ 
LANGUAGE plpgsql;

CREATE TRIGGER es_udpate_repository_lang
    AFTER INSERT OR UPDATE OR DELETE ON public.language_stat
    FOR EACH ROW EXECUTE PROCEDURE udpate_repository_lang();

ALTER TABLE public.language_stat ENABLE ALWAYS TRIGGER es_udpate_repository_lang;