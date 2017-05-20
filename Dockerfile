FROM alpine:latest

ADD build/imagetoascii.docker /usr/local/bin/imagetoascii
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/imagetoascii"]
CMD []
