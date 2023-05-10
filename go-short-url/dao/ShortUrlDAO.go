package dao

import (
	"go-short-url/config"
	"go-short-url/model/entity"
	"go-short-url/model/param"
	"time"
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
	config.DB.Find(&urlEntity, shortUrl)
	return &urlEntity
}

func (shortDAO ShortUrlDAO) AddOrUpdate(update *param.UpdateShortUrlParam) {
	var urlEntity entity.ShortUrlEntity
	urlEntity.Deleted = 0
	urlEntity.UpdatedAt = time.Now()
	urlEntity.NativeUrl = update.NativeUrl
	urlEntity.ShortUrl = update.ShortUrl
	urlEntity.ValidTime = update.ValidTime
	if update.Id != 0 {
		urlEntity.ID = update.Id
		config.DB.Updates(&urlEntity)
		return
	}
	urlEntity.CreatedAt = time.Now()
	config.DB.Save(&urlEntity)
}
