package main

import (
	"fmt"
	"os"
	"cluster-exploration/workloadExplore"
)

func main() {
	clArgs := os.Args
	if len(clArgs) < 2 {
		fmt.Println("No option entered. Exiting program...")
		os.Exit(0)
	}

	exploreWorkloadType := clArgs[1]

	switch exploreWorkloadType {
	case "pods":
		workloadExplore.ExplorePods()
	}

}