package entities

type SubscriptionResponse struct {
	TotalSize     int            `json:"totalSize"`
	Done          bool           `json:"done"`
	Subscriptions []Subscription `json:"records"`
}

type Subscription struct {
	Attributes                   Attributes  `json:"attributes"`
	ID                           string      `json:"Id"`
	Name                         string      `json:"Name"`
	SBQQProductNameC             string      `json:"SBQQ__ProductName__c"`
	SBQQContractC                string      `json:"SBQQ__Contract__c"`
	SBQQProductIDC               string      `json:"SBQQ__ProductId__c"`
	SBQQStartDateC               string      `json:"SBQQ__StartDate__c"`
	SBQQEndDateC                 string      `json:"SBQQ__EndDate__c"`
	SBQQTerminatedDateC          string      `json:"SBQQ__TerminatedDate__c"`
	SBQQQuantityC                float64     `json:"SBQQ__Quantity__c"`
	SBQQProductSubscriptionTypeC string      `json:"SBQQ__ProductSubscriptionType__c"`
	SBQQProductC                 string      `json:"SBQQ__Product__c"`
	SBQQAccountC                 string      `json:"SBQQ__Account__c"`
	SBQQOptionTypeC              interface{} `json:"SBQQ__OptionType__c"`
	SBQQOrderProductC            string      `json:"SBQQ__OrderProduct__c"`
	QuantityBillableHoursC       float64     `json:"Quantity_Billable_Hours__c"`
	QuantityInHourC              float64     `json:"Quantity_In_Hour__c"`
	QuantityInHourUsedC          float64     `json:"Quantity_In_Hour_Used__c"`
}
