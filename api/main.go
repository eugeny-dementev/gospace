package main

import (
	"fmt"
	"log"

	"example.com/database"
)

func main() {
	albumsDB := database.AlbumsDB{}
	albumsDB.Connect()

	albums, err := albumsDB.AlbumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found %v\n", albums)

	album, err := albumsDB.AlbumByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found album:", album)

	newAlbum := database.Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	}
	albumId, err := albumsDB.AddAlbum(newAlbum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of new album: %v", albumId)
}
