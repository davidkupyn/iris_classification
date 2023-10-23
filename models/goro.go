package models

import (
	"fmt"
	"os"

	m "github.com/aunum/goro/pkg/v1/model"
	"github.com/go-gota/gota/dataframe"
)

func getDataset(name string) dataframe.DataFrame {
	datasetFile, err := os.Open("iris.csv")
	if err != nil {
		panic(err)
	}
	defer datasetFile.Close()

	dataset := dataframe.ReadCSV(datasetFile)

	return dataset
}

func generateXY(dataset dataframe.DataFrame) (X, Y dataframe.DataFrame) {
	X = dataset.Drop([]string{"Id", "Species"})
	Y = dataset.Select("Species")
	return X, Y
}

func NN() {
	data := getDataset("iris.csv")
	X, Y := generateXY(data)
	fmt.Println(X)
	fmt.Println(Y)
	model, err := m.NewSequential("iris")
	if err != nil {
		panic(err)
	}

	model.AddLayers(
		// layer.NewReLU(),
	)

	// model.Fit(nil, nil)

	// prediction, err := model.Predict(nil)
	// if err != nil {
	// 	panic(err)
	// }
	
	// fmt.Println(prediction)
}
