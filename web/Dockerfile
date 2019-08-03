FROM node:latest
RUN npm install -g @vue/cli

ADD ./ /web
WORKDIR /web

RUN npm ci
RUN npm run build
ENTRYPOINT node index.js