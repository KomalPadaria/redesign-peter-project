-- +migrate Up
ALTER TABLE public.frameworks ADD CONSTRAINT uk_name UNIQUE ("name");

ALTER TABLE public.company_frameworks ADD CONSTRAINT uk_company_framework UNIQUE (company_uuid, frameworks_uuid);

ALTER TABLE public.company_questionnaires ADD CONSTRAINT uk_company_category UNIQUE (company_uuid, category);

ALTER TABLE public.company_questionnaires DROP COLUMN stats;

CREATE TABLE public.questionnaire_answers_options (
    questionnaire_options_uuid uuid NOT NULL,
    questionnaire_answers_uuid uuid NOT NULL
);

ALTER TABLE public.questionnaire_answers_options ADD CONSTRAINT fk_questionnaire_answers FOREIGN KEY (questionnaire_answers_uuid) REFERENCES public.questionnaire_answers(questionnaire_answers_uuid);
ALTER TABLE public.questionnaire_answers_options ADD CONSTRAINT fk_questionnaire_options FOREIGN KEY (questionnaire_options_uuid) REFERENCES public.questionnaire_options(questionnaire_options_uuid);

ALTER TABLE public.questionnaire_answers DROP COLUMN questionnaire_options_uuid;

ALTER TABLE public.questionnaire_options DROP COLUMN value;
ALTER TABLE public.questionnaire_options ADD COLUMN value smallint;

ALTER TABLE public.questionnaire_options ADD CONSTRAINT uk_question_label UNIQUE (questionnaires_uuid, label);

ALTER TABLE public.questionnaire_answers ADD CONSTRAINT uk_company_questionnaire UNIQUE (company_uuid, questionnaires_uuid);

-- +migrate StatementBegin
-- this function questionnaire_options_before_insert() increments the value of the option by 1 for the options added under a question
-- the value column can be used as the index for the UI
CREATE OR REPLACE FUNCTION questionnaire_options_before_insert() RETURNS TRIGGER AS $BODY$
DECLARE
v smallint;
begin
select coalesce(max(value) + 1,1) into v from questionnaire_options where questionnaires_uuid = NEW.questionnaires_uuid;
NEW.value = v;
RETURN NEW;
END;
$BODY$ LANGUAGE 'plpgsql';
-- +migrate StatementEnd

CREATE TRIGGER questionnaire_options_before_insert_trigger
    BEFORE
        INSERT ON questionnaire_options
    FOR EACH ROW EXECUTE PROCEDURE questionnaire_options_before_insert();

-- +migrate Down
ALTER TABLE public.frameworks DROP CONSTRAINT IF EXISTS uk_name;

ALTER TABLE public.company_frameworks DROP CONSTRAINT IF EXISTS uk_company_framework;

ALTER TABLE public.company_questionnaires DROP CONSTRAINT IF EXISTS uk_company_category;

ALTER TABLE public.company_questionnaires ADD COLUMN stats jsonb;

ALTER TABLE public.questionnaire_answers_options DROP CONSTRAINT IF EXISTS fk_questionnaire_answers;
ALTER TABLE public.questionnaire_answers_options DROP CONSTRAINT IF EXISTS fk_questionnaire_options;

DROP TABLE IF EXISTS public.questionnaire_answers_options;

ALTER TABLE public.questionnaire_answers ADD COLUMN questionnaire_options_uuid uuid NULL;

ALTER TABLE public.questionnaire_options DROP CONSTRAINT uk_question_label;

ALTER TABLE public.questionnaire_answers DROP CONSTRAINT uk_company_questionnaire;
