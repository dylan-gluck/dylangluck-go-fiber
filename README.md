# Dylan Navajas Gluck Personal Website

A modern personal website and blog built with Go Fiber, featuring server-side rendering with HTML templates, HTMX for dynamic interactions, and a custom blog system with Markdown support.

## Overview

This is a personal portfolio and blog website for Dylan Navajas Gluck, a full-stack engineer, designer, and AI researcher. The application features a clean, responsive design with dark/light theme support, a custom blog system with tag filtering, and optimized performance through caching and static file serving.

## Tech Stack

- **Backend**: Go 1.22+ with Fiber web framework
- **Frontend**: HTML templates with HTMX and Hyperscript
- **Styling**: PicoCSS framework with custom CSS modules
- **Content**: Markdown-based blog system with YAML frontmatter
- **Deployment**: Docker containerization with Fly.io configuration

## Features

- **Server-Side Rendered**: HTML templates with embedded Go template engine
- **Dynamic Navigation**: HTMX-powered SPA-like navigation without full page reloads
- **Blog System**: File-based blog with Markdown content and YAML metadata
- **Tag Filtering**: Browse blog posts by tags
- **Theme Support**: Automatic dark/light theme detection and switching
- **Static Asset Optimization**: Cached static file serving with custom fonts
- **Responsive Design**: Mobile-first responsive layout
- **SEO Optimized**: Meta tags, Open Graph, and structured navigation

## Project Structure

```
dylangluck-com/
├── web.go                  # Main application entry point
├── go.mod                  # Go module dependencies
├── devenv.nix             # Dev shell, `devenv up` process, deploy script
├── devenv.yaml / devenv.lock # devenv inputs (pinned)
├── .envrc                 # direnv: loads the devenv shell
├── .air.toml              # Air hot reload configuration
├── Dockerfile              # Docker build configuration
├── fly.toml               # Fly.io deployment configuration
├── blog/                  # Blog system package
│   ├── blog.go           # Blog collection and post management
│   └── articles/         # Markdown blog posts
│       ├── example-article/
│       │   ├── meta.yml  # Post metadata (title, date, tags)
│       │   └── post.md   # Markdown content
│       └── example-article-2/
├── handlers/              # HTTP request handlers
│   └── handlers.go       # Route handlers for pages and blog
├── util/                  # Utility functions
│   └── util.go           # Markdown to HTML conversion
├── views/                 # HTML templates
│   ├── layouts/
│   │   └── main.html     # Main page layout template
│   ├── partials/         # Reusable template components
│   │   ├── header.html   # Site header with navigation
│   │   ├── footer.html   # Site footer
│   │   ├── nav-items.html # Navigation menu items
│   │   ├── theme-toggle.html # Dark/light theme toggle
│   │   └── home/         # Homepage specific partials
│   ├── index.html        # Homepage template
│   ├── blog.html         # Blog listing template
│   └── post.html         # Individual blog post template
└── public/               # Static assets
    ├── styles/           # CSS stylesheets
    │   ├── main.css      # Main stylesheet with imports
    │   ├── global/       # Global CSS modules
    │   └── pages/        # Page-specific styles
    ├── fonts/            # Custom font files
    ├── *.png, *.ico      # Favicon and app icons
    ├── 404.html          # Custom 404 error page
    ├── robots.txt        # Search engine directives
    └── site.webmanifest  # Web app manifest
```

## Installation and Setup

### Prerequisites

