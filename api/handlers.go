package api

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"io"

	"salias/model"
	"salias/tokenizer"
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
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	result := model.Result{
		Status: true, Error: "", Solution: task.Text,
	}
	successJson(w, result);
}

func Index(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func Analyze(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text != "" {
		tokens := tokenizer.TokenizeToStrings(text)
		res := model.SuccessResult(strings.Join(tokens, ", "))

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
	}else{
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		res := model.Result{"", false, "Bad request. Text property not set or empty"}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
	}
}

func successJson(w http.ResponseWriter, result model.Result) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}