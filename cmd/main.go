package main

import (
	"html/template"
	"log"

	"green-api-test-assignment/internal/handlers"

	"github.com/valyala/fasthttp"
)

var tmpl *template.Template

func init() {
	var err error
	tmpl = template.New("main")
	tmpl, err = tmpl.ParseGlob("web/templates/partials/*.html")
	if err != nil {
		log.Fatalf("Ошибка загрузки partials: %v", err)
	}
	tmpl, err = tmpl.ParseGlob("web/templates/index.html")
	if err != nil {
		log.Fatalf("Ошибка загрузки index: %v", err)
	}
	log.Println("Загруженные шаблоны:")
}

func main() {
	router := func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		method := string(ctx.Method())

		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type")

		if method == "OPTIONS" {
			ctx.SetStatusCode(fasthttp.StatusNoContent)
			return
		}

		if method == "POST" {
			switch path {
			case "/api/state":
				handlers.GetStateInstanceHandler(ctx)
				return
			case "/api/settings":
				handlers.GetSettingsHandler(ctx)
				return
			case "/api/message":
				handlers.SendMessageHandler(ctx)
				return
			case "/api/file":
				handlers.SendFileByUrlHandler(ctx)
				return
			}
		}

		if path == "/" {
			ctx.SetContentType("text/html; charset=utf-8")
			data := map[string]interface{}{
				"Title": "Green API Test",
			}
			if err := tmpl.ExecuteTemplate(ctx, "index", data); err != nil {
				ctx.SetStatusCode(fasthttp.StatusInternalServerError)
				ctx.WriteString("Ошибка рендеринга: " + err.Error())
				log.Printf("Template error: %v", err)
			}
			return
		}
		if len(path) > 8 && path[:8] == "/static/" {
			fasthttp.ServeFile(ctx, "./web"+path)
			return
		}
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	}
	log.Println("Server on http://localhost:8080")
	log.Fatal(fasthttp.ListenAndServe(":8080", router))
}
