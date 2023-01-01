FROM golang:latest

WORKDIR /app

COPY . .
ENV GOPROXY=https://goproxy.io

RUN go mod download
RUN go build -o app .

EXPOSE 8081

CMD ["./app"]
