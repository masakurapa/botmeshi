package shopsearch

type body struct {
	Callback string `json:"callback"`
	Query    string `json:"query"`
}
