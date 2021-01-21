FROM scratch
COPY example /usr/bin/example
ENTRYPOINT ["/usr/bin/example"]
