FROM node:14.15.4-alpine3.12
RUN apk update && apk add --update python3 make g++
RUN npm install -g http-server
WORKDIR /app
COPY package*.json ./
#RUN npm install axios
#RUN npm install deasync
RUN npm install
COPY . .
RUN npm run build
EXPOSE 8080
CMD [ "http-server", "dist" ]