package service

import (
	"astraProj/domain"
	"astraProj/model"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

func SaveData(body []byte) error {
	var apiStruct model.ApiStruct

	err := json.Unmarshal(body, &apiStruct)
	if err != nil {
		fmt.Println("could not unmarshal the given request")
		return nil
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	apiStruct.TimeStamp = formattedTime

	jsonBody, err := json.Marshal(apiStruct)
	if err != nil {
		fmt.Println("could not marshal the given request")
		return nil
	}

	// Write JSON data to file
	filename := "data.json"
	directory := filepath.Join("C:", "astraProj")

	err = ioutil.WriteFile(directory+filename, jsonBody, 0644)
	if err != nil {
		return errors.New("could not store the data in desired location")
		//c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//return
	}

	go func() {
		ch := make(chan error)
		err = domain.SaveDataInDB(directory + filename)
		ch <- err
	}()

	if err != nil {
		return errors.New("could not save data in db")
	}

	return nil
}
