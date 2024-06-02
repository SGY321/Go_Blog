// Package article 应用的文章模型
package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"strconv"
)

// Article 文章模型
type Article struct {
	models.BaseModel

	Title string
	Body  string
}

// 为 Article 添加一个 Link 方法，生成文章的链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))

}
