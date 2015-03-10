FROM gliderlabs/alpine:latest
MAINTAINER Patrick M. Slattery <pslattery@mywebgrocer.com>

RUN apk-install go

# Add files from local storage
ADD ./config /

EXPOSE 8080
CMD ["/hello_world"]
