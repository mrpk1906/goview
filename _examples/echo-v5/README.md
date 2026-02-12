# Echo v5 Example

This example shows how to integrate goview with Echo v5.

## Key Changes from Echo v4 to v5

- **Context Type**: Handlers now receive `*echo.Context` (pointer) instead of `echo.Context` (interface)
- **Renderer Interface**: Signature changed to `Render(c *Context, w io.Writer, templateName string, data any) error`
- **Logger Middleware**: Removed - use RequestLogger or custom logging
- **echo.Map**: Removed - use `map[string]any` instead
- **Server Start**: `e.Logger.Fatal()` pattern changed - use standard `log.Fatal()` or handle errors directly

## Run

```bash
go run main.go
```

Then visit `http://localhost:9090` in your browser.

## Project Structure

```
views/
├── index.html
├── page.html
└── layouts/
    ├── footer.html
    └── master.html
```
