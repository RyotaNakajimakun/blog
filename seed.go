package main

import (
	"flag"
	"github.com/RyotaNakajimakun/blog/database/seeds"
	"github.com/k0kubun/pp"
)

func main() {
	flag.Parse()
	functionNames := flag.Args()
	pp.Print(functionNames)
	var record []interface{}
	for i := 0; i < len(functionNames); i++ {
		switch functionNames[i] {
		case "AddPermission":
			record = seeds.AddPermissions()
		case "AddRole":
			record = seeds.AddRole()
		}
	}
	pp.Print(record)
	seeds.CreateNewRecord(record)
}
