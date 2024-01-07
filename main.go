package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/MashyBasker/zlib-kindle-automate/scraper"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Enter search keyword: ")
	line, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	scraper.ScrapeBookNames(strings.TrimSpace(line))
}