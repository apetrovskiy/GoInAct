package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	fmt.Println("Hello, playground")
	search.Run("president")
}
