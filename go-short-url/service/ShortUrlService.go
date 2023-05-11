package service

import (
	"go-short-url/dao"
	"go-short-url/model/entity"
	"go-short-url/model/param"
	"go-short-url/tools"
	"time"
)

type ShortUrlService struct {
}

func (service ShortUrlService) AddOrUpdate(update *param.UpdateShortUrlParam) string {
	var urlEntity entity.ShortUrlEntity
	urlEntity.Deleted = 0
	urlEntity.ID = update.Id
	urlEntity.UpdatedAt = time.Now()
	urlEntity.NativeUrl = update.NativeUrl
	urlEntity.ShortUrl = tools.RandAllString(8)
	urlEntity.ValidTime = update.ValidTime
	urlEntity.CreatedAt = time.Now()
	dao.ShortUrlDAO{}.AddOrUpdate(&urlEntity)
	return urlEntity.ShortUrl
}

func (service ShortUrlService) GetByShortUrl(shortUrl string) string {
	shortEntity := dao.ShortUrlDAO{}.GetBySHortUrl(shortUrl)
	return shortEntity.NativeUrl
}
