package models

type Surah struct {
	Id       uint `json:"id" gorm:"primary_key"`
	NamaSurah    string `json:"nama_surah"`
	JuzId uint `json:"juz_id" gorm:"not null"`
  }