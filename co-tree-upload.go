package main

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func establishClient(){
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	uploader := s3manager.NewUploader(cfg)
	
}

func uploadOutput(outPutfolder string){

	//GET ALL THE FILES
	//fOR EACH, UPLOAD THEM. 

	//investigate batch uploads!	

}



