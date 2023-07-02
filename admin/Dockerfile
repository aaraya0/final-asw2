FROM golang:1.18


ADD . /admin

WORKDIR /admin
RUN go mod init admin
RUN go mod tidy
RUN go build -o admin .
RUN chmod +x /admin

ENTRYPOINT ["/admin/admin"]