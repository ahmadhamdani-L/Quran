package models


type Juz struct {
	Id        uint  `json:"id" gorm:"primary_key"`
	NamaJuz    string `json:"nama_juz"`
	Juz string `json:"juz"`
	Surah []Surah
}
