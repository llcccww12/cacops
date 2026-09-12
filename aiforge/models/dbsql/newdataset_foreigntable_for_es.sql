INSERT INTO public.dataset_es(
	    uuid, 
        name,
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
        updated_unix,
        file_name,
        file_desc
        )
	    SELECT 
        b.id, 
        b.name, 
        b.alias, 
        0, 
        b.tags, 
        '', 
        b.download_count, 
        b.license, 
        b.tasks, 
        0, 
        b.owner_id, 
        0, 
        b.created_unix, 
        b.updated_unix,
        '',
		''
		 FROM public.dataset_registry b where is_private=false;


DROP TRIGGER  IF EXISTS es_insert_dataset_new on public.dataset_registry;

CREATE OR REPLACE FUNCTION public.insert_dataset_new() RETURNS trigger AS 
$def$
    BEGIN
        if not NEW.is_private  then
            INSERT INTO public.dataset_es(
                uuid, 
                name,
                title, 
                category, 
                download_times, 
                license, 
                task, 
                user_id, 
                created_unix, 
                updated_unix)
                VALUES (
                NEW.id, 
                NEW.name,
                NEW.alias, 
                NEW.tags, 
                NEW.download_count, 
                NEW.license, 
                NEW.tasks, 
                NEW.owner_id, 
                NEW.created_unix, 
                NEW.updated_unix
                );
        end if;
        RETURN NEW;
    END;
$def$ 
LANGUAGE plpgsql;



CREATE TRIGGER es_insert_dataset_new
    AFTER INSERT ON public.dataset_registry
    FOR EACH ROW EXECUTE PROCEDURE insert_dataset_new();

ALTER TABLE public.dataset_registry ENABLE ALWAYS TRIGGER es_insert_dataset_new;



DROP TRIGGER  IF EXISTS es_update_dataset_new on public.dataset_registry;

CREATE OR REPLACE FUNCTION public.update_dataset_new() RETURNS trigger AS 
$def$
    DECLARE
       privateValue text=NEW.id;
       tName text = NEW.name;
       newName text = NEW.alias;
       newTags text = NEW.tags;
       newtasks text = NEW.tasks;
       newDcount integer = NEW.download_count;
       newUuinx bigint = NEW.updated_unix;
    BEGIN
       if OLD.is_private != NEW.is_private then
		  if NEW.is_private then
             EXECUTE format('delete from public.dataset_es where uuid=%L',privateValue);
		  elsif not NEW.is_private then
             INSERT INTO public.dataset_es(
                uuid, 
                name,
                title, 
                category, 
                download_times, 
                license, 
                task, 
                user_id, 
                created_unix, 
                updated_unix)
                VALUES (
                NEW.id, 
                NEW.name,
                NEW.alias, 
                NEW.tags, 
                NEW.download_count, 
                NEW.license, 
                NEW.tasks, 
                NEW.owner_id, 
                NEW.created_unix, 
                NEW.updated_unix
             );
          end if;
	  end if;
      if not NEW.is_private then
            EXECUTE format('UPDATE public.dataset_es 
            SET title=%L,
            name=%L,
            category=%L,
            task=%L,
            download_times=%s,
            updated_unix=%s
            where uuid=%L',newName,tName,newTags,newtasks,newDcount,newUuinx,privateValue);
       end if;
       return new;
    END
$def$ 
LANGUAGE plpgsql;

CREATE TRIGGER es_update_dataset_new
    AFTER UPDATE  ON public.dataset_registry
    FOR EACH ROW EXECUTE PROCEDURE update_dataset_new();

ALTER TABLE public.dataset_registry ENABLE ALWAYS TRIGGER es_update_dataset_new;

DROP TRIGGER IF EXISTS es_delete_dataset_new on public.dataset_registry;

CREATE OR REPLACE FUNCTION public.delete_dataset_new() RETURNS trigger AS 
$def$
    declare
       privateValue text=OLD.id;
    BEGIN
	   EXECUTE format('DELETE FROM public.dataset_es  where uuid=%L',privateValue);
       return new;
    END
$def$ 
LANGUAGE plpgsql;


CREATE TRIGGER es_delete_dataset_new
    AFTER DELETE ON public.dataset_registry
    FOR EACH ROW EXECUTE PROCEDURE delete_dataset_new();

ALTER TABLE public.dataset_registry ENABLE ALWAYS TRIGGER es_delete_dataset_new;
