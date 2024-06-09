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
	Title  string
	Artist string
	ID     int64
	Price  float32
}

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

  pingErr := db.Ping()
  if pingErr != nil {
    log.Fatal(pingErr);
  }

  fmt.Println("Connected", cfg.FormatDSN())

  albums, err := albumsByArtist("John Coltrane")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Albums found %v\n", albums)

  album, err := albumByID(1)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Found album:", album)

  newAlbum := Album{
    Title: "The Modern Sound of Betty Carter",
    Artist: "Betty Carter",
    Price: 49.99,
  }
  albumId, err := addAlbum(newAlbum)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("ID of new album: %v", albumId)
}

func albumsByArtist(name string) ([]Album, error) {
  var albums []Album

  rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
  if err != nil {
    return nil, fmt.Errorf("albumsByArtists %q, %v", name, err)
  }

  defer rows.Close()

  for rows.Next() {
    var album Album
    if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
      return nil, fmt.Errorf("albumsByArtist %q, %v", name, err)
    }
    albums = append(albums, album)
  }

  if err := rows.Err(); err != nil {
    return nil, fmt.Errorf("albumsByArtist %q, %v", name, err)
  }

  return albums, nil
}

func albumByID(id int64) (Album, error) {
  var album Album

  row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
  if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
    if err == sql.ErrNoRows {
      return album, fmt.Errorf("albumByID %d: no such album", id)
    }
    return album, fmt.Errorf("albumByID %d, %v", id, err)
  }

  return album, nil
}

func addAlbum(album Album) (int64, error) {
  result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
  if err != nil {
    return 0, fmt.Errorf("allAlbum: %v", err)
  }

  id, err := result.LastInsertId()
  if err != nil {
    return 0, fmt.Errorf("addAlbum: %v", err)
  }

  return id, nil
}
