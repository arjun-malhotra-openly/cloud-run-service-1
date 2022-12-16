FROM golang:1.15-alpine3.14

WORKDIR /app
COPY . .
RUN go mod download
RUN go build main.go
CMD ./main
