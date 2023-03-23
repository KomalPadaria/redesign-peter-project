-- +migrate Up
CREATE TABLE public.company_subscriptions
(
    company_subscriptions_uuid uuid,
    company_uuid               uuid NOT NULL,
    name                       text,
    sf_subscription_id         text,
    sf_product_id              text,
    type                       text,
    sub_type                   text,
    status                     text,
    start_date                 timestamptz NULL,
    end_date                   timestamptz NULL,
    created_at                 timestamptz NULL DEFAULT now(),
    updated_at                 timestamptz NULL,
    created_by                 uuid NULL,
    updated_by                 uuid NULL,
    PRIMARY KEY (company_subscriptions_uuid),
    CONSTRAINT fk_companies FOREIGN KEY (company_uuid) REFERENCES companies (company_uuid)
);

ALTER TABLE public.company_subscriptions
    ADD CONSTRAINT cs_company_sf_subscription_id_ukey UNIQUE (company_uuid, sf_subscription_id);

-- +migrate Down
