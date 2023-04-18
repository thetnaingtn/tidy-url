FROM node:16-alpine3.16 as UI
WORKDIR /ui
COPY ./ui/ .
ENV VITE_API_ENDPOINT=${VITE_API_ENDPOINT}
RUN yarn install
RUN yarn build

FROM golang:1.19-alpine as APP
WORKDIR /tidyurl
COPY . .
COPY --from=UI /ui/dist /tidyurl/ui/dist
ENV DB_CONNECTION_URL=${DB_CONNECTION_URL}
ENV BASE_URL=${BASE_URL}
RUN go mod tidy
RUN go build .

FROM alpine:3.14
WORKDIR /tidyurl
COPY --from=UI /ui/dist /tidyurl/ui/dist
COPY --from=APP /tidyurl/tidy-url /tidyurl/tidy-url
CMD [ "./tidy-url" ]
