FROM golang:1.12.0-alpine3.9 as builder
WORKDIR /app
ADD . .
RUN apk add --update git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w -extldflags "-static"' -o /go/bin/app

FROM scratch
WORKDIR /app/
USER 1000
COPY --from=builder /go/bin/app .
CMD ["./app"]
