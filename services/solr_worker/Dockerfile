FROM golang:1.18

ADD . /solr_worker

WORKDIR /solr_worker
RUN go mod tidy
RUN go build -o solr_worker .
RUN chmod +x /solr_worker

ENTRYPOINT ["/solr_worker/solr_worker"]