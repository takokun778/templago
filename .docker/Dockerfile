ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION}-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
