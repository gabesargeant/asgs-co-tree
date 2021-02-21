package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//Region region struct
type Region struct {
	RegionID       string `json:"RegionID,omitempty"`
	parentRegionID string
	RegionName     string            `json:"RegionName,omitempty"`
	LevelType      string            `json:"LevelType,omitempty"`
	ChildRegions   map[string]Region `json:"ChildRegions,omitempty"`
}

//LevelName, Level Code
var asgsRegionMap = make(map[string]Region)
var lgaRegionMap = make(map[string]Region)
var gccsaRegionMap = make(map[string]Region)
var sscRegionMap = make(map[string]Region)
var poaRegionMap = make(map[string]Region)

var regionMap = map[string]map[string]Region{
	"asgs":  asgsRegionMap,
	"lga":   lgaRegionMap,
	"gccsa": gccsaRegionMap,
	"ssc":   sscRegionMap,
	"poa":   poaRegionMap,
}

//regionsetsmap[parentcode][childcode][child]
var asgsRegionSetsMap = make(map[string]map[string]Region)
var lgaRegionSetsMap = make(map[string]map[string]Region)
var gccsaRegionSetsMap = make(map[string]map[string]Region)
var sscRegionSetsMap = make(map[string]map[string]Region)
var poaRegionSetsMap = make(map[string]map[string]Region)

var regionSetsMaps = map[string]map[string]map[string]Region{
	"asgs":  asgsRegionSetsMap,
	"lga":   lgaRegionSetsMap,
	"gccsa": gccsaRegionSetsMap,
	"ssc":   sscRegionSetsMap,
	"poa":   poaRegionSetsMap,
}

var totalRegions float64 = 0

var asgsLevelSeq = []string{
	"AUS",
	"STE",
	"SA4",
	"SA3",
	"SA2",
	"SA1",
	"MB",
}

var gccsaLevelSeq = []string{
	"AUS",
	"STE",
	"GCCSA",
}

var lgaLevelSeq = []string{
	"AUS",
	"STE",
	"LGA",
}

var sscLevelSeq = []string{
	"AUS",
	"STE",
	"SSC",
	//"MB",
}

var poaLevelSeq = []string{
	"AUS",
	"STE",
	"POA",
	//"MB",
}

var levelSequences = map[string][]string{
	"asgs":  asgsLevelSeq,
	"lga":   lgaLevelSeq,
	"ssc":   sscLevelSeq,
	"poa":   poaLevelSeq,
	"gccsa": gccsaLevelSeq,
}

var levelCodeMap = map[string]string{

	"MB":    "MB_CODE_2016",
	"SA1":   "SA1_MAINCODE_2016",
	"SA2":   "SA2_MAINCODE_2016",
	"SA3":   "SA3_CODE_2016",
	"SA4":   "SA4_CODE_2016",
	"STE":   "STATE_CODE_2016",
	"AUS":   "AUS_CODE_2016",
	"LGA":   "LGA_CODE_2020",
	"POA":   "POA_CODE_2016",
	"SSC":   "SSC_CODE_2016",
	"GCCSA": "GCCSA_CODE_2016",
}

var levelNameMap = map[string]string{

	"MB":    "MB_CATEGORY_NAME_2016",
	"SA1":   "SA1_NAME_2016",
	"SA2":   "SA2_NAME_2016",
	"SA3":   "SA3_NAME_2016",
	"SA4":   "SA4_NAME_2016",
	"STE":   "STATE_NAME_2016",
	"AUS":   "AUS_NAME_2016",
	"LGA":   "LGA_NAME_2020",
	"POA":   "POA_NAME_2016",
	"SSC":   "SSC_NAME_2016",
	"GCCSA": "GCCSA_NAME_2016",
}

var asgsParentLevel = map[string]string{
	"MB":  "SA1",
	"SA1": "SA2",
	"SA2": "SA3",
	"SA3": "SA4",
	"SA4": "STE",
	"STE": "AUS",
	"AUS": "",
}

var lgaParentLevel = map[string]string{
	"LGA": "STE",
	"STE": "AUS",
	"AUS": "",
}

var gccsaParentLevel = map[string]string{
	"GCCSA": "STE",
	"STE":   "AUS",
	"AUS":   "",
}

var sscParentLevel = map[string]string{
	"MB":  "SSC",
	"SSC": "STE",
	"STE": "AUS",
	"AUS": "",
}
var poaParentLevel = map[string]string{
	"MB":  "POA",
	"POA": "STE",
	"STE": "AUS",
	"AUS": "",
}

