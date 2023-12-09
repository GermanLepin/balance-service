FROM alpine:latest

RUN mkdir /app

COPY binary_file/balanceServiceApp /app
COPY docker_env/.env /

CMD [ "/app/balanceServiceApp" ]
