package model

type Task struct  {
	Id string `json:"id"`
	Text string `json:"text"`
	Context string `json:"context"`
}