package entity

type ModelsManagerEvent struct {
	EventType string `json:"eventType"`
	Data map[string]interface{} `json:"data"`
}

type ModelsManagerLocation struct {
	Id int `json:"id"`
	ManaLevel int `json:"manaLevel"`
}
