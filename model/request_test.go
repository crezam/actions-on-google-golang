package model_test

import (
	"encoding/json"
	"github.com/crezam/actions-on-google-golang/model"
	"os"
	"testing"
)

func TestRequestParsing(t *testing.T) {

	var req model.ApiAiRequest

	file, _ := os.Open("./data/sample_request1.json")
	dec := json.NewDecoder(file)

	if err := dec.Decode(&req); err != nil {
		t.Fatal(err)
	}

	if resolvedQuery := req.Result.ResolvedQuery; resolvedQuery != "Hi, my name is Sam!" {
		t.Fatal("Failed parsing resolved query")
	}

}
