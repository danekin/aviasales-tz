package anagram

import (
	"sort"
	"strings"
)

type FindAnagramUsecase struct {
	dict Dictionary
}

func NewFindAnagramUsecase(dict Dictionary) *FindAnagramUsecase {
	return &FindAnagramUsecase{
		dict: dict,
	}
}

func (u *FindAnagramUsecase) Do(word string) (anagrams []string) {
	dict := u.dict.Get()

	for _, dictWord := range dict {
		if sortString(word) == sortString(dictWord) {
			anagrams = append(anagrams, dictWord)
		}
	}

	return anagrams
}

func sortString(str string) string {
	strs := strings.Split(strings.ToLower(str), "")
	sort.Strings(strs)

	return strings.Join(strs, "")
}
