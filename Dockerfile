FROM golang:1.19.2-bullseye AS build-env
WORKDIR /opt
COPY go.mod go.sum server/server.go ./
RUN go install ./server.go

FROM public.ecr.aws/amazonlinux/amazonlinux:2
RUN yum install iproute -y
WORKDIR /app
COPY --from=build-env /go/bin/server ./
COPY run.sh ./
RUN chmod +x /app/run.sh
CMD ["/app/run.sh"]
