package main

import (
	"github.com/leonscriptcc/cyguy/cyguy"
	"log"
)

type Client struct {
	Name   string `json:"name" cypher:"names"`
	Gender string `json:"gender" cypher:"gender"`
}

func (c Client) NodeInfo() (string, string) {
	return "Person", "Client:Test"
}

func main() {
	c := Client{
		Name:   "莽夫贼",
		Gender: "1",
	}

	cyGuy := cyguy.NewCypherGuy()
	cql, err := cyGuy.Node(c).Create()
	if err != nil {
		log.Panic(err)
	}
	log.Println(cql)
}
