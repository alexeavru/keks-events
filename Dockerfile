FROM golang:alpine3.19
ENV PORT=8080

RUN apk add --no-cache alpine-conf curl && \
    setup-timezone -z Europe/Moscow

COPY keksevents .env ./
CMD ["./keksevents"]