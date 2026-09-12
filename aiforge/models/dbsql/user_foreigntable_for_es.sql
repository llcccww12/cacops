DELETE FROM public.user_es;
DROP FOREIGN table if exists public.user_es;
DROP TRIGGER  IF EXISTS es_insert_user on public.user;
DROP FUNCTION  public.insert_user_data;
DROP TRIGGER  IF EXISTS es_update_user on public.user;
DROP FUNCTION public.update_user;

DROP TRIGGER IF EXISTS es_delete_user on public.user;
DROP FUNCTION public.delete_user;

CREATE FOREIGN TABLE public.user_es
(
   id bigint NOT NULL ,
    lower_name character varying(255)  NULL,
    name character varying(255)  NULL,
    full_name character varying(255), 
    email character varying(255), 
    keep_email_private boolean,
    email_notifications_preference character varying(20) ,
    passwd character varying(255) ,
    passwd_hash_algo character varying(255) ,
    must_change_password boolean NOT NULL DEFAULT false,
    login_type integer,
    login_source bigint NOT NULL DEFAULT 0,
    login_name character varying(255) ,
    type integer,
    location character varying(255),
    website character varying(255),
    rands character varying(10),
    salt character varying(10),
    language character varying(5),
    description character varying(255),
    created_unix bigint,
    updated_unix bigint,
    last_login_unix bigint,
    last_repo_visibility boolean,
    max_repo_creation integer,
    is_active boolean,
    is_admin boolean,
    is_restricted boolean NOT NULL DEFAULT false,
    allow_git_hook boolean,
    allow_import_local boolean,
    allow_create_organization boolean DEFAULT true,
    prohibit_login boolean NOT NULL DEFAULT false,
    avatar character varying(2048) ,
    avatar_email character varying(255),
    use_custom_avatar boolean,
    num_followers integer,
    num_following integer NOT NULL DEFAULT 0,
    num_stars integer,
    num_repos integer,
    num_teams integer,
    num_members integer,
    visibility integer NOT NULL DEFAULT 0,
    repo_admin_change_team_access boolean NOT NULL DEFAULT false,
    diff_view_style character varying(255),
    theme character varying(255),
    token character varying(1024) ,
    public_key character varying(255),
    private_key character varying(255),
    is_operator boolean NOT NULL DEFAULT false,
    num_dataset_stars integer NOT NULL DEFAULT 0
) SERVER multicorn_es
OPTIONS
    (
        host '192.168.207.94',
        port '9200',
        index 'user-es-index',
        rowid_column 'id',
        default_sort '_id'
    )
;
delete from public.user_es;
 INSERT INTO public.user_es(
	    id, 
        lower_name,
        name, 
        full_name, 
        email, 
        keep_email_private, 
        email_notifications_preference, 
        must_change_password, 
        login_type, 
        login_source, 
        login_name, 
        type, 
        location, 
        website, 
        rands, 
        language, 
        description, 
        created_unix, 
        updated_unix, 
        last_login_unix, 
        last_repo_visibility, 
        max_repo_creation, 
        is_active, 
        is_restricted, 
        allow_git_hook, 
        allow_import_local, 
        allow_create_organization, 
        prohibit_login, 
        avatar, 
        avatar_email, 
        use_custom_avatar, 
        num_followers, 
        num_following, 
        num_stars, 
        num_repos, 
        num_teams, 
        num_members, 
        visibility, 
        repo_admin_change_team_access, 
        diff_view_style, 
        theme, 
        is_operator,
        num_dataset_stars)
	    SELECT 
        id, 
        lower_name,
        name, 
        full_name, 
        email, 
        keep_email_private, 
        email_notifications_preference, 
        must_change_password, 
        login_type, 
        login_source, 
        login_name, 
        type, 
        location, 
        website, 
        rands, 
        language, 
        description, 
        created_unix, 
        updated_unix, 
        last_login_unix, 
        last_repo_visibility, 
        max_repo_creation, 
        is_active, 
        is_restricted, 
        allow_git_hook, 
        allow_import_local, 
        allow_create_organization, 
        prohibit_login, 
        avatar, 
        avatar_email, 
        use_custom_avatar, 
        num_followers, 
        num_following, 
        num_stars, 
        num_repos, 
        num_teams, 
        num_members, 
        visibility, 
        repo_admin_change_team_access, 
        diff_view_style, 
        theme, 
        is_operator,
        num_dataset_stars
        FROM public.user;

