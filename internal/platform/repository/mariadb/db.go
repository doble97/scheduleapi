package mariadb

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func InitDB(dsn string) (*sql.DB, error) {
	// 1. Abrir la conexión (sql.Open no prueba la conexión inmediatamente)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la conexión: %w", err)
	}

	// 2. Configuración del Pooling de Conexiones
	// Esto es crucial para el rendimiento en Go.

	// Máximo de conexiones abiertas a la vez
	db.SetMaxOpenConns(100)

	// Máximo de conexiones inactivas (en la piscina)
	db.SetMaxIdleConns(25)

	// Tiempo máximo que una conexión puede estar abierta
	db.SetConnMaxLifetime(5 * time.Minute)

	// 3. Probar la conexión (Ping)
	if err = db.Ping(); err != nil {
		db.Close() // Cerrar la conexión si el ping falla
		return nil, fmt.Errorf("error al hacer ping a la base de datos: %w", err)
	}

	log.Println("Conexión a MariaDB establecida y verificada.")
	return db, nil
}
