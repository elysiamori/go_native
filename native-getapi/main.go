package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/elysiamori/go_native/native-getapi/models"
)

/*
	# Learn Go Native : Fetch API/Get API

	Valdy Ramadhan

	Fetching data from API using Go
	Step:
		0. define struct data with respinse json from api
		1. get data from url api
		2. close response body
		3. read response body
		4. create variable to read JSON response
		5. unmarshal body to content variable
		6. get url api
		7. call function to read reading data from api

	Note:
		- get data by id
		- get all data
*/

func GetAPI(apiUrl string) (models.Contents, error) {

	// 1 get data from url api
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
		return models.Contents{}, err
	}
	// 2. close response body
	defer resp.Body.Close()

	// 3. read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return models.Contents{}, err
	}

	// 4. create variable to read JSON response
	var content models.Contents

	// 5. unmarshal body to content variable
	err = json.Unmarshal(body, &content)
	if err != nil {
		fmt.Println(err)
		return models.Contents{}, err
	}

	return content, nil
}

func main() {
	// 6. get url api
	apiUrl := "http://localhost:3000/api/content"

	content, err := GetAPI(apiUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 7.1 call function to read reading data from api {get by id}
	// fmt.Println("ID\t:" + fmt.Sprint(content.Contents.ID))
	// fmt.Println("Title\t:" + content.Contents.Title)
	// fmt.Println("Desc\t:" + content.Contents.Desc)
	// fmt.Println("Date\t:" + content.Contents.Date)

	// 7.2 call function to read reading data from api {get all}
	for _, v := range content.Contents {
		fmt.Println("ID\t:" + fmt.Sprint(v.ID))
		fmt.Println("Title\t:" + v.Title)
		fmt.Println("Desc\t:" + v.Desc)
		fmt.Println("Date\t:" + v.Date)
		fmt.Println("====================================")
	}
}
