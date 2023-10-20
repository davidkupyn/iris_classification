package models

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

type IrisClassifer struct {
}

func (c *IrisClassifer) Fit(X, y interface{}) {

}

func NN() {
	datasetFile, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer datasetFile.Close()

	dataset := dataframe.ReadCSV(datasetFile)

	fmt.Println(dataset)
}
