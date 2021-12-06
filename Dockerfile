FROM golang:latest

WORKDIR /js_statistic

ADD . /js_statistic

COPY . .

RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o js_statistic /js_statistic

EXPOSE 9090

CMD ["js_statistics"]
