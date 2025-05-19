package DB

import (
	"APIdinyPodkluchiny/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// GetMeme возвращает список всех мемов
// @Summary Получить все мемы
// @Description Возвращает список всех мемов из базы данных
// @Tags Memes
// @Accept json
// @Produce json
// @Success 200 {array} models.Meme
// @Failure 500 {object} map[string]string
// @Router /memes [get]
func GetMeme(c *gin.Context, db *gorm.DB) {
	var memes []models.Meme
	if err := db.Find(&memes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, memes)
}

// GetMemeByID возвращает мем по ID
// @Summary Получить мем по ID
// @Description Возвращает мем с указанным ID
// @Tags Memes
// @Accept json
// @Produce json
// @Param id path int true "ID мема"
// @Success 200 {object} models.Meme
// @Failure 404 {object} map[string]string
// @Router /memes/{id} [get]
func GetMemeByID(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var meme models.Meme

	if err := db.First(&meme, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Meme not found"})
		return
	}
	c.JSON(http.StatusOK, meme)
}

// CreateMeme создает новый мем
// @Summary Создать новый мем
// @Description Добавляет новый мем в базу данных
// @Tags Memes
// @Accept json
// @Produce json
// @Param meme body models.Meme true "Данные нового мема"
// @Success 201 {object} models.Meme
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /memes [post]
func CreateMeme(c *gin.Context, db *gorm.DB) {
	var NewMeme models.Meme
	if err := c.ShouldBindJSON(&NewMeme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&NewMeme).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, NewMeme)
}

// UpdateMeme обновляет данные мема
// @Summary Обновить мем
// @Description Обновляет данные мема по ID
// @Tags Memes
// @Accept json
// @Produce json
// @Param id path int true "ID мема"
// @Param updates body map[string]interface{} true "Обновляемые поля"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /memes/{id} [patch]
func UpdateMeme(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var updatedData map[string]interface{}
	if err := c.BindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Model(&models.Meme{}).Where("idmeme = ?", id).Updates(updatedData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Meme not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Meme updated"})
}

// DeleteMeme удаляет мем
// @Summary Удалить мем
// @Description Удаляет мем по указанному ID
// @Tags Memes
// @Accept json
// @Produce json
// @Param id path int true "ID мема"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /memes/{id} [delete]
func DeleteMeme(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	result := db.Where("idmeme = ?", id).Delete(&models.Meme{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Meme not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "meme deleted"})
}

// CreateTag создает новый тег
// @Summary Создать новый тег
// @Description Добавляет новый тег в базу данных
// @Tags Tags
// @Accept json
// @Produce json
// @Param meme body models.Tag true "Данные нового тега"
// @Success 201 {object} models.Tag
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tags [post]
func CreateTag(c *gin.Context, db *gorm.DB) {
	var NewTag models.Tag
	if err := c.ShouldBindJSON(&NewTag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&NewTag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, NewTag)
}

// UpdateTag обновляет тег
// @Summary Обновть тег
// @Description Обновляет данные тега по ID
// @Tags Tags
// @Accept json
// @Produce json
// @Param id path int true "ID тега"
// @Param updates body map[string]interface{} true "Обновляемые поля"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tags/{id} [patch]
func UpdateTag(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var updatedData map[string]interface{}
	if err := c.BindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.Model(&models.Tag{}).Where("idtag = ? ", id).Updates(updatedData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Tag not found"})
		return
	}
	c.JSON(http.StatusOK, updatedData)
}

// DeleteTag удаляет тег
// @Summary Удалить тег
// @Description Удаляет тег по указанному ID
// @Tags Tags
// @Accept json
// @Produce json
// @Param id path int true "ID тега"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tags/{id} [delete]
func DeleteTag(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	result := db.Where("idtag = ? ", id).Delete(&models.Tag{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Tag not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "tag deleted"})
}

// CreateCategory создает новую категорию мемов
// @Summary Создать новую категорию мемов
// @Description Добавляет новую категорию мемов в базу данных
// @Tags Categories
// @Accept json
// @Produce json
// @Param meme body models.Category true "Данные новой категории"
// @Success 201 {object} models.Category
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories [post]
func CreateCategory(c *gin.Context, db *gorm.DB) {
	var NewCategory models.Category
	if err := c.ShouldBindJSON(&NewCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&NewCategory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, NewCategory)
}

// UpdateCategory обновляет категорию мемов
// @Summary Обновить категорию мемов
// @Description Обновляет категорию мемов по указанному ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "ID категории"
// @Param updates body map[string]interface{} true "Обновляемые поля"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories [patch]
func UpdateCategory(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var updatedData map[string]interface{}
	if err := c.BindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Where("idcategory = ? ", id).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, updatedData)
}

// DeleteCategory удаляет категорию мемов
// @Summary Удалить категорию мемов
// @Description Удаляет категорию мемов по указанному ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "ID категории"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	result := db.Where("idcategory = ? ", id).Delete(&models.Category{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}
