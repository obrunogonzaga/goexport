package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	"os"
	"sync"
)

var (
	s3Client *s3.Client
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(
			aws.NewCredentialsCache(
				credentials.NewStaticCredentialsProvider(
					"",
					"",
					""))),
	)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	s3Client = s3.NewFromConfig(cfg)
	s3Bucket = "goexpert-s3-bucket"

}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		log.Fatalf("Error opening directory: %v", err)
	}
	defer dir.Close()

	uploadControl := make(chan struct{}, 100)
	errorFileUpload := make(chan string, 10)

	go func() {
		for {
			select {
			case filename := <-errorFileUpload:
				wg.Add(1)
				uploadControl <- struct{}{}
				fmt.Printf("Error uploading file: %s\n", filename)
				uploadFile(filename, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Println("Uploading file: ", completeFileName)
	file, err := os.Open(completeFileName)
	if err != nil {
		<-uploadControl // remove from the channel
		errorFileUpload <- filename
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		<-uploadControl // remove from the channel
		errorFileUpload <- filename
		fmt.Printf("Error opening file: %v\n", err)
	}
	fmt.Printf("File uploaded: %s\n", filename)
	<-uploadControl // remove from the channel
}
