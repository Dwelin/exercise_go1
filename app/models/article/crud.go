package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/types"
)

// Get 通过 ID 获取文章
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idstr)
	if err := model.ConnectDB().First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

// GetAll 获取全部文章
func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.ConnectDB().Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

// Create 创建文章，通过 article.ID 来判断是否创建成功
func (article *Article) Create() (err error) {
	result := model.ConnectDB().Create(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// Update 更新文章
func (article Article) Update() (rowsAffected int64, err error) {
	result := model.ConnectDB().Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}
