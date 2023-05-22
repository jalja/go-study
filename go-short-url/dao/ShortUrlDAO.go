package dao

import (
	"go-short-url/config"
	"go-short-url/model/entity"
	"go-short-url/model/param"
	"go-short-url/model/response"
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
func (shortDAO ShortUrlDAO) PageList(param *param.ShortUrlParam) *response.Pagination {
	short := entity.ShortUrlEntity{
		ShortUrl: param.ShortUrl,
	}
	count := shortDAO.Count(param)
	pagination := response.NewPagination(param.PageSize, param.CurrentPage, param.Sort, int(count))
	var shortUrlEntity []*entity.ShortUrlEntity
	config.DB.Model(&entity.ShortUrlEntity{}).Where(short).Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Find(&shortUrlEntity)
	pagination.SetRows(shortUrlEntity)
	return pagination
}

func (shortDAO ShortUrlDAO) Count(param *param.ShortUrlParam) int64 {
	var count int64
	short := entity.ShortUrlEntity{
		ShortUrl: param.ShortUrl,
	}
	config.DB.Model(&entity.ShortUrlEntity{}).Where(short).Count(&count)
	return count
}

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
