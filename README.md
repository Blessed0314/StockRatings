# Proyecto de Consulta de Stock Ratings

Este proyecto es una aplicación desarrollada en **Go** para el backend y **Vue.js** para el frontend. Su propósito es consultar una API externa para obtener una lista actualizada de stock ratings y almacenarla en una base de datos. La aplicación está diseñada para ser eficiente, escalable y fácil de mantener.

---

## Tecnologías Utilizadas

### Backend
- **Go**: Lenguaje principal para el desarrollo del backend.
- **GORM**: ORM utilizado para interactuar con la base de datos.
- **CockroachDB**: Base de datos distribuida utilizada para almacenar los datos.
- **godotenv**: Para cargar variables de entorno desde un archivo `.env`.

### Frontend
- **Vue.js**: Framework utilizado para construir la interfaz de usuario.

---

## Funcionalidades

1. **Consulta de API Externa**:
   - El backend realiza una solicitud a una API externa para obtener una lista de stock ratings.
   - La API utilizada es: `https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list?format=truora`.

2. **Almacenamiento en Base de Datos**:
   - Los datos obtenidos de la API se almacenan en **CockroachDB** utilizando GORM.

3. **Interfaz de Usuario**:
   - El frontend desarrollado en Vue.js permite visualizar los datos almacenados de manera amigable.

---

## Configuración del Proyecto

### Variables de Entorno
El proyecto utiliza un archivo `.env` para configurar las variables de entorno necesarias. Asegúrate de crear este archivo en la raíz del proyecto y definir las siguientes variables:

```env
# Configuración de la API externa
API_TOKEN=tu_token_de_api

# Configuración de la base de datos
DB_HOST=localhost
DB_USER=usuario
DB_PASSWORD=contraseña
DB_NAME=nombre_base_datos
DB_PORT=26257
