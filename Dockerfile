FROM golang:alpine as build
RUN apk add --no-cache --update git
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get ./...
RUN go build -o fdrv main.go

FROM alpine
COPY --from=build /go/src/app/fdrv /usr/bin/fdrv

EXPOSE 1510
#root folder
VOLUME /var/lib/fedorov

ENTRYPOINT ["/usr/bin/fdrv"]
CMD ["serve","-port", "1510", "-stderr"]
