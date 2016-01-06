package model

type Result struct {
	Solution string `json:"solution"`
	Status   bool `json:"status"`
	Error    string `json:"error"`
	Info	 *Info `json:"info"`
}


func SuccessResult(data string) Result {
	return Result{data, true, "", nil}
}