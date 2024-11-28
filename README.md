# Go Database 🚀  

Un proyecto en Go que implementa conexión y manipulación de bases de datos utilizando **GORM** y **MySQL**. Este repositorio incluye un ejemplo básico de configuración y uso de bases de datos relacionales con el objetivo de servir como base para proyectos más avanzados.  

## Características  

- Configuración de conexión a MySQL usando GORM.  
- CRUD básico para gestionar datos en la base de datos.  
- Organización de código en capas para facilitar la escalabilidad y el mantenimiento.  
- Manejo de rutas HTTP con **httprouter**.  

## Requisitos previos  

Antes de comenzar, asegúrate de tener instalados los siguientes elementos en tu entorno:  

- [Go](https://golang.org/) 1.20 o superior.  
- [MySQL](https://www.mysql.com/) 8.0 o superior.  
- Conexión a internet para instalar dependencias.  

## Instalación  

1. Clona este repositorio:  
   ```bash  
   git clone https://github.com/langermanaxel/go-database.git  
   cd go-database  
2. Instala las dependencias del proyecto:
   ```bash  
   go mod tidy
3. Configura la base de datos.
- Crea una base de datos MySQL local.
- Modifica el archivo main.go con los datos de conexión:
  const (
    DBUser     = "tu_usuario"
    DBPassword = "tu_password"
    DBName     = "nombre_base_de_datos"
  )
4. Ejecuta la migración para crear las tablas necesarias:
  go run main.go

## Uso

1. Inicia el servidor:
   go run main.go
2. Accede a las rutas disponibles para interactuar con la base de datos (por ejemplo, usando Postman o curl):
- GET /items: Obtiene todos los registros.
- POST /items: Crea un nuevo registro.
- PUT /items/:id: Actualiza un registro existente.
- DELETE /items/:id: Elimina un registro.

Aquí tienes un ejemplo del archivo README.md para tu proyecto go-database en formato profesional y bien estructurado:

markdown
Copiar código
# Go Database 🚀  

Un proyecto en Go que implementa conexión y manipulación de bases de datos utilizando **GORM** y **MySQL**. Este repositorio incluye un ejemplo básico de configuración y uso de bases de datos relacionales con el objetivo de servir como base para proyectos más avanzados.  

## Características  

- Configuración de conexión a MySQL usando GORM.  
- CRUD básico para gestionar datos en la base de datos.  
- Organización de código en capas para facilitar la escalabilidad y el mantenimiento.  
- Manejo de rutas HTTP con **httprouter**.  

## Requisitos previos  

Antes de comenzar, asegúrate de tener instalados los siguientes elementos en tu entorno:  

- [Go](https://golang.org/) 1.20 o superior.  
- [MySQL](https://www.mysql.com/) 8.0 o superior.  
- Conexión a internet para instalar dependencias.  

## Instalación  

1. Clona este repositorio:  
   ```bash  
   git clone https://github.com/langermanaxel/go-database.git  
   cd go-database  
Instala las dependencias del proyecto:

bash
Copiar código
go mod tidy  
Configura la base de datos.

Crea una base de datos MySQL local.
Modifica el archivo main.go con los datos de conexión:
go
Copiar código
const (
    DBUser     = "tu_usuario"
    DBPassword = "tu_password"
    DBName     = "nombre_base_de_datos"
)  
Ejecuta la migración para crear las tablas necesarias:

bash
Copiar código
go run main.go  
Uso
Inicia el servidor:

bash
Copiar código
go run main.go  
Accede a las rutas disponibles para interactuar con la base de datos (por ejemplo, usando Postman o curl):

GET /items: Obtiene todos los registros.
POST /items: Crea un nuevo registro.
PUT /items/:id: Actualiza un registro existente.
DELETE /items/:id: Elimina un registro.

## Tecnologías usadas
- Golang: Lenguaje de programación.
- GORM: ORM para manejar la base de datos.
- MySQL: Base de datos relacional.
- httprouter: Enrutador eficiente para manejar rutas HTTP.

## Contribuciones
¡Las contribuciones son bienvenidas! Si deseas colaborar:
- Haz un fork del proyecto.
- Crea una rama con tu función o mejora (git checkout -b feature/nueva-funcionalidad).
- Realiza un pull request explicando tus cambios.

## Autor
Creado y mantenido por Axel Langerman.
