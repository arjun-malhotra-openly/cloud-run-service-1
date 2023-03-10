# FROM golang:1.19-alpine

# WORKDIR /app
# COPY . .
# RUN go mod download
# RUN go build main.go

# CMD ./main


FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main

FROM alpine:latest
RUN apk add --no-cache libc6-compat 
WORKDIR /root/
COPY --from=builder /app/main .


CMD ["./main"]




# FROM golang:1.19-alpine as builder

# # Install git
# RUN apk update && apk upgrade && \
#     apk add --no-cache bash git openssh

# WORKDIR /app
# COPY . .
# # Install dependencies, recursively
# RUN go get -d -v ./...
# RUN go install -v ./...

# COPY / /app
# RUN go build -o appservice1

# FROM alpine:3.12

# RUN apk update && apk upgrade && \
#     apk add --no-cache openssh curl ca-certificates

# WORKDIR /app
# COPY --from=builder /app/appservice1 /app/appservice1
# EXPOSE 8080

# CMD ["./appservice1"]