package dao

import (
	"go-short-url/config"
	"go-short-url/model/entity"
)

type ShortUrlDAO struct {
}

/*
	func init() {
		if (!config.DB.Migrator().HasTable(&entity.ShortUrlEntity{})) {
			err := config.DB.AutoMigrate(&entity.ShortUrlEntity{})
			if err != nil {
				return
			}
		}
	}
*/
func (shortDAO ShortUrlDAO) GetBySHortUrl(shortUrl string) *entity.ShortUrlEntity {
	var urlEntity entity.ShortUrlEntity
	config.DB.Where("short_url=?", shortUrl).First(&urlEntity)
	return &urlEntity
}

func (shortDAO ShortUrlDAO) AddOrUpdate(shortUrl *entity.ShortUrlEntity) {
	if shortUrl.ID == 0 {
		config.DB.Save(shortUrl)
	} else {
		config.DB.Updates(shortUrl)
	}
}
