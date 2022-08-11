# Criação aplicação GO

1 - Subir um servidor

Arquivo `cmd/web/main.go`

Conteúdo:
```go
package main

import "net/http"

const portNumber = ":8080"

func main() {
	http.ListenAndServe(portNumber, nil)
}

```

2 - Implementar rotas

Arquivo `internal/render/render.go`

Conteúdo:
```go
package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
```


Arquivo `templates/home.page.tmpl`

Conteúdo:
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>Home</h1>
</body>
</html>
```


Arquivo `internal/handlers/handlers.go`

Conteúdo:
```go
package handlers

import (
	"net/http"
	"xcasluw/golang-from-scratch/internal/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl")
}
```


Arquivo `cmd/web/routes.go`

```go
package main

import (
	"net/http"
	"xcasluw/golang-from-scratch/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	return mux
}
```
