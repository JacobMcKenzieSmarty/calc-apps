package calc_cli

import (
	"log"
	"os"

	"github.com/JacobMcKenzieSmarty/calc-apps/handlers"
)

func main() {
	err := handler.NewCLIHandler(os.Stdout, os.Args[1:]).Handle()
	if err != nil {
		log.Fatal(err)
	}
}
