package main

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"

	g "github.com/serpapi/google-search-results-golang"
)

type Config struct {
	APIKey string `json:"api_key"`
}

func get_api_key() string {
	data, err := ioutil.ReadFile("api.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)

	}
	return config.APIKey
}

func get_parameters(keywords string, region string, date_range string) map[string]string {
	parameters := map[string]string{
		"engine":    "google_trends",
		"q":         keywords,
		"data_type": "TIMESERIES",
		"geo":       region,
		"date":      date_range,
	}
	return parameters

}

func get_data(data map[string]interface{}) map[string]interface{} {
	timelineRaw := data["timeline_data"]
	timeline_data, ok := timelineRaw.([]interface{})
	if !ok {
		fmt.Println("Not ok")
		return nil
	}

	for _, value := range timeline_data {
		value_cast, ok := value.(map[string]interface{})
		if !ok {
			fmt.Println("Not ok")
			return nil
		}
		date := value_cast["date"]
		timestamp := value_cast["timestamp"]
		values := value_cast["values"]
		values_list, ok := values.([]interface{})
		if !ok {
			return nil
		}
		fmt.Println(date, timestamp)
		for _, value_insideraw := range values_list {
			value_inside, ok := value_insideraw.(map[string]interface{})
			if !ok {
				fmt.Println("Not ok")
				return nil
			}
			extracted_value := value_inside["extracted_value"]
			query := value_inside["query"]
			quant := value_inside["value"]
			fmt.Println(extracted_value, query, quant)
		}

	}
	return nil
}

func main() {
	api_key := get_api_key()
	parameters := get_parameters("cat", "US", "2025-01-01 2025-02-02")
	fmt.Println(parameters)
	search := g.NewGoogleSearch(parameters, api_key)
	results, _ := search.GetJSON()
	organic_results := results["interest_over_time"]
	fmt.Println(organic_results)
	get_data(organic_results.(map[string]interface{}))
}
