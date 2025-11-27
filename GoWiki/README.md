# Go Wiki Application

A simple wiki web application written in Go that allows users to view, edit, and save pages.

## Features

- View wiki pages
- Edit page content
- Save changes automatically
- Simple file-based storage

## Installation

1. Ensure you have Go installed
2. Clone this repository
3. Run the application:

```go
go run main.go
```

# Usage
Access the application at http://localhost:8080

Navigate to /view/[pagename] to view a page

Use /edit/[pagename] to edit pages

Changes are automatically saved as text files

# Template Files
The application requires two HTML templates:

edit.html - Edit page template

view.html - View page template