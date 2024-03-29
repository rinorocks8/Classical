package service

import (
	obj "Classical/Backend/model"
	"database/sql"
	f "fmt"

	_ "github.com/go-sql-driver/mysql"
)

// classesByName queries for albums that have the specified class name.
func ClassesByName(name string) ([]obj.Class, error) {
	// An albums slice to hold data from returned rows
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var classes []obj.Class

	rows, err := db.Query("SELECT * FROM class WHERE className = ?", name)
	if err != nil {
		return nil, f.Errorf("classesByName %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var cla obj.Class
		if err := rows.Scan(&cla.ClassName); err != nil {
			return nil, f.Errorf("classesByName %q: %v", name, err)
		}
		classes = append(classes, cla)
	}
	if err := rows.Err(); err != nil {
		return nil, f.Errorf("classesByName %q: %v", name, err)
	}
	return classes, nil
}
