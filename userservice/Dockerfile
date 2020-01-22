FROM golang:alpine
 
RUN apk update && apk add git
 
COPY . /go/src/app/
 
WORKDIR /go/src/app

 
RUN go get -d -v
RUN go install -v
 
CMD ["app"]