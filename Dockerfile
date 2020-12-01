ARG ARCH
FROM ${ARCH}/alpine
COPY example /usr/bin/example
ENTRYPOINT ["/usr/bin/example"]
