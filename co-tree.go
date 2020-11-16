package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

//Arguments for the program
type Arguments struct {
	DynamoDBTableName *string
	InputFile *string
	OutputDir *string
	S3Bucket  *string
}

//Global Variable!
var outfolder string

func main() {
	//establish args.
	//fetch them//
	fmt.Println("Start")

	args := setArgs()
	flag.Parse()

	if *args.InputFile == "" {
		fmt.Print("No input specified.")
		os.Exit(9)
	}

	inputFile := getFile(*args.InputFile)

	outfolder = *args.OutputDir

	createOutDir(outfolder)

	mp := readCSV(inputFile)

	//summarizeRegions(mp)

	pushToDatabase(mp)

	if *args.S3Bucket != "" {
		files := getFiles(outfolder)
		uploadOutput(files, *args.S3Bucket)
	}

}

func readCSV(file *os.File) map[string]AsgsRegionNode {

	br := bufio.NewReader(file)
	r := csv.NewReader(br)

	firstLine, err := r.Read()
	if err != nil {
		panic(err)
	}

	headMap := getHeaderMap(firstLine)

	buildNodes(headMap, r)
	

	mp := mergeLevels()

	return mp
	

}

func getFile(file string) *os.File {

	fmt.Printf("Attempting to read %s \n", file)
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}
	return f

}

func setArgs() Arguments {
	a := Arguments{}
	a.DynamoDBTableName = flag.String("n", "test", "Name of the DynamoDB Table")
	a.InputFile = flag.String("i", "cat.csv", "Input File for building tree")
	a.OutputDir = flag.String("o", "./out/", "Output folder, if not set defaults to pwd ./out/ .")
	a.S3Bucket = flag.String("s", "", "If not set, no upload attempted. Assumes sdk v2 configured")
	return a
}
