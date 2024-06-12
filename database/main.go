package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type AlbumsDB struct {
	db *sql.DB
}

func (c *AlbumsDB) Connect() error {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error
	c.db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := c.db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected", cfg.FormatDSN())

	return nil
}

type Album struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	ID     int64   `json:"id"`
	Price  float32 `json:"price"`
}

func (c *AlbumsDB) Albums() ([]Album, error) {
	var albums []Album

	rows, err := c.db.Query("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("albums %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %v", err)
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %v", err)
	}

	return albums, nil
}

func (c *AlbumsDB) AlbumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := c.db.Query("SELECT * FROM album WHERE artist = ?", name)
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

func (c *AlbumsDB) AlbumByID(id int64) (Album, error) {
	var album Album

	row := c.db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("albumByID %d: no such album", id)
		}
		return album, fmt.Errorf("albumByID %d, %v", id, err)
	}

	return album, nil
}

func (c *AlbumsDB) AddAlbum(album Album) (int64, error) {
	result, err := c.db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, fmt.Errorf("allAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}
