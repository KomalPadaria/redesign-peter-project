-- +migrate Up
ALTER TABLE public.company_signatures ADD CONSTRAINT uk_signature_uuid_company_uuid UNIQUE (signature_uuid,company_uuid);
ALTER TABLE signatures RENAME COLUMN "data" TO document_url;

-- +migrate Down
ALTER TABLE public.company_signatures DROP CONSTRAINT IF EXISTS uk_signature_uuid_company_uuid;
ALTER TABLE signatures RENAME COLUMN document_url TO "data";
