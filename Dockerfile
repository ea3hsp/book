FROM debian:stable-slim
RUN mkdir -p /app/
COPY ./bin/book /app
WORKDIR /app
ENTRYPOINT [ "./book"]