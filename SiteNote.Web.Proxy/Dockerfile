FROM node:20-alpine

WORKDIR /frontend

COPY . .

WORKDIR /frontend/SiteNote.Web

RUN npm install

RUN npm run build

EXPOSE 5173

WORKDIR /frontend

RUN npm install

CMD ["node", "app.js"]