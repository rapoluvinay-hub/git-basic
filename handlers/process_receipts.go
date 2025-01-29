package handlers

import (
    "encoding/json"
    "math/rand"
    "net/http"
    "receipt-processor/models"
    "receipt-processor/utils"
    "sync"
    "time"
)

var receiptStore = make(map[string]models.Receipt)
var mutex = &sync.Mutex{}

func ProcessReceipts(w http.ResponseWriter, r *http.Request) {
    var receipt models.Receipt
    err := json.NewDecoder(r.Body).Decode(&receipt)
    if err != nil {
        http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
        return
    }

    // Generate unique ID
    receipt.ID = generateID()
    receipt.Points = utils.CalculatePoints(receipt)

    // Store receipt
    mutex.Lock()
    receiptStore[receipt.ID] = receipt
    mutex.Unlock()

    // Respond with ID
    response := map[string]string{"id": receipt.ID}
    json.NewEncoder(w).Encode(response)
}

func generateID() string {
    rand.Seed(time.Now().UnixNano())
    return fmt.Sprintf("%x", rand.Int())
}
