// test
package main

import (
	"fmt"
	"golore/historygen"
)

func main() {
	h := &historygen.History{}

	fmt.Println(h.CreateWorld().Text)
	fmt.Println(h.CreateWater().Text)
	fmt.Println(h.CreateHeavens().Text)
}
