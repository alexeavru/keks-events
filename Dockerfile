FROM golang:alpine3.19

RUN apk add --no-cache alpine-conf && \
    setup-timezone -z Europe/Moscow

COPY keksevents /
CMD ["/keksevents"]