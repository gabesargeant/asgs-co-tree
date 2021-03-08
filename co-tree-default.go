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
	parentRegionID map[string]string
	RegionName     string                       `json:"RegionName,omitempty"`
	LevelType      string                       `json:"LevelType,omitempty"`
	ChildRegions   map[string]map[string]Region `json:"ChildRegions,omitempty"`
}

//LevelName, Level Code
var asgsRegionMap = make(map[string]Region)
var lgaRegionMap = make(map[string]Region)
var gccsaRegionMap = make(map[string]Region)
var sscRegionMap = make(map[string]Region)
var poaRegionMap = make(map[string]Region)

var regionMap = make(map[string]map[string]Region)
var regionSetsMap = make(map[string]map[string]map[string]Region)

var childrenregionsMap = make(map[string]map[string]Region)

var totalRegions float64 = 0

var asgsLevelSeq = map[string][]string{
	"AUS":   {"STE"},
	"STE":   {"SA4", "SOS", "GCCSA", "IREG", "SUA", "RA", "POA", "CED", "SED", "SSC", "ADD", "NRMR", "LGA"},
	"SA4":   {"SA3"},
	"SA3":   {"SA2"},
	"SA2":   {"SA1"},
	"SA1":   {"MB"},
	"MB":    {},
	"SOS":   {"SOSR"},
	"SOSR":  {"UCL"},
	"UCL":   {"SA1"},
	"GCCSA": {"SA4"},
	"SUA":   {"SA3"},
	"RA":    {"SA1"},
	"IRGE":  {"IARE"},
	"IARE":  {"ILOC"},
	"ILOC":  {"SA1"},
	"POA":   {},
	"CED":   {},
	"SED":   {},
	"SSC":   {},
	"ADD":   {},
	"NRMR":  {},
	"LGA":   {},
}
var asgsParentSeq = map[string][]string{
	"AUS":   {},
	"STE":   {"AUS"},
	"SA4":   {"STE"},
	"SA3":   {"SA3"},
	"SA2":   {"SA3", "SUA"},
	//"SA1":   {"SA2", "RA", "UCL", "ILOC"},
	//"MB":    {"SA1"},
	"SOS":   {"STE"},
	"SOSR":  {"SOS"},
	"UCL":   {"SOSR"},
	"GCCSA": {"STE"},
	"SUA":   {"STE"},
	"RA":    {"STE"},
	"IREG":  {"STE"},
	"IARE":  {"IREG"},
	"ILOC":  {"IARE"},
	"POA":   {"STE"},
	"CED":   {"STE"},
	"SED":   {"STE"},
	"SSC":   {"STE"},
	"ADD":   {"STE"},
	"NRMR":  {"STE"},
	"LGA":   {"STE"},
}

var parentLevel = make(map[string]map[string]string)

var levelCodeMap = map[string]string{
	"MB":    "MB_CODE_2016",
	"SA1":   "SA1_MAINCODE_2016",
	"SA2":   "SA2_MAINCODE_2016",
	"SA3":   "SA3_CODE_2016",
	"SA4":   "SA4_CODE_2016",
	"GCCSA": "GCCSA_CODE_2016",
	"STE":   "STATE_CODE_2016",
	"AUS":   "AUS_CODE_2016",
	"DZN":   "DZN_CODE_2016",
	"LGA":   "LGA_CODE_2016",
	"POA":   "POA_CODE_2016",
	"ADD":   "ADD_CODE_2016",
	"NRMR":  "NRMR_CODE_2016",
	"SSC":   "SSC_CODE_2016",
	"TR":    "TR_CODE_2016",
	"RA":    "RA_CODE_2016",
	"ILOC":  "ILOC_CODE_2016",
	"IARE":  "IARE_CODE_2016",
	"IREG":  "IREG_CODE_2016",
	"UCL":   "UCL_CODE_2016",
	"SOSR":  "SOSR_CODE_2016",
	"SOS":   "SOS_CODE_2016",
	"SUA":   "SUA_CODE_2016",
	"SED":   "SED_CODE_2016",
	"CED":   "CED_CODE_2016",
}

