FROM golang:1.19-bullseys AS builder

ENV CGO_ENABLED=0
ADD . /dist
WORKDIR /dist
RUN go get -v all
RUN go build -o main .

FROM ubuntu:22.04

COPY --from=builder /dist/main /
ENV LANG=C.UTF-8
ENTRYPOINT ["/main"]
