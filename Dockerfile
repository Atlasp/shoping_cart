FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    max_basket_size=100 \
    discount_threshold=20 \
    discount=1 \
    free_item_threshold=5

WORKDIR /build

RUN apk update
RUN	apk add git

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

EXPOSE 4141

CMD ["/dist/main"]