package basics

import (
	"fmt"
	"net/http"

)

func PackageExample() {
	fmt.Println("This is an example of a package in Go.")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	defer resp.Body.Close()
	/*
			What is defer?
				Executes a function at the end of the current function
				Clean-up mechanism (Do this cleanup work at the end, no matter what happens)
		        Runs in LIFO order
			Why needed?
				resp.Body holds:
				Network connection
				Memory resources
			**/

	fmt.Println("Response Status:", resp.Status)
}
