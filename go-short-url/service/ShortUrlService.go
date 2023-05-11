package service

import (
	"errors"
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

func (service ShortUrlService) Redirect(shortUrl string) (string, error) {
	shortEntity := dao.ShortUrlDAO{}.GetBySHortUrl(shortUrl)
	if shortEntity == nil {
		return "", errors.New("该短地址不存在")
	}
	validTimeSecond := uint64(shortEntity.UpdatedAt.Unix()) + shortEntity.ValidTime
	if validTimeSecond < uint64(time.Now().Unix()) {
		return "", errors.New("该短地址已过期")
	}
	return shortEntity.NativeUrl, nil
}
