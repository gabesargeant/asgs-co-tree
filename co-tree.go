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

	mb := levels["MB"]

	fmt.Println("No MB Levels ")
	fmt.Println(len(mb))

	//printLevels(aust.ChildRegions)

	mp := mergeLevels()

	writeOutFile(mp)

}

func writeOutFile(region map[string]AsgsRegionNode) {

	var jsonData []byte
	jsonData, err := json.Marshal(&region)
	fmt.Println(len(jsonData))
	fmt.Println("length of json data")
	dataFile, err := os.Create("asgsjsonFile.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(9)
	}

	bw := bufio.NewWriter(dataFile)
	bw.Write(jsonData)
	bw.Flush()
	dataFile.Close()

}



//DFS
func printLevels(l []*AsgsRegionNode) {

	for _, v := range l {

		fmt.Println("region :" + v.RegionName + ", level :" + v.LevelIDName)
		if len(v.ChildRegions) != 0 {
			printLevels(v.ChildRegions)
		}

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
