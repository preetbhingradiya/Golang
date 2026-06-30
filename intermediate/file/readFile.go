package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile() {

	file, err := os.Open("example.txt")

	if err != nil {
		fmt.Print(err)
	}

	defer func() {
		fmt.Println("Closing file...")
		file.Close()
	}()

	scnner := bufio.NewScanner(file)

	keyword := "sample"

	for scnner.Scan() {
		line := scnner.Text()
		fmt.Println(line)

		if strings.Contains(line, keyword) {
			line = strings.ReplaceAll(line, keyword, "REPLACED")
			fmt.Println(line)
		}
	}
}
