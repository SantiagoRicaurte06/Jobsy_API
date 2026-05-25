# Jobsy API

Sistema de microservicios desarrollado con Go y Beego para la plataforma **Jobsy**.

El proyecto está dividido en múltiples APIs independientes organizadas por dominio funcional.

---

# Arquitectura del Proyecto

El sistema utiliza una arquitectura basada en microservicios.

Cada módulo contiene:

* Controllers
* Models
* Routers
* Configuración propia
* Swagger UI
* Scripts SQL
* Conexión independiente a PostgreSQL

---

# Tecnologías Utilizadas

* Go
* Beego Framework
* PostgreSQL
* Swagger UI
* REST API
* JSON

---

# Estructura General

```bash
Jobsy_API/
│
├── Catalogo_api/
├── Core_api/
├── Empleo_API/
├── Soporte_API/
├── suscripciones_API/
├── Tienda_APi/
├── Usuarios_Api/
│
└── README.md
```

---

# Microservicios Detectados

## 1. Catalogo_api

### Endpoints 

```http
GET    /v1/dias_semana
POST   /v1/dias_semana
PUT    /v1/dias_semana/:id
DELETE /v1/dias_semana/:id

GET    /v1/estados
POST   /v1/estados
PUT    /v1/estados/:id
DELETE /v1/estados/:id

GET    /v1/generos
POST   /v1/generos
PUT    /v1/generos/:id
DELETE /v1/generos/:id

GET    /v1/modalidades_suscripcion
POST   /v1/modalidades_suscripcion
PUT    /v1/modalidades_suscripcion/:id
DELETE /v1/modalidades_suscripcion/:id

GET    /v1/modalidades_trabajo
POST   /v1/modalidades_trabajo
PUT    /v1/modalidades_trabajo/:id
DELETE /v1/modalidades_trabajo/:id
```

---

## 2. Core_api

### Endpoints 

```http
GET    /v1/credenciales
POST   /v1/credenciales
PUT    /v1/credenciales/:id
DELETE /v1/credenciales/:id

GET    /v1/documentos_identidad
POST   /v1/documentos_identidad
PUT    /v1/documentos_identidad/:id
DELETE /v1/documentos_identidad/:id

GET    /v1/roles_usuarios
POST   /v1/roles_usuarios
PUT    /v1/roles_usuarios/:id
DELETE /v1/roles_usuarios/:id

GET    /v1/sesiones
POST   /v1/sesiones
PUT    /v1/sesiones/:id
DELETE /v1/sesiones/:id

GET    /v1/tipo_documento
POST   /v1/tipo_documento
PUT    /v1/tipo_documento/:id
DELETE /v1/tipo_documento/:id
```

---

## 3. Empleo_API

### Endpoints 

```http
GET    /v1/certificaciones
POST   /v1/certificaciones
PUT    /v1/certificaciones/:id
DELETE /v1/certificaciones/:id

GET    /v1/contrataciones
POST   /v1/contrataciones
PUT    /v1/contrataciones/:id
DELETE /v1/contrataciones/:id

GET    /v1/dias_trabajados
POST   /v1/dias_trabajados
PUT    /v1/dias_trabajados/:id
DELETE /v1/dias_trabajados/:id

GET    /v1/disponibilidad_empleado
POST   /v1/disponibilidad_empleado
PUT    /v1/disponibilidad_empleado/:id
DELETE /v1/disponibilidad_empleado/:id

GET    /v1/documentos_usuario
POST   /v1/documentos_usuario
PUT    /v1/documentos_usuario/:id
DELETE /v1/documentos_usuario/:id
```

---

## 4. Soporte_API

### Endpoints 

```http
GET    /v1/reportes
POST   /v1/reportes
PUT    /v1/reportes/:id
DELETE /v1/reportes/:id

GET    /v1/reporte_mensajes
POST   /v1/reporte_mensajes
PUT    /v1/reporte_mensajes/:id
DELETE /v1/reporte_mensajes/:id
```

---

## 5. suscripciones_API

### Endpoints 

```http
GET    /v1/planes
POST   /v1/planes
PUT    /v1/planes/:id
DELETE /v1/planes/:id

GET    /v1/plan_features
POST   /v1/plan_features
PUT    /v1/plan_features/:id
DELETE /v1/plan_features/:id

GET    /v1/usuarios_suscripciones
POST   /v1/usuarios_suscripciones
PUT    /v1/usuarios_suscripciones/:id
DELETE /v1/usuarios_suscripciones/:id
```

---

## 6. Tienda_APi

### Endpoints 

```http
GET    /v1/carritos
POST   /v1/carritos
PUT    /v1/carritos/:id
DELETE /v1/carritos/:id

GET    /v1/categorias_tienda
POST   /v1/categorias_tienda
PUT    /v1/categorias_tienda/:id
DELETE /v1/categorias_tienda/:id

GET    /v1/cupones
POST   /v1/cupones
PUT    /v1/cupones/:id
DELETE /v1/cupones/:id

GET    /v1/facturacion
POST   /v1/facturacion
PUT    /v1/facturacion/:id
DELETE /v1/facturacion/:id

GET    /v1/metodos_pago_guardados
POST   /v1/metodos_pago_guardados
PUT    /v1/metodos_pago_guardados/:id
DELETE /v1/metodos_pago_guardados/:id
```

---

## 7. Usuarios_Api

### Endpoints

```http
GET    /v1/configuraciones
POST   /v1/configuraciones
PUT    /v1/configuraciones/:id
DELETE /v1/configuraciones/:id

GET    /v1/notificaciones
POST   /v1/notificaciones
PUT    /v1/notificaciones/:id
DELETE /v1/notificaciones/:id
```

---

# Configuración Local

## PostgreSQL

Puerto:

```bash
5432
```

## APIs

Puerto base:

```bash
8080
```

---

# Instalación

## 1. Clonar el proyecto

```bash
git clone https://github.com/SantiagoRicaurte06/Jobsy_API.git
```

---

## 2. Entrar al microservicio

Ejemplo:

```bash
cd Catalogo_api
```

---

## 3. Instalar dependencias

```bash
go mod tidy
```

---

## 4. Configurar PostgreSQL

Editar:

```bash
conf/app.conf
```

---

## 5. Ejecutar API

```bash
bee run
```

O:

```bash
go run main.go
```

---

# Swagger

Cada microservicio incluye Swagger UI.

Archivos :

```bash
swagger/
```

acceso:

```bash
http://localhost:8080/swagger
```

---

# Estructura Interna de Cada API

```bash
API_NAME/
│
├── conf/
├── controllers/
├── database/
├── models/
├── routers/
├── scripts/
├── swagger/
├── go.mod
├── go.sum
└── main.go
```

---

# Características del Proyecto

* Arquitectura modular
* APIs desacopladas
* Swagger integrado
* PostgreSQL
* CRUD automático con Beego
* Separación por dominios
* Escalable
* Microservicios independientes

---

# Recomendaciones

## Desarrollo

Instalar:

```bash
go install github.com/beego/bee/v2@latest
```

Verificar:

```bash
bee version
```

---

# Ejemplo de Flujo CRUD

```http
POST   /v1/generos
GET    /v1/generos
GET    /v1/generos/:id
PUT    /v1/generos/:id
DELETE /v1/generos/:id
```

# Autor

Proyecto Jobsy API.

Arquitectura basada en microservicios utilizando Go + Beego + PostgreSQL.
