FROM golang:1.19.2-bullseye AS build-env
WORKDIR /opt
RUN mkdir -p server2 forwarder
COPY go.mod go.sum ./
COPY server2/ server2/
COPY forwarder/ forwarder/
RUN go install ./server2 ./forwarder

FROM public.ecr.aws/amazonlinux/amazonlinux:2
RUN yum install iproute -y
WORKDIR /app
COPY --from=build-env /go/bin/server2 /go/bin/forwarder ./
COPY run2.sh ./
RUN chmod +x /app/run2.sh
CMD ["/app/run2.sh"]
