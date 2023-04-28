package main

import (
	"net/http"

	docs "go_study/api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	docs.SwaggerInfo.Title = "Album API"
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// docs.SwaggerInfo.Description = "This is a sample server for Swagger."
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "petstore.swagger.io"
	router := setRouter()

	router.Run(":8080")
}

func setRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", getAlbums) // curl http://localhost:8080/albums --header "Content-Type: application/json" --request "GET"
	// router.POST("/albums", postAlbums) // curl http://localhost:8080/albums --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4", "title": "Panda", "artist": "Designer", "price": 49.99}'
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}

// getAlbums responds with the list of all albums as JSON.

// getAlbums godoc
// @Router /albums [get]
// @Summary get albums list json
// @Schemes
// @Description description
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Success 200 {json} albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
	// IndentedJSON -> struct를 JSON 형태로 serialize 해주고 response에 추가해준다.
	// StatusOK -> 200 OK
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
	// StatusCreated -> 201
}
