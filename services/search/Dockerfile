FROM golang:1.18

ADD . /search

WORKDIR /search
RUN go mod tidy
RUN go build -o solr .
RUN chmod +x /search

ENTRYPOINT ["/search/solr"]