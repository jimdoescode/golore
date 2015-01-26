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
	Power    float32
	Elements []string
}

func CreateGods(ngen namegen.NameType, count int) []Diety {
	ds := make([]Diety, count)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	elms := []string{"earth", "sky", "love", "fire", "death", "birth", "water", "fear", "war", "moon", "sun", "sport", "night", "dawn", "harvest", "seasons", "fate", "time"}
	pwr := float32(len(elms)) / float32(count)

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
		d.Power += pwr
	}

	return ds
}

func DietyByElement(dieties []Diety, elm string) Diety {
	for _, d := range dieties {
		for _, e := range d.Elements {
			if e == elm {
				return d
			}
		}
	}
	return dieties[0]
}

func randElm(elms *[]string, r *rand.Rand) string {
	i := r.Intn(len(*elms))
	elm := (*elms)[i]
	*elms = append((*elms)[:i], (*elms)[i+1:]...)
	return elm
}
