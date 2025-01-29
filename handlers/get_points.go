package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    mutex.Lock()
    receipt, exists := receiptStore[id]
    mutex.Unlock()

    if !exists {
        http.Error(w, "Receipt not found", http.StatusNotFound)
        return
    }

    response := map[string]int{"points": receipt.Points}
    json.NewEncoder(w).Encode(response)
}
