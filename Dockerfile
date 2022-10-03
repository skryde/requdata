FROM golang:1.19-alpine AS builder

COPY . /app/src
WORKDIR /app/src

RUN go get -v -d

RUN go build -o out/requdata


FROM alpine:3.16

ENV TZ America/Montevideo
ARG PORT 8080

RUN adduser -u 1000 -g 1000 --disabled-password requdata

COPY --from=builder /app/src/out/requdata /app/

WORKDIR /app

RUN chown -R requdata:requdata .

USER requdata

ENTRYPOINT ["/app/requdata"]
