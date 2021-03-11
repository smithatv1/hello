package hello
import (
	"fmt"
	"rsc.io/quote"
)
func Hello() string {
	fmt.Println("This is hello from Hello in package go-hello-module");
	return quote.Hello()
}
