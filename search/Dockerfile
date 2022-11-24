FROM golang:1.18

ADD . /search

WORKDIR /search
RUN go mod init wesolr
RUN go mod tidy
RUN go build -o wesolr .
RUN chmod +x /search

ENTRYPOINT ["/search/solr"]