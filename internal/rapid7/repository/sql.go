package repository

var (
	lastAndNextVulnerabilityScan = `select (select ds.started from dim_scan ds
join dim_site ds2 on ds2.last_scan_id = ds.scan_id 
where ds.site_id in ? limit 1) as last_vulnerability_scan,
(select ds.started from dim_scan ds
join dim_site ds2 on ds2.site_id  = ds.site_id 
where ds.site_id in ? and ds.started > now() limit 1) as next_vulnerability_scan;`

	sql = `
select DATE_TRUNC('month',date_added) as date,severity,count(*) as count  from dim_vulnerability dv where dv.vulnerability_id in (
select vulnerability_id from (
select vulnerability_id from dim_vulnerability_exception where asset_id in (select asset_id from dim_site_asset dsa where dsa.site_id in ?)
union
select vulnerability_id from dim_asset_validated_vulnerability where asset_id in (select asset_id from dim_site_asset dsa where dsa.site_id in ?)
union
select vulnerability_id from dim_asset_vulnerability_finding_solution where asset_id in (select asset_id from dim_site_asset dsa where dsa.site_id in ?)
union
select vulnerability_id from dim_asset_vulnerability_finding_rollup_solution where asset_id in (select asset_id from dim_site_asset dsa where dsa.site_id in ?)
union
select vulnerability_id from fact_asset_vulnerability_instance where asset_id in (select asset_id from dim_site_asset dsa where dsa.site_id in ?)
union
select vulnerability_id from fact_asset_vulnerability_finding where asset_id in (select asset_id from dim_site_asset dsa where dsa.site_id in ?)
union
select vulnerability_id from fact_asset_vulnerability_finding_date where asset_id in (select asset_id from dim_site_asset dsa where dsa.site_id in ?)
union
select vulnerability_id from fact_asset_vulnerability_remediation_date where asset_id in (select asset_id from dim_site_asset dsa where dsa.site_id in ?)
union
select vulnerability_id from fact_asset_vulnerability_finding_remediation where asset_id in (select asset_id from dim_site_asset dsa where dsa.site_id in ?)
union
select vulnerability_id from fact_asset_vulnerability_finding_exploit_remediation where asset_id in (select asset_id from dim_site_asset dsa where dsa.site_id in ?)
union
select vulnerability_id from fact_asset_vulnerability_finding_exploit where asset_id in (select asset_id from dim_site_asset dsa where dsa.site_id in ?)
) sq) group by DATE_TRUNC('month',date_added),severity order by DATE_TRUNC('month',date_added)
`

	topRemediationSql = `
select s.solution_id,s.vulnerability_id,s.risk_score, dv2.title, htmltotext(dv2.description) issue_description, ds.summary as remediation_name, 
dvr.reference remediation_link, htmltotext(ds.fix) solution_steps  from (
select distinct dvs.solution_id, dv.vulnerability_id,dv.risk_score from dim_vulnerability dv 
join dim_asset_vulnerability_finding_rollup_solution davfrs on davfrs.vulnerability_id =dv.vulnerability_id 
join dim_vulnerability_solution dvs on dvs.solution_id = davfrs.solution_id 
where davfrs.asset_id in (select asset_id from dim_site_asset dsa where dsa.site_id in ?)
order by dv.risk_score desc) s
join dim_vulnerability dv2 on dv2.vulnerability_id = s.vulnerability_id
join dim_solution ds on ds.solution_id = s.solution_id
join dim_vulnerability_reference dvr on dvr.vulnerability_id = dv2.vulnerability_id limit 10`

	getSolutionSupercedenceStepsSql = `
select distinct(htmltotext(ds.fix)) from dim_solution ds where ds.solution_id in (select solution_id from dim_solution_supercedence dss 
where dss.superceding_solution_id =?);
`
	listRemeditation = `
select distinct dvr.reference as recommendation_url,
    htmltotext(dv.description) as issue_description,
    dvr."source",
    dv.vulnerability_id,
    ds.solution_id,
    dv.risk_score,
    dv.severity,
    fv.vulnerability_instances as instances,
    ds.summary as issue_name,
    htmltotext(ds.fix) as recommendation
from dim_asset_vulnerability_finding_solution davfs
    join dim_site_asset dsa on dsa.asset_id = davfs.asset_id
    join dim_solution ds on ds.solution_id = davfs.solution_id
    join dim_vulnerability dv on dv.vulnerability_id = davfs.vulnerability_id
    join dim_vulnerability_reference dvr on dvr.vulnerability_id = davfs.vulnerability_id
    join fact_vulnerability fv on fv.vulnerability_id = davfs.vulnerability_id
where dsa.site_id in ?
    and fv.vulnerability_instances <> 0
order by dv.risk_score desc
limit ?
`
)
