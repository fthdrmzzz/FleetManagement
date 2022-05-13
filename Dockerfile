FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/DevelopmentHiring/FatihDurmaz/
WORKDIR /go/src/github.com/DevelopmentHiring/FatihDurmaz
RUN go mod download
COPY . /go/src/github.com/DevelopmentHiring/FatihDurmaz
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/FatihDurmaz github.com/DevelopmentHiring/FatihDurmaz

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/DevelopmentHiring/FatihDurmaz /usr/bin/FatihDurmaz
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/FatihDurmaz"]