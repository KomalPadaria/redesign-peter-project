-- +migrate Up
alter table users
    add is_active boolean default false;
-- +migrate Down
alter table users drop is_active;