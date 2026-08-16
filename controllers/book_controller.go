package controllers

import (
	"book-management-api/database"
	"book-management-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book

	if err := database.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "buku tidak tersedia",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func GetBookByID(c *gin.Context) {
	bookId := c.Param("id")
	var book models.Book

	if err := database.DB.First(&book, bookId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Buku tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func CreateBook(c *gin.Context) {
	var bookData *models.Book
	var books []models.Book

	if err := c.ShouldBindJSON(&bookData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "judul harus dilengkapi",
		})
		return
	}

	newBook := []models.Book{
		{
			Title:       bookData.Title,
			Author:      bookData.Author,
			ReleaseYear: bookData.ReleaseYear,
			CategoryID:  bookData.CategoryID,
		},
	}

	database.DB.Create(&newBook)
	database.DB.Preload("Category").Last(&books)

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func EditBook(c *gin.Context) {
	bookId := c.Param("id")
	var newBookData *models.Book
	var books []models.Book

	if err := c.ShouldBindJSON(&newBookData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "judul harus dilengkapi",
		})
		return
	}

	database.DB.Model(&books).Where("id = ?", bookId).Updates(models.Book{
		Title:       newBookData.Title,
		Author:      newBookData.Author,
		ReleaseYear: newBookData.ReleaseYear,
		CategoryID:  newBookData.CategoryID,
	})

	database.DB.First(&books, bookId)
	c.JSON(http.StatusCreated, gin.H{
		"data": books,
	})
}

func DeleteBook(c *gin.Context) {
	bookId := c.Param("id")
	var books models.Book

	if err := database.DB.First(&books, bookId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Buku tidak tersedia",
		})
		return
	}

	database.DB.Delete(&books)
	c.JSON(http.StatusNoContent, gin.H{})
}
