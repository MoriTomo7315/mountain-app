FROM golang:1.14

WORKDIR /api
COPY . .
ENV GO114MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]