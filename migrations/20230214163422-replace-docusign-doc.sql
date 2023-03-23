-- +migrate Up
WITH item AS (
    SELECT ('{' || index - 1 || ',"status"}')::TEXT[] AS path,
            company_uuid,
           'DRAFT' as val
    from companies,
         jsonb_array_elements(onboarding) WITH ORDINALITY arr(item, index)
    where item->>'url' = '/security-awareness'
    )
UPDATE companies
SET onboarding = jsonb_set(onboarding, item.path, to_json(item.val)::JSONB)
    FROM item;

INSERT INTO public.signatures (signature_uuid, "name", document_url, company_types) VALUES(gen_random_uuid(), 'Authorization', 'https://www.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=57095a85-1849-4b7e-8240-bc9da5b382f6&env=na1&acct=7171320a-7100-4375-97c7-3266016165d9&v=2', '{Entertainment}');

delete from company_signatures cs where signature_uuid = (select signature_uuid from signatures where "name" = 'Cybersecurity Questionnaire & Authorization');
delete from signatures where "name" = 'Cybersecurity Questionnaire & Authorization';
-- +migrate Down
