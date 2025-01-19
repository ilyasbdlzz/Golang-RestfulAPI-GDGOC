package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
}
	var albums = []album {
		{ID: "1", Title: "Red", Price: 56.99},
		{ID: "2", Title: "Blue", Price: 56.99},
		{ID: "3", Title: "Green", Price: 56.99},
	}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if error :=c.BindJSON(& newAlbum); error != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H {"message" :"Invalid Request" })
		return
	}

	if newAlbum.ID == "" || newAlbum.Title == "" || newAlbum.Price <= 0 {
        c.IndentedJSON(http.StatusBadRequest, gin.H{
            "message": "Kolom ID, Title, & Price harus diisi dan dengan format yg valid",
        })
        return
    }

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context){
	id := c.Param("id")
	
	for _, album := range albums{
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Not Found"})
} 
func updateAlbumByID(c *gin.Context){
	id := c.Param("id")
	var updatedAlbum album

	if error := c.BindJSON(&updatedAlbum) ; error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H {"message" : "Invalid request"})
		return
	}
	for i, album := range albums {
		if album.ID == id {
			// Update the album details
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Album deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}


func main() {
	router := gin.Default()

	router.GET ("/albums", getAlbums)
	router.POST ("/albums", postAlbums)
	router.GET ("/albums/:id", getAlbumById)
	router.PUT ("/albums/:id", updateAlbumByID)
	router.DELETE ("/albums/:id", deleteAlbumByID)

	router.Run ("localhost:8080")
}