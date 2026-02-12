/*
 * Copyright 2018 Foolin.  All rights reserved.
 *
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/mrpk1906/goview/supports/echoview-v5"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	// Note: Logger middleware was removed in v5, use RequestLogger or other alternatives
	e.Use(middleware.Recover())

	//Set Renderer
	e.Renderer = echoview.Default()

	// Routes - Note: In Echo v5, handlers use *echo.Context instead of echo.Context
	e.GET("/", func(c *echo.Context) error {
		//render with master
		// Note: echo.Map was removed in v5, use map[string]any instead
		return c.Render(http.StatusOK, "index", map[string]any{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	e.GET("/page", func(c *echo.Context) error {
		//render only file, must full name with extension
		return c.Render(http.StatusOK, "page.html", map[string]any{"title": "Page file title!!"})
	})

	// Start server
	// Note: e.Logger is now *slog.Logger and doesn't have Fatal method
	if err := e.Start(":9090"); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
