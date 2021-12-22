#####################################
#   STEP 1 build executable binary  #
#####################################
FROM golang:1.16.2-alpine AS builder

WORKDIR /opt/app/
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify
COPY . .

RUN go build -o /go/bin/ ./cmd/...

#####################################
#   STEP 2 build a small image      #
#####################################
FROM alpine:3.12.3

ARG SERVICE_NAME

COPY --from=builder /go/bin/${SERVICE_NAME} main

ENTRYPOINT ["./main"]