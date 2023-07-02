FROM golang:1.18


ADD . /items

WORKDIR /items
RUN go mod init github.com/aaraya0/final-asw2/services/items
RUN go mod tidy
RUN go build -o items .
RUN chmod +x /items

ENTRYPOINT ["/items/items"]