package main

import (
	"fmt"
	"net/http"

	"github.com/elysiamori/go_native/native-train/routers"
)

/*
	# Learn Go Native : Rest API using JSON file
	Valdy Ramadhan

	Rest API using JSON file:
	1. Make dummy data JSON file {data.json}
	2. Define struct for JSON file {models.go}
	3. Create reads json file {handler.go}
	4. Rest API Logic {handler.go} : - Get all datas
	                    			 - Get data by id/uuid
									 - Add data local
									 - Update data local
									 - Delete data local
	5. Create router {routes.go}
	6. Create main.go

	Note: This is not the best practice, but this is the easiest way to learn
*/

// Main function
func main() {
	routers.InitRouter()
	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", nil)
}
