FROM golang:1.6.2

WORKDIR /go/src/github.com/shogonir/zzz

COPY . .

RUN go install github.com/shogonir/zzz

EXPOSE 7999

ENTRYPOINT ["zzz"]

