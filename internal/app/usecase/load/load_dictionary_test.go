package load

import (
	"testing"
)

func TestUsecase_Do(t *testing.T) {
	tests := []struct {
		name   string
		dict   *MockDictionary
		input  []string
		dictIn []string
	}{
		{
			name:   "one simple anagram",
			dict:   NewMockDictionary(),
			input:  []string{"asd", "zxc"},
			dictIn: []string{"asd", "zxc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Usecase{
				dict: tt.dict,
			}

			tt.dict.On("Set", tt.dictIn)

			u.Do(tt.input)

			tt.dict.AssertNumberOfCalls(t, "Set", 1)
		})
	}
}
