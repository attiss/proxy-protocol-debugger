FROM alpine
EXPOSE 8080/tcp
COPY proxy-protocol-debugger /
ENTRYPOINT ["/proxy-protocol-debugger"]
