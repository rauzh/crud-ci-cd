FROM golang:1.20-alpine AS build_stage
COPY ./notestorage /go/src/notestorage
WORKDIR /go/src/notestorage
RUN go install .

FROM alpine AS run_stage
WORKDIR /app_binary
COPY --from=build_stage /go/bin/notestorage /app_binary/
RUN chmod +x ./notestorage
EXPOSE 8080/tcp
ENTRYPOINT ./notestorage

EXPOSE 8080/tcp
CMD [ "notestorage" ]