- [devenv](https://devenv.sh) (provides Go, gopls, Air, and flyctl)
- [direnv](https://direnv.net) (optional, auto-loads the shell on `cd`)
- Docker (optional, for local container builds)

### Local Development

1. **Enter the dev shell**:
   ```bash
   direnv allow   # first time only; afterwards the shell loads on cd
   # or, without direnv:
   devenv shell
   ```

2. **Start the dev server with hot reload**:
   ```bash
   devenv up
   ```

   Air (`.air.toml`) rebuilds and restarts the server whenever Go, template, blog content, or static asset files change. Views are embedded at build time and responses are cached in memory, so a restart is required for any change to show up.

   Or run once without hot reload:
   ```bash
   go run .
   ```

3. **Access the application**:
   Open your browser to `http://localhost:3000` (devenv sets `PORT=3000`).

### Docker Deployment

1. **Build the Docker image**:
   ```bash
   docker build -t dylangluck-go-fiber .
   ```

2. **Run the container**:
   ```bash
   docker run -p 8080:8080 dylangluck-go-fiber
   ```

### Fly.io Deployment

The application is configured for deployment on Fly.io. `flyctl` is provided by the dev shell.

#### Initial Setup

```bash
# Authenticate with Fly.io
fly auth login
```

#### Deploy

```bash
# Compiles locally to fail fast, then runs `flyctl deploy` (remote Docker build).
# Extra arguments are passed through to flyctl, e.g. `deploy --local-only`.
deploy
```

#### Essential Commands

```bash
# Deploy changes
fly deploy

# Check deployment status
fly status

# View application logs
fly logs
# Stream logs in real-time
fly logs --tail

# Manage secrets (environment variables)
fly secrets set DATABASE_URL=postgres://...
fly secrets list
fly secrets unset KEY_NAME

# Scale application
fly scale count 2          # Scale to 2 machines
fly scale memory 1gb       # Set memory per machine
fly scale vm shared-cpu-1x # Change VM size

# Multi-region deployment
fly regions add lax sea    # Add US West regions
fly regions list          # View all regions
fly regions set lax       # Set primary region

# Debugging and maintenance
fly ssh console           # SSH into running container
fly machine list         # View all machines
fly machine restart <id> # Restart specific machine
fly apps restart         # Restart all machines

# Monitor resources
fly status --all         # Detailed status
fly checks list         # Health check status
```

#### Production Tips

- The app is configured to listen on port 8080 (Fly.io default)
- Use `fly secrets` for sensitive configuration instead of environment variables
- Monitor logs with `fly logs --tail` during deployment
- Use `fly ssh console` to debug issues in production
- Configure health checks in fly.toml for better reliability

## API Endpoints

- `GET /` - Homepage with recent blog posts and profile information
- `GET /blog` - Blog listing page with all posts (paginated)
- `GET /blog/:name` - Individual blog post page
- `GET /blog/tag/:tag` - Blog posts filtered by specific tag
- `GET /static/*` - Static asset serving (CSS, fonts, images)

## Blog System

### Adding New Blog Posts

1. **Create a new directory** in `blog/articles/` with your post slug:
   ```bash
   mkdir blog/articles/my-new-post
   ```

2. **Create metadata file** (`meta.yml`):
   ```yaml
   title: Your Post Title
   short: Brief description for previews
   date: 2024-01-01
   tags:
     - tag1
     - tag2
     - tag3
   ```

3. **Create content file** (`post.md`):
   ```markdown
   # Your Post Title

   Your markdown content here...
   ```

### Blog Post Features

- **Markdown Support**: Full CommonMark syntax with extensions
- **Auto-Generated IDs**: Automatic heading ID generation for navigation
- **External Link Handling**: External links open in new tabs
- **Tag System**: Categorize posts with multiple tags
- **Date Sorting**: Posts automatically sorted by date (newest first)
- **Pagination Ready**: Backend supports pagination (frontend TODO)

## Configuration

### Environment Variables

The application uses default configurations but can be customized:

- **Port**: `PORT` environment variable, defaults to `8080` (devenv sets `3000` for local development)
- **Static Cache**: 1 hour cache for static assets
- **Template Caching**: Automatic template reloading in development

### Performance Features

- **Middleware Stack**:
  - Cache middleware for response caching
  - Recovery middleware for error handling
  - Static file serving with cache headers
- **Template Optimization**: Embedded file system for production builds
- **Asset Optimization**: Compressed fonts and optimized images

## Dependencies

### Core Dependencies

- `github.com/gofiber/fiber/v2` - Fast HTTP web framework
- `github.com/gofiber/template/html/v2` - HTML template engine
- `github.com/gomarkdown/markdown` - Markdown to HTML conversion
- `github.com/goccy/go-yaml` - YAML parsing for blog metadata

### Frontend Dependencies (CDN)

- PicoCSS - Modern CSS framework
- HTMX - Dynamic HTML interactions
- Hyperscript - Client-side scripting

## Architecture Decisions

### File-Based Blog System
The blog uses a file-based approach rather than a database for simplicity and portability. Each post consists of a directory containing `meta.yml` (metadata) and `post.md` (content).

### Server-Side Rendering
All pages are server-side rendered for optimal SEO and performance, with HTMX providing enhanced interactivity without full page reloads.

### Embedded Assets
Static assets and templates are embedded in the binary for simplified deployment and distribution.

## Current Status

### Completed Features
- Homepage structure and styling
- Blog system with post collection and management
- Individual post rendering with Markdown support
- Tag-based filtering system
- Responsive design with theme support
- Docker containerization
- Fly.io deployment configuration

### Planned Features
- Pagination for blog listing
- Search functionality across blog posts
- RSS feed generation
- XML sitemap generation
- Additional blog content

## Contributing

When adding new features or making changes:

1. Ensure all Go code follows standard formatting (`go fmt`)
2. Update documentation for any API changes
3. Test locally before deploying
4. Update this README if project structure changes

## License

Personal project - All rights reserved to Dylan Navajas Gluck.
