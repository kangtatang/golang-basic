package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"be_karyawan/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SetupRoutes sets up the routes for employee management
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/employees", func(c *gin.Context) { CreateEmployeeHandler(c, db) })
	r.GET("/employees", func(c *gin.Context) { GetAllEmployeesHandler(c, db) })
	r.GET("/employees/:id", func(c *gin.Context) { GetEmployeeByIDHandler(c, db) })
	r.PUT("/employees/:id", func(c *gin.Context) { UpdateEmployeeHandler(c, db) })
	r.DELETE("/employees/:id", func(c *gin.Context) { DeleteEmployeeHandler(c, db) })
}

// CreateEmployeeHandler adalah handler untuk membuat karyawan baru
func CreateEmployeeHandler(c *gin.Context, db *gorm.DB) {
	var employee models.Employee

	// Bind data dari request body ke struct Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi data karyawan
	if err := validateEmployeeData(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set waktu pembuatan dan pembaruan
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()

	// Simpan karyawan ke dalam database
	if err := db.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan karyawan"})
		return
	}

	// Beri respons sukses
	c.JSON(http.StatusCreated, employee)
}

// GetAllEmployeesHandler adalah handler untuk menampilkan semua karyawan
func GetAllEmployeesHandler(c *gin.Context, db *gorm.DB) {
	var employees []models.Employee

	if err := db.Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data karyawan"})
		return
	}

	c.JSON(http.StatusOK, employees)
}

// GetEmployeeByIDHandler adalah handler untuk menampilkan detail karyawan berdasarkan ID
func GetEmployeeByIDHandler(c *gin.Context, db *gorm.DB) {
	employeeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID karyawan tidak valid"})
		return
	}

	var employee models.Employee
	if err := db.First(&employee, employeeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Karyawan tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// UpdateEmployeeHandler adalah handler untuk memperbarui informasi karyawan berdasarkan ID
func UpdateEmployeeHandler(c *gin.Context, db *gorm.DB) {
	employeeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID karyawan tidak valid"})
		return
	}

	var employee models.Employee
	if err := db.First(&employee, employeeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Karyawan tidak ditemukan"})
		return
	}

	// Bind data dari request body ke struct Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set waktu pembaruan
	employee.UpdatedAt = time.Now()

	// Simpan perubahan karyawan ke dalam database
	if err := db.Save(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui informasi karyawan"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// DeleteEmployeeHandler adalah handler untuk menghapus karyawan berdasarkan ID
func DeleteEmployeeHandler(c *gin.Context, db *gorm.DB) {
	employeeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID karyawan tidak valid"})
		return
	}

	var employee models.Employee
	if err := db.First(&employee, employeeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Karyawan tidak ditemukan"})
		return
	}

	// Hapus karyawan dari database
	if err := db.Delete(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus karyawan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Karyawan berhasil dihapus"})
}

// validateEmployeeData adalah fungsi untuk validasi data karyawan
func validateEmployeeData(employee *models.Employee) error {
	// Lakukan validasi data di sini
	// Misalnya, validasi email, tanggal masuk, NIK, dll.

	// Contoh validasi sederhana: pastikan Nama tidak kosong
	if employee.Name == "" {
		return errors.New("Nama karyawan tidak boleh kosong")
	}

	return nil
}
