-- +migrate Up
ALTER TYPE company_user_role RENAME VALUE 'manager' TO 'basic';
-- +migrate Down
