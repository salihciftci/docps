FROM node:alpine AS base
LABEL MAINTAINER="Salih Çiftçi"

WORKDIR /liman

COPY . .

RUN yarn install --production

FROM node:alpine

COPY --from=base /liman /liman

RUN apk add -U --no-cache ca-certificates docker

EXPOSE 5000

CMD /liman/bin/www