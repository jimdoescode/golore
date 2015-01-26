// historygen project historygen.go
package historygen

import (
	"fmt"
	"golore/namegen"
)

type Event struct {
	Text string
}

type History struct {
	Dieties   []Diety
	Events    []Event
	WorldName string
}

func (h *History) CreateWorld() Event {
	ngen := namegen.RogueLike{}

	h.WorldName = ngen.Female().String()
	h.Dieties = CreateGods(ngen, 7)

	d := DietyByElement(h.Dieties, "earth")
	e := Event{fmt.Sprintf("The %s %s created the world %s", d.Gender, d.Name, h.WorldName)}
	h.Events = append(h.Events, e)

	return e
}

func (h *History) CreateWater() Event {
	d := DietyByElement(h.Dieties, "water")
	e := Event{fmt.Sprintf("The %s %s filled the rivers and the seas of %s", d.Gender, d.Name, h.WorldName)}
	h.Events = append(h.Events, e)
	return e
}

func (h *History) CreateHeavens() Event {
	s := DietyByElement(h.Dieties, "sun")
	m := DietyByElement(h.Dieties, "moon")
	t := ""

	if s.Name != m.Name {
		if s.Gender != m.Gender {
			t = fmt.Sprintf("The %s %s and %s %s ignited the sun and moon to light %s", s.Gender, s.Name, m.Gender, m.Name, h.WorldName)
		} else {
			t = fmt.Sprintf("The %s's %s and %s ignited the sun and moon to light %s", s.Gender, s.Name, m.Name, h.WorldName)
		}
	} else {
		t = fmt.Sprintf("The %s %s ignited the sun and moon to light %s", s.Gender, s.Name, h.WorldName)
	}

	e := Event{t}
	h.Events = append(h.Events, e)
	return e
}
