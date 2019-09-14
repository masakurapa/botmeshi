package api

// EventBody struct
type EventBody struct {
	Token     string     `json:"token"`
	Type      string     `json:"type"`
	Event     innerEvent `json:"event"`
	Challenge string     `json:"challenge"`
}

// innerEvent struct
type innerEvent struct {
	Type    string `json:"type"`
	Text    string `json:"text"`
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// ToParameter is convert EventBody to Parameter
func (b *EventBody) ToParameter() *Parameter {
	return &Parameter{
		Type:      b.Type,
		Token:     b.Token,
		Challenge: b.Challenge,
		ChannelID: b.Event.Channel,
		UserID:    b.Event.User,
		Event: EventParameter{
			Type: b.Event.Type,
			Text: b.Event.Text,
		},
	}
}
