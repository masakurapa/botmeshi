package api

// Parameter struct
type Parameter struct {
	Type      string
	Token     string
	Challenge string
	ChannelID string
	UserID    string
	Event     EventParameter
	Action    ActionParameter
}

// EventParameter struct
type EventParameter struct {
	Type string
	Text string
}

// ActionParameter struct
type ActionParameter struct {
	Name            string
	Value           string
	SelectedOptions []string
}
