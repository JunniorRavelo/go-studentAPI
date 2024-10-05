### README.md

---

# Student Management API (Go)

## Description

This is a simple API for managing students built with Go. The API stores all student data in memory (without a database) and provides basic CRUD (Create, Read, Update, Delete) functionality for student records. Each student record includes fields like `ID`, `username`, `name`, `age`, `grade`, `background`, `admission date`, `status`, `data creation date`, and `data update date`.

### Features:
- In-memory data storage (no external database required)
- CRUD operations for student management
- Student states: New, Old, Suspended, Graduated
- Auto-generated ID, created and updated timestamps for each student

## Technologies:
- **Go** (Programming Language)
- **Gorilla Mux** (HTTP routing and URL matching)

## Endpoints

### List all students
- **Method**: `GET`
- **URL**: `/estudiantes`
- **Response**: JSON array of students
```bash
curl http://localhost:8080/estudiantes
```

### Get a single student by ID
- **Method**: `GET`
- **URL**: `/estudiantes/{id}`
- **Response**: JSON of the requested student
```bash
curl http://localhost:8080/estudiantes/1
```

### Add a new student
- **Method**: `POST`
- **URL**: `/estudiantes`
- **Body**: JSON
```json
{
    "nombre_usuario": "jdoe",
    "nombre": "John Doe",
    "edad": 20,
    "grado": "3rd Year",
    "antecedentes": "No records",
    "estado": "new student"
}
```
```bash
curl -X POST http://localhost:8080/estudiantes -H "Content-Type: application/json" -d '{
    "nombre_usuario": "jdoe",
    "nombre": "John Doe",
    "edad": 20,
    "grado": "3rd Year",
    "antecedentes": "No records",
    "estado": "new student"
}'
```

### Update an existing student
- **Method**: `PUT`
- **URL**: `/estudiantes/{id}`
- **Body**: JSON of updated student data
```bash
curl -X PUT http://localhost:8080/estudiantes/1 -H "Content-Type: application/json" -d '{
    "nombre_usuario": "jdoe",
    "nombre": "John Doe",
    "edad": 21,
    "grado": "4th Year",
    "antecedentes": "No records",
    "estado": "old student"
}'
```

### Delete a student
- **Method**: `DELETE`
- **URL**: `/estudiantes/{id}`
```bash
curl -X DELETE http://localhost:8080/estudiantes/1
```

## How to Run

1. **Install Go**: Make sure you have Go installed on your machine. Check by running:
    ```bash
    go version
    ```

2. **Clone the repository**:
    ```bash
    git clone https://github.com/JunniorRavelo/go-studentAPI.git
    cd go-studentAPI
    ```

3. **Initialize the Go module**:
    ```bash
    go mod init go-studentAPI
    ```

4. **Get the Gorilla Mux dependency**:
    ```bash
    go get github.com/gorilla/mux
    ```

5. **Run the API**:
    ```bash
    go run main.go
    ```

6. The API will be available at `http://localhost:8080`.

## Example

To test the API, you can use `curl` or a tool like Postman. Below is an example of creating and listing students using `curl`.

1. **Create a student**:
    ```bash
    curl -X POST http://localhost:8080/estudiantes -H "Content-Type: application/json" -d '{
        "nombre_usuario": "jdoe",
        "nombre": "John Doe",
        "edad": 20,
        "grado": "3rd Year",
        "antecedentes": "No records",
        "estado": "new student"
    }'
    ```

2. **List all students**:
    ```bash
    curl http://localhost:8080/estudiantes
    ```

---

# API de Gestión de Estudiantes (Go)

## Descripción

Esta es una API sencilla para la gestión de estudiantes construida con Go. La API almacena todos los datos de los estudiantes en memoria (sin base de datos) y proporciona funcionalidad básica CRUD (Crear, Leer, Actualizar, Eliminar) para los registros de estudiantes. Cada registro de estudiante incluye campos como `ID`, `nombre de usuario`, `nombre`, `edad`, `grado`, `antecedentes`, `fecha de ingreso`, `estado`, `fecha de creación de datos` y `fecha de actualización de datos`.

