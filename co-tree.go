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

	obj := levels["STATE"]
	fmt.Println(len(obj))
	aust := levels["AUS"]["AUS"]

	printLevels(aust.ChildRegions)

}

//DFS
func printLevels(l []*AsgsRegionNode) {

	for _, v := range l {
		 
		fmt.Println("region :" + v.RegionName +", level :" + v.LevelIDName)
		printLevels(v.ChildRegions)


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

	a.InputFile = flag.String("i", "", "Input File for building tree")
	a.OutputDir = flag.String("o", "./output/", "Output folder, if not set defaults to creating a folder ./output in pwd.")
	return a
}

//TODO protobuff the output.
