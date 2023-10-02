FROM golang:1.21.1-alpine3.18
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY api/* datastore/* public/* ./
COPY . ./
COPY .env ./
COPY config.json ./
RUN go mod download 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/jourism_ai -ldflags="-s -w" main.go
EXPOSE 8080
CMD ["./bin/jourism_ai"]
