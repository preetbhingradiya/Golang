package intermediate

import (
	"fmt"
	"net/url"
)

func UrlParsing() {

	rawUrl := "https://www.example.com:8080/path/?query=golang&section=url#section1"

	// Parsing the URL
	parsedURL, err := url.Parse(rawUrl)

	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	queryParams := parsedURL.Query()
	fmt.Println("Query Parameters:", queryParams)
	fmt.Println("Query :", queryParams.Get("query"))
	fmt.Println("Section :", queryParams.Get("section"))
}
