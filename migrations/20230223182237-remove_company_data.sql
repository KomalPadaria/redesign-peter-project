-- +migrate Up
-- +migrate StatementBegin
create or replace function remove_company_data(cname VARCHAR)
returns void
language plpgsql
as
$$
declare
    company_id uuid;
begin
    select company_uuid into company_id from companies where "name" = cname;

    delete from application_envs where company_uuid = company_id;
    delete from tech_info_applications where company_uuid = company_id;
    delete from tech_info_ip_ranges where company_uuid = company_id;
    delete from tech_info_wireless where company_uuid = company_id;
    delete from company_facilities where company_uuid = company_id;
    delete from company_address where company_uuid = company_id;
    delete from company_meetings where company_uuid = company_id;
    delete from company_signatures where company_uuid = company_id;
    delete from policy_histories ph where policy_uuid in (select policy_uuid from policies p where company_uuid = company_id);
    delete from policies p where company_uuid = company_id;
    delete from questionnaire_answers where company_uuid = company_id;

    update companies set onboarding=jsonb_set(onboarding::jsonb, '{0,status}', '"DRAFT"') where company_uuid = company_id;
    update companies set onboarding=jsonb_set(onboarding::jsonb, '{1,status}', '"DRAFT"') where company_uuid = company_id;
    update companies set onboarding=jsonb_set(onboarding::jsonb, '{2,status}', '"DRAFT"') where company_uuid = company_id;
    update companies set onboarding=jsonb_set(onboarding::jsonb, '{3,status}', '"DRAFT"') where company_uuid = company_id;
    update companies set onboarding=jsonb_set(onboarding::jsonb, '{4,status}', '"DRAFT"') where company_uuid = company_id;
    update companies set onboarding=jsonb_set(onboarding::jsonb, '{5,status}', '"DRAFT"') where company_uuid = company_id;
    update companies set onboarding=jsonb_set(onboarding::jsonb, '{6,status}', '"DRAFT"') where company_uuid = company_id;
end;
$$;
-- +migrate StatementEnd

-- +migrate Down
