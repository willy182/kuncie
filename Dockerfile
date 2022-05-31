# Stage 1
FROM golang:1.16.4-alpine3.13 AS dependency_builder

WORKDIR /go/src
ENV GO111MODULE=on

RUN apk update
RUN apk add --no-cache bash ca-certificates git make

COPY go.mod .
COPY go.sum .

RUN go mod download

# Stage 2
FROM dependency_builder AS service_builder

WORKDIR /usr/app

COPY . .
RUN make migration
RUN CGO_ENABLED=0 GOOS=linux go build -o bin

# Stage 3
FROM alpine:latest  

ARG BUILD_NUMBER
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /usr/app/
ENV BUILD_NUMBER=$BUILD_NUMBER

RUN mkdir -p /usr/app/api
COPY --from=service_builder /usr/app/bin bin
COPY --from=service_builder /usr/app/.env .env
COPY --from=service_builder /usr/app/api /usr/app/api

ENTRYPOINT ["./bin"]
