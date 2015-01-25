// test
package main

import (
	"fmt"
	"golore/historygen"
	"golore/namegen"
	"strings"
)

func main() {
	ngen := namegen.RogueLike{}

	gods := historygen.CreateGods(ngen, 7)

	for _, g := range gods {
		fmt.Println(g.Name + " " + g.Gender + " of " + strings.Join(g.Elements, " and "))
	}
}
