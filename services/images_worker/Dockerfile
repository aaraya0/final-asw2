FROM golang:1.20

ADD . /images_worker

WORKDIR /images_worker
RUN go mod init github.com/aaraya0/final-asw2/services/images_worker
RUN go mod tidy
RUN go build -o images_worker .
RUN chmod +x /images_worker

ENTRYPOINT ["/images_worker/images_worker"]