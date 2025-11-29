package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type GlobalConfig struct {
	ServerPort int
	DBDsn      string
	DBStage    string // Para diferenciar entornos de DB (local, test, etc.)
	SecretKey  string
	// Configuración del Entorno
	Environment string // development, staging, production
}

var Config *GlobalConfig

func LoadConfig() error {
	// 1. Inicialización de la configuración
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or error loading it: %v", err)
		return err
	}
	cfg := &GlobalConfig{}

	// --- Configuración del Servidor ---

	// Se lee el puerto y se convierte de string a int
	portStr := os.Getenv("PORT_SERVER")
	if portStr == "" {
		portStr = "8080" // Valor por defecto si no está seteado
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return logAndError("PORT_SERVER debe ser un número entero. Valor actual: " + portStr)
	}
	cfg.ServerPort = port

	// --- Configuración de la Base de Datos ---

	// MARIADB_DSN es crítica, debe estar presente
	cfg.DBDsn = os.Getenv("DBDsn")
	if cfg.DBDsn == "" {
		return logAndError("MARIADB_DSN no puede estar vacío.")
	}

	cfg.DBStage = os.Getenv("DB_STAGE")

	// --- Configuración del Entorno ---

	cfg.Environment = os.Getenv("ENVIRONMENT")
	if cfg.Environment == "" {
		cfg.Environment = "development" // Valor por defecto
	}
	cfg.SecretKey = os.Getenv("SECRET_KEY")
	if cfg.SecretKey == "" {
		cfg.SecretKey = "Pruebas123@Secret"
	}
	// Almacenar la configuración cargada
	Config = cfg
	log.Printf("Configuración cargada. Entorno: %s, Puerto: %d", Config.Environment, Config.ServerPort)

	return nil
}

// Helper para loggear y devolver un error
func logAndError(msg string) error {
	log.Printf("[ERROR] Configuración fallida: %s", msg)
	return os.ErrInvalid
}
