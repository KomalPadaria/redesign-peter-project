package entities

type AccountWebhookData struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	KnowBe4 string `json:"KnowBe4"`
	Rapid7  string `json:"Rapid7"`
	Jira    string `json:"Jira"`
}

type AccountSubscription struct {
	SubscriptionId string   `json:"SubscriptionId"`
	Status         string   `json:"Status"`
	AccSubId       []string `json:"AccSubId"`
	AccountId      string   `json:"AccountId"`
}

type AccountSubscriptionWebhookData []AccountSubscription
