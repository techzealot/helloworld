package main

import (
	"fmt"
	"helloworld/downloader/inf"
	"helloworld/downloader/infra"
	"helloworld/downloader/testing"
)

func main() {
	url := "https://www.imooc.com"
	retriever := getRetriever()
	fmt.Println(retriever.Get(url))

	typeSwitch(retriever)
	typeAssertion(retriever)
}

func typeAssertion(retriever inf.Retriever) {
	if testRetriever, ok := retriever.(*testing.Retriever); ok {
		fmt.Printf("test : %T ,%v\n", testRetriever, testRetriever)
	}
	if infraRetriever, ok := retriever.(infra.Retriever); ok {
		fmt.Printf("infra : %T ,%v\n", infraRetriever, infraRetriever)
	}
}

func typeSwitch(retriever inf.Retriever) {
	switch retriever.(type) {
	case *testing.Retriever:
		fmt.Printf("%T ,%v\n", retriever, retriever)
	case infra.Retriever:
		fmt.Printf("value receiver: %T\n,%v\n", retriever, retriever)
	case *infra.Retriever:
		fmt.Printf("pointer receiver: %T\n,%v\n", retriever, retriever)
	default:
		fmt.Println("other retriever")
	}
}

func getRetriever() inf.Retriever {
	return &infra.Retriever{}
	// return infra.Retriever{}
	// return &testing.Retriever{}
}
