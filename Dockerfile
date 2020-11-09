FROM alpine
COPY proxy-protocol-debugger /
ENTRYPOINT ["/proxy-protocol-debugger"]
