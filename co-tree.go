package main

//AsgsRegionNode One region in all the regions, Let the nodes begin!
type AsgsRegionNode struct{
	RegionID string
	RegionName string
	ParentRegionID *AsgsRegionNode
	ChildRegions []string
}

func main(){
	

}