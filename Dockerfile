#
# Testing & Building
#
FROM golang:1.18 as builder
WORKDIR /go/src/app
COPY . .
RUN go mod download \
    && go test ./... \
    && CGO_ENABLED=0 GOOS=linux go build cmd/main.go

#
# Final Stage
#
FROM alpine:3.10
RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Bangkok
WORKDIR /src
COPY --from=builder \
    /go/src/app ./
CMD ["./main"]