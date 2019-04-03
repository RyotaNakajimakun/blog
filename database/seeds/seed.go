package seeds

import "github.com/RyotaNakajimakun/blog/database/migration"

type SeedName struct {
	Name       string ``
	Permission string `"action", "permission level", "scene", "access"`
}

func CreateNewRecord(records []interface{}) {
	db := migration.InitDatabase()
	for i := 0; i < len(records); i++ {
		db.Create(&records[i])
		db.NewRecord(records[i])
	}
}
