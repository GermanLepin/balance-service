FROM alpine:latest

RUN mkdir /app

COPY binary_file/balanceServiceApp /app

CMD [ "/app/balanceServiceApp" ]
