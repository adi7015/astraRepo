package domain

import (
	"astraProj/model"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func SaveDataInDB(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.New("could not open file for the given filePath")
	}

	defer file.Close()

	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var apiStruct model.ApiStruct
	err = json.Unmarshal(byteData, &apiStruct)
	if err != nil {
		log.Fatalf("Error parsing JSON data: %v", err)
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/astra")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO user (id, name,age, timestamp) VALUES (?, ?, ?, ?)",
		apiStruct.Id, apiStruct.Name, apiStruct.Age, apiStruct.TimeStamp)
	if err != nil {
		log.Printf("Error inserting item into database: %v", err)
	}

	return nil
}
