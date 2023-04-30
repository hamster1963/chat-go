FROM golang:1.20-buster AS builder

ARG VERSION=dev

WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go build -o service -ldflags=-X=main.version=${VERSION} main.go

FROM loads/alpine:3.8
LABEL maintainer="Hamster <liaolaixin@gmail.com>"
COPY --from=builder /go/src/app/service /go/bin/
RUN chmod +x /go/bin/service
ENV PATH="/go/bin"
CMD ["service"]





