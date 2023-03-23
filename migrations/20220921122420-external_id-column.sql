-- +migrate Up
alter table companies
    add external_id text unique;

alter table users
    add external_id text unique;

-- +migrate Down
alter table companies drop external_id;

alter table users drop external_id;