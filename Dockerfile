# Imagen de cualquiera de las 7 APIs CRUD. Cada una es un modulo Go aparte,
# asi que se elige con el argumento SERVICIO:
#
#   docker build --build-arg SERVICIO=Empleo_API -t jobsy-empleo .
#
# Beego lee conf/app.conf relativo al directorio de trabajo y sirve swagger/
# como estatico, por eso ambos se copian junto al binario. La configuracion
# (puerto, base, esquema) llega por variables de entorno, no por .env.

FROM golang:1.26-alpine AS compilacion
ARG SERVICIO
WORKDIR /src
COPY ${SERVICIO}/go.mod ${SERVICIO}/go.sum ./
RUN go mod download
COPY ${SERVICIO}/ ./
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/api .

FROM alpine:3.22
RUN apk add --no-cache ca-certificates tzdata \
    && adduser -D -H -u 10001 jobsy
WORKDIR /app
COPY --from=compilacion /out/api ./api
COPY --from=compilacion /src/conf ./conf
COPY --from=compilacion /src/swagger ./swagger
USER jobsy
ENTRYPOINT ["./api"]
