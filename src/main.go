package main

import (
	// "fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	// "html/template"
	// "log"
	// "net/http"
)

func index(r render.Render) {
	r.HTML(200, "index", "") // index is the template's name
}

func about(r render.Render) {
	r.HTML(200, "about", "关于") // about is the template's name
}

func faqs(r render.Render) {
	r.HTML(200, "faqs", "常见问题") // faqs is the template's name
}

func blog(r render.Render) {
	r.HTML(200, "blog", "团队博客") // faqs is the template's name
}

func contact(r render.Render) {
	r.HTML(200, "contact", "联系我们") // hello is the template's name
}

func jobs(r render.Render) {
	r.HTML(200, "jobs", "加入我们") // hello is the template's name
}

func jobdetail(r render.Render) {
	// log.Println("hello world ------------------ ")
	r.HTML(200, "jobdetail", "") // hello is the template's name
}

func ico(r render.Render) {
	r.HTML(200, "ico", "ico")
}

func main() {
	m := martini.Classic()
	// m.Get("/", func() string {
	// 	return "Hello world!"
	// })

	m.Use(martini.Static("assets"))
	m.Use(martini.Static("version"))
	m.Use(martini.Logger())
	// render html templates from templates directory
	m.Use(render.Renderer(render.Options{
		Layout: "layout", // Specify a layout template. Layouts can call {{ yield }} to render the current template.
	}))

	m.Get("/", index)
	m.Get("/index", index)
	m.Get("/ico", ico)
	// m.Get("/about", about)
	// m.Get("/faqs", faqs)
	// m.Get("/blog", blog)
	// m.Get("/contact", contact)
	// m.Get("/jobs", jobs)
	// m.Get("/jobdetail", jobdetail)

	// m.Group("/demos/restaurant_picker", func(r martini.Router) {
	// 	r.Get("", restaurants)
	// 	r.Get("/choose_town", choose_town)
	// 	r.Get("/choose_restaurant", choose_restaurant)
	// 	r.Get("/restaurant", restaurant)
	// })

	// This will set the Content-Type header to "application/json; charset=UTF-8"
	m.Get("/api", func(r render.Render) {
		r.JSON(200, map[string]interface{}{"hello": "world"})
	})

	m.Run()
	// log.Fatal(http.ListenAndServe(":3000", m))
}
