FROM golang:1.15 as builder

WORKDIR /src
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on

RUN apt-get update

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /tmp/cleaner

FROM scratch

ENV CLEANER_PODS
ENV CLEANER_NAMESPACE

ENTRYPOINT [ "/cleaner" ]
COPY --from=builder /tmp/cleaner /
