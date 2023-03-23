package entities

type ContractResponse struct {
	TotalSize int        `json:"totalSize"`
	Done      bool       `json:"done"`
	Contracts []Contract `json:"records"`
}

type Contract struct {
	Attributes     Attributes `json:"attributes"`
	Id             string     `json:"Id"`
	ProjectC       string     `json:"Project__c"`
	SeatsLicensedC *float64   `json:"Seats_Licensed__c"`
	SeatsActiveC   *float64   `json:"Seats_Active__c"`
	StartDate      string     `json:"StartDate"`
	EndDate        string     `json:"EndDate"`
	ContractNumber string     `json:"ContractNumber"`
	Type           string     `json:"Type__c"`
}
