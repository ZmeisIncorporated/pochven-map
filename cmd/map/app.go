package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ZmeisIncorporated/pochven-map/pkg/dscan"
)


func main() {
	fmt.Println("pochven-map v 0.1.0")

	var shipsYML = flag.String("ships", "ships.yaml", "path to file with ships list")
	flag.Parse()

	fmt.Println("ships config:", *shipsYML)
	_, err := dscan.NewDscan(*shipsYML)
	if err != nil {
		log.Fatalln(err)
	}

}
