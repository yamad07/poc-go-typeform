package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/yamad07/poc-go-typeform/api"
)

func main() {
	c := api.DefaultClient("EUqaoi488LNQ2tCp5hNWx7kmDCaKSUUeaVVQESoYsimC")

	r, err := c.RetrieveForms()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("%#v", r)

	for _, f := range r.Items {
		spew.Dump(f)
	}

}
