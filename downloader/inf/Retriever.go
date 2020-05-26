package inf

type Retriever interface {
	Get(url string) string
}
