FROM scratch

WORKDIR /app
COPY main .
# set normal user
USER 1000:1000
# env interval list to merge
ARG INTERVALS
ENV INTERVALS=${INTERVALS}

CMD [ "./main" ]