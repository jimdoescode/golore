// historygen project historygen.go
package historygen

import (
	"golore/namegen"
	"math"
	"math/rand"
	"time"
)

type Diety struct {
	Name     string
	Gender   string
	power    int
	Elements []string
}

func CreateGods(ngen namegen.NameType, count int) []Diety {
	ds := make([]Diety, count)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pwr := 100
	elms := []string{"earth", "sky", "love", "fire", "death", "birth", "water", "fear", "war", "moon", "sun", "sport", "night", "dawn", "harvest", "seasons", "fate", "time"}

	for i := 0; i < count; i++ {
		elm := randElm(&elms, r)
		if int(math.Mod(float64(i), 2.0)) == 0 {
			ds[i] = Diety{
				ngen.Male().Generate(r),
				"God",
				pwr,
				[]string{elm},
			}
		} else {
			ds[i] = Diety{
				ngen.Female().Generate(r),
				"Goddess",
				pwr,
				[]string{elm},
			}
		}
	}

	//Randomly assign remaining elements
	for len(elms) > 0 {
		d := &ds[r.Intn(len(ds))]
		d.Elements = append(d.Elements, randElm(&elms, r))
	}

	return ds
}

func randElm(elms *[]string, r *rand.Rand) string {
	i := r.Intn(len(*elms))
	elm := (*elms)[i]
	*elms = append((*elms)[:i], (*elms)[i+1:]...)
	return elm
}
