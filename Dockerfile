FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o mood-journal-api

EXPOSE 8080

CMD [ "./mood-journal-api" ]