FROM alpine:latest

LABEL Author="Dat Nguyen <it.tandat@gmail.com>"

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

COPY dist/stably_test /bin/
RUN chmod +x /bin/stably_test
RUN mkdir -p /stably_test
WORKDIR /stably_test

VOLUME /stably_test
EXPOSE 3100

CMD [ "stably_test", "" ]
