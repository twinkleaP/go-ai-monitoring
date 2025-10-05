package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"go-service/metrics"
)

func SendMetricstoAI(m metrics.Metric) {

	//url := "http://ai-service:8000/predict"
	//  // Python FastAPI endpoint

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error marshaling metrics:", err)
	}
	for i := 0; i < 5; i++ {
		resp, err := http.Post("http://ai-service:8000/predict", "application/json", (bytes.NewBuffer(data)))
		if err != nil {
			fmt.Println("Error sending metrics to AI service:", err)
			return
		}
		time.Sleep(2 * time.Second)

		//defer resp.Body.Close()

		responseData, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Response from AI service:", string(responseData))

	}
}
