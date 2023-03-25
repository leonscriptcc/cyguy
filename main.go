package main

import (
	"github.com/leonscriptcc/cyguy/cyguy"
	"log"
)

type Client struct {
	Name   string `json:"name" cypher:"names"`
	Gender int    `json:"gender" cypher:"gender"`
}

func main() {
	c := Client{
		Name:   "莽夫贼",
		Gender: 1,
	}

	cyGuy := cyguy.NewCypherGuy()
	log.Println(cyGuy.Node("person", "client", "foo").
		SetProperties(c).Create())
}
