FROM golang:1.20

ADD . /solr_worker

WORKDIR /solr_worker
RUN go mod init github.com/aaraya0/final-asw2/services/solr_worker
RUN go mod tidy
RUN go build -o solr_worker .
RUN chmod +x /solr_worker

ENTRYPOINT ["/solr_worker/solr_worker"]