FROM golang:latest

WORKDIR $GOPATH/src/code-camp-2019

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

# RUN go build -o main .

EXPOSE 80

CMD ["ls"]
