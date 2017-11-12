FROM alpine:3.5

RUN apk add --update bash curl && rm -rf /var/cache/apk/*

ADD Database-linux /bin/Database-linux

EXPOSE 80

CMD ["/bin/Database-linux"]
