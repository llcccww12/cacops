DELETE FROM public.dataset_es;
DROP FOREIGN TABLE public.dataset_es;
DROP TRIGGER  IF EXISTS es_insert_dataset on public.dataset;
DROP FUNCTION  public.insert_dataset_data();
DROP TRIGGER IF EXISTS es_udpate_dataset_file_name on public.attachment;
DROP FUNCTION public.udpate_dataset_file_name;

DROP TRIGGER  IF EXISTS es_update_dataset on public.dataset;
DROP FUNCTION public.update_dataset;

DROP TRIGGER IF EXISTS es_delete_dataset on public.dataset;
DROP FUNCTION public.delete_dataset;


CREATE FOREIGN TABLE public.dataset_es
(
    id serial  not null,
	uuid character varying(255),
	name character varying(255),
    title character varying(255),
    status integer,
    category character varying(255),
    description text,
    download_times bigint,
    license character varying(255),
    task character varying(255),
    release_id bigint,
    user_id bigint,
    repo_id bigint,
    created_unix bigint,
    updated_unix bigint,
    file_name text,
    file_desc text
)SERVER multicorn_es
OPTIONS
    (
        host '192.168.207.116',
        port '9200',
        index 'dataset-es-index',
        rowid_column 'id',
        default_sort '_id'
    )
;
DELETE FROM public.dataset_es;
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
        updated_unix,
        file_name,
        file_desc
        )
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
        b.updated_unix,
        (select array_to_string(array_agg(name order by created_unix desc),'-#,#-') from public.attachment a where a.dataset_id=b.id and a.is_private=false),
        (select array_to_string(array_agg(description order by created_unix desc),'-#,#-') from public.attachment a where a.dataset_id=b.id and a.is_private=false)
        FROM public.dataset b,public.repository c where b.repo_id=c.id and c.is_private=false;


DROP TRIGGER  IF EXISTS es_insert_dataset on public.dataset;

CREATE OR REPLACE FUNCTION public.insert_dataset_data() RETURNS trigger AS 
$def$
    DECLARE
       privateValue boolean=false;
    BEGIN
        select into privateValue is_private from public.repository where id=NEW.repo_id;
        if not privateValue then
            INSERT INTO public.dataset_es(
                id, 
                title, 
                status, 
                category, 
                description, 
                download_times, 
                license, 
                task, 
                release_id, 
                user_id, 
                repo_id, 
                created_unix, 
                updated_unix)
                VALUES (
                NEW.id, 
                NEW.title, 
                NEW.status, 
                NEW.category, 
                NEW.description, 
                NEW.download_times, 
                NEW.license, 
                NEW.task, 
                NEW.release_id, 
                NEW.user_id, 
                NEW.repo_id, 
                NEW.created_unix, 
                NEW.updated_unix
                );
        end if;
        RETURN NEW;
    END;
$def$ 
LANGUAGE plpgsql;



CREATE TRIGGER es_insert_dataset
    AFTER INSERT ON public.dataset
    FOR EACH ROW EXECUTE PROCEDURE insert_dataset_data();

ALTER TABLE public.dataset ENABLE ALWAYS TRIGGER es_insert_dataset;


DROP TRIGGER IF EXISTS es_udpate_dataset_file_name on public.attachment;

CREATE OR REPLACE FUNCTION public.udpate_dataset_file_name() RETURNS trigger AS 
$def$
    BEGIN
        if (TG_OP = 'UPDATE') then
           update public.dataset_es  SET file_desc=(select array_to_string(array_agg(description order by created_unix desc),'-#,#-') from public.attachment where dataset_id=NEW.dataset_id and is_private=false) where id=NEW.dataset_id;
	    elsif (TG_OP = 'INSERT') then
           update public.dataset_es  SET file_name=(select array_to_string(array_agg(name order by created_unix desc),'-#,#-') from public.attachment where dataset_id=NEW.dataset_id and is_private=false) where id=NEW.dataset_id;
        elsif (TG_OP = 'DELETE') then
           update public.dataset_es  SET file_name=(select array_to_string(array_agg(name order by created_unix desc),'-#,#-') from public.attachment where dataset_id=OLD.dataset_id and is_private=false) where id=OLD.dataset_id;
           update public.dataset_es  SET file_desc=(select array_to_string(array_agg(description order by created_unix desc),'-#,#-') from public.attachment where dataset_id=OLD.dataset_id and is_private=false) where id=OLD.dataset_id;
       end if;
       return NEW;
    END;
$def$ 
LANGUAGE plpgsql;


CREATE TRIGGER es_udpate_dataset_file_name
     AFTER INSERT OR UPDATE OR DELETE ON public.attachment
    FOR EACH ROW EXECUTE PROCEDURE udpate_dataset_file_name();

ALTER TABLE public.attachment ENABLE ALWAYS TRIGGER es_udpate_dataset_file_name;

DROP TRIGGER  IF EXISTS es_update_dataset on public.dataset;

CREATE OR REPLACE FUNCTION public.update_dataset() RETURNS trigger AS 
$def$
    BEGIN
       if (NEW.status=0) then
           delete from public.dataset_es where id=NEW.id;
       elsif (NEW.status=1) then
            UPDATE public.dataset_es 
            SET description=NEW.description, 
            title=NEW.title,
            category=NEW.category,
            task=NEW.task,
            download_times=NEW.download_times,
            updated_unix=NEW.updated_unix,
            file_name=(select array_to_string(array_agg(name order by created_unix desc),'-#,#-') from public.attachment where dataset_id=NEW.id and is_private=false),
            file_desc=(select array_to_string(array_agg(description order by created_unix desc),'-#,#-') from public.attachment where dataset_id=NEW.id and is_private=false)
            where id=NEW.id;
       end if;
       return new;
    END
$def$ 
LANGUAGE plpgsql;

CREATE TRIGGER es_update_dataset
    AFTER UPDATE  ON public.dataset
    FOR EACH ROW EXECUTE PROCEDURE update_dataset();

ALTER TABLE public.dataset ENABLE ALWAYS TRIGGER es_update_dataset;

DROP TRIGGER IF EXISTS es_delete_dataset on public.dataset;

CREATE OR REPLACE FUNCTION public.delete_dataset() RETURNS trigger AS 
$def$
    declare
    BEGIN
	   DELETE FROM public.dataset_es  where id=OLD.id;
       return new;
    END
$def$ 
LANGUAGE plpgsql;


CREATE TRIGGER es_delete_dataset
    AFTER DELETE ON public.dataset
    FOR EACH ROW EXECUTE PROCEDURE delete_dataset();

ALTER TABLE public.dataset ENABLE ALWAYS TRIGGER es_delete_dataset;
