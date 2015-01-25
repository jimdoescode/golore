package namegen

type RogueLike struct{}

func (r RogueLike) Male() Name {
	return Name{
		map[string][]string{
			"vowel":  []string{"a", "e", "i", "o", "u", "y"},
			"start":  []string{"Aer", "Al", "Am", "An", "Ar", "Arm", "Arth", "B", "Bal", "Bar", "Be", "Bel", "Ber", "Bok", "Bor", "Bran", "Breg", "Bren", "Brod", "Cam", "Chal", "Cham", "Ch", "Cuth", "Dag", "Daim", "Dair", "Del", "Dr", "Dur", "Duv", "Ear", "Elen", "Er", "Erel", "Erem", "Fal", "Ful", "Gal", "G", "Get", "Gil", "Gor", "Grin", "Gun", "H", "Hal", "Han", "Har", "Hath", "Hett", "Hur", "Iss", "Khel", "K", "Kor", "Lel", "Lor", "M", "Mal", "Man", "Mard", "N", "Ol", "Radh", "Rag", "Relg", "Rh", "Run", "Sam", "Tarr", "T", "Tor", "Tul", "Tur", "Ul", "Ulf", "Unr", "Ur", "Urth", "Yar", "Z", "Zan", "Zer"},
			"middle": []string{"de", "do", "dra", "du", "duna", "ga", "go", "hara", "kaltho", "la", "latha", "le", "ma", "nari", "ra", "re", "rego", "ro", "rodda", "romi", "rui", "sa", "to", "ya", "zila"},
			"end":    []string{"bar", "bers", "blek", "chak", "chik", "dan", "dar", "das", "dig", "dil", "din", "dir", "dor", "dur", "fang", "fast", "gar", "gas", "gen", "gorn", "grim", "gund", "had", "hek", "hell", "hir", "hor", "kan", "kath", "khad", "kor", "lach", "lar", "ldil", "ldir", "leg", "len", "lin", "mas", "mnir", "ndil", "ndur", "neg", "nik", "ntir", "rab", "rach", "rain", "rak", "ran", "rand", "rath", "rek", "rig", "rim", "rin", "rion", "sin", "sta", "stir", "sus", "tar", "thad", "thel", "tir", "von", "vor", "yon", "zor"},
		},
		[]string{"start|vowel|middle|middle|end", "start|vowel|middle|end", "start|vowel|end"},
		nil,
	}
}

func (r RogueLike) Female() Name {

	return Name{
		map[string][]string{
			"vowel":  []string{"a", "e", "i", "o", "u", "y"},
			"start":  []string{"Ad", "Aer", "Ar", "Bel", "Bet", "Beth", "Ce'N", "Cyr", "Eilin", "El", "Em", "Emel", "G", "Gl", "Glor", "Is", "Isl", "Iv", "Lay", "Lis", "May", "Ner", "Pol", "Por", "Sal", "Sil", "Vel", "Vor", "X", "Xan", "Xer", "Yv", "Zub"},
			"middle": []string{"bre", "da", "dhe", "ga", "lda", "le", "lra", "mi", "ra", "ri", "ria", "re", "se", "ya"},
			"end":    []string{"ba", "beth", "da", "kira", "laith", "lle", "ma", "mina", "mira", "na", "nn", "nne", "nor", "ra", "rin", "ssra", "ta", "th", "tha", "thra", "tira", "tta", "vea", "vena", "we", "wen", "wyn"},
		},
		[]string{"start|vowel|middle|middle|end", "start|vowel|middle|end", "start|vowel|end"},
		nil,
	}

}
