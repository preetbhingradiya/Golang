package file

import (
	"fmt"
	"os"
)

func WriteFile() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Ensure the file is closed after writing
	defer file.Close()

	// Create a buffered writer
	// writer := bufio.NewWriter(file)
	data := "Hello, this is a sample text written to the file.\n"
	n, err := file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("Wrote %d bytes to file.\n", n)
	// writer.Flush() // Ensure all data is written to the file
}
