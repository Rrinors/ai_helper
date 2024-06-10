FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN chmod +x build.sh \
    && ./build.sh

EXPOSE 8888

CMD ["./output/bootstrap.sh"]