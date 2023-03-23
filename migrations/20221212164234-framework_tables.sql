-- +migrate Up
CREATE TYPE public.questionnaire_option_type AS ENUM(
    'radio', 'multi'
);

CREATE TYPE public.questionnaire_comment_type AS ENUM(
    'mandatory', 'optional', 'none'
);

CREATE TABLE public.frameworks (
    frameworks_uuid uuid NOT NULL,
    "name" text NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT frameworks_pkey PRIMARY KEY (frameworks_uuid)
);

CREATE TABLE public.questionnaires (
    questionnaires_uuid uuid NOT NULL,
    category text NOT NULL,
    question text,
    option_type questionnaire_option_type,
    comment_type questionnaire_comment_type,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT questionnaires_pkey PRIMARY KEY (questionnaires_uuid)
);

CREATE TABLE public.frameworks_questionnaires (
    frameworks_questionnaires_uuid uuid NOT NULL,
    questionnaires_uuid uuid NOT NULL,
    frameworks_uuid uuid NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT frameworks_questionnaires_pkey PRIMARY KEY (frameworks_questionnaires_uuid)
);

ALTER TABLE public.frameworks_questionnaires ADD CONSTRAINT fk_frameworks_uuid FOREIGN KEY (frameworks_uuid) REFERENCES public.frameworks(frameworks_uuid);
ALTER TABLE public.frameworks_questionnaires ADD CONSTRAINT fk_questionnaires_uuid FOREIGN KEY (questionnaires_uuid) REFERENCES public.questionnaires(questionnaires_uuid);

CREATE TABLE public.questionnaire_options (
    questionnaire_options_uuid uuid NOT NULL,
    questionnaires_uuid uuid NOT NULL,
    value text,
    label text,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT questionnaire_options_pkey PRIMARY KEY (questionnaire_options_uuid)
);

ALTER TABLE public.questionnaire_options ADD CONSTRAINT fk_questionnaires_uuid FOREIGN KEY (questionnaires_uuid) REFERENCES public.questionnaires(questionnaires_uuid);

CREATE TABLE public.framework_controls (
    framework_control_uuid uuid NOT NULL,
    frameworks_uuid uuid NOT NULL,
    topic text,
    "name" text,
    best_practices text,
    solution text,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT framework_controls_pkey PRIMARY KEY (framework_control_uuid)
);

ALTER TABLE public.framework_controls ADD CONSTRAINT fk_frameworks FOREIGN KEY (frameworks_uuid) REFERENCES public.frameworks(frameworks_uuid);


CREATE TABLE public.control_remediations (
    control_remediation_uuid uuid NOT NULL,
    company_uuid uuid NOT NULL,
    frameworks_uuid uuid NOT NULL,
    framework_control_uuid uuid NOT NULL,
    severity text,
    comment text,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT control_remediations_pkey PRIMARY KEY (control_remediation_uuid)
);

ALTER TABLE public.control_remediations ADD CONSTRAINT fk_frameworks FOREIGN KEY (frameworks_uuid) REFERENCES public.frameworks(frameworks_uuid);
ALTER TABLE public.control_remediations ADD CONSTRAINT fk_companies FOREIGN KEY (company_uuid) REFERENCES public.companies(company_uuid);
ALTER TABLE public.control_remediations ADD CONSTRAINT fk_framework_controls FOREIGN KEY (control_remediation_uuid) REFERENCES public.control_remediations(control_remediation_uuid);

CREATE TABLE public.company_frameworks (
    company_frameworks_uuid uuid NOT NULL,
    company_uuid uuid NOT NULL,
    frameworks_uuid uuid NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT company_frameworks_pkey PRIMARY KEY (company_frameworks_uuid)
);

ALTER TABLE public.company_frameworks ADD CONSTRAINT fk_frameworks FOREIGN KEY (frameworks_uuid) REFERENCES public.frameworks(frameworks_uuid);
ALTER TABLE public.company_frameworks ADD CONSTRAINT fk_companies FOREIGN KEY (company_uuid) REFERENCES public.companies(company_uuid);

CREATE TABLE public.company_questionnaires (
    company_questionnaires_uuid uuid NOT NULL,
    company_uuid uuid NOT NULL,
    category text,
    stats jsonb NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT company_questionnaires_pkey PRIMARY KEY (company_questionnaires_uuid)
);

ALTER TABLE public.company_questionnaires ADD CONSTRAINT fk_companies FOREIGN KEY (company_uuid) REFERENCES public.companies(company_uuid);

CREATE TABLE public.questionnaire_answers (
    questionnaire_answers_uuid uuid NOT NULL,
    company_uuid uuid NOT NULL,
    questionnaires_uuid uuid NOT NULL,
    questionnaire_options_uuid uuid NOT NULL,
    comment text NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT questionnaire_answers_pkey PRIMARY KEY (questionnaire_answers_uuid)
);

ALTER TABLE public.questionnaire_answers ADD CONSTRAINT fk_companies FOREIGN KEY (company_uuid) REFERENCES public.companies(company_uuid);
ALTER TABLE public.questionnaire_answers ADD CONSTRAINT fk_questionnaire_options FOREIGN KEY (questionnaire_options_uuid) REFERENCES public.questionnaire_options(questionnaire_options_uuid);
ALTER TABLE public.questionnaire_answers ADD CONSTRAINT fk_questionnaires FOREIGN KEY (questionnaires_uuid) REFERENCES public.questionnaires(questionnaires_uuid);

-- +migrate Down
DROP TABLE IF EXISTS public.questionnaire_answers;
DROP TABLE IF EXISTS public.company_questionnaires;
DROP TABLE IF EXISTS public.company_frameworks;
DROP TABLE IF EXISTS public.control_remediations;
DROP TABLE IF EXISTS public.framework_controls;
DROP TABLE IF EXISTS public.questionnaire_options;
DROP TABLE IF EXISTS public.frameworks_questionnaires;
DROP TABLE IF EXISTS public.questionnaires;
DROP TABLE IF EXISTS public.frameworks;

DROP TYPE IF EXISTS questionnaire_option_type;
DROP TYPE IF EXISTS questionnaire_comment_type;
