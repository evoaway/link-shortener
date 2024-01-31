FROM golang:alpine AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /link-shortener

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=build-stage /link-shortener .
EXPOSE 3000
ENTRYPOINT ["./link-shortener"]
