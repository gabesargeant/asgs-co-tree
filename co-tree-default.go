package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//AsgsRegionNode One region in all the regions, Let the nodes begin!
// This is a doubly linked node. ie it points up and down.
//Maps not arrays for the pointers, this to avoid duplicates
type AsgsRegionNode struct {
	RegionID      string
	RegionName    string
	LevelType     string
	LevelIDName   string
	ParentRegions map[string]*AsgsRegionNode
	ChildRegionType string
	ChildRegions  map[string]string
}

//LevelName, Level Code
var levels = map[string]map[string]AsgsRegionNode{
	"MB":    {},
	"SA1":   {},
	"SA2":   {},
	"SA3":   {},
	"SA4":   {},
	"STATE": {},
	"AUS":   {},
	"LGA":   {},
	"POA":   {},
	"SSC":   {},
}

func mergeLevels() map[string]AsgsRegionNode {

	mp := make(map[string]AsgsRegionNode)
	for k, v := range levels {

		fmt.Printf("Region Type: %s, size: %d \n", k, len(v))

		for kk, vv := range v {
			mp[kk] = vv
		}

	}
	fmt.Println(len(mp))
	return mp
}

//The order of regions in this array is significant.
//I need to build the nodes from the MeshBlocks up.
//as each 'higher' region will require the lower to be there
//to link it. And each child gets linked it's parent when
//its parent gets its child.
var levelSequence = []string{
	"MB",
	"SA1",
	"SA2",
	"SA3",
	"SA4",
	"STATE",
	"AUS",
		"LGA",
	"POA",
	"SSC",
}



var levelCodeMap = map[string]string{

	"MB":    "MB_CODE_2016",
	"SA1":   "SA1_MAINCODE_2016",
	"SA2":   "SA2_MAINCODE_2016",
	"SA3":   "SA3_CODE_2016",
	"SA4":   "SA4_CODE_2016",
	"STATE": "STATE_CODE_2016",
	"AUS":   "AUS_CODE_2016",
	"LGA":   "LGA_CODE_2020",
	"POA":   "POA_CODE_2016",
	"SSC":   "SSC_CODE_2016",
}

var levelNameMap = map[string]string{

	"MB":    "MB_CATEGORY_NAME_2016",
	"SA1":   "SA1_NAME_2016",
	"SA2":   "SA2_NAME_2016",
	"SA3":   "SA3_NAME_2016",
	"SA4":   "SA4_NAME_2016",
	"STATE": "STATE_NAME_2016",
	"AUS":   "AUS_NAME_2016",
	"LGA":   "LGA_NAME_2020",
	"POA":   "POA_NAME_2016",
	"SSC":   "SSC_NAME_2016",
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
	"STATE": "SA4",
	"AUS":   "STATE",
	"LGA":   "MB",
	"POA":   "MB",
	"SSC":   "MB",
}

// //The many to many is a problem....hmm.
// var nonAsgsChildLevel = map[string][]string{
// 	"MB":    {},
// 	"STATE": {"LGA", "POA", "SSC"},
// 	"AUS":   {"STATE"},
// 	"LGA":   {"MB"},
// 	"POA":   {"MB"},
// 	"SSC":   {"MB"},
// }

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

		for _, currentLevel := range levelSequence {

			levelCode := levelCodeMap[currentLevel]
			levelName := levelNameMap[currentLevel]

			//instanceLevelCode
			iLevelCode := row[headerMap[levelCode]]
			iLevelName := row[headerMap[levelName]]

			region := levels[currentLevel][iLevelCode]

			if region.RegionID == "" {

				region.LevelIDName = levelCode
				region.LevelType = currentLevel
				region.RegionName = iLevelName
				region.RegionID = iLevelCode
				region.ChildRegions = make(map[string]string)
				region.ParentRegions = make(map[string]*AsgsRegionNode)

			}

			//Add child element
			child := asgsChildLevel[currentLevel]

			if child == "" {
				levels[currentLevel][iLevelCode] = region
				continue
			}

			childLevelCode := levelCodeMap[child]

			childRegionCode := row[headerMap[childLevelCode]]

			childRegion := levels[child][childRegionCode]

			//Establish Relationships
			childRegion.ParentRegions[region.RegionID] = &region

			region.ChildRegions[childRegion.RegionID] = childRegion.RegionID

			//Set objects
			levels[currentLevel][iLevelCode] = region
			levels[child][childRegionCode] = childRegion

		}
	}
}


// //need to cut up.
// func summarizeRegions(regions map[string]AsgsRegionNode) {

// 	fmt.Println("starting region output build")
// 	for _, v := range regions {

// 		if skipLevel[v.LevelType] == "SKIP" {
// 			continue
// 		}

// 		out := OutputAsgsRegionNode{}
// 		out.LevelType = v.LevelType
// 		out.RegionID = v.RegionID
// 		out.RegionName = v.RegionName

// 		out.ChildRegions = make(map[string]*ChildRegion)

// 		for _, val := range v.ChildRegions {

// 			cr := ChildRegion{}
// 			cr.LevelType = v.LevelType
// 			cr.RegionID = v.RegionName
// 			cr.RegionName = v.RegionName

// 			out.ChildRegions[val.RegionID] = &cr
// 		}

// 		out.ParentRegions = buildParentTree(v.ParentRegions, regions)

// 		printRegion(out.RegionID, out)

// 	}

// }


//Not needed
func createOutDir(outDir string) {

	err := os.Mkdir(outDir, 0777)
	if os.IsNotExist(err) {
		fmt.Printf("Error with creating output folder %s", err)
	}
}

//only for testing
func printRegion(id string, out AsgsRegionNode) {

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

// //not needed.
// func buildParentTree(parents map[string]*AsgsRegionNode, regions map[string]AsgsRegionNode) map[string]*ParentRegion {

// 	prArr := make(map[string]*ParentRegion)

// 	//usually there's 1 parent
// 	for _, v := range parents {

// 		pr := ParentRegion{}
// 		pr.LevelType = v.LevelType
// 		pr.RegionID = v.RegionName
// 		pr.RegionName = v.RegionName
// 		pr.ParentRegions = make(map[string]*ParentRegion)

// 		if v.RegionID == "AUS" {
// 			prArr[pr.RegionID] = &pr
// 			return prArr
// 		}

// 		//for each parent in v
// 		pt := buildParentTree(v.ParentRegions, regions)
// 		for _, val := range pt {
// 			pr.ParentRegions[val.RegionID] = val
// 		}
// 		prArr[pr.RegionID] = &pr

// 	}
// 	return prArr
// }
