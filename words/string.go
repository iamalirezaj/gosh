package words

import (
	"strings"
)

type String struct {
	Type string
	Value string
}

func AlreadyPluralized(test string) bool {
	var singularDictionary = Singular{}.SetDictionary().Dictionary
	var pluralDictionary = Plural{}.SetDictionary().Dictionary

	if len(test) != 1 {

		// to handle words like genii, data and etc.
		if singularDictionary[test] != "" {
			return true
		}

		// put in some exceptions
		//if (string(test[len(test)-1]) != "s") || (string(test[len(test)-2]) != "ii") {
		if string(test[len(test)-1]) != "s" {

			if (string(test[len(test)-1:]) != "e") || (string(test[len(test)-1:]) != "y") {
				return false
			}

			if (string(test[len(test)-2:]) == "ch") || (string(test[len(test)-2:]) == "sh") || (string(test[len(test)-3:]) == "nes") {
				return false
			}

			if string(test[len(test)-3:]) == "ius" {
				return false
			}

			if pluralDictionary[test] == "" {
				return true
			}

		} else {
			return true
		}
	}
	return false

}

func inVowels(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func (s String) ToPlural() Plural {
	return Plural{String: s}.Convert()
}

func (s String) ToLowercase() String {
	s.Value = strings.ToLower(s.Value)
	return s
}

func (s String) ToUppercase() String {
	s.Value = strings.ToUpper(s.Value)
	return s
}

func (s String) ToString() string {
	return s.Value
}