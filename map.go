package main

import "fmt"

func main() {
	trk := map[string]string{}
	fmt.Println(trk)
	//trk2 := map[string]string
	//fmt.Println(trk2)
	// trk2 is nil map, cannot add key-value pairs to it until it is initialized
	//so trk has {} which makes trk an empty map , but trk2 is nil map which is not initialized and it will give error if we try to add key-value pair to it
	trk["subham"] = "20"
	trk["sai"] = "19"
	fmt.Println(trk)
	fmt.Println(trk["subham"])
	trk["sai"] = "21"
	fmt.Println(trk)
	delete(trk, "subham")
	fmt.Println(trk)
}
