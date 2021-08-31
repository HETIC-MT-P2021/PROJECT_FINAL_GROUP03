FROM node:lts-alpine

WORKDIR /app

COPY front/package*.json ./

RUN npm install

COPY front/ .

EXPOSE 8080

CMD ["npm", "run", "serve"]
