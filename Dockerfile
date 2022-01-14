FROM golang:alpine

MAINTAINER Raven Labs


ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src/checkout

COPY . /go/src/checkout

# Run the two commands below to install git and dependencies for the project. 
RUN go get .

#COPY dependencies /go/src #if you dont want to pull dependencies from git 

RUN go build .

EXPOSE $PORT

ENTRYPOINT ["./checkout"]
