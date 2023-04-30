package fetcher

type Fetcher interface {
	Fetch() (string, error)
}
