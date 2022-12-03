package main

import (
	"bookshop/protos"
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type BookService struct {
	protos.UnsafeBookServiceServer
}

func (c *BookService) GetBookById(ctx context.Context, in *protos.BookID) (*protos.Book, error) {
	db := ConnectDB()
	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM books where id = %d`, in.ID))
	defer db.Close()
	CheckError(err)
	defer rows.Close()
	var book protos.Book
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.TotalPages, &book.Author, &book.GenreID)
		CheckError(err)
	}

	CheckError(err)

	return &book, nil
}

func (c *BookService) CreateBook(ctx context.Context, in *protos.Book) (*protos.BookID, error) {
	db := ConnectDB()
	insertDynStmt := `insert into books (title, id, genre_id, total_pages, author) values($1, $2, $3, $4, $5)`
	_, err := db.Exec(insertDynStmt, in.Title, in.ID, in.GenreID, in.TotalPages, in.Author)
	defer db.Close()
	CheckError(err)
	return &protos.BookID{ID: in.ID}, nil
}

func (c *BookService) DeleteBook(ctx context.Context, in *protos.BookID) (*protos.BookID, error) {
	db := ConnectDB()
	deleteStmt := `delete from books where id=$1`
	_, err := db.Exec(deleteStmt, in.ID)
	defer db.Close()
	CheckError(err)
	return &protos.BookID{ID: in.ID}, nil
}

func (c *BookService) EditBook(ctx context.Context, in *protos.Book) (*protos.Book, error) {
	db := ConnectDB()
	updateStmt := `update books set id=$1, total_pages=$2, genre_id=$3, title=$4, author=$5 where "id"=$1`
	_, err := db.Exec(updateStmt, in.ID, in.TotalPages, in.GenreID, in.Title, in.Author)
	defer db.Close()
	CheckError(err)
	return &protos.Book{ID: in.ID, TotalPages: in.TotalPages, GenreID: in.GenreID, Title: in.Title, Author: in.Author}, nil
}

func (c *BookService) CreateGenre(ctx context.Context, in *protos.Genre) (*protos.Genre, error) {
	db := ConnectDB()
	insertDynStmt := `insert into genres (id, name) values($1, $2)`
	_, err := db.Exec(insertDynStmt, in.GenreId, in.Name)
	defer db.Close()
	CheckError(err)
	return &protos.Genre{GenreId: in.GenreId, Name: in.Name}, nil
}

func (c *BookService) GetGenres(ctx context.Context, in *protos.Empty) (*protos.GenreResponse, error) {
	db := ConnectDB()
	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM genres limit %d`, in.Limit))
	defer db.Close()
	CheckError(err)
	var genreResponse protos.GenreResponse
	defer rows.Close()
	for rows.Next() {
		var single protos.Genre
		err = rows.Scan(&single.GenreId, &single.Name)
		CheckError(err)
		genreResponse.Genres = append(genreResponse.Genres, &single)
	}

	return &genreResponse, nil
}

func main() {
	log := hclog.Default()
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	protos.RegisterBookServiceServer(grpcServer, &BookService{})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error("Unable to listen 8080", err)
	}
	grpcServer.Serve(lis)
}
