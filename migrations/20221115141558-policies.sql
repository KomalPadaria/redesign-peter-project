-- +migrate Up
CREATE TABLE public.policy_templates (
    policy_template_uuid uuid NOT NULL,
    "name" text NOT NULL,
    document text NULL,
    industry_type industry_type[] NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT policy_templates_pkey PRIMARY KEY (policy_template_uuid)
);

ALTER TABLE public.policy_templates ADD CONSTRAINT fk_created_by_users FOREIGN KEY (created_by) REFERENCES public.users(user_uuid);
ALTER TABLE public.policy_templates ADD CONSTRAINT fk_updated_by_users FOREIGN KEY (updated_by) REFERENCES public.users(user_uuid);

CREATE TYPE public.policies_status AS ENUM(
    'Draft', 'Approved', 'Submitted', 'Rejected'
);

CREATE TABLE public.policies (
    policy_uuid uuid NOT NULL,
    company_uuid uuid NOT NULL,
    policy_template_uuid uuid NULL,
    "name" text NOT NULL,
    status policies_status NOT NULL DEFAULT 'Draft',
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT policy_policies_pkey PRIMARY KEY (policy_uuid)
);

ALTER TABLE public.policies ADD CONSTRAINT fk_created_by_users FOREIGN KEY (created_by) REFERENCES public.users(user_uuid);
ALTER TABLE public.policies ADD CONSTRAINT fk_updated_by_users FOREIGN KEY (updated_by) REFERENCES public.users(user_uuid);
ALTER TABLE public.policies ADD CONSTRAINT fk_company FOREIGN KEY (company_uuid) REFERENCES public.companies(company_uuid);
ALTER TABLE public.policies ADD CONSTRAINT fk_policy_templates FOREIGN KEY (policy_template_uuid) REFERENCES public.policy_templates(policy_template_uuid);

CREATE TABLE public.policy_histories (
    policy_history_uuid uuid NOT NULL,
    policy_uuid uuid NOT NULL,
    document text NULL,
    comment text NULL,
    version smallint NOT NULL DEFAULT 0,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT policy_histories_pkey PRIMARY KEY (policy_history_uuid)
);

ALTER TABLE public.policy_histories ADD CONSTRAINT fk_created_by_users FOREIGN KEY (created_by) REFERENCES public.users(user_uuid);
ALTER TABLE public.policy_histories ADD CONSTRAINT fk_updated_by_users FOREIGN KEY (updated_by) REFERENCES public.users(user_uuid);
ALTER TABLE public.policy_histories ADD CONSTRAINT fk_policies FOREIGN KEY (policy_uuid) REFERENCES public.policies(policy_uuid);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION policy_histories_before_insert() RETURNS TRIGGER AS $BODY$
DECLARE
v smallint;
begin
select coalesce(max(version) + 1,1) into v from policy_histories where policy_uuid = NEW.policy_uuid;
NEW.version = v;
RETURN NEW;
END;
$BODY$ LANGUAGE 'plpgsql';
-- +migrate StatementEnd

CREATE TRIGGER policy_histories_before_insert_trigger
    BEFORE
        INSERT ON policy_histories
    FOR EACH ROW EXECUTE PROCEDURE policy_histories_before_insert();

-- +migrate Down
ALTER TABLE public.policies DROP CONSTRAINT IF EXISTS fk_created_by_users;
ALTER TABLE public.policies DROP CONSTRAINT IF EXISTS fk_updated_by_users;
ALTER TABLE public.policies DROP CONSTRAINT IF EXISTS fk_company;
ALTER TABLE public.policies DROP CONSTRAINT IF EXISTS fk_policy_templates;

ALTER TABLE public.policy_histories DROP CONSTRAINT IF EXISTS fk_created_by_users;
ALTER TABLE public.policy_histories DROP CONSTRAINT IF EXISTS fk_updated_by_users;
ALTER TABLE public.policy_histories DROP CONSTRAINT IF EXISTS fk_policies;

DROP TABLE IF EXISTS public.policy_templates;
DROP TABLE IF EXISTS public.policy_histories;
DROP TABLE IF EXISTS public.policies;

DROP TYPE IF EXISTS policies_status;

