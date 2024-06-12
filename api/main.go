package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/database"
	"github.com/gin-gonic/gin"
)

func getAlbumsFrom(db *database.AlbumsDB) func(c *gin.Context) {
	return func(c *gin.Context) {
		albums, err := db.Albums()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Albums found %v\n", albums)
		c.IndentedJSON(http.StatusOK, albums)
	}
}

func postAlbumTo(db *database.AlbumsDB) func (c *gin.Context) {
  return func(c *gin.Context) {
    var newAlbum database.Album

    if err := c.BindJSON(&newAlbum); err != nil {
      log.Print("Failed to extract album from payload")
      return;
    }

    id, err := db.AddAlbum(newAlbum)
    if err != nil {
      log.Fatal(err)
      return
    }
    newAlbum.ID = id

    c.IndentedJSON(http.StatusCreated, newAlbum)
  }
}

func main() {
	albumsDB := database.AlbumsDB{}
	albumsDB.Connect()

	router := gin.Default()
	router.GET("/", getAlbumsFrom(&albumsDB))
  router.POST("/albums", postAlbumTo(&albumsDB))

	router.Run("localhost:8888")

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
