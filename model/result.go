package model

type Result struct  {
	Solution string `json:"solution"`
	Status bool `json:"status"`
	Error string `json:"error"`
}