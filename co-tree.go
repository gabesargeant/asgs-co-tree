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

	readCSV(inputFile)

	if *args.S3Bucket != "" {
		files := getFiles(outfolder)
		uploadOutput(files, *args.S3Bucket)
	}

}

func readCSV(file *os.File) {

	br := bufio.NewReader(file)
	r := csv.NewReader(br)

	firstLine, err := r.Read()
	if err != nil {
		panic(err)
	}

	headMap := getHeaderMap(firstLine)

	buildASGSLevels(headMap, r)
	buildNonASGSLevels(headMap, r)

	mp := mergeLevels()

	//createOutputRegions(mp)
	summarizeRegions(mp)

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

	a.InputFile = flag.String("i", "cat.csv", "Input File for building tree")
	a.OutputDir = flag.String("o", "./out/", "Output folder, if not set defaults to pwd ./out/ .")
	a.S3Bucket = flag.String("s", "", "If not set, no upload attempted. Assumes sdk v2 configured")
	return a
}
