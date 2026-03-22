package main

import (
"fmt"
"log"
"net/http"

"github.com/gorilla/mux"
"github.com/zaferertasglobal-hub/go-microservice-boilerplate/config"
"github.com/zaferertasglobal-hub/go-microservice-boilerplate/db"
"github.com/zaferertasglobal-hub/go-microservice-boilerplate/handlers"
"github.com/zaferertasglobal-hub/go-microservice-boilerplate/middleware"
)

func main() {
// Load configuration
cfg, err := config.LoadConfig()
if err != nil {
log.Fatalf("❌ Failed to load config: %v", err)
}

// Initialize database
err = db.InitDB(cfg)
if err != nil {
log.Fatalf("❌ Failed to initialize database: %v", err)
}
defer db.CloseDB()

// Create router
router := mux.NewRouter()

// Apply middleware
router.Use(middleware.LoggerMiddleware)

// Register routes
router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

// Health check endpoint
router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
w.Write([]byte(`{"status":"healthy"}`))
}).Methods("GET")

// Start server
addr := fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort)
log.Printf("🚀 Server starting on http://%s", addr)
err = http.ListenAndServe(addr, router)
if err != nil {
log.Fatalf("❌ Failed to start server: %v", err)
}
}
