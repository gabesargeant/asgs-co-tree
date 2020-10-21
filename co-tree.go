package main

import (
	"flag"
	"fmt"
	"os"
)


//AsgsRegionNode One region in all the regions, Let the nodes begin!
// This is a doubly linked node. ie it points up and down. 
// I reserve the right to decide if I'm going to change this.
type AsgsRegionNode struct{
	RegionID string
	RegionName string
	LevelType string
	LevelIDName string
	ParentRegionID *AsgsRegionNode
	ChildRegions []*AsgsRegionNode
}

//Arguments for the program
type Arguments struct {
	InputFile *string
	OutputDir *string
}

func main(){
 //establish args.
 //fetch them//
  fmt.Print("Start");
  
  args := setArgs()
  flag.Parse();

  if *args.InputFile == "" {
	  fmt.Print("No input specified.")
	  os.Exit(9)
  }

  readFile(buildFile(*args.InputFile))

}

func buildFile(filePath string) string {
	s:= "x"
	return s
}

func readFile(file string){

}

func setArgs() Arguments {
	a := Arguments{}

	a.InputFile = flag.String("i", "", "Input File for building tree")
	a.OutputDir = flag.String("o", "./output/", "Output folder, if not set defaults to creating a folder ./output in pwd.")
	return a
}

//

//cmdline args, input file, output location.

//structures. 

//Australia,
//States
//SA4
//SA3
//SA2
//SA1 -- // CED  // SED
//MB -- LGA -- POA -- SSC

