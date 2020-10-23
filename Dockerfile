# test & build app
FROM golang AS build-env

ADD . /app

WORKDIR /app

RUN go build -o server ./cmd/server

# safe image
FROM muninn/debian

COPY --from=build-env /app/server /usr/bin/server

EXPOSE 8080

CMD ["server"]