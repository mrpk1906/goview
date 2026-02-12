# Echo v5 Multiple Template Engines Example

This example demonstrates using multiple template engines with different view directories in Echo v5.

## Key Changes from Echo v4 to v5

- **Context Type**: Handlers now receive `*echo.Context` (pointer) instead of `echo.Context` (interface)
- **echo.Map**: Removed - use `map[string]any` instead
- **Logger Middleware**: Removed in v5
- **Server Start**: Changed error handling pattern

## Features

- Frontend templates in `views/frontend/`
- Backend admin templates in `views/backend/`
- Different layouts and partials per section
- Middleware-based template engine selection

## Run

```bash
go run main.go
```

Then visit:
- Frontend: `http://localhost:9090/`
- Backend: `http://localhost:9090/admin/`

## Project Structure

```
views/
├── frontend/
│   ├── index.html
│   ├── layouts/
│   └── partials/
└── backend/
    ├── index.html
    └── layouts/
```
