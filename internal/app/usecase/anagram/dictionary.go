package anagram

import (
	"github.com/stretchr/testify/mock"
)

type Dictionary interface {
	Get() []string
}

type MockDictionary struct {
	mock.Mock
}

func NewMockDictionary() *MockDictionary {
	return &MockDictionary{}
}

func (m *MockDictionary) Get() []string {
	return m.Called().Get(0).([]string)
}
