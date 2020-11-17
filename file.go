// File handling
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//Open a new file
	emptyFile, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)

	//info printing
	fileStat, err := os.Stat("empty.txt")

	if err != nil {
		log.Fatal(err)
	}
	emptyFile.Close()
	fmt.Println("File Name:", fileStat.Name())        // Base name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", fileStat.IsDir())   // Abbreviation for Mode().IsDir()

	// Open file for reading.
	var file, err = os.OpenFile("empty.txt", os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Println("Reading from file.")
	fmt.Println(string(text))

	//Writing a file

	var file, err = os.OpenFile("empty.txt", os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString("Hello \n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("World \n")
	if isError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("File Updated Successfully.")

	//rename a file
	err := os.Rename("empty.txt", "test.txt")
	if err != nil {
		log.Fatal(err)
	}
}
