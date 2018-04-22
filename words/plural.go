package words

type Plural struct {
	String
	Dictionary map[string] string
}

var vowels = []string{"a", "e", "i", "o", "u"}

func (p Plural) SetDictionary() Plural {
	p.Dictionary = map[string]string{
		"is":         "are",
		"analysis":   "analyses",
		"alumnus":    "alumni",
		"alumnae":    "alumni",
		"atlas":      "atlases",
		"appendix":   "appendices",
		"barrack":    "barracks",
		"beef":       "beefs",
		"brother":    "brothers",
		"cafe":       "cafes",
		"corpus":     "corpuses",
		"cow":        "cows",
		"ganglion":   "ganglions",
		"genus":      "genera",
		"graffito":   "graffiti",
		"loaf":       "loaves",
		"money":      "monies",
		"mongoose":   "mongooses",
		"move":       "moves",
		"mythos":     "mythoi",
		"niche":      "niches",
		"numen":      "numina",
		"octopus":    "octopuses",
		"opus":       "opuses",
		"ox":         "oxen",
		"penis":      "penises",
		"vagina":     "vaginas",
		"sex":        "sexes",
		"testis":     "testes",
		"turf":       "turfs",
		"tooth":      "teeth",
		"foot":       "feet",
		"cactus":     "cacti",
		"child":      "children",
		"criterion":  "criteria",
		"news":       "news",
		"datum":      "data",
		"deer":       "deer",
		"echo":       "echoes",
		"elf":        "elves",
		"embargo":    "embargoes",
		"foe":        "foes",
		"focus":      "foci",
		"fungus":     "fungi",
		"goose":      "geese",
		"hero":       "heroes",
		"hoof":       "hooves",
		"index":      "indices",
		"knife":      "knives",
		"leaf":       "leaves",
		"life":       "lives",
		"man":        "men",
		"mouse":      "mice",
		"nucleus":    "nuclei",
		"person":     "people",
		"phenomenon": "phenomena",
		"potato":     "potatoes",
		"self":       "selves",
		"syllabus":   "syllabi",
		"tomato":     "tomatoes",
		"torpedo":    "torpedoes",
		"veto":       "vetoes",
		"woman":      "women",
		"zero":       "zeroes",
	}
	return p
}

func (p Plural) GetFromDictionary() string {
	plural := p.SetDictionary()
	return plural.Dictionary[p.Value]
}

func (p Plural) SimpleConvert() Plural {
	p.Value = p.Value + "s"
	return p
}

func (p Plural) HandleWordWithSisCharacter(root string, suffix string) (string, string, string) {
	if string(p.Value[len(p.Value)-3:]) == "sis" {
		root = string(p.Value[:len(p.Value)-3])
		suffix = "ses"
		p.Value = root + suffix
	}
	return root, suffix, p.Value
}

func (p Plural) HandleWordWithZCharacter(root string, suffix string) (string, string, string) {
	if string(p.Value[len(p.Value)-1:]) == "z" {
		root = string(p.Value[:len(p.Value)-1])
		suffix = "zes"
		p.Value = root + suffix
	}
	return root, suffix, p.Value
}

func (p Plural) HandleWordWithExCharacter(root string, suffix string) (string, string, string) {
	if string(p.Value[len(p.Value)-2:]) == "ex" {
		root = string(p.Value[:len(p.Value)-2])
		suffix = "ices"
		p.Value = root + suffix
	}
	return root, suffix, p.Value
}

func (p Plural) HandleWordWithIxCharacter(root string, suffix string) (string, string, string) {
	if string(p.Value[len(p.Value)-2:]) == "ix" {
		root = string(p.Value[:len(p.Value)-2])
		suffix = "ices"
		p.Value = root + suffix
	}
	return root, suffix, p.Value
}

func (p Plural) HandleWordWithUsCharacter(root string, suffix string) (string, string, string) {
	if string(p.Value[len(p.Value)-2:]) == "us" {
		root = string(p.Value[:len(p.Value)-2])
		suffix = "uses"
		p.Value = root + suffix
	}
	return root, suffix, p.Value
}

func (p Plural) HandleWordSuchAsDolly(root string, suffix string) (string, string, string) {
	if (string(p.Value[len(p.Value)-1]) == "y") && !inVowels(string(p.Value[len(p.Value)-2]), vowels) {
		root = string(p.Value[:len(p.Value)-1])
		suffix = "ies"

	} else if string(p.Value[len(p.Value)-1]) == "s" {
		if inVowels(string(p.Value[len(p.Value)-2]), vowels) {
			if string(p.Value[len(p.Value)-3:]) == "ius" {
				root = string(p.Value[:len(p.Value)-2])
				suffix = "i"
			} else {
				root = string(p.Value[:len(p.Value)-1])
				suffix = "ses"
			}
		} else {
			suffix = "es"
		}
	} else if (string(p.Value[len(p.Value)-2:]) == "ch") || (string(p.Value[len(p.Value)-2:]) == "sh") {
		suffix = "es"
	} else {
		suffix = "s"
	}
	p.Value = root + suffix
	return root, suffix, p.Value
}

func (p Plural) Convert() Plural {

	plural := ""
	suffix := ""

	if p.Value != "" {

		root := p.Value

		// are we dealing with a single character?
		if len(p.Value) == 1 { return p.SimpleConvert() }

		if plural = p.GetFromDictionary(); plural != "" {
			p.Value = plural
			return p
		}

		root, suffix, plural = p.HandleWordWithSisCharacter(root, suffix)
		root, suffix, plural = p.HandleWordWithZCharacter(root, suffix)
		root, suffix, plural = p.HandleWordWithExCharacter(root, suffix)
		root, suffix, plural = p.HandleWordWithIxCharacter(root, suffix)
		root, suffix, plural = p.HandleWordWithUsCharacter(root, suffix)
		root, suffix, plural = p.HandleWordSuchAsDolly(root, suffix)

		// sanity check
		if AlreadyPluralized(p.Value) {
			return p
		} else {
			p.Value = plural
			return p
		}
	}

	return p
}