package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

//Arguments for the program
type Arguments struct {
	DynamoDBTableName *string
	InputFile         *string
	OutputDir         *string
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

	readCSV(inputFile)

}

func readCSV(file *os.File) {

	for name , levelSequence := range levelSequenceSets{
	file.Seek(0, io.SeekStart)
	br := bufio.NewReader(file)
	r := csv.NewReader(br)

	firstLine, err := r.Read()
	if err != nil {
		panic(err)
	}

	headMap := getHeaderMap(firstLine)

	buildNodes(name, headMap, r, levelSequence)


	}

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
	return a
}
