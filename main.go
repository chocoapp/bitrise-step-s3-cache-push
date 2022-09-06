package main

import (
	"log"
	"os"
)

const (
	CACHE_AWS_ACCESS_KEY_ID     = "cache_aws_access_key_id"
	CACHE_AWS_SECRET_ACCESS_KEY = "cache_aws_secret_access_key"
	CACHE_AWS_REGION            = "cache_aws_region"
	CACHE_BUCKET_NAME           = "cache_bucket_name"
	CACHE_KEY                   = "cache_key"
	CACHE_PATH                  = "cache_path"
)

func main() {
	awsAccessKeyId := GetEnvOrExit(CACHE_AWS_ACCESS_KEY_ID)
	awsSecretAccessKey := GetEnvOrExit(CACHE_AWS_SECRET_ACCESS_KEY)
	awsRegion := GetEnvOrExit(CACHE_AWS_REGION)
	bucketName := GetEnvOrExit(CACHE_BUCKET_NAME)
	cacheKey := GetEnvOrExit(CACHE_KEY)
	cachePath := GetEnvOrExit(CACHE_PATH)

	s3 := NewAwsS3(
		awsRegion,
		awsAccessKeyId,
		awsSecretAccessKey,
		bucketName,
	)

	log.Printf("Checking if cache exists for key '%s'\n", cacheKey)
	cacheExists := s3.CacheExists(cacheKey)

	if cacheExists {
		log.Println("Cache found! Skiping...")
		return
	}

	log.Println("Cache not found, trying to upload the folder.")

	err := s3.UploadToAws(
		cacheKey,
		cachePath,
	)

	if err != nil {
		log.Printf("Failed to upload! Failing gracefully. Error: %s\n", err)
		return
	}

	log.Println("Upload was successful!")

	os.Exit(0)
}
