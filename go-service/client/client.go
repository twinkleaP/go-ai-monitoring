package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/you/go-service/metrics"
)

func SendMetricstoAI(m metrics.Metric) {

	url := "http://localhost:8000/predict" // Python FastAPI endpoint

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error marshaling metrics:", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error sending metrics to AI service:", err)
		return
	}
	defer resp.Body.Close()

	responseData, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response from AI service:", string(responseData))

}
