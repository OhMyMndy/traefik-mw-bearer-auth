FROM golang:1.24.0-bookworm AS builder

ENV CGO_ENABLED=0
ADD . /dist
WORKDIR /dist
RUN go get -v all
RUN go build -o main .

FROM alpine

COPY --from=builder /dist/main /
ENV LANG=C.UTF-8
ENTRYPOINT ["/main"]
