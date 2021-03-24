package testing

type Retriever struct{}

// 指针接收者
func (*Retriever) Get(url string) string {
	return "this is a test website"
}
