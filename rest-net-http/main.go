package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
type Album struct {
	ID int64
	Title string
	Artist string
	Price float32
}

func albumByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumByArtist %q: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price);err != nil {
			return nil, fmt.Errorf("albumByArtist %q: %v", err)
		}

		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumByArtist %q: %v", err)
	}

	return albums, nil
}

func main() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Error on opening db:%v", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalf("error on db ping:%v", pingErr)
	}

	fmt.Println("connected :)")

	albums, err := albumByArtist("weekend")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album found: %v\n", albums)
}
