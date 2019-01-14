package models

import (
	"github.com/jinzhu/gorm"
)

type Diary struct {
	gorm.Model
	Content string `valid:"required,length(1|130)"`
}

type Person struct {
	gorm.Model
	FirstName string
	LastName  string
}

//func (diary Diary) Validate(db *gorm.DB) {
//	if diary.Content == "" {
//		db.AddError(validations.NewError(diary, "Content", "Content is required"))
//	}
//	if diary.Content > 130 {
//		db.AddError(validations.NewError(diary, "Content", "Your story is too long, my friend"))
//	}
//}

func GetDiaries(db *gorm.DB) []Diary {
	var diaries []Diary

	if err := db.Find(&diaries).Error; err != nil {
		panic(err)
	} else {
		return diaries
	}
}

// old method and test
//func CreateDiary(db *gorm.DB, diary Diary) (Diary) {
//	if err := db.Create(&diary).GetErrors(); err != nil {
//		panic(err)
//	} else {
//		return diary
//	}
//}

func CreateDiary(db *gorm.DB, diary Diary) error {
	err := db.Create(&diary)
	return err.Error
}
