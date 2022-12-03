package handlers

import (
	"bookshopV1/protos"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Empty struct {
}

type HandlerV1 struct {
	cclient protos.BookServiceClient
}

func New() *HandlerV1 {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("error while Dial gRPC on 8080")
	}
	cclient := protos.NewBookServiceClient(conn)
	return &HandlerV1{cclient: cclient}
}

// @Summary Recives book by id number
// @Param id path number true "bookid"
// @Success 200 {object} protos.Book
// @Router /{id} [get]
// @Tags book
func (h *HandlerV1) GetBookById(c *gin.Context) {

	var id int32
	i, err := strconv.Atoi(c.Param("id"))
	if HaveError(err, http.StatusBadRequest, c, "incorrect id number for book") {
		return
	}
	id = int32(i)

	book, err := h.cclient.GetBookById(context.Background(), &protos.BookID{ID: id})
	if HaveError(err, http.StatusBadRequest, c, "erro while geting book") {
		return
	}
	c.JSON(http.StatusOK, book)
}

// @Summary Creates book
// @Param book body protos.Book true "book"
// @Success 200 {object} protos.BookID
// @Router / [put]
// @Tags book
func (h *HandlerV1) CreateBook(c *gin.Context) {
	var requestBody protos.Book
	var BookID *protos.BookID
	err := c.ShouldBindJSON(&requestBody)
	if HaveError(err, http.StatusBadRequest, c, "") {
		return
	}
	BookID, err = h.cclient.CreateBook(context.Background(), &requestBody)
	if HaveError(err, http.StatusBadRequest, c, "erro while creating book") {
		return
	}
	c.JSON(http.StatusOK, BookID)
}

// @Summary Deletes book
// @Param id header number true "book id"
// @Success 200 {object} protos.BookID
// @Router / [delete]
// @Tags book
func (h *HandlerV1) DeleteBook(c *gin.Context) {
	var requestBody protos.BookID
	err := c.ShouldBindJSON(&requestBody)
	if HaveError(err, http.StatusBadRequest, c, "") {
		return
	}
	BookID, err := h.cclient.DeleteBook(context.Background(), &requestBody)
	if HaveError(err, http.StatusBadRequest, c, "erro while deleting book") {
		return
	}
	c.JSON(http.StatusOK, BookID)
}

// @Summary Updates book
// @Param book body protos.Book true "book"
// @Sucsess 200 {object} protos.Book
// @Router / [post]
// @Tags book
func (h *HandlerV1) EditBook(c *gin.Context) {
	var requestBody protos.Book
	err := c.ShouldBindJSON(&requestBody)
	if HaveError(err, http.StatusBadRequest, c, "") {
		return
	}
	Book, err := h.cclient.EditBook(context.Background(), &requestBody)
	if HaveError(err, http.StatusBadRequest, c, "erro while editing book") {
		return
	}
	c.JSON(http.StatusOK, Book)
}

// @Summary Creates new genre
// @Param genre body protos.Genre true "genre"
// @Success 200 {object} protos.Genre
// @Router /genre [put]
// @Tags genre
func (h *HandlerV1) CreateGenre(c *gin.Context) {

	var requestBody protos.Genre
	err := c.ShouldBindJSON(&requestBody)

	if HaveError(err, http.StatusBadRequest, c, "") {
		return
	}

	Genre, err := h.cclient.CreateGenre(context.Background(), &requestBody)

	if HaveError(err, http.StatusBadRequest, c, "erro while creating genre") {
		return
	}

	c.JSON(http.StatusOK, Genre)
}

// @Summary Recives list of genres
// @Param limit path number true "Limit count of genres"
// @Success 200 {object} protos.GenreResponse
// @Router /genre/{limit} [get]
// @Tags genre
func (h *HandlerV1) GetGenres(c *gin.Context) {

	var limit int32
	i, _ := strconv.ParseInt(c.Param("limit"), 10, 32)
	limit = int32(i)

	Genres, err := h.cclient.GetGenres(context.Background(), &protos.Empty{Limit: limit})
	if HaveError(err, http.StatusBadRequest, c, "erro while getting genres list") {
		return
	}
	c.JSON(http.StatusOK, Genres)

}

func HaveError(err error, code int, c *gin.Context, message string) bool {
	if err != nil {
		c.JSON(code, gin.H{
			"error": message,
		})
		return true
	}
	return false
}

func GetInt32(key string, c *gin.Context) (i int32) {
	if val, ok := c.Get(key); ok && val != nil {
		i, _ = val.(int32)
	}
	return
}
