FROM node:16-alpine3.16 as UI
WORKDIR /ui
COPY ./ui/ .
RUN yarn install
RUN yarn build

FROM golang:1.19-alpine as APP
WORKDIR /tidyurl
COPY . .
COPY --from=UI /ui/dist /tidyurl/ui/dist
RUN go mod tidy
RUN go build .

FROM alpine:3.14
WORKDIR /tidyurl
COPY --from=UI /ui/dist /tidyurl/ui/dist
COPY --from=APP /tidyurl/tidy-url /tidyurl/tidy-url
CMD [ "./tidy-url" ]
