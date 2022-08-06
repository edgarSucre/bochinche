FROM golang:1.18

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build


EXPOSE 8080
CMD ["./bochinche"]
