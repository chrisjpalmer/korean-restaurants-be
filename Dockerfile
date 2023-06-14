FROM golang:alpine

WORKDIR ["/app"]

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/serve ./cmd/serve 

ENTRYPOINT ["/app/serve"]