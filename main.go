package main

import (
	"flag"
	"fmt"

	"github.com/davidkupyn/iris_classification/models"
)

func main() {
	var goLearn, goro bool

	flag.BoolVar(&goLearn, "gl", false, "Set goLearn to true")
	flag.BoolVar(&goro, "gr", false, "Set Goro to true")
	flag.Parse()

	if goLearn {
		models.GoLearn()
	}

	if goro || !goLearn {
		models.NN()
	}
	
	if !goLearn && !goro {
		fmt.Println("Neither goLearn nor Goro is set.")
		fmt.Println()
		flag.PrintDefaults()
	}
}
