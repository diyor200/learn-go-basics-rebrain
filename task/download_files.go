package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	var (
		eg = new(errgroup.Group)

		fileURLs = []string{
			"https://unsplash.com/photos/XMFZqrGyV-Q/download?ixid=M3wxMjA3fDB8MXxzZWFyY2h8NHx8cHJvZ3JhbW1pbmd8ZW58MHx8fHwxNzQwNjg3NDQ0fDA&force=true",
			"https://unsplash.com/photos/4hbJ-eymZ1o/download?ixid=M3wxMjA3fDB8MXxzZWFyY2h8Mnx8cHJvZ3JhbW1pbmd8ZW58MHx8fHwxNzQwNjg3NDQ0fDA&force=true",
			"https://unsplash.com/photos/EJMTKCZ00I0/download?ixid=M3wxMjA3fDB8MXxzZWFyY2h8N3x8cHJvZ3JhbW1pbmd8ZW58MHx8fHwxNzQwNjg3NDQ0fDA&force=true",
			"https://unsplash.com/photos/DuHKoV44prg/download?ixid=M3wxMjA3fDB8MXxzZWFyY2h8OHx8cHJvZ3JhbW1pbmd8ZW58MHx8fHwxNzQwNjg3NDQ0fDA&force=true",
		}
	)

	for i := range fileURLs {
		eg.Go(func() error {
			return downloadFile(fmt.Sprintf("file_%d", i+1), fileURLs[i])
		})
	}

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("files downloaded successfully!")
}

func downloadFile(fileName, fileURL string) error {
	resp, err := http.Get(fileURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	filePath, err := os.Create(fileName + ".jpg")
	if err != nil {
		return err
	}

	_, err = io.Copy(filePath, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
