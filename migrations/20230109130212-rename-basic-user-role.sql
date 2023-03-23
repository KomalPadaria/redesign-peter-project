-- +migrate Up
ALTER TYPE company_user_role RENAME VALUE 'basic' TO 'user';
-- +migrate Down
