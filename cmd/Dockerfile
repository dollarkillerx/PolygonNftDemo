FROM alpine:latest
WORKDIR /app
COPY ./cmd/main /app
COPY ./configs /app/configs
COPY ./polygon_demo /app/polygon_demo
CMD ["/app/main"]