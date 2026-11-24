# Despliegue del backend

Todo el backend (Postgres, las 7 APIs, el MID y Caddy) corre en un solo servidor con
Docker Compose. El front va aparte, en Vercel.

```
front (Vercel) ──HTTPS──> caddy :443 ──> mid :9090 ──> 7 APIs ──> db
                          (lo unico publico)          (red interna de Docker)
```

## Requisitos

- Un servidor Linux con Docker y el plugin compose (1 vCPU / 2 GB alcanza).
- Un dominio con un registro **A** apuntando a la IP del servidor (ej. `api.tudominio.com`).
- Puertos **80 y 443** abiertos. No abrir 5432 ni 8081-8087.

## Primera vez

```bash
mkdir ~/jobsy && cd ~/jobsy
git clone https://github.com/SantiagoRicaurte06/Jobsy_API.git
git clone https://github.com/juanezzzzz/MID_JOBSY.git
# Hoy se despliega develop: main de MID_JOBSY solo tiene el README y el de
# Jobsy_API no tiene estos archivos. Cambiar a main cuando se fusione.
git -C Jobsy_API checkout develop && git -C MID_JOBSY checkout develop

cd Jobsy_API/deploy
cp .env.ejemplo .env
nano .env                  # dominio, URL del front y claves (openssl rand -base64 48)
docker compose up -d --build
docker compose ps          # todo "running"; mid y db "healthy"
curl https://api.tudominio.com/api/salud
```

La primera vez Postgres carga `initdb/` (esquema y datos de ejemplo) y crea el usuario
`jobsy_app`, que solo lee y escribe datos. Eso pasa **solo con el volumen vacio**.

Despues, en Vercel: `environment.prod.ts` del front debe tener
`apiUrl: 'https://api.tudominio.com/api'`, y `URL_FRONT` en `.env` el dominio exacto del front.

## Actualizar

```bash
cd ~/jobsy
git -C Jobsy_API pull && git -C MID_JOBSY pull
cd Jobsy_API/deploy && docker compose up -d --build
```

Los cambios de esquema **no** se aplican solos: `initdb/` solo corre la primera vez.
Cada migracion nueva se aplica a mano:

```bash
docker compose exec -T db psql -U postgres -d jobsy_api -v ON_ERROR_STOP=1 < cambio.sql
```

## Operacion

| Que | Comando |
|---|---|
| Logs de un servicio | `docker compose logs -f mid` |
| Consola SQL | `docker compose exec db psql -U postgres -d jobsy_api` |
| Backup | `docker compose exec -T db pg_dump -U postgres -Fc jobsy_api > jobsy_$(date +%F).dump` |
| Restaurar | `docker compose exec -T db pg_restore -U postgres -d jobsy_api --clean < archivo.dump` |
| Reiniciar uno | `docker compose restart empleo` |

Programar el backup con cron y copiarlo **fuera** del servidor. Los archivos que suben los
usuarios estan en el volumen `archivos_mid`: tambien entran en el backup.

## Pendiente antes de abrirlo al publico

- `initdb/01_jobsy.sql` trae datos de ejemplo (usuarios laura, carlos, vale...). Quitarlos
  si la base de produccion debe arrancar vacia.
- Las credenciales de esos usuarios solo las crea `sembrar-credenciales`, que es de desarrollo.
- La pasarela de pago necesita `MID_SECRETO_PASARELA` y que el proveedor apunte a
  `https://<DOMINIO_API>`.
