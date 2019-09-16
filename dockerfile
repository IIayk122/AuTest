#######################################

FROM golang:alpine as  builder
ADD . /go/src/github.com/IIayk122/AuTest
RUN go get -d -v /go/src/github.com/IIayk122/AuTest
RUN go install /go/src/github.com/IIayk122/AuTest && cp -av /go/src/github.com/IIayk122/AuTest/templates /go/bin/templates


FROM alpine
WORKDIR /go/bin
COPY --from=builder /go/bin/AuTest .
COPY --from=builder /go/bin/templates /go/bin/templates 

CMD [ "./AuTest" ]

#######################################
#docker build -t autest .
#docker run -ti -p 8080:8080 autest
#docker load -i autest.tar 
