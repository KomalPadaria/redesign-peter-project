-- +migrate Up
CREATE TYPE public.control_remediations_status_type AS ENUM(
    'pending', 'completed'
);

ALTER TABLE public.control_remediations ADD status control_remediations_status_type;

-- +migrate Down
ALTER TABLE public.control_remediations DROP status;
DROP TYPE IF EXISTS control_remediations_status_type;