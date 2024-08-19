ackage FechaHoraAutomatica

import (
 "encoding/json"
 "fmt"
 "net/http"
 "time"
)

type TimeResponse struct {
 Timestamp int64 `json:"timestamp"`
}

func ObtenerFecha(w http.ResponseWriter, r *http.Request) {
 timestamp := time.Now().Unix()
 response := TimeResponse{Timestamp: timestamp}

 w.Header().Set("Content-Type", "application/json")
 json.NewEncoder(w).Encode(response)
}

func main() {
 http.HandleFunc("/", GetTime)
 fmt.Println("Funcion Iniciada...") // Cambiado para Cloud Functions
}
