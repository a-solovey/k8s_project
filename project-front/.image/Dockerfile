FROM node:16.16.0 AS build
WORKDIR /project
COPY . .
RUN npm install
RUN npm run build

FROM alpine:latest AS dist
RUN apk update && apk add bash
WORKDIR /dist
COPY --from=build /project/dist .
COPY ./.image/keep-alive.sh /usr/bin/keep-alive.sh
RUN chmod +x /usr/bin/keep-alive.sh
ENTRYPOINT [ "keep-alive.sh" ]

