package main

import (
	"gorm.io/gen"

	"test-echo/internal/db"
)

func main() {
	db := db.Get().DB

	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/database",
	})

	// Use the database connection
	g.UseDB(db)

	// Generate basic type-safe DAO API for all tables
	g.ApplyBasic(g.GenerateAllTable()...)

	// Execute the code generation
	g.Execute()
}
