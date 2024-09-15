FROM node:22-alpine
WORKDIR /app
COPY /web .
RUN npm install
RUN npm run build
EXPOSE 3000
CMD ["npm", "run", "preview"]