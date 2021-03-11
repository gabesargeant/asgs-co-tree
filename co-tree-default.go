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
var totalRegions float64 = 0

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

func buildNodeSet(headerMap map[string]int, row []string, parentSeq map[string][]string, regionMap map[string]Region) map[string]Region {

	for currentLevel, parent := range parentSeq {

		levelCode := levelCodeMap[currentLevel]
		levelName := levelNameMap[currentLevel]

		//instanceLevelCode
		iLevelCode := row[headerMap[levelCode]]
		iLevelName := row[headerMap[levelName]]

		region := regionMap[iLevelCode]

		if region.RegionID == "" {

			region.LevelType = currentLevel			
			region.RegionName = iLevelName
			region.RegionID = iLevelCode
			region.ChildRegions = make(map[string]Region)

		}

		//Add parent relationship
		if (len(parent)!=0){
			parentlevel := levelCodeMap[parent[0]]
			parentRegionID := row[headerMap[parentlevel]]
			region.parentRegionID = parentRegionID
		}
		

		regionMap[iLevelCode] = region
		//fmt.Println("length ", (regionMap[currentLevel]))

	}
	return regionMap;
}

func buildNodes(name string, headerMap map[string]int, r *csv.Reader, parentSeq map[string][]string) {

	regionMap := make(map[string]Region)

	for {
		row, err := r.Read()
		if err == io.EOF {
			//for the last record, write the buffer without the commaNewLine
			break
		}
		regionMap = buildNodeSet(headerMap, row, parentSeq, regionMap)	
	}

	buildTree(name, regionMap)
}

func buildTree(name string, regionMap map[string]Region) {

	fmt.Printf("Attempting to read %d \n", len(regionMap))

	sortedParentRegions := sortNodes(regionMap);

	rootNode := regionMap["AUS"]
	node := getChild(rootNode, sortedParentRegions)

	printRegion(name, node)

}

func getChild(root Region, parentMap map[string]map[string]Region) Region {
	//fmt.Println(len(regionSetsMap))

	children := parentMap[root.RegionID]
	for i, c := range children {
		root.ChildRegions[i] = getChild(c, parentMap)
	}
	root.ChildRegions = children
	return root
}

func sortNodes(regionMap map[string]Region) map[string]map[string]Region{

	parentMap := make(map[string]map[string]Region)

	for _, region := range regionMap {

		ipmap := parentMap[region.parentRegionID]

		if(ipmap == nil){
			ipmap = make(map[string]Region)
		}

		ipmap[region.RegionID] = region;
		parentMap[region.parentRegionID] = ipmap;
	}

	return parentMap;

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
	//jsonData, err = json.Marshal(out)
	if err != nil {
		fmt.Println(err)
		os.Exit(9)
	}
	bw.Write(jsonData)
	bw.Flush()
	dataFile.Close()

}
