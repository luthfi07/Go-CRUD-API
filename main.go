package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"pustaka-api/book"
	"pustaka-api/handler"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=user1 password=user1 dbname=book port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	fmt.Println("database connection success")

	db.AutoMigrate(&book.Book{})
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// bookRequest := book.BookRequest{
	// 	Title: "Buku Baru",
	// 	//		Description: "ini adalah buku baru yang ditambahkan",
	// 	Price: "80990",
	// 	//Rating:      4,
	// }

	// bookService.Create(bookRequest)

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run(":8888")
}

//main
//handler
//service
//repository
//db
//postgresSQL