var parentLevel = map[string]map[string]string{
	"asgs":  asgsParentLevel,
	"lga":   lgaParentLevel,
	"gccsa": gccsaParentLevel,
	"ssc":   sscParentLevel,
	"poa":   poaParentLevel,
}

var asgsRegionArray = []string{
	"MB_CODE_2016",
	"MB_CATEGORY_NAME_2016",
	"SA1_MAINCODE_2016",
	"SA1_NAME_2016",
	"SA2_MAINCODE_2016",
	"SA2_NAME_2016",
	"SA3_CODE_2016",
	"SA3_NAME_2016",
	"SA4_CODE_2016",
	"SA4_NAME_2016",
	"STATE_CODE_2016",
	"STATE_NAME_2016",
	"AUS_CODE_2016",
	"AUS_NAME_2016",
	"LGA_CODE_2020",
	"LGA_NAME_2020",
	"POA_CODE_2016",
	"POA_NAME_2016",
	"SSC_CODE_2016",
	"SSC_NAME_2016",
	"GCCSA_CODE_2016",
	"GCCSA_NAME_2016",
}

//getHeaderMap -
func getHeaderMap(firstLine []string) map[string]int {

	m := make(map[string]int)

	for i := 0; i < len(asgsRegionArray); i++ {

		for j := 0; j < len(firstLine); j++ {

			areaValue := asgsRegionArray[i]

			header := firstLine[j]

			if areaValue == header {
				m[areaValue] = j
			}

		}
	}

	return m
}

func buildNodeSet(headerMap map[string]int, row []string, level string, levelSequence []string) {

	for _, currentLevel := range levelSequence {

		levelCode := levelCodeMap[currentLevel]
		levelName := levelNameMap[currentLevel]

		//instanceLevelCode
		iLevelCode := row[headerMap[levelCode]]
		iLevelName := row[headerMap[levelName]]

		region := regionMap[level][iLevelCode]

		if region.RegionID == "" {

			//region.LevelIDName = levelCode
			region.LevelType = currentLevel

			region.RegionName = iLevelName
			region.RegionID = iLevelCode
			region.ChildRegions = make(map[string]Region)

		}

		//Add parent relationship
		parent := parentLevel[level][currentLevel]
		//fmt.Printf("Parent %d ", parent)
		if parent == "" {

			regionMap[level][iLevelCode] = region
			continue
		}
		parentRegionID := row[headerMap[levelCodeMap[parent]]]

		//fmt.Printf("region: %s \n", region.RegionID)
		//fmt.Printf("parent: %s \n", parentRegionID)

		region.parentRegionID = parentRegionID

		regionMap[level][iLevelCode] = region

	}

}

func buildNodes(headerMap map[string]int, r *csv.Reader) {

	//outter loop, read a row.
	for {
		row, err := r.Read()
		if err == io.EOF {
			//for the last record, write the buffer without the commaNewLine
			break
		}

		for level, levelSequence := range levelSequences {

			buildNodeSet(headerMap, row, level, levelSequence)
		}

	}
}

func buildTree() {

	fmt.Printf("Attempting to read %d \n", len(regionMap))

	sortNodes()

	for level, rgnMap := range regionMap {
		fmt.Printf("Level %s \n", level)
		rootNode := rgnMap["AUS"]
		getChild(level, rootNode)
		printRegion(level, rootNode)
	}

}

var tick int = 0

func sortNodes() {

	for level, rgnMap := range regionMap {

		for _, region := range rgnMap {

			childrenset := regionSetsMaps[level][region.parentRegionID]
			if childrenset == nil {
				childrenset = make(map[string]Region)
			}
			//fmt.Println(len(childrenset))
			childrenset[region.RegionID] = region
			regionSetsMaps[level][region.parentRegionID] = childrenset

		}
	}

}

func getChild(level string, root Region) {

	childRegionSet := regionSetsMaps[level][root.RegionID]

	for _, childRegion := range childRegionSet {

		root.ChildRegions[childRegion.RegionID] = childRegion
		getChild(level, childRegion)
	}
}

func printRegion(id string, out Region) {

	dataFile, err := os.Create(outfolder + id + ".json")
	bw := bufio.NewWriter(dataFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(9)
	}

	var jsonData []byte
	jsonData, err = json.MarshalIndent(out, "", "\t")

	if err != nil {
		fmt.Println(err)
		os.Exit(9)
	}
	bw.Write(jsonData)
	bw.Flush()
	dataFile.Close()

}
