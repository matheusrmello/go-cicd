FROM ubuntu:latest

LABEL version="1.0"
LABEL description="imagem GO para CI"

RUN apt-get update && apt-get install -y

EXPOSE 8000

WORKDIR /app

ENV HOST=localhost PORT=5432

ENV USER=root PASSWORD=root DBNAME=root

COPY ./main.exe main

CMD [ "./main" ]