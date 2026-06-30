package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BufioExample() {

	info := bufio.NewReader(strings.NewReader("HELLOOOOO package bufio!\n"))
	data := make([]byte, 20)

	n, _ := info.Read(data)
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	// bufio.NewScanner(info) creates a new Scanner that reads from the info reader. The Scanner is used to read the input line by line.
	// scanner := bufio.NewScanner(info)

	//After read 20 bytes
	process(info)

	writer := bufio.NewWriter(os.Stdout)
	writeData := []byte("Hello, World!\n")
	w, _ := writer.Write(writeData)
	fmt.Printf("Wrote %d bytes %s\n", w, writeData[:w])
	writer.Flush() // Flush the buffer to ensure all data is written to the underlying writer (os.Stdout in this case).

}

func process(r *bufio.Reader) {
	line, _ := r.ReadString('\n')
	fmt.Println("line: ", line)
}
