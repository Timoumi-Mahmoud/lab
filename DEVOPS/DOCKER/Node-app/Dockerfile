FROM node:22-alpine

ENV MONGO_DB_USERNAME=admin \
    MONGO_DB_PASSWORD=password

RUN mkdir -p /home/app
COPY package*.json /home/app
RUN npm install
copy /app /home/app

CMD ["node", "/home/app/server.js"]
