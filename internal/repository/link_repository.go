package repository

import (
	"database/sql"
	"log"

	"github.com/AlderFurtado/passlink/internal/domain/entity"
	"github.com/AlderFurtado/passlink/internal/utils"
	_ "modernc.org/sqlite" // driver SQLite
)

func FindAll() []entity.Link {
	db := setupDB()
	defer db.Close()

	rows, err := db.Query(`SELECT id, origin, destiny, validate FROM links`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []entity.Link
	for rows.Next() {
		var id int
		var origin, destiny, validate string
		if err := rows.Scan(&id, &origin, &destiny, &validate); err != nil {
			log.Fatal(err)
		}
		validateConverted := utils.ConvertStringToTimestamp(validate)
		newLink := entity.Link{Origin: origin, Destiny: destiny, Validate: validateConverted.Local()}
		links = append(links, newLink)
	}
	return links
}

func InsertNewLink(l entity.Link) {
	db := setupDB()
	// INSERT
	_, err := db.Exec(`INSERT INTO links (origin, destiny,validate) VALUES (?, ?, ?)`, l.Origin, l.Destiny, utils.ConvertTimestampToString(l.Validate))
	if err != nil {
		log.Fatal("Erro ao inserir:", err)
	}
}

// func findB

func setupDB() *sql.DB {
	db, err := sql.Open("sqlite", "./mylinks.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS links (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		origin TEXT NOT NULL,
		destiny TEXT NOT NULL,
		validate TEXT
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}
	return db
}
