package main

import (
	"fmt"
	"time"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main2() {
	start := time.Now()

	rawData, err := base.ParseCSVToInstances("iris.csv", true)
	if err != nil {
		panic(err)
	}

	fmt.Println(rawData)

	cls := knn.NewKnnClassifier("euclidean", "linear", 2)

	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	cls.Fit(trainData)

	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}

	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))

	fmt.Println("Time elapsed:", time.Since(start))
}
