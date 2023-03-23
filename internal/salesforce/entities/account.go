package entities

type Account struct {
	ID            string `json:"Id"`
	Name          string `json:"Name"`
	Industry      string `json:"Industry"`
	PhotoUrl      string `json:"PhotoUrl"`
	Knowbe4Token  string `json:"KnowBe4_Access_Token__c"`
	JiraEpicId    string `json:"Jira_Epic_Id__c"`
	Rapid7SiteIds string `json:"Rapid7_InsightVM_Site_ID__c"`
}
