package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

//Region region struct
type Region struct {
	RegionID       string `json:"RegionID,omitempty"`
	parentRegionID string
	RegionName     string             `json:"RegionName,omitempty"`
	LevelType      string             `json:"LevelType,omitempty"`
	ChildRegions   map[string]Region `json:"ChildRegions,omitempty"`
}

//LevelName, Level Code
var regionMap = make(map[string]Region)
//regionsetsmap[parentcode][childcode][child]
var regionSetsMap = make(map[string]map[string]Region)

var totalRegions float64 = 0

//The order of regions in this array is significant.
//I need to build the nodes from the MeshBlocks up.
//as each 'higher' region will require the lower to be there
//to link it. And each child gets linked it's parent when
//its parent gets its child.
var asgsLevelSequence = []string{
	"AUS",
	"STE",
	"SA4",
	"SA3",
	"SA2",
	"SA1",
	"MB",
}

var gccsaLevelSeq = []string{
	"GCCSA",
	"STE",
	"AUS",
}

var lgaLevelSeq = []string{
	"LGA",
	"AUS",
	"STATE",
}

var levelCodeMap = map[string]string{

	"MB":    "MB_CODE_2016",
	"SA1":   "SA1_MAINCODE_2016",
	"SA2":   "SA2_MAINCODE_2016",
	"SA3":   "SA3_CODE_2016",
	"SA4":   "SA4_CODE_2016",
	"STE": "STATE_CODE_2016",
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
	"STE": "STATE_NAME_2016",
	"AUS":   "AUS_NAME_2016",
	"LGA":   "LGA_NAME_2020",
	"POA":   "POA_NAME_2016",
	"SSC":   "SSC_NAME_2016",
	"GCCSA": "GCCSA_NAME_2016",
}

//Australia,
//States
//SA4
//SA3
//SA2
//SA1
//MB -- LGA -- POA -- SSC
//ChildLevles Key = the current region, value = it's child region
var asgsChildLevel = map[string]string{
	"MB":    "",
	"SA1":   "MB",
	"SA2":   "SA1",
	"SA3":   "SA2",
	"SA4":   "SA3",
	"STE": "SA4",
	"AUS":   "STATE",
}

var asgsParentLevel = map[string]string{
	"MB":   "SA1",
	"SA1":  "SA2",
	"SA2":  "SA3",
	"SA3":  "SA4",
	"SA4":  "STE",
	"STE": "AUS",
	"AUS":  "",
}

var lgaChildLevels = map[string]string{
	"LGA":   "",
	"STATE": "LGA",
	"AUS":   "STATE",
}

var gccsaChildLevels = map[string]string{
	"GCCSA": "",
	"STATE": "GCCSA",
	"AUS":   "STATE",
}

var skipLevel = map[string]string{
	"MB":  "SKIP",
	"SA1": "SKIP",
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

func buildNodes(headerMap map[string]int, r *csv.Reader) {

	//outter loop, read a row.
	for {
		row, err := r.Read()
		if err == io.EOF {
			//for the last record, write the buffer without the commaNewLine
			break
		}

		for _, currentLevel := range asgsLevelSequence {

			levelCode := levelCodeMap[currentLevel]
			levelName := levelNameMap[currentLevel]

			//instanceLevelCode
			iLevelCode := row[headerMap[levelCode]]
			iLevelName := row[headerMap[levelName]]

			region := regionMap[iLevelCode]

			if region.RegionID == "" {

				//region.LevelIDName = levelCode
				region.LevelType = currentLevel

				region.RegionName = iLevelName
				region.RegionID = iLevelCode
				region.ChildRegions = make(map[string]Region)

			}

			//Add parent relationship
			parent := asgsParentLevel[currentLevel]
			//fmt.Printf("Parent %d ", parent)
			if parent == "" {
				regionMap[iLevelCode] = region
				continue
			}
			parentRegionID :=  row[headerMap[levelCodeMap[parent]]]  
			
			//fmt.Printf("region: %s \n", region.RegionID)
			//fmt.Printf("parent: %s \n", parentRegionID)
			
			region.parentRegionID = parentRegionID

			regionMap[iLevelCode] = region

		}
	}
}

var wg sync.WaitGroup

func buildTree() {
	rootNode := regionMap["AUS"]
	fmt.Printf("Attempting to read %d \n", len(regionMap))

	sortNodes()
	
	// ste1 := regionMap["1"]
	// ste2 := regionMap["2"]
	// ste3 := regionMap["3"]
	// ste4 := regionMap["4"]
	// ste5 := regionMap["5"]
	// ste6 := regionMap["6"]
	// ste7 := regionMap["7"]
	// ste8 := regionMap["8"]
	// ste9 := regionMap["9"]
	
	// go getChildren(ste1)
	// go getChildren(ste2)
	// go getChildren(ste3)
	// go getChildren(ste4)
	// go getChildren(ste5)
	// go getChildren(ste6)
	// go getChildren(ste7)
	// go getChildren(ste8)
	// go getChildren(ste9)
	// go getRootChildren(rootNode)
	// wg.Wait()
	getChild(rootNode)
	printRegion("AUS", rootNode)

}

var tick int = 0;

func sortNodes(){

	for _, region := range regionMap {

		childrenset := regionSetsMap[region.parentRegionID]
		if childrenset == nil{
			childrenset = make(map[string]Region)
		}
		//fmt.Println(len(childrenset))
		childrenset[region.RegionID] = region;
		regionSetsMap[region.parentRegionID] = childrenset

	}
}

func getChild(root Region){

	childRegionSet := regionSetsMap[root.RegionID];

	for _, childRegion := range childRegionSet {

		root.ChildRegions[childRegion.RegionID] = childRegion
		getChild(childRegion);
	}
}

func getRootChildren(parent Region) {
	wg.Add(1)
	for _, childregion := range regionMap {
		if childregion.parentRegionID == parent.RegionID {				
			parent.ChildRegions[childregion.RegionID] = childregion
			
			tick++;
			if(tick % 1000 == 0){
				fmt.Println(tick)
				//fmt.Printf("Region map %d \n", len(regionMap));
			}
		}
	}
	wg.Done()
}

func getChildren(parent Region) {
	//wg.Add(1)
	for _, childregion := range regionMap {
		if childregion.parentRegionID == parent.RegionID {			
			
			parent.ChildRegions[childregion.RegionID] = childregion
			delete(regionMap, childregion.RegionID)
			tick++;
			if(tick % 1000 == 0){
				fmt.Println(tick)
				//fmt.Printf("Region map %d \n", len(regionMap));
			}
			getChildren(parent.ChildRegions[childregion.RegionID])

		}
	}
	//wg.Done()

}

//only for testing
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
