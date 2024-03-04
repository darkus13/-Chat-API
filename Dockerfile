FROM golang:1.20.3-alpine AS bilder

COPY . /github.com/darks13/-Chat-API/source
WORKDIR /github.com/darks13/-Chat-API/source

RUN go mod download
RUN go mod tidy -e
RUN go build -o ./bin/server app/cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/darkus13/-Chat-API/source/bin/server .

CMD ["./server"]
