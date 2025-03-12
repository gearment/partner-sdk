FROM node:18-alpine
RUN apk add --no-cache curl
WORKDIR /app
COPY package.json package-lock.json /app/
RUN npm install
RUN npm install -g @redocly/cli@latest
COPY download-openapi.sh /app/
RUN chmod +x /app/download-openapi.sh
RUN /bin/sh /app/download-openapi.sh
COPY . /app/
CMD ["npx", "run", "start"]