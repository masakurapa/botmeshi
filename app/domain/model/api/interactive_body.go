package api

// InteractiveBody struct
type InteractiveBody struct {
	Type    string              `json:"type"`
	Token   string              `json:"token"`
	Channel interactiveChannel  `json:"channel"`
	Actions []interactiveAction `json:"actions"`
}

type interactiveChannel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type interactiveAction struct {
	Name            string           `json:"name"`
	Value           string           `json:"value"`
	SelectedOptions []selectedOption `json:"selected_options"`
}

type selectedOption struct {
	Value string `json:"value"`
}

// ToParameter is convert InteractiveBody to Parameter
func (b *InteractiveBody) ToParameter() *Parameter {
	// 先頭のActionだけ使う
	var action ActionParameter
	if len(b.Actions) > 0 {
		a := b.Actions[0]
		opts := make([]string, len(a.SelectedOptions))
		for i, op := range a.SelectedOptions {
			opts[i] = op.Value
		}

		action = ActionParameter{
			Name:            a.Name,
			Value:           a.Value,
			SelectedOptions: opts,
		}
	}

	return &Parameter{
		Type:      b.Type,
		Token:     b.Token,
		ChannelID: b.Channel.ID,
		Action:    action,
	}
}
