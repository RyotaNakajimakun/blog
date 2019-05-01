package main

import (
	"github.com/RyotaNakajimakun/blog/database/migration"
	"github.com/RyotaNakajimakun/blog/database/seeds"
	"github.com/k0kubun/pp"
)

func adders() {
	db := migration.InitDatabase()

	permissions := seeds.AddPermissions()
	for i := 0; i < len(permissions); i++ {
		pp.Print(permissions[i])
		db.Create(&permissions[i])
		db.NewRecord(permissions[i])
	}

	roles := seeds.AddRole()
	for i := 0; i < len(roles); i++ {
		pp.Print(roles[i])
		db.Create(&roles[i])
		db.NewRecord(roles[i])
	}
}

func main() {
	adders()
}