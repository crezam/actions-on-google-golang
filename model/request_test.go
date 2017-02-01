package model_test

import (
	"encoding/json"
	"github.com/crezam/actions-on-google-golang/internal/test"
	"github.com/crezam/actions-on-google-golang/model"
	"os"
	"testing"
)

func TestRequestParsing(t *testing.T) {

	var req model.ApiAiRequest

	file, _ := os.Open("./data/sample_request1.json")
	dec := json.NewDecoder(file)

	err := dec.Decode(&req)

	// test if any issues decoding file
	test.Ok(t, err)

	// assert values in fields
	test.Equals(t, "Hi, my name is Sam!", req.Result.ResolvedQuery)
}
