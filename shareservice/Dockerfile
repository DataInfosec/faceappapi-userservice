FROM golang:alpine
 
RUN apk update && apk add git
 
COPY . /go/src/app/
 
WORKDIR /go/src/app

 
RUN go get -d -v
RUN go install -v
 
CMD ["app"]


# FROM golang:latest 
# RUN mkdir /app 
# ADD . /app/ 
# WORKDIR /app 
# RUN go build -o main . 
# CMD ["/app/main"]