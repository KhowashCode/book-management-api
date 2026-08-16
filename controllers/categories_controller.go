package controllers

import (
	"book-management-api/database"
	"book-management-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category

	if err := database.DB.Find(&categories).Error; err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})

}

func CreateCategory(c *gin.Context) {
	var category *models.Category
	var categories []models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nama kategori harus diisi",
		})
		return
	}

	newCategory := []models.Category{
		{
			Name: category.Name,
		},
	}

	database.DB.Create(&newCategory)
	database.DB.Last(&categories)

	c.JSON(http.StatusCreated, gin.H{
		"data": categories,
	})
}

func EditCategory(c *gin.Context) {
	categoryId := c.Param("id")
	var category []models.Category
	var newCategoryData *models.Category

	if err := c.ShouldBindJSON(&newCategoryData); err != nil {
		return
	}

	database.DB.Model(&category).Where("id = ?", categoryId).Updates(models.Category{
		Name: newCategoryData.Name,
	})

	database.DB.First(&category, categoryId)
	c.JSON(http.StatusCreated, gin.H{
		"data": category,
	})
}

func DeleteCategory(c *gin.Context) {
	categoryId := c.Param("id")
	var category models.Category

	if err := database.DB.First(&category, categoryId).Error; err != nil {
		return
	}

	database.DB.Delete(&category)
	c.JSON(http.StatusNoContent, gin.H{})
}
