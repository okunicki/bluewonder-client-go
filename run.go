package main

import (
	"fmt"

	"example.com/client/getme"
)

func main() {

	// Call GetMe with the common bearer token and path (default is "me")
	getmeData, getmeErr := getme.GetMe("me")
	if getmeErr != nil {
		fmt.Println("Error calling GetMe:", getmeErr)
		return
	}

	// Print the name from the GetMe JSON
	fmt.Println("GetMe Name:", getmeData.Name)

}
