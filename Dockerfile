FROM golang:1.22.2-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./...
RUN go build -o fish .

CMD ["./fish"]
