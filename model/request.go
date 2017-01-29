package model

import "time"

type Iso8601Time struct {
	time.Time
}

const iso8601Layout = "2006-01-02T15:04:05-0700"

func (isoTime *Iso8601Time) UnmarshalJSON(b []byte) (err error) {
	return
}

func (isoTime *Iso8601Time) MarshalJSON() ([]byte, error) {
	return nil, nil
}

type ApiAiRequest struct {
	Id        string      `json:"id"`
	Timestamp Iso8601Time `json:"timestamp"`
	Result    struct {
		Source           string `json:"source"`
		ResolvedQuery    string `json:"resolvedQuery"`
		Speech           string `json:"speech"`
		Action           string `json:"action"`
		ActionIncomplete bool   `json:"actionIncomplete"`
		Parameters       struct {
			Parameters map[string]string
		} `json:"parameters"`
		Contexts []struct {
			Name       string `json:name`
			Parameters struct {
				Parameters map[string]string
			} `json:parameters`
			Lifespan int `json:lifespan`
		} `json:"contexts"`
		Metadata struct {
			IntentID                  string `json:"intentId"`
			WebhookUsed               string `json:"webhookUsed"`
			WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
			IntentName                string `json:"intentName"`
		} `json:"metadata"`
		Fulfillment struct {
			Speech   string `json:"speech"`
			Messages []struct {
				Type   int    `json:"type"`
				Speech string `json:"speech"`
			} `json:"messages"`
		} `json:"fulfillment"`
		Score float64 `json:"score"`
	} `json:"result"`
	Status struct {
		Code      int    `json:"code"`
		ErrorType string `json:"errorType"`
	} `json:"status"`
	SessionID       string `json:"sessionId"`
	OriginalRequest struct {
		Source string `json:"source"`
		Data   struct {
			Inputs []struct {
				Arguments []struct {
					RawText   string `json:"raw_text"`
					TextValue string `json:"text_value"`
					Name      string `json:"name"`
				} `json:"arguments"`
				Intent    string `json:"intent"`
				RawInputs []struct {
					Query     string `json:"query"`
					InputType int    `json:"input_type"`
				} `json:"raw_inputs"`
			} `json:"inputs"`
			User struct {
				UserID  string `json:"user_id"`
				Profile struct {
					DisplayName string `json:display_name`
					GivenName   string `json:given_name`
					FamilyName  string `json:family_name`
				} `json:profile`
				AccessToken string `json:access_token`
			} `json:"user"`
			Device struct {
				Location struct {
					Coordinates struct {
						Latitude  string `json:"latitude`
						Longitude string `json:"longitude`
					} `json:"coordinates`
					FormattedAddress string `json:"formatted_address`
					ZipCode          int    `json:"zip_code"`
					City             string `json:city`
				} `json:"location"`
			} `json:"device"`
			Conversation struct {
				ConversationToken string `json:"conversation_token"`
				ConversationID    string `json:"conversation_id"`
				Type              int    `json:"type"`
			} `json:"conversation"`
		} `json:"data"`
	} `json:"originalRequest"`
}
