FROM scratch

WORKDIR /app
COPY main main
USER 1000:1000

ENTRYPOINT [ "/app/main" ]