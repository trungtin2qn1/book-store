FROM golang:1.11-alpine as builder
WORKDIR /go/src/book-store
COPY . .
RUN apk add --update git make
RUN go get -u github.com/Masterminds/glide
RUN glide install
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/book-store/book-store ./book-store
CMD ["/book-store"]
EXPOSE 2013 