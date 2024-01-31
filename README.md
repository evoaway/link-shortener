# Link shortener
## About
Rest API service for shortening links on Go using [Chi](https://github.com/go-chi/chi) router, `MongoDB` to store data and `ngninx` proxy server. There are two endpoints:

| Method | Endpoint      | Request body      | Response                                           |
|--------|---------------|-------------------|----------------------------------------------------|
| POST   | `/api`        | `{ link: 'url'} ` | ` { short: 'short_url' original: 'original_url' }` |
| GET    | `/api/{link}` |                   | ` { short: 'short_url' original: 'original_url' }` |


## Usage

First, make sure that the environment variables in `.env` file are set correctly (I use the [syntax described for Docker Compose](https://docs.docker.com/compose/environment-variables/env-file/)).

Then create and start containers in detach using:

```bash
docker compose up -d
```
Now the service is available at http://localhost/api

To run server separately use:
```bash
docker build -t link-shortener .
docker run -p 3000:3000 --env-file .env link-shortener
```