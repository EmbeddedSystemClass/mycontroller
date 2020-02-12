FROM alpine:3.11

LABEL maintainer="Jeeva Kandasamy <jkandasa@gmail.com>"

EXPOSE 8080

COPY mcserver-linux /app/mcserver-linux

ENTRYPOINT ["/app/mcserver-linux"]