package model

type Result struct {
	Solution string `json:"solution"`
	Status   bool `json:"status"`
	Error    string `json:"error"`
}


func SuccessResult(data string) Result {
	return Result{data, true, ""}
}