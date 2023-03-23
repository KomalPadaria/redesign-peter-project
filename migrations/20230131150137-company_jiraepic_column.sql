-- +migrate Up
ALTER TABLE public.companies ADD jira_epic_id TEXT;

-- +migrate Down
ALTER TABLE public.companies DROP jira_epic_id;
