package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFetcher struct {
	mock.Mock
}

func (m *MockFetcher) FetchData() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func TestGetData(t *testing.T) {
	fetcher := new(MockFetcher)

	fetcher.On("FetchData").Return("mocked data", nil)

	// GetData関数をテスト
	data, err := GetData(fetcher)

	// アサーション
	assert.NoError(t, err)
	assert.Equal(t, "mocked data", data)

	// FetchDataメソッドが1回呼ばれたことを確認
	fetcher.AssertNumberOfCalls(t, "FetchData", 1)
}
