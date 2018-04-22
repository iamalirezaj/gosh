package words

type Singular struct {
	String
	Dictionary map[string] string
}

func (s Singular) SetDictionary() Singular {
	s.Dictionary = map[string]string{
		"are":      "is",
		"analyses": "analysis",
		"alumni":   "alumnus",
		//"alumni": "alumnae", // for female - cannot have duplicate in map

		"genii":      "genius",
		"data":       "datum",
		"atlases":    "atlas",
		"appendices": "appendix",
		"barracks":   "barrack",
		"beefs":      "beef",
		"brothers":   "brother",
		"cafes":      "cafe",
		"corpuses":   "corpus",
		"cows":       "cow",
		"ganglions":  "ganglion",
		"genera":     "genus",
		"graffiti":   "graffito",
		"loaves":     "loaf",
		"monies":     "money",
		"mongooses":  "mongoose",
		"moves":      "move",
		"mythoi":     "mythos",
		"niches":     "niche",
		"numina":     "numen",
		"octopuses":  "octopus",
		"opuses":     "opus",
		"oxen":       "ox",
		"penises":    "penis",
		"vaginas":    "vagina",
		"sexes":      "sex",
		"testes":     "testis",
		"turfs":      "turf",
		"teeth":      "tooth",
		"feet":       "foot",
		"cacti":      "cactus",
		"children":   "child",
		"criteria":   "criterion",
		"news":       "news",
		"deer":       "deer",
		"echoes":     "echo",
		"elves":      "elf",
		"embargoes":  "embargo",
		"foes":       "foe",
		"foci":       "focus",
		"fungi":      "fungus",
		"geese":      "goose",
		"heroes":     "hero",
		"hooves":     "hoof",
		"indices":    "index",
		"knifes":     "knife",
		"leaves":     "leaf",
		"lives":      "life",
		"men":        "man",
		"mice":       "mouse",
		"nuclei":     "nucleus",
		"people":     "person",
		"phenomena":  "phenomenon",
		"potatoes":   "potato",
		"selves":     "self",
		"syllabi":    "syllabus",
		"tomatoes":   "tomato",
		"torpedoes":  "torpedo",
		"vetoes":     "veto",
		"women":      "woman",
		"zeroes":     "zero",
	}
	return s
}

func (s Singular) GetFromDictionary() string {
	plural := s.SetDictionary()
	return plural.Dictionary[s.Value]
}