package main

import (
	"github.com/nmiculinic/go-generate-talk/gRPC/summer"
)
func main() {
	summer.Serve("[::]:8080", "localhost:50051", summer.DefaultHtmlStringer)
}
