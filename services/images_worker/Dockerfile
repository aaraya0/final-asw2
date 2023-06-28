FROM golang:1.18

ADD . /images_worker

WORKDIR /images_worker
RUN go mod tidy
RUN go build -o images_worker .
RUN chmod +x /images_worker

ENTRYPOINT ["/images_worker/images_worker"]