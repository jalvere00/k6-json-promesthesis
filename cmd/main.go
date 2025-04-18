package main

import (
	"fmt"

	"github.com/grafana/xk6-output-prometheus-remote/pkg/remotewrite"
	"github.com/sirupsen/logrus"
	"go.k6.io/k6/output"
)

func main() {
	output, err := remotewrite.New(output.Params{
		Logger: logrus.New(),
	})

	if err != nil {
		fmt.Printf("An error occurred: %v", err)
	}

	output.AddMetricSamples()
	fmt.Println("Testing")

}
