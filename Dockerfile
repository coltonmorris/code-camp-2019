# -------- GO STEP -----------       
FROM golang:latest

WORKDIR $GOPATH/src/github.com/coltonmorris/code-camp-2019

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN CGO_ENABLED=0 go build -o main .

# -------- JS STEP -----------       
FROM node:12.2.0-alpine

WORKDIR /frontend

# add `node_modules/.bin` to $PATH
ENV PATH /frontend/node_modules/.bin:$PATH

# copy over js files from first step
COPY --from=0 /go/src/github.com/coltonmorris/code-camp-2019/package.json .
COPY --from=0 /go/src/github.com/coltonmorris/code-camp-2019/public public/
COPY --from=0 /go/src/github.com/coltonmorris/code-camp-2019/src src/

RUN npm install --silent
RUN npm install react-scripts@3.2.0 -g --silent

# build production bundle
RUN ["npm", "run", "build"]

# -------- FINAL RUN STEP -----------       

FROM alpine:latest

# RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=0 /go/src/github.com/coltonmorris/code-camp-2019/main .
COPY --from=1 /frontend/build build

# heroku creates this env var automagically
EXPOSE $PORT

CMD ["./main"]
# RUN ls build
# CMD ["ls"]
