FROM golang:1.22.3

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

EXPOSE 3000

CMD ["air"]
