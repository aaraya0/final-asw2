FROM golang:1.18

ADD . /search

WORKDIR /search
RUN go mod init github.com/aaraya0/final-asw2/services/search
RUN go mod tidy
RUN go build -o search .
RUN chmod +x /search

ENTRYPOINT ["/search/search"]