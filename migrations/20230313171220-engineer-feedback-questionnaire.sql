-- +migrate Up
ALTER TABLE public.questionnaire_answers ADD feedback jsonb;

-- +migrate Down
ALTER TABLE public.questionnaire_answers DROP feedback;