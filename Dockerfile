FROM golang:1.9 as gobuild
WORKDIR /go/src/kylehodgetts.com/go-quiz
ENV CGO_ENABLED=0 GOOS=linux
COPY . .
RUN go build -a

FROM alpine:3.7
LABEL maintainer="kyle.hodgetts@icloud.com"
WORKDIR /
COPY --from=gobuild /go/src/kylehodgetts.com/go-quiz/go-quiz .
RUN mkdir data
COPY --from=gobuild /go/src/kylehodgetts.com/go-quiz/data/problems.csv /data
ENTRYPOINT ["/go-quiz"]
