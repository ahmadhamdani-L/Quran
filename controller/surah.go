package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang-api1/models"
)

type SurahInput struct {
	Id       uint `json:"id" gorm:"primary_key"`
	NamaSurah    string `json:"nama_surah"`
	JuzId uint `json:"juz_id"`
}

func SurahTampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var surah []models.Surah
	db.Find(&surah)
	
	
	c.JSON(http.StatusOK, gin.H{"surah": surah})

}




func SurahChange(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var surah models.Surah
	if err := db.Where("id = ?", c.Param("id")).First(&surah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "surah tidak ditemukan"})
		return
	}

	var surahInput SurahInput
	if err := c.ShouldBindJSON(&surahInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//proses ubah nya
	db.Model(&surah).Update(surahInput)

	c.JSON(http.StatusOK, gin.H{"data": surah})
}

func SurahDelete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var surah models.Surah
	if err := db.Where("id = ?", c.Param("id")).First(&surah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "surah tidak ditemukan"})
		return
	}

	//proses delete nya
	db.Delete(&surah)

	c.JSON(http.StatusOK, gin.H{"data": "berhasil di hapus"})
}

//surah

func SurahAdd(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var surahInput SurahInput
	if err := c.ShouldBindJSON(&surahInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	surah := models.Surah{
		NamaSurah: surahInput.NamaSurah,
		JuzId: surahInput.JuzId,
	}

	var juz models.Juz
	db.Find(&juz)
	if  (surahInput.JuzId != juz.Id){
		c.JSON(http.StatusBadRequest, gin.H{"error":"juz tidak ditemukan"})
	}else{
		db.Create(&surah)

		c.JSON(http.StatusOK, gin.H{"status": "data berhasil di tambhakan"})
	}	
	
}

func FindByIdsurah(c *gin.Context) { // Get model if exist
	var surah models.Surah

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&surah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Surah Not Found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": surah})
}
