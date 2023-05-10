package service

import (
	"go-short-url/dao"
	"go-short-url/model/param"
)

type ShortUrlService struct {
}

func (service ShortUrlService) AddOrUpdate(update *param.UpdateShortUrlParam) {
	dao.ShortUrlDAO{}.AddOrUpdate(update)
}
