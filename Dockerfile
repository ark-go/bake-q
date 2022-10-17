# FROM node:16
FROM node:16.17.0-bullseye-slim
MAINTAINER  Arkadii
RUN npm install -g npm@latest
# нужно для передачи всех сигналов нашему процессу node
RUN apt-get update && apt-get install -y --no-install-recommends dumb-init && \
    apt-get install -y telnet && \
    apt-get install -y iputils-ping && \
    apt-get update && apt-get upgrade -y && apt-get clean
ENV NODE_ENV production
# наш пользователь, login его node
USER node
# Create app directory
WORKDIR /usr/src/app
# Install app dependencies
#COPY --chown=node:node package*.json ./
#RUN npm ci --only=production
#RUN npm ci --omit=dev
# Bundle app source
COPY --chown=node:node ./server ./server
COPY --chown=node:node ./quasar/dist/spa ./quasar/dist/spa
# из node читаем номер версии поэтому пока копируем
COPY --chown=node:node ./quasar/package.json ./quasar/

WORKDIR /usr/src/app/server
RUN npm ci --omit=dev
WORKDIR /usr/src/app

EXPOSE 8877
# CMD [ "node", "./server/index.js" ]
CMD ["dumb-init", "node", "./server/index.js"]
#CMD ["dumb-init","./server/waitStartDocker.sh"] 

# Сборка из корневого каталога с указанием пути к Docerfile
# docker build . -t arkadii/bake -f server/Dockerfile     -f укажем путь к Dockerfile, если он не в текущ. каталоге
# docker build . -t arkadii/bake --squash    сжимает до одного слоя Экcпеиментальный 
#  параметр --net="host"  localhost в док-контейнере будет указывать на ваш док-хост. только Linux
# Запуск docker run --name bakeserver --hostname bakesrv --net="host" --env-file ./server/.env -p 8877:8877 arkadii/bake

# запуск с параметром --net="host" переназначение портов не используется ?
# docker run --rm --name bakeserver --hostname bakesrv --net="host" --env-file ./server/.env arkadii/bake

# Запуск docker run --env-file ./server/.env_docker -p 8877:8877 --add-host="host.docker.internal:host-gateway" arkadii/bake

# docker images
# docker rmi xxxxx -f  удаляем
# docker ps
# docker ps -a  и остановленные
# docker container prune - Удаляетвсе контейнеры 

# docker kill xxx  останавливаем

# это пример для environment  --env-file .env
# docker run --env-file ./env.list ubuntu bash


# Флаги -dit означают запуск контейнера отсоединенным (в фоновом режиме), интерактивным (с возможностью ввода в него текста) 
# и с TTY (чтобы вы могли видеть ввод и вывод).

# docker attach alpine1 Подключениек работающему в фоне контейнеру