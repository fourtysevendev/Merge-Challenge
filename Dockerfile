FROM scratch

WORKDIR /app
COPY main main
# set normal user
USER 1000:1000

ENTRYPOINT [ "/app/main" ]