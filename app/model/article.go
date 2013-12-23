package model

import (
	"github.com/fuxiaohei/GoBlog/app"
	"fmt"
)

type Article struct {
	Id         int
	Title      string
	Slug       string
	Summary    string
	Content    string
	CreateTime int64
	EditTime   int64
	CategoryId int
	AuthorId   int
	Format     string
	Status     string
	IsComment  bool
	IsFeed     bool
	Comments   int
	Views      int
}

func (this *Article) Author() *User {
	return UserM.GetUserById(this.AuthorId)
}

func (this *Article) Category() *Category {
	return CategoryM.GetCategoryById(this.CategoryId)
}

type ArticleModel struct {
	article map[string]*Article
	idIndex map[int]string
	pagedCache map[string][]*Article
}

func (this *ArticleModel) cacheArticle(a *Article) {
	if a == nil {
		return
	}
	this.article[a.Slug] = a
	this.idIndex[a.Id] = a.Slug
}

func (this *ArticleModel) nocacheArticle(a *Article) {
	if a == nil {
		return
	}
	delete(this.article, a.Slug)
	delete(this.idIndex, a.Id)
}

func (this *ArticleModel) GetArticleBySlug(slug string) *Article {
	a := this.article["slug"]
	if a == nil {
		sql := "SELECT * FROM blog_content WHERE type = ? AND slug = ?"
		res, _ := app.Db.Query(sql, "article", slug)
		a = new(Article)
		res.One(a)
		if a.Slug != slug {
			return nil
		}
		this.cacheArticle(a)
	}
	return a
}

func (this *ArticleModel) GetArticleById(id int) *Article {
	slug := this.idIndex[id]
	if slug != "" {
		return this.GetArticleBySlug(slug)
	}
	sql := "SELECT * FROM blog_content WHERE type = ? AND id = ?"
	res, _ := app.Db.Query(sql, "article", id)
	a := new(Article)
	res.One(a)
	if a.Id != id {
		return nil
	}
	this.cacheArticle(a)
	return a
}

func (this *ArticleModel) nocachePagedAll() {
	this.pagedCache = make(map[string][]*Article)
}

func (this *ArticleModel) GetPaged(page, size int, noDraft bool) []*Article {
	page = (page-1) * size
	key := fmt.Sprintf("%d-%d-draft-%t", page, size, noDraft)
	if this.pagedCache[key] == nil {
		sql := "SELECT * FROM blog_content WHERE type = ?"
		args := []interface {}{"article"};
		if noDraft {
			sql +=" AND status != ?"
			args = append(args, "draft")
		}
		sql +=" ORDER BY id DESC LIMIT " + fmt.Sprintf("%d,%d", page, size)
		res, _ := app.Db.Query(sql, args...)
		if len(res.Data) < 1 {
			return nil
		}
		articles := make([]*Article, 0)
		res.All(&articles)
		this.pagedCache[key] = articles
	}
	return this.pagedCache[key]
}

func (this *ArticleModel) GetCategoryPaged(categoryId, page, size int, noDraft bool) []*Article {
	page = (page-1) * size
	key := fmt.Sprintf("%d-%d-draft-%t-category-%d", page, size, noDraft, categoryId)
	if this.pagedCache[key] == nil {
		sql := "SELECT * FROM blog_content WHERE type = ? AND category_id = ?"
		args := []interface {}{"article", categoryId};
		if noDraft {
			sql +=" AND status != ?"
			args = append(args, "draft")
		}
		sql +=" ORDER BY id DESC LIMIT " + fmt.Sprintf("%d,%d", page, size)
		res, _ := app.Db.Query(sql, args...)
		if len(res.Data) < 1 {
			return nil
		}
		articles := make([]*Article, 0)
		res.All(&articles)
		this.pagedCache[key] = articles
	}
	return this.pagedCache[key]
}
