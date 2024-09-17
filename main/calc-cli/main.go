package calc_cli

import (
	"log"
	"os"

	"github.com/JacobMcKenzieSmarty/calc-apps/handlers"
	"github.com/JacobMcKenzieSmarty/calc-lib/calc"
)

func main() {
	err := handlers.NewCLIHandler(os.Stdout, calc.Addition{}).Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
