package main

import (
	"log"
	"os"

	"github.com/JacobMcKenzieSmarty/calc-apps/handlers"
)

func main() {
	//os.Stdin, os.Stdout, log.Default()
	myCsv, err := os.Open("/Users/jacob/src/github.com/JacobMcKenzieSmarty/calc-apps/main/calc-csv/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	handler := handler.NewCSVHandler(myCsv, os.Stdout, log.Default())
	err = handler.Handle()
	if err != nil {
		log.Fatal(err)
	}
}
