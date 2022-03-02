FROM golang:1.17.0 as build

ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/akhilerm/eywa/
COPY . .

RUN go build

FROM alpine:3.15

COPY --from=build /go/src/github.com/akhilerm/eywa/eywa /usr/bin/eywa

CMD /usr/bin/eywa