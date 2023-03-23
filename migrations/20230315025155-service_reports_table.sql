-- +migrate Up
CREATE TYPE public.service_evidences_status_type AS ENUM(
    'ACKNOWLEDGED',
    'REQUIRED_ACKNOWLEDGEMENT'
);

CREATE TABLE public.service_evidences
(
    service_evidences_uuid     uuid NOT NULL,
    company_subscriptions_uuid uuid NOT NULL,
    completed_on               timestamptz NULL,
    acknowledged_at            timestamptz NULL,
    acknowledged_by            uuid NULL,
    data                       jsonb NULL,
    status                     service_evidences_status_type,
    PRIMARY KEY (service_evidences_uuid),
    CONSTRAINT fk_company_subscriptions FOREIGN KEY (company_subscriptions_uuid) REFERENCES company_subscriptions (company_subscriptions_uuid)
);

ALTER TABLE public.service_evidences
    ALTER COLUMN status SET DEFAULT 'REQUIRED_ACKNOWLEDGEMENT';

-- +migrate Down
