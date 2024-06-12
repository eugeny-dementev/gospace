package main

import "example.com/database"

func main() {
	albumsDB := database.AlbumsDB{}
	albumsDB.Connect()
}
