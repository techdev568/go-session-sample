package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Httpp() {
	fmt.Println("Hello")
	fileName := "hello.pdf"
	URL := ""
	err := downloadFile(URL, fileName)
	if err != nil {
		log.Fatal(err)
	}
}

func downloadFile(URL, fileName string) error {
	// Getting data from URL
	response, err := http.Get(URL)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("Not Received 200 response")
	}

	// create the file
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	// write the data to file
	_, err1 := io.Copy(file, response.Body)

	if err1 != nil {
		return err1
	}
	return nil
}
