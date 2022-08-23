package main

import(
	"bookstore/store/factory"
)

func main(){
	s, err := factory.New("mem")

	if err != nil {
		panic(err)
	}
}