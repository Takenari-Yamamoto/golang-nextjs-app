package fetcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	fetcherMock := new(FetcherMock)
	fetcherMock.On("Fetch").Return("test data", nil)

	result, err := fetcherMock.Fetch()
	assert.NoError(t, err)
	assert.Equal(t, "test data", result)
	fetcherMock.AssertExpectations(t)
}
