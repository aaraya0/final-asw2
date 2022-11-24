FROM golang:1.18

ADD . /search

WORKDIR /search
RUN go mod init worker
RUN go mod tidy
RUN go build -o worker .
RUN chmod +x /search

ENTRYPOINT ["/search/worker"]