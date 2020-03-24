package article

import (
	"errors"
	"github.com/dengzii/blog_server/controllers"
	"github.com/dengzii/blog_server/models/article"
	"github.com/kataras/iris/context"
)

func GetArticleLatest(ctx context.Context) (err error) {

	articles := article.GetArticleLatest(10)
	responseJson := controllers.ErrorResponse(200, "success", articles)
	_, err = ctx.JSON(responseJson)
	return err
}

func GetArticle(ctx context.Context) (err error) {

	id, err := ctx.Params().GetInt("id")
	if err != nil {
		return
	}
	articles := article.GetArticle(id)
	responseJson := controllers.ErrorResponse(200, "success", articles)
	_, err = ctx.JSON(responseJson)
	return
}

func AddArticle(ctx context.Context) (err error) {

	articleJson := &article.ArticleJson{}
	err = ctx.ReadJSON(articleJson)
	if err != nil {
		return err
	}
	art := article.AddArticle(articleJson)
	if art == nil {
		return errors.New("create article failure")
	}
	_, err = ctx.JSON(controllers.SuccessResponse("success", art))
	return err
}
