FROM golang:1.17

ADD . /users

WORKDIR /users
RUN go mod init github.com/aaraya0/final-asw2/services/users
RUN go mod tidy
RUN go build -o users .
RUN chmod +x /users

ENTRYPOINT ["/users/users"]