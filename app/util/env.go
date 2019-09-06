package util

import "os"

const (
	// APIVerificationTokenKey APIトークンのキー
	APIVerificationTokenKey = "API_VERIFICATION_TOKEN"
)

// APIVerificationToken returns BOT_VERIFICATION_TOKEN
func APIVerificationToken() string {
	return os.Getenv(APIVerificationTokenKey)
}

// BotAccessToken returns BOT_ACCESS_TOKEN
func BotAccessToken() string {
	return os.Getenv("BOT_ACCESS_TOKEN")
}

// SearchEngineID returns SEARCH_ENGINE_ID
func SearchEngineID() string {
	return os.Getenv("SEARCH_ENGINE_ID")
}

// PlaceAPIKey returns PLACE_API_KEY
func PlaceAPIKey() string {
	return os.Getenv("PLACE_API_KEY")
}

// CustomSearchAPIKey returns CUSTOM_SEARCH_API_KEY
func CustomSearchAPIKey() string {
	return os.Getenv("CUSTOM_SEARCH_API_KEY")
}

// InvokeLambdaArn returns INVOKE_LAMBDA_ARN
func InvokeLambdaArn() string {
	return os.Getenv("INVOKE_LAMBDA_ARN")
}
