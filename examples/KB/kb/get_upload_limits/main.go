// Print the instance's upload limits with GetUploadLimits: the most files one
// upload request may carry and the largest file size accepted.
//
// Usage (from examples/):
//
//	go run ./KB/kb/get_upload_limits <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"enterprise_search/auth"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./KB/kb/get_upload_limits <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	sdk, err := auth.NewOAuthClient()
	if err != nil {
		log.Fatal(err)
	}

	res, err := sdk.KnowledgeBase.GetUploadLimits(context.Background())
	if err != nil {
		log.Fatalf("get upload limits: %v", err)
	}
	if res.UploadLimitsResponseSchema == nil {
		log.Fatal("get upload limits: empty response")
	}
	limits := res.UploadLimitsResponseSchema

	fmt.Printf("Max files per upload request: %d\n", limits.MaxFilesPerRequest)
	fmt.Printf("Max file size:                %d bytes (%.1f MiB)\n", limits.MaxFileSizeBytes, float64(limits.MaxFileSizeBytes)/(1<<20))
}
