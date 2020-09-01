FROM golang:1.15

WORKDIR /app
COPY . .
RUN go build -o ./bin/quickkv
WORKDIR /app/bin
RUN chmod +x quickkv

CMD ./quickkv