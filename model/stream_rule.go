package model

type StreamModelRequest struct {
	Add []StreamModel `json:"add"`
}

type StreamModel struct {
	Value string `json:"value"`
	Tag   string `json:"tag"`
}
