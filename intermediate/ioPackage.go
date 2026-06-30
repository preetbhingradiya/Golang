package intermediate

import (
	"fmt"
	"io"
	"os"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("Read bytes:", n)
	fmt.Println("Data:", string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		panic(err)
	}
}

func pipeExample() {
	r, w := io.Pipe()

	go func() {
		w.Write([]byte("PIPE Example: Writing data to the pipe"))
		w.Close()
	}()

	data, _ := io.ReadAll(r)

	fmt.Println(string(data))
}

func IoPackage() {

	cf, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer cf.Close()

	readFromReader(cf)

	text := "\nOpen same file and add some new text"
	file, err := os.OpenFile(
		"example.txt",
		os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writeToWriter(file, text)

	pipeExample()
}
