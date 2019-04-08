package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/tachiniererin/malog"
)

func run(item malog.Response) {
	cmd := exec.Command("echo", item.URL)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("scraper finished with error: %v", err)
	}
}

func main() {
	u, errs := malog.Fetch()
	for {
		select {
		case item, ok := <-u:
			if !ok {
				return
			}
			run(item)
		case err, ok := <-errs:
			if !ok {
				return
			}
			log.Printf("error occured while scraping updates: %v", err)
		}
	}
}
