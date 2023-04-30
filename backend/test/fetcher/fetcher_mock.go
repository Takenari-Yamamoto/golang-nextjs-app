package fetcher

import "github.com/stretchr/testify/mock"

type FetcherMock struct {
	mock.Mock
}

func (f *FetcherMock) Fetch() (string, error) {
	args := f.Called()
	return args.String(0), args.Error(1)
}
