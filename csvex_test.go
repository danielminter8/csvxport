package main

import (
	"os"
	"testing"
)

func TestExportMap(t *testing.T) {
	var weatherForcast []map[string]interface{}
	var data = make(map[string]interface{})
	data["Monday"] = "28C"
	data["Tuesday"] = "15C"
	data["Wednesday"] = "21C"
	data["Thursday"] = "17C"
	data["Friday"] = "20C"

	weatherForcast = append(weatherForcast, data)

	filename := "weather_forcast"
	_, err := ExportMapToCSV(filename, weatherForcast)
	if err != nil {
		t.Error("Error : ", err)
	}

	os.Remove(filename + ".csv")
}

func TestExportStruct(t *testing.T) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var users []User
	var user User
	user.Name = "John"
	user.Age = 25
	users = append(users, user)
	user.Name = "Piet"
	user.Age = 24
	users = append(users, user)

	filename := "user_data"
	_, err := ExportStructToCSV(filename, users)
	if err != nil {
		t.Error("Error : ", err)
	}

	os.Remove(filename + ".csv")
}
