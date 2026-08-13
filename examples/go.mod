module enterprise_search

go 1.25.10

require (
	github.com/joho/godotenv v1.5.1
	github.com/pipeshub-ai/pipeshub-sdk-go v1.6.0
)

require golang.org/x/sync v0.8.0 // indirect

// The examples always build against the SDK source in this repo, so they stay
// correct between releases. Drop this to build against a published version.
replace github.com/pipeshub-ai/pipeshub-sdk-go => ../
