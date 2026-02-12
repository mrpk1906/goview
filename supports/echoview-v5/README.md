# EchoView v5

[![GoDoc Widget]][GoDoc]

goview support for echo template v5 version.

## Install
```bash

go get -u github.com/mrpk1906/goview

go get -u github.com/mrpk1906/goview/supports/echoview-v5

```

### Example

```go

package main

import (
	"github.com/mrpk1906/goview/supports/echoview-v5"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"net/http"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Set Renderer
	e.Renderer = echoview.Default()

	// Routes - Note: In Echo v5, handlers use *echo.Context instead of echo.Context
	e.GET("/", func(c *echo.Context) error {
		//render with master
		return c.Render(http.StatusOK, "index", echo.Map{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	e.GET("/page", func(c *echo.Context) error {
		//render only file, must full name with extension
		return c.Render(http.StatusOK, "page.html", echo.Map{"title": "Page file title!!"})
	})

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}

```

Project structure:
```go
|-- app/views/
    |--- index.html
    |--- page.html
    |-- layouts/
        |--- footer.html
        |--- master.html


See in "examples/basic" folder
```

[Echo example](https://github.com/mrpk1906/goview/tree/master/_examples/echo-v5)

## Key Changes from Echo v4 to v5

- **Context Type**: Handlers now receive `*echo.Context` (pointer) instead of `echo.Context` (interface)
- **Renderer Interface**: Signature changed to `Render(c *Context, w io.Writer, templateName string, data any) error`
- **Logger**: Now uses `*slog.Logger` from standard library instead of custom Logger interface
- **Response**: `c.Response()` returns `http.ResponseWriter` directly instead of `*Response`

## More examples

See [_examples/](https://github.com/mrpk1906/goview/blob/master/_examples/) for a variety of examples.

[GoDoc]: https://godoc.org/github.com/mrpk1906/goview/supports/echoview-v5
[GoDoc Widget]: https://godoc.org/github.com/mrpk1906/goview/supports/echoview-v5?status.svg

