package handler

import (
	"github.com/fuxiaohei/GoInk/Core"
	"github.com/fuxiaohei/GoBlog/app/model"
	"strconv"
	"github.com/fuxiaohei/GoBlog/app"
)

func AdminCategory(context *Core.Context) interface {} {
	context.Render("admin:admin/category.html", map[string]interface {}{
			"Title":"分类",
			"IsMeta":true,
			"IsCategory":true,
			"Categories":model.CategoryM.GetAll(),
		})
	return nil
}

func AdminCategoryEdit(context *Core.Context) interface {} {
	cid, _ := strconv.Atoi(context.Param(3))
	if cid < 1 {
		context.Redirect("/admin/meta/")
		return nil
	}
	context.Render("admin:admin/category_edit.html", map[string]interface {}{
			"Title":"分类",
			"IsMeta":true,
			"IsCategory":true,
			"Category":model.CategoryM.GetCategoryById(cid),
		})
	return nil
}

func AdminCategoryEditPost(context *Core.Context) interface {} {
	cid, _ := strconv.Atoi(context.Param(3))
	if cid < 1 {
		context.Redirect("/admin/meta/")
		return nil
	}
	data := context.Input()
	// check slug exist
	c := model.CategoryM.GetCategoryBySlug(data["slug"])
	if c.Id != cid {
		context.Redirect(context.Referer+"?err=1")
		return nil
	}
	model.CategoryM.SaveCategory(cid, data["name"], data["slug"], data["desc"])
	context.Redirect("/admin/meta/?update=1")
	app.Ink.Listener.EmitAll("model.category.update", cid)
	return nil
}

func AdminCategoryNew(context *Core.Context) interface {} {
	context.Render("admin:admin/category_new.html", map[string]interface {}{
			"Title":"分类",
			"IsMeta":true,
			"IsCategory":true,
		})
	return nil
}

func AdminCategoryNewPost(context *Core.Context) interface {} {
	data := context.Input()
	c := model.CategoryM.GetCategoryBySlug(data["slug"])
	if c.Id > 0 {
		context.Redirect(context.Referer+"?err=1")
		return nil
	}
	c = model.CategoryM.CreateCategory(data["name"], data["slug"], data["desc"])
	context.Redirect("/admin/meta/?new=1")
	app.Ink.Listener.EmitAll("model.category.new", c)
	return nil
}

func AdminCategoryDelete(context *Core.Context) interface {} {
	data := context.Input()
	if data["category"] == data["move"] {
		context.Redirect("/admin/meta/?err=1")
		return nil
	}
	category , _ := strconv.Atoi(data["category"])
	move, _ := strconv.Atoi(data["move"])
	model.CategoryM.DeleteCategory(category, move)
	context.Redirect("/admin/meta/?move=1")
	app.Ink.Listener.EmitAll("mode.category.move", category, move)
	return nil
}