DROP TRIGGER  IF EXISTS es_insert_user on public.user;

CREATE OR REPLACE FUNCTION public.insert_user_data() RETURNS trigger AS 
$def$
    BEGIN
        INSERT INTO public."user_es"(
	    id, 
        lower_name,
        name, 
        full_name, 
        email, 
        keep_email_private, 
        email_notifications_preference, 
        must_change_password, 
        login_type, 
        login_source, 
        login_name, 
        type, 
        location, 
        website, 
        rands, 
        language, 
        description, 
        created_unix, 
        updated_unix, 
        last_login_unix, 
        last_repo_visibility, 
        max_repo_creation, 
        is_active, 
        is_restricted, 
        allow_git_hook, 
        allow_import_local, 
        allow_create_organization, 
        prohibit_login, 
        avatar, 
        avatar_email, 
        use_custom_avatar, 
        num_followers, 
        num_following, 
        num_stars, 
        num_repos, 
        num_teams, 
        num_members, 
        visibility, 
        repo_admin_change_team_access, 
        diff_view_style, 
        theme, 
        is_operator,
        num_dataset_stars)
	    VALUES (
        NEW.id, 
        NEW.lower_name,
        NEW.name, 
        NEW.full_name, 
        NEW.email, 
        NEW.keep_email_private, 
        NEW.email_notifications_preference, 
        NEW.must_change_password, 
        NEW.login_type, 
        NEW.login_source, 
        NEW.login_name, 
        NEW.type, 
        NEW.location, 
        NEW.website, 
        NEW.rands, 
        NEW.language, 
        NEW.description, 
        NEW.created_unix, 
        NEW.updated_unix, 
        NEW.last_login_unix, 
        NEW.last_repo_visibility, 
        NEW.max_repo_creation, 
        NEW.is_active, 
        NEW.is_restricted, 
        NEW.allow_git_hook, 
        NEW.allow_import_local, 
        NEW.allow_create_organization, 
        NEW.prohibit_login, 
        NEW.avatar, 
        NEW.avatar_email, 
        NEW.use_custom_avatar, 
        NEW.num_followers, 
        NEW.num_following, 
        NEW.num_stars, 
        NEW.num_repos, 
        NEW.num_teams, 
        NEW.num_members, 
        NEW.visibility, 
        NEW.repo_admin_change_team_access, 
        NEW.diff_view_style, 
        NEW.theme, 
        NEW.is_operator,
        NEW.num_dataset_stars
        );

        RETURN NEW;
    END;
$def$ 
LANGUAGE plpgsql;



CREATE TRIGGER es_insert_user
    AFTER INSERT ON public.user
    FOR EACH ROW EXECUTE PROCEDURE insert_user_data();

ALTER TABLE public.user ENABLE ALWAYS TRIGGER es_insert_user;

DROP TRIGGER  IF EXISTS es_update_user on public.user;

CREATE OR REPLACE FUNCTION public.update_user() RETURNS trigger AS 
$def$
    BEGIN
	   UPDATE public.user_es  
       SET description=NEW.description,
       name=NEW.name,
       full_name=NEW.full_name,
       location=NEW.location,
       website=NEW.website,
       email=NEW.email,
       num_dataset_stars=NEW.num_dataset_stars,
       updated_unix=NEW.updated_unix
       where id=NEW.id;
       return new;
    END
$def$ 
LANGUAGE plpgsql;



CREATE TRIGGER es_update_user
    AFTER UPDATE  ON public.user
    FOR EACH ROW EXECUTE PROCEDURE update_user();

ALTER TABLE public.user ENABLE ALWAYS TRIGGER es_update_user;

DROP TRIGGER IF EXISTS es_delete_user on public.user;

CREATE OR REPLACE FUNCTION public.delete_user() RETURNS trigger AS 
$def$
    declare
    BEGIN
	   DELETE FROM public.user_es  where id=OLD.id;
       return new;
    END
$def$ 
LANGUAGE plpgsql;


CREATE TRIGGER es_delete_user
    AFTER DELETE ON public.user
    FOR EACH ROW EXECUTE PROCEDURE delete_user();
	
ALTER TABLE public.user ENABLE ALWAYS TRIGGER es_delete_user;