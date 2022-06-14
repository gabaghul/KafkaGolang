package model

import "time"

type SetStreamRuleRequest struct {
	Add []SetStreamRule `json:"add"`
}

type SetStreamRule struct {
	Value string `json:"value"`
	Tag   string `json:"tag"`
}

type GetStreamRuleResponse struct {
	Data []StreamRule `json:"data"`
	Meta RuleMetadata `json:"meta"`
}

type StreamRule struct {
	ID    string `json:"id"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
}

type RuleMetadata struct {
	Sent        time.Time `json:"sent"`
	ResultCount int64     `json:"result_count"`
}

type DeleteStreamRuleRequest struct {
	Delete StremRuleDelete `json:"delete"`
}

type StremRuleDelete struct {
	IDs []string `json:"ids"`
}
