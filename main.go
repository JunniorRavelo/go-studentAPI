package main

import (
    "encoding/json"
    "log"
    "math/rand"
    "net/http"
    "strconv"
    "sync"
    "time"

    "github.com/gorilla/mux"
)

// Estado del estudiante
type Estado string

const (
    NuevoEstudiante   Estado = "nuevo estudiante"
    AntiguoEstudiante Estado = "antiguo estudiante"
    Suspendido        Estado = "suspendido"
    Graduado          Estado = "graduado"
)

// Estructura del estudiante
type Estudiante struct {
    ID            int       `json:"id"`
    NombreUsuario string    `json:"nombre_usuario"`
    Nombre        string    `json:"nombre"`
    Edad          int       `json:"edad"`
    Grado         string    `json:"grado"`
    Antecedentes  string    `json:"antecedentes"`
    FechaIngreso  time.Time `json:"fecha_ingreso"`
    Estado        Estado    `json:"estado"`
    DataCreate    time.Time `json:"data_create"`
    DataUpdate    time.Time `json:"data_update"`
}

// Almacén de estudiantes en memoria
var estudiantes = make(map[int]Estudiante)
var mu sync.Mutex

// Función para obtener un nuevo ID
func getNewID() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(1000000)
}

// Listar todos los estudiantes
func obtenerEstudiantes(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    var listaEstudiantes []Estudiante
    for _, estudiante := range estudiantes {
        listaEstudiantes = append(listaEstudiantes, estudiante)
    }

    json.NewEncoder(w).Encode(listaEstudiantes)
}

// Obtener un estudiante por ID
func obtenerEstudiante(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    estudiante, existe := estudiantes[id]
    if !existe {
        http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(estudiante)
}

// Crear un nuevo estudiante
func crearEstudiante(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    var nuevoEstudiante Estudiante
    json.NewDecoder(r.Body).Decode(&nuevoEstudiante)

    nuevoEstudiante.ID = getNewID()
    nuevoEstudiante.FechaIngreso = time.Now()
    nuevoEstudiante.DataCreate = time.Now()
    nuevoEstudiante.DataUpdate = time.Now()

    estudiantes[nuevoEstudiante.ID] = nuevoEstudiante

    json.NewEncoder(w).Encode(nuevoEstudiante)
}

// Actualizar estudiante
func actualizarEstudiante(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    estudiante, existe := estudiantes[id]
    if !existe {
        http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
        return
    }

    var estudianteActualizado Estudiante
    json.NewDecoder(r.Body).Decode(&estudianteActualizado)

    estudiante.NombreUsuario = estudianteActualizado.NombreUsuario
    estudiante.Nombre = estudianteActualizado.Nombre
    estudiante.Edad = estudianteActualizado.Edad
    estudiante.Grado = estudianteActualizado.Grado
    estudiante.Antecedentes = estudianteActualizado.Antecedentes
    estudiante.Estado = estudianteActualizado.Estado
    estudiante.DataUpdate = time.Now()

    estudiantes[id] = estudiante

    json.NewEncoder(w).Encode(estudiante)
}

// Eliminar estudiante
func eliminarEstudiante(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    _, existe := estudiantes[id]
    if !existe {
        http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
        return
    }

    delete(estudiantes, id)

    w.WriteHeader(http.StatusNoContent)
}

// Configurar rutas
func configurarRutas() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/estudiantes", obtenerEstudiantes).Methods("GET")
    router.HandleFunc("/estudiantes/{id}", obtenerEstudiante).Methods("GET")
    router.HandleFunc("/estudiantes", crearEstudiante).Methods("POST")
    router.HandleFunc("/estudiantes/{id}", actualizarEstudiante).Methods("PUT")
    router.HandleFunc("/estudiantes/{id}", eliminarEstudiante).Methods("DELETE")

    return router
}

func main() {
    router := configurarRutas()

    log.Println("Servidor escuchando en http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
