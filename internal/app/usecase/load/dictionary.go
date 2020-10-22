package load

import (
	"github.com/stretchr/testify/mock"
)

type Dictionary interface {
	Set(dict []string)
}

type MockDictionary struct {
	mock.Mock
}

func NewMockDictionary() *MockDictionary {
	return &MockDictionary{}
}

func (m *MockDictionary) Set(dict []string) {
	m.Called(dict)
}
