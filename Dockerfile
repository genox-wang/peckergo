FROM busybox

RUN echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app

CMD [ "./peckergo" ]