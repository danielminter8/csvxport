package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// ExportMapToCSV - utility to export go map to csv
func ExportMapToCSV(filename string, data []map[string]interface{}) (string, error) {

	filename += ".csv"
	csvFile, err := os.Create("./" + filename)
	if err != nil {
		return "", err
	}
	csvwriter := csv.NewWriter(csvFile)
	if len(data) == 0 {
		err := errors.New("map slice empty")
		return "", err
	}

	csvHeaderValues := getMapKeys(data[0]) // gets key from first map in slice
	csvWriteErr := csvwriter.Write(csvHeaderValues)
	if csvWriteErr != nil {
		return "", csvWriteErr
	}

	values := make([][]string, len(data)) // contains all map values
	for i := 0; i < len(data); i++ {
		val := make([]string, len(csvHeaderValues))
		for j := 0; j < len(csvHeaderValues); j++ {
			str := fmt.Sprintf("%v", data[i][csvHeaderValues[j]])
			val[j] = str
		}
		values[i] = val
	}

	csvWriteAllErr := csvwriter.WriteAll(values) // writes map values to csv

	if csvWriteAllErr != nil {
		return "", csvWriteErr
	}

	csvwriter.Flush()
	csvFile.Close()
	return "", nil
}

// ExportStructToCSV - export struct to csv
func ExportStructToCSV(filename string, data interface{}) (string, error) {

	js, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	var dataInMap []map[string]interface{}
	json.Unmarshal(js, &dataInMap)

	ExportMapToCSV(filename, dataInMap)

	return "", nil
}

// getMapKeys - build csv headers from map keys
func getMapKeys(data map[string]interface{}) []string {
	keys := make([]string, len(data))
	i := 0
	for k := range data {
		keys[i] = k
		i++
	}

	return keys
}
