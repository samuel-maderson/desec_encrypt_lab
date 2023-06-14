package readfile

import (
	"bufio"
	"io"
	"log"
	"os"
)

func Reader(filename string) []string {

	fileMode := os.O_RDONLY
	filePermission := os.FileMode(644)
	file, err := os.OpenFile(filename, fileMode, filePermission)

	if err != nil {
		log.Panic("Error opening file:", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(file)
	var lines []string
	for {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			log.Panic("Error reading:", err)
		}

		lines = append(lines, line)

		if err == io.EOF {
			return lines
		}

	}
}
