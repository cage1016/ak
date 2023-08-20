/*
Copyright © 2023 KAI CHU CHUNG
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