### Características:
- Almacenamiento de datos en memoria (no se requiere base de datos externa)
- Operaciones CRUD para la gestión de estudiantes
- Estados del estudiante: Nuevo, Antiguo, Suspendido, Graduado
- ID generado automáticamente, timestamps de creación y actualización

## Tecnologías:
- **Go** (Lenguaje de programación)
- **Gorilla Mux** (Enrutamiento HTTP y coincidencia de URLs)

## Endpoints

### Listar todos los estudiantes
- **Método**: `GET`
- **URL**: `/estudiantes`
- **Respuesta**: Array JSON de estudiantes
```bash
curl http://localhost:8080/estudiantes
```

### Obtener un estudiante por ID
- **Método**: `GET`
- **URL**: `/estudiantes/{id}`
- **Respuesta**: JSON del estudiante solicitado
```bash
curl http://localhost:8080/estudiantes/1
```

### Agregar un nuevo estudiante
- **Método**: `POST`
- **URL**: `/estudiantes`
- **Cuerpo**: JSON
```json
{
    "nombre_usuario": "jdoe",
    "nombre": "John Doe",
    "edad": 20,
    "grado": "3er Año",
    "antecedentes": "Sin antecedentes",
    "estado": "nuevo estudiante"
}
```
```bash
curl -X POST http://localhost:8080/estudiantes -H "Content-Type: application/json" -d '{
    "nombre_usuario": "jdoe",
    "nombre": "John Doe",
    "edad": 20,
    "grado": "3er Año",
    "antecedentes": "Sin antecedentes",
    "estado": "nuevo estudiante"
}'
```

### Actualizar un estudiante existente
- **Método**: `PUT`
- **URL**: `/estudiantes/{id}`
- **Cuerpo**: JSON de los datos actualizados del estudiante
```bash
curl -X PUT http://localhost:8080/estudiantes/1 -H "Content-Type: application/json" -d '{
    "nombre_usuario": "jdoe",
    "nombre": "John Doe",
    "edad": 21,
    "grado": "4to Año",
    "antecedentes": "Sin antecedentes",
    "estado": "estudiante antiguo"
}'
```

### Eliminar un estudiante
- **Método**: `DELETE`
- **URL**: `/estudiantes/{id}`
```bash
curl -X DELETE http://localhost:8080/estudiantes/1
```

## Cómo Ejecutar

1. **Instalar Go**: Asegúrate de tener Go instalado en tu máquina. Compruébalo ejecutando:
    ```bash
    go version
    ```

2. **Clonar el repositorio**:
    ```bash
    git clone https://github.com/JunniorRavelo/go-studentAPI.git
    cd go-studentAPI
    ```

3. **Inicializa el módulo Go**:
    ```bash
    go mod init go-studentAPI
    ```

4. **Obtén la dependencia Gorilla Mux**:
    ```bash
    go get github.com/gorilla/mux
    ```

5. **Ejecuta la API**:
    ```bash
    go run main.go
    ```

6. La API estará disponible en `http://localhost:8080`.

## Ejemplo

Para probar la API, puedes usar `curl` o una herramienta como Postman. A continuación se muestra un ejemplo de cómo crear y listar estudiantes usando `curl`.

1. **Crear un estudiante**:
    ```bash
    curl -X POST http://localhost:8080/estudiantes -H "Content-Type: application/json" -d '{
        "nombre_usuario": "jdoe",
        "nombre": "John Doe",
        "edad": 20,
        "grado": "3er Año",
        "antecedentes": "Sin antecedentes",
        "estado": "nuevo estudiante"
    }'
    ```

2. **Listar todos los estudiantes**:
    ```bash
    curl http://localhost:8080/estudiantes
    ```

---