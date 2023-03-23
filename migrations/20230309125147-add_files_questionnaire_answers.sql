-- +migrate Up
ALTER TABLE public.questionnaire_answers ADD files jsonb;

-- +migrate Down
ALTER TABLE public.questionnaire_answers DROP files;
