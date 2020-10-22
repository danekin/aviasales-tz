package anagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagramUsecase_Do(t *testing.T) {
	tests := []struct {
		name    string
		dict    *MockDictionary
		word    string
		dictOut []string
		result  []string
	}{
		{
			name: "one simple anagram",
			dict: NewMockDictionary(),
			word: "asd",

			dictOut: []string{"sad"},
			result:  []string{"sad"},
		},
		{
			name: "none anagrams",
			dict: NewMockDictionary(),
			word: "asd",

			dictOut: []string{"sada", "zxc"},
			result:  nil,
		},
		{
			name: "multiple case-insensitive anagrams",
			dict: NewMockDictionary(),
			word: "asd",

			dictOut: []string{"sad", "zxc", "ADs", "asda"},
			result:  []string{"sad", "ADs"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &FindAnagramUsecase{
				dict: tt.dict,
			}

			tt.dict.On("Get").Return(tt.dictOut)

			result := u.Do(tt.word)

			assert.Equal(t, tt.result, result)
			tt.dict.AssertNumberOfCalls(t, "Get", 1)
		})
	}
}
