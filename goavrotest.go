package main


import (
	"fmt"
	"github.com/linkedin/goavro/v2"
	"io/ioutil"
)

func main() {
	fmt.Println("hello")
	schemaFile, _ := ioutil.ReadFile("schema.avsc")
	codec, err := goavro.NewCodec(string(schemaFile))
	if err != nil {
		fmt.Print("----------: ",err)
	}
	schema := codec.Schema()
	fmt.Print(schema)
}