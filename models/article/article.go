package article

import (
	"github.com/dengzii/blog_server/db"
	"time"
)

func (that *Article) Insert() (article *Article, err error) {
	desc := that.Content
	if len(that.Content) > 50 {
		desc = that.Content[:50]
	}
	that.Description = desc
	that.CreatedAt = time.Now().Unix()
	db.Insert(that)
	return that, nil
}

func Insert(newArticle *Article) (article *Article, err error) {
	desc := newArticle.Content
	if len(newArticle.Content) > 50 {
		desc = newArticle.Content[:50]
	}
	newArticle.Description = desc
	//newArticle.CreatedAt = time.Now().Unix()
	//newArticle.UpdatedAt = newArticle.CreatedAt
	db.Insert(newArticle)
	return newArticle, nil
}

func GetArticleLatest(from int64, count int) (articles []*ArticleBase) {
	var article []*Article

	println(from)
	db.Mysql.
		Find(&article).
		Where("updated_at < ?", from).
		Order("created_at", true).
		Limit(count)

	articles = make([]*ArticleBase, len(article))
	for i, v := range article {
		articles[i] = v.toArticleBase()
	}
	return articles
}

func GetArticle(id int) *ArticleBase {
	var article ArticleBase
	db.Mysql.Where("id = ?", id).Find(&article).Limit(1)
	return &article
}

func DelArticle(id int) {

}

func UpdateArticle() {

}

func LikeArticle() {

}

func CommentArticle() {

}