var levelNameMap = map[string]string{
	"MB":    "MB_CATEGORY_NAME_2016",
	"SA1":   "SA1_NAME_2016",
	"SA2":   "SA2_NAME_2016",
	"SA3":   "SA3_NAME_2016",
	"SA4":   "SA4_NAME_2016",
	"GCCSA": "GCCSA_NAME_2016",
	"STE":   "STATE_NAME_2016",
	"AUS":   "AUS_NAME_2016",
	"DZN":   "DZN_NAME_2016",
	"LGA":   "LGA_NAME_2016",
	"POA":   "POA_NAME_2016",
	"ADD":   "ADD_NAME_2016",
	"NRMR":  "NRMR_NAME_2016",
	"SSC":   "SSC_NAME_2016",
	"TR":    "TR_NAME_2016",
	"RA":    "RA_NAME_2016",
	"ILOC":  "ILOC_NAME_2016",
	"IARE":  "IARE_NAME_2016",
	"IREG":  "IREG_NAME_2016",
	"UCL":   "UCL_NAME_2016",
	"SOSR":  "SOSR_NAME_2016",
	"SOS":   "SOS_NAME_2016",
	"SUA":   "SUA_NAME_2016",
	"SED":   "SED_NAME_2016",
	"CED":   "CED_NAME_2016",
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
	"GCCSA_CODE_2016",
	"GCCSA_NAME_2016",
	"STATE_CODE_2016",
	"STATE_NAME_2016",
	"AUS_CODE_2016",
	"AUS_NAME_2016",
	"DZN_CODE_2016",
	"DZN_NAME_2016",
	"LGA_NAME_2015",
	"LGA_CODE_2015",
	"LGA_CODE_2016",
	"LGA_NAME_2016",
	"POA_CODE_2016",
	"POA_NAME_2016",
	"ADD_CODE_2016",
	"ADD_NAME_2016",
	"NRMR_CODE_2016",
	"NRMR_NAME_2016",
	"SSC_CODE_2016",
	"SSC_NAME_2016",
	"TR_CODE_2016",
	"TR_NAME_2016",
	"RA_CODE_2016",
	"RA_NAME_2016",
	"ILOC_CODE_2016",
	"ILOC_NAME_2016",
	"IARE_CODE_2016",
	"IARE_NAME_2016",
	"IREG_CODE_2016",
	"IREG_NAME_2016",
	"UCL_CODE_2016",
	"UCL_NAME_2016",
	"SOSR_CODE_2016",
	"SOSR_NAME_2016",
	"SOS_CODE_2016",
	"SOS_NAME_2016",
	"SUA_CODE_2016",
	"SUA_NAME_2016",
	"SED_CODE_2016",
	"SED_NAME_2016",
	"CED_CODE_2016",
	"CED_NAME_2016",
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

func buildNodeSet(headerMap map[string]int, row []string, parentSeq map[string][]string) {

	for currentLevel, parentRegions := range parentSeq {

		levelCode := levelCodeMap[currentLevel]
		levelName := levelNameMap[currentLevel]

		//instanceLevelCode
		iLevelCode := row[headerMap[levelCode]]
		iLevelName := row[headerMap[levelName]]

		if regionMap[currentLevel] == nil {
			regionMap[currentLevel] = make(map[string]Region)
		}

		region := regionMap[currentLevel][iLevelCode]

		if region.RegionID == "" {

			//region.LevelIDName = levelCode
			region.LevelType = currentLevel
			region.parentRegionID = make(map[string]string)
			region.RegionName = iLevelName
			region.RegionID = iLevelCode
			region.ChildRegions = make(map[string]map[string]Region)

		}

		//Add parent relationship
		//parent := parentLevel[level][currentLevel]

		for _, parentRegion := range parentRegions {
			parentRegionID := row[headerMap[levelCodeMap[parentRegion]]]

			region.parentRegionID[parentRegion] = parentRegionID
		}

		regionMap[currentLevel][iLevelCode] = region
		//fmt.Println("length ", (regionMap[currentLevel]))

	}

}

func buildNodes(headerMap map[string]int, r *csv.Reader) {

	//outter loop, read a row.
	tick = 0
	for {

		row, err := r.Read()
		if err == io.EOF {
			//for the last record, write the buffer without the commaNewLine
			break
		}

		buildNodeSet(headerMap, row, asgsParentSeq)
		//fmt.Println(tick)
		tick++
	}
}

func buildTree() {

	fmt.Printf("Attempting to read %d \n", len(regionMap))

	sortNodes()

	fmt.Printf("Level %s \n", "AUS")
	rootNode := regionMap["AUS"]["AUS"]
	node := getChild(rootNode)
	printRegion("AUS", node)

}

var tick int = 0

func sortNodes() {

	for _, rgnMap := range regionMap {

		for _, region := range rgnMap {

			for lvl, pRegion := range region.parentRegionID {

				//childregion := regionSetsMap[level][pregion]

				set := regionSetsMap[pRegion]
				if set == nil {
					set = make(map[string]map[string]Region)
					regionSetsMap[pRegion] = set
				}
				

				rSet := regionSetsMap[pRegion][lvl]
				if(rSet == nil){
					rSet = make(map[string]Region)
					regionSetsMap[pRegion][lvl] = rSet
				}

				regionSetsMap[pRegion][lvl][region.RegionID] = region

			}
		}
	}
}

func getChild(root Region) Region {
	//fmt.Println(len(regionSetsMap))

	children := regionSetsMap[root.RegionID]

	root.ChildRegions = children

	for i, childMap := range children{

		for ii, child := range childMap{
			 childMap[ii] = getChild(child)
		} 
		children[i] = childMap
	}
	return root
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
