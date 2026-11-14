FROM node:22-alpine AS runner

WORKDIR /app

COPY  package.json  ./

RUN npm install

COPY . .

EXPOSE 8081

CMD [ "npm", "run", "dev", "--", "--hostname", "0.0.0.0" ]