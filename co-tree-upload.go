package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func getFiles(dir string) []string {

	fileInfos, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Printf("Error in accessing directory: %s", err)
	}
	var filePaths []string
	for _, file := range fileInfos {
		//fmt.Println(filepath.Join(dir, file.Name()))
		filePaths = append(filePaths, filepath.Join(dir, file.Name()))
	}

	return filePaths

}

func uploadOutput(outputfolder []string, s3BucketName string) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		//SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{Region: aws.String("ap-southeast-2")},
	}))

	uploader := s3manager.NewUploader(sess)

	for i, outputFile := range outputfolder {

		if i > 0 && i%1000 == 0 {
			percentUploaded := (len(outputfolder) / i) * 100
			fmt.Printf("Percent Uploaded %d", percentUploaded)

		}

		file, err := os.Open(outputFile)
		if err != nil {
			fmt.Printf("Error with upload of file %s" + outputFile)
		}

		defer file.Close()

		upi := s3manager.UploadInput{}

		upi.Body = file
		upi.Bucket = &s3BucketName
		upi.Key = &outputFile

		out, err := uploader.Upload(&upi)

		if err != nil {
			fmt.Printf("Error with upload of file %s" + outputFile)
			fmt.Print(out)
		}

	}

}

//Push to dynamoDB table.
func pushToDatabase(nodeSet []AsgsRegionNode) dynamodb.BatchWriteItemInput {

	

	av, err := dynamodbattribute.MarshalMap(nodeSet)

	if err != nil {
		fmt.Println("Error with unmarhalling list of nodesets")
		fmt.Println(err.Error())
		os.Exit(1)
		}

	pr := dynamodb.PutRequest{}
	pr.SetItem(av)	
	wr := dynamodb.WriteRequest{}
	bwi := dynamodb.BatchWriteItemInput{}
	bwi.SetRequestItems(&wr)

	return bwi

}

func buildKeySchema() []*dynamodb.KeySchemaElement {

	arr := make([]*dynamodb.KeySchemaElement, 2)

	//partition key
	pid := "PartitionID"
	hash := "HASH"

	//sort key
	rid := "RegionID"
	rng := "RANGE"

	kse1 := dynamodb.KeySchemaElement{}
	kse1.AttributeName = &pid
	kse1.KeyType = &hash

	kse2 := dynamodb.KeySchemaElement{}
	kse2.AttributeName = &rid
	kse2.KeyType = &rng

	arr[0] = &kse1
	arr[1] = &kse2

	return arr

}
