package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

//Arguments for the program
type Arguments struct {
	InputFile *string
	OutputDir *string
}

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

	f := getFile(*args.InputFile)
	readCSV(f)

}

func readCSV(file *os.File) {

	br := bufio.NewReader(file)
	r := csv.NewReader(br)

	firstLine, err := r.Read()
	if err != nil {
		panic(err)
	}

	headMap := getHeaderMap(firstLine)

	buildLevels(headMap, r)

	//aust := levels["AUS"]["AUS"]

	// mb := levels["MB"]

	// fmt.Println("No MB Levels ")
	// fmt.Println(len(mb))

	//printLevels(aust.ChildRegions)

	mp := mergeLevels()

	createOutputRegions(mp)

	

}

func writeOutFile(region map[string]OutputAsgsRegionNode) {

	dataFile, err := os.Create("asgsjsonFile.json")
	bw := bufio.NewWriter(dataFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(9)
	}

	for _, v := range region {

		var jsonData []byte
		jsonData, err := json.MarshalIndent(v, "", "\t")
		//fmt.Println(len(jsonData))
		if err != nil {
			fmt.Println(err)
			os.Exit(9)
		}
		bw.Write(jsonData)
		bw.Flush()
	}
	// //var jsonData []byte
	// jsonData, err := json.MarshalIndent(&region, "", "\t")
	// //fmt.Println(len(jsonData))
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(9)
	// }
	// bw.Write(jsonData)
	// bw.Flush()

	dataFile.Close()

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
	a.OutputDir = flag.String("o", "", "Output folder, if not set defaults to pwd ./ .")
	return a
}

//TODO protobuff the output.
