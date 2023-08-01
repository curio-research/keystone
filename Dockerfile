FROM golang:latest

RUN mkdir /build
WORKDIR /build

COPY . .
RUN go build -o ./main .

EXPOSE 9000 9001

ENTRYPOINT ["/build/main"]