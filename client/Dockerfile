FROM node:14.0.0-alpine3.11

WORKDIR /client

COPY package.json ./
COPY package-lock.json ./

RUN npm install

COPY . ./

