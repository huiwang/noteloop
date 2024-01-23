package main

import (
	"fmt"
	"os"

	"truelaurel.com/noteloop/notion"
	"truelaurel.com/noteloop/weread"
)

func main() {

	// Read the first argument as an app type
	if len(os.Args) < 2 {
		fmt.Println("App type is missing")
		return
	}

	appType := os.Args[1]

	switch appType {
	case "notion":
		notion.UploadPage()
	case "weread":
		if len(os.Args) < 3 {
			fmt.Println("Cookie is missing")
			return
		}
		cookie := os.Args[2]
		fetchWereadBooks(cookie)
	default:
		fmt.Println("Invalid app type")
	}

}

func fetchWereadBooks(cookie string) bool {
	client := weread.NewClient(cookie)
	books, err := client.GetBooks()
	if err != nil {
		fmt.Println("Error while getting books:", err)
		return true
	}

	for _, book := range books {
		fmt.Println(book)
	}
	return false
}
