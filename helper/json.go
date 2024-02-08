package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfErr(err)
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("content-type", "Application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfErr(err)
}
