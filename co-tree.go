package main

//AsgsRegionNode One region in all the regions, Let the nodes begin!
// This is a doubly linked node. ie it points up and down. 
// I reserve the right to decide if I'm going to change this.
type AsgsRegionNode struct{
	RegionID string
	RegionName string
	ParentRegionID *AsgsRegionNode
	ChildRegions []*AsgsRegionNode
}

func main(){


}