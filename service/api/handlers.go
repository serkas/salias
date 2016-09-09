package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/serkas/salias/core/objects"
	"github.com/serkas/salias/core/tokenizer"
	"github.com/serkas/salias/service/model"
	"strings"
)

func Classify(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &task); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	model.InitStorage().SaveTask(&task)

	doc := objects.MakeDoc(task.Text)
	info := model.Info{len(doc.Tokens)}
	result := model.Result{
		Status: true, Error: "", Solution: "", Info: &info,
	}
	successJson(w, result)
}

func ShowTasks(w http.ResponseWriter, r *http.Request) {
	s := model.InitStorage()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s.Tasks()); err != nil {
		panic(err)
	}

}

func Analyze(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	text := r.URL.Query().Get("text")
	var res model.Result
	if text != "" {
		tokens := tokenizer.TokenizeToStrings(text)
		res = model.SuccessResult(strings.Join(tokens, ", "))
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		res = model.Result{"", false, "Bad request. Text property not set or empty", nil}
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func successJson(w http.ResponseWriter, result model.Result) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}
