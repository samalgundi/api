FROM golang:alpine
WORKDIR /app
COPY ./ ./

RUN go build -o /app/main

# Copy the exe into a smaller base image
FROM alpine
WORKDIR /app
COPY --from=0 /app/main /app/main
COPY --from=0 /app/config /app/config
COPY --from=0 /app/lib/cos/config /app/lib/cos/config


EXPOSE 8000
CMD /app/main
