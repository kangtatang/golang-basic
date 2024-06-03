package main

import (
	"log"

	"be_karyawan/config"
	"be_karyawan/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
    // Koneksi ke database
    db, err := config.ConnectDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer func() {
        if err := db.Close(); err != nil {
            log.Fatalf("Failed to close database connection: %v", err)
        }
    }()

    // Inisialisasi router Gin
    r := gin.Default()

    // Setup rute dengan database connection
    controllers.SetupRoutes(r, db)

    // Jalankan server pada port 8080
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
