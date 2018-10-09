package models

import (
	"fmt"
	"crypto/sha256"
	"time"

	"github.com/samuelhornsey/file-upload/config"
)

// Define file struct
type file struct {
	Id int64 `json:"id"`
	FileName string `json:"fileName"`
	Hash string `json:"hash"`
}

// Create new file
func Create(fileName string) (f *file) {
	currentTime := time.Now()
	f = new(file)
	h := sha256.New()
	h.Write([]byte(fileName))
	h.Write([]byte(currentTime.String()))
	f.Hash = fmt.Sprintf("%x", h.Sum(nil))
	f.FileName = fileName

	return f
}

// Insert file to db
func (f *file) Insert() {
	// Init db and defer close
	db := config.DB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO files (filename, hash) VALUES (?, ?)")

	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(f.FileName, f.Hash)

	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	f.Id = id
}

// Update file
func (f *file) Update(filename string) {
	// Init db and defer close
	db := config.DB()
	defer db.Close()

	f.FileName = filename

	stmt, err := db.Prepare("UPDATE files SET filename = ? WHERE ID = ?")

	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(f.FileName, f.Id)

	if err != nil {
		panic(err)
	}
}

// Delete Model
func (f *file) Delete() {
	// Init db and defer close
	db := config.DB()
	defer db.Close()

	stmt, err := db.Prepare("DELET FROM files WHERE ID = ?")

	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(f.FileName, f.Id)

	if err != nil {
		panic(err)
	}

	f = nil
}

// Get file by id
func Get(id int) (f *file) {
	// Init db and defer close
	db := config.DB()
	defer db.Close()

	f = new(file)

	row, err := db.Query("SELECT * FROM files WHERE ID = ?", id)

	if err != nil {
		panic(err)
	}

	if row.Next() {
		err = row.Scan(&f.Id, &f.FileName, &f.Hash)

		if err != nil {
			panic(err)
		}
	}

	return f
}

// Get all files
func GetAll() (f []*file) {
	// Init db and defer close
	db := config.DB()
	defer db.Close()
	
	f = []*file{}

	row, err := db.Query("SELECT * FROM files")

	if err != nil {
		panic(err)
	}

	for row.Next() {
		nextFile := new(file)

		err = row.Scan(&nextFile.Id, &nextFile.FileName, &nextFile.Hash)

		if err != nil {
			panic(err)
		}

		f = append(f, nextFile)
	}

	return f
}