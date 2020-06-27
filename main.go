package main

import (
	"log"

	"github.com/oscarDAN553/GOTWITER/bd"
	"github.com/oscarDAN553/GOTWITER/handlers"
)

//  import (
// 	"log"

// 	"github.com/oscarDAN553/GOTWITER/bd"
// 	"github.com/oscarDAN553/GOTWITER/handlers"
// )

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("sin conexxion a la BD")
		return
	}
	handlers.Manejadores()

}
