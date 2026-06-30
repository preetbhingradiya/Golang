package intermediate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type myError struct {
	message string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error : %s", m.message)
}

func eProcess() error {
	return &myError{"Custom error message"}
}

type APIError struct {
	Code    int
	Message string
	Data    interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func fetchProdut(postID int) (*Post, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postID)
	// url := fmt.Sprintf("https://invalid-url.typicode.com/posts/%d", postID)

	res, err := http.Get(url)
	if err != nil {
		return nil, &APIError{
			Code:    500,
			Message: "Failed to call external API",
			Data:    err,
		}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &APIError{
			Code:    500,
			Message: "Failed to read response",
			Data:    err,
		}
	}

	var post Post
	if err := json.Unmarshal(body, &post); err != nil {
		return nil, &APIError{
			Code:    500,
			Message: "Invalid JSON response",
			Data:    err.Error(),
		}
	}

	return &post, nil

}

func CustomError() {

	// err := eProcess()
	// if err != nil {
	// 	fmt.Print(err)
	// }

	post, err := fetchProdut(99)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Success Response:\n%+v\n", *post)
}
