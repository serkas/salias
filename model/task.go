package model


type Task struct  {
	Id string `json:"id" bson:"id"`
	Text string `json:"text"`
	Context string `json:"context"`
}