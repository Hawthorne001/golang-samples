FROM golang:1.22.7-alpine

WORKDIR /go/src
COPY ./go.mod ./
COPY ./go.sum ./
RUN mkdir ./hotapp/
COPY ./hotapp/*.go ./hotapp/

RUN go install ./...

CMD ["hotapp"]
