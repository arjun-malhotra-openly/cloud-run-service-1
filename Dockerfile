FROM golang:1.19-alpine

WORKDIR /app
COPY . .
RUN go mod download
RUN go build main.go
CMD ./main
