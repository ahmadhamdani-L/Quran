package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang-api1/models"
)

type JuzInput struct {
	NamaJuz string `json:"nama_juz"`
	Juz    string `json:"juz"`
	SurahId  uint     `json:"surah_id"`
	Surah []models.Surah `json:"surah" gorm:"foreignKey:SurahId"`
}

func JuzTampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var juz []models.Juz
	db.Preload("Surah").Find(&juz)
	
	
	c.JSON(http.StatusOK, gin.H{"juz": juz})

}

func JuzAdd(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var juzInput JuzInput
	if err := c.ShouldBindJSON(&juzInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	juz := models.Juz{
		NamaJuz: juzInput.NamaJuz,
		Juz: juzInput.Juz,
	}

	db.Create(&juz)

	c.JSON(http.StatusOK, gin.H{"status": "data berhasil di tambhakan"})
}

func JuzChange(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var juz models.Juz
	if err := db.Where("id = ?", c.Param("id")).First(&juz).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "juz tidak ditemukan"})
		return
	}

	var juzInput JuzInput
	if err := c.ShouldBindJSON(&juzInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//proses ubah nya
	db.Model(&juz).Update(juzInput)

	c.JSON(http.StatusOK, gin.H{"data": juz})
}

func JuzDelete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var juz models.Juz
	if err := db.Where("id = ?", c.Param("id")).First(&juz).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "juz tidak ditemukan"})
		return
	}

	//proses delete nya
	db.Delete(&juz)

	c.JSON(http.StatusOK, gin.H{"data": "berhasil di hapus"})
}



func FindByIdJuz(c *gin.Context) { // Get model if exist
	var juz models.Juz

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&juz).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Juz Not Found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": juz})
}
