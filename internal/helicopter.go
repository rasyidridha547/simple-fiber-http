package internal

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	"github.com/rasyidridha547/simple-fiber-http/model"
)

func getDataHelicopter() (chopper *model.Helicopter) {
	jsonFile, err := os.Open("helicopter.json")
	if err != nil {
		log.Fatalf("JSON File error: %s", err)
	}

	defer jsonFile.Close()

	read, _ := io.ReadAll(jsonFile)

	json.Unmarshal(read, &chopper)

	return chopper
}

func TypeHelicopter(helicopterType string) model.Helicopter {
	if helicopterType == "" {
		return nil
	}

	helicopters := getDataHelicopter()

	result := model.Helicopter{}

	for _, helicopter := range *helicopters {
		lowercase := strings.ToLower(helicopter.Type)

		if strings.Contains(lowercase, helicopterType) {
			result = append(result, helicopter)
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result
}

func NameHelicopter(helicopterName string) model.Helicopter {
	if helicopterName == "" {
		return nil
	}

	helicopters := getDataHelicopter()

	result := model.Helicopter{}

	for _, helicopter := range *helicopters {
		lowercase := strings.ToLower(helicopter.Name)

		if strings.Contains(lowercase, helicopterName) {
			result = append(result, helicopter)
			break
		} else {
			return nil
		}
	}

	return result
}
