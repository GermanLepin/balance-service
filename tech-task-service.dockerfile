FROM alpine:latest

RUN mkdir /app

COPY binary_file/avitoTechTaskServiceApp /app

CMD [ "/app/avitoTechTaskServiceApp" ]

