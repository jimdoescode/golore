package namegen

type Fantasy struct{}

func (f Fantasy) Male() Name {
	return Name{
		map[string][]string{
			"start":  []string{"Aer", "An", "Ar", "Ban", "Bar", "Ber", "Beth", "Bett", "Cut", "Dan", "Dar", "Dell", "Der", "Edr", "Er", "Eth", "Ett", "Fin", "Ian", "Iarr", "Ill", "Jed", "Kan", "Kar", "Ker", "Kurr", "Kyr", "Man", "Mar", "Mer", "Mir", "Tsal", "Tser", "Tsir", "Van", "Var", "Yur", "Yyr"},
			"middle": []string{"al", "an", "ar", "el", "en", "ess", "ian", "onn", "or"},
			"end":    []string{"ai", "an", "ar", "ath", "en", "eo", "ian", "is", "u", "or"},
		},
		[]string{"start|middle|end", "start|end"},
		[]string{"orar", "arrar", "ianian"},
	}
}

func (f Fantasy) Female() Name {
	return Name{
		map[string][]string{
			"start":  []string{"Aer", "An", "Ar", "Ban", "Bar", "Ber", "Beth", "Bett", "Cut", "Dan", "Dar", "Dell", "Der", "Edr", "Er", "Eth", "Ett", "Fin", "Ian", "Iarr", "Ill", "Jed", "Kan", "Kar", "Ker", "Kurr", "Kyr", "Man", "Mar", "Mer", "Mir", "Tsal", "Tser", "Tsir", "Van", "Var", "Yur", "Yyr"},
			"middle": []string{"al", "an", "ar", "el", "en", "ess", "ian", "onn", "or"},
			"end":    []string{"a", "ae", "aelle", "ai", "ea", "i", "ia", "u", "wen", "wyn"},
		},
		[]string{"start|middle|end", "start|end"},
		[]string{"arrar"},
	}
}
