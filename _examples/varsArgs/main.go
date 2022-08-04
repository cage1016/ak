/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com> (https://kaichu.io)

*/
package main

import (
	"flag"

	aw "github.com/deanishe/awgo"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	flag.Parse()

	av := aw.NewArgVars()
	av.Var("far", "bar")
	av.Arg(flag.Args()...)
	if err := av.Send(); err != nil {
		panic(err)
	}
}
