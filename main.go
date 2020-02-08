package main

import (
	"fmt"
	_ "net/http"

	_ "github.com/gin-gonic/gin"
	"github.com/lonlng/go-learn/abc"
	"github.com/lonlng/go-learn/panictest"
)

func main() {
	fmt.Println("hello world!")
	abc.Abc()

	panictest.Test_01()
	abc.Abc()

}
