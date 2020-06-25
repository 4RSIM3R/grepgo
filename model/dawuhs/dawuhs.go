package dawuhs

import (
	"log"

	db "github.com/ilzam/hackernews/db"
	"github.com/ilzam/hackernews/model/users"
)

//Dawuh is exported struct
type Dawuh struct {
	ID     string      `json:"id"`
	Dawuh  string      `json:"dawuh"`
	Qoil   string      `json:"qoil"`
	Author *users.User `json:"author"`
}

// Save is exported function
func (dawuh Dawuh) Save() int64 {
	// save to DB using prepare statement
	statement, err := db.Db.Prepare("INSERT INTO dawuhs(Dawuh, Qoil, UserID) VALUES (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := statement.Exec(dawuh.Dawuh, dawuh.Qoil, 1)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	log.Print("Row Inserted!")
	return id
}

//GetAll is exported function
func GetAll() []Dawuh {
	var dawuhs []Dawuh
	statement, err := db.Db.Prepare("SELECT ID, Dawuh, Qoil FROM dawuhs")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var dawuh Dawuh
		err := rows.Scan(&dawuh.ID, &dawuh.Dawuh, &dawuh.Qoil)
		if err != nil {
			log.Fatal(err)
		}
		dawuhs = append(dawuhs, dawuh)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return dawuhs
}
