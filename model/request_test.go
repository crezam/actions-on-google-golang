package model_test

import (
	"encoding/json"
	"github.com/crezam/actions-on-google-golang/internal/test"
	"github.com/crezam/actions-on-google-golang/model"
	"os"
	"testing"
	"time"
)

func TestRequest1Parsing(t *testing.T) {

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

	test.Equals(t, "greetings", req.Result.Contexts[0].Name)
	test.Equals(t, "Sam", req.Result.Contexts[0].Parameters["user_name"])
	test.Equals(t, "Sam!", req.Result.Contexts[0].Parameters["user_name.original"])

	test.Equals(t, "373a354b-c15a-4a60-ac9d-a9f2aee76cb4", req.Result.Metadata.IntentID)
	test.Equals(t, "true", req.Result.Metadata.WebhookUsed)
	test.Equals(t, "greetings", req.Result.Metadata.IntentName)

	test.Equals(t, "Nice to meet you, Sam!", req.Result.Fulfillment.Speech)

	test.Equals(t, float64(1), req.Result.Score)

	test.Equals(t, "...", req.OriginalRequest.Data.User.UserID)
	test.Equals(t, "Sam", req.OriginalRequest.Data.User.Profile.DisplayName)
	test.Equals(t, "Sam", req.OriginalRequest.Data.User.Profile.GivenName)
	test.Equals(t, "Johnson", req.OriginalRequest.Data.User.Profile.FamilyName)

	test.Equals(t, "...", req.OriginalRequest.Data.User.AccessToken)

	test.Equals(t, 123.456, req.OriginalRequest.Data.Device.Location.Coordinates.Latitude)
	test.Equals(t, -123.456, req.OriginalRequest.Data.Device.Location.Coordinates.Longitude)

	test.Equals(t, "1234 Random Road, Anytown, CA 12345, United States", req.OriginalRequest.Data.Device.Location.FormattedAddress)
	test.Equals(t, "12345", req.OriginalRequest.Data.Device.Location.ZipCode)
	test.Equals(t, "Anytown", req.OriginalRequest.Data.Device.Location.City)

	test.Equals(t, 200, req.Status.Code)
	test.Equals(t, "success", req.Status.ErrorType)

	test.Equals(t, "37151f7c-a409-48b8-9890-cd980cd2548e", req.SessionID)
}

func TestRequest2Parsing(t *testing.T) {

	var req model.ApiAiRequest

	file, _ := os.Open("./data/sample_request2.json")
	dec := json.NewDecoder(file)

	err := dec.Decode(&req)

	// test if any issues decoding file
	test.Ok(t, err)

	// assert correct parsing
	test.Equals(t, "f4b72ee9-cabb-4acd-af9b-2d2cb6ff53d2", req.Id)

	expectedTimestamp, _ := time.Parse(time.RFC3339Nano, "2016-12-24T07:23:46.64Z")
	test.Equals(t, expectedTimestamp, req.Timestamp)

	test.Equals(t, "i wonder if you are ready for me?", req.Result.ResolvedQuery)
	test.Equals(t, "agent", req.Result.Source)
	test.Equals(t, "get_status", req.Result.Action)
	test.Equals(t, false, req.Result.ActionIncomplete)
	test.Equals(t, map[string]string{}, req.Result.Parameters)

	//TODO test Context
	//test.Equals(t, []interface{}, req.Result.Contexts)

	test.Equals(t, "dd7f8d83-6f10-474f-b16c-d3dd3d071730", req.Result.Metadata.IntentID)
	test.Equals(t, "true", req.Result.Metadata.WebhookUsed)
	test.Equals(t, "false", req.Result.Metadata.WebhookForSlotFillingUsed)
	test.Equals(t, "get_status", req.Result.Metadata.IntentName)

	test.Equals(t, "", req.Result.Fulfillment.Speech)
	test.Equals(t, 0, req.Result.Fulfillment.Messages[0].Type)
	test.Equals(t, "", req.Result.Fulfillment.Messages[0].Speech)
	test.Equals(t, 0.31, req.Result.Score)

	test.Equals(t, 200, req.Status.Code)
	test.Equals(t, "success", req.Status.ErrorType)

	test.Equals(t, "1482564192341", req.SessionID)

	test.Equals(t, "google", req.OriginalRequest.Source)
	test.Equals(t, "i wonder if you are ready for me?", req.OriginalRequest.Data.Inputs[0].Arguments[0].RawText)
	test.Equals(t, "i wonder if you are ready for me?", req.OriginalRequest.Data.Inputs[0].Arguments[0].TextValue)
	test.Equals(t, "text", req.OriginalRequest.Data.Inputs[0].Arguments[0].Name)




}

