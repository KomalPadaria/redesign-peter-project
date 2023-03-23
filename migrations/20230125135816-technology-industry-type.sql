-- +migrate Up
ALTER TYPE public.industry_type ADD VALUE IF NOT EXISTS 'Technology';
-- +migrate Down
