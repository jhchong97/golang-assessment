package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Genre       string    `json:"genre"`
	Description string    `json:"description"`
	ISBN        string    `json:"isbn"`
	Image       string    `json:"image"`
	Published   string    `json:"published"` // can also use time.Time if needed
	Publisher   string    `json:"publisher"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
}

var db *sql.DB
var err error

func main() {
	// Change to your DB credentials
	dsn := "user1:password@tcp(127.0.0.1:3306)/emb?parseTime=true"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}
	defer db.Close()

	router := gin.Default()
	router.Use(cors.Default()) // Allow all origins for development

	// CRUD Routes
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.POST("/books", createBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)

	router.Run(":3000")
}

func getBooks(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(
			&b.ID, &b.Title, &b.Author, &b.Genre, &b.Description, &b.ISBN,
			&b.Image, &b.Published, &b.Publisher, &b.CreatedAt, &b.UpdatedAt,
			&b.CreatedBy, &b.UpdatedBy,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, b)
	}
	c.JSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	var b Book
	err := db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(
		&b.ID, &b.Title, &b.Author, &b.Genre, &b.Description, &b.ISBN,
		&b.Image, &b.Published, &b.Publisher, &b.CreatedAt, &b.UpdatedAt,
		&b.CreatedBy, &b.UpdatedBy,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, b)
}

func createBook(c *gin.Context) {
	var b Book
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec(
		`INSERT INTO books (title, author, genre, description, isbn, image, published, publisher, created_by, updated_by)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		b.Title, b.Author, b.Genre, b.Description, b.ISBN, b.Image, b.Published, b.Publisher, "system", "system",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	b.ID = int(id)
	c.JSON(http.StatusCreated, b)
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var b Book
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(
		`UPDATE books SET title=?, author=?, genre=?, description=?, isbn=?, image=?, published=?, publisher=?, updated_by='system'
		 WHERE id=?`,
		b.Title, b.Author, b.Genre, b.Description, b.ISBN, b.Image, b.Published, b.Publisher, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
