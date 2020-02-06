package main

import (
	"fmt"
	_ "net/http"

	"github.com/lonlng/go-learn/abc"

	_ "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world!")
	abc.Abc()

}
