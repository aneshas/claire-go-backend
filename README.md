# Claire entry project - Go backend

This repo contains a simple golang json api server written from scratch in Go.
It uses a simple repository pattern whcich makes use of go interfaces, mysql database through go database/sql package.

# Api

The app containse three entities as follows: **make** (Car make), **model** (car model) and **tag** (car model tags)

## Currently provided api endpoinst

### Make
**GET /api/make** - *list all makes*

**GET /api/make/[id]** - *view specific car make with id of [id]*

### Model
**GET /api/model** - *list all car models*

**GET /api/model/[id]** - *view specific car model info with id of [id]*

### Tag
**GET /api/tag** - *list all tags*

**GET /api/tag/[id]** - *view specific tag with id of [id]*

# Response

Response format:
```
{
  "head": {
    "status": int_status,
    "message: "status message"
  },
  "body": {
    resource
  }
}
```
