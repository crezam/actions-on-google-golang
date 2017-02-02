package model_test

import (
	"encoding/json"
	"github.com/crezam/actions-on-google-golang/internal/test"
	"github.com/crezam/actions-on-google-golang/model"
	"os"
	"testing"
	"time"
)

func TestRequestParsing(t *testing.T) {

	var req model.ApiAiRequest

	file, _ := os.Open("./data/sample_request1.json")
	dec := json.NewDecoder(file)

	err := dec.Decode(&req)

	// test if any issues decoding file
	test.Ok(t, err)

	// assert correct parsing
	test.Equals(t, "209eefa7-adb5-4d03-a8b9-9f7ae68a0c11", req.Id)

	expectedTimestamp, _ := time.Parse(time.RFC3339Nano, "2016-10-10T07:41:40.098Z")
	test.Equals(t, expectedTimestamp, req.Timestamp)

	test.Equals(t, "Hi, my name is Sam!", req.Result.ResolvedQuery)
	test.Equals(t, "agent", req.Result.Source)
	test.Equals(t, "greetings", req.Result.Action)
	test.Equals(t, false, req.Result.ActionIncomplete)
	test.Equals(t, "Sam", req.Result.Parameters["user_name"])
	test.Equals(t, "", req.Result.Parameters["school"])



}
