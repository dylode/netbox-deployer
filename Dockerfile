FROM golang:1.21

WORKDIR /app

RUN adduser --uid 1001 app
USER app

COPY --chown=app:app . /app