-- +migrate Up
ALTER TYPE hosting_provider RENAME VALUE 'Others' TO 'Other';
ALTER TYPE ids_ips_solution RENAME VALUE 'Others' TO 'Other';
ALTER TYPE mfa RENAME VALUE 'Others' TO 'Other';

-- +migrate Down
