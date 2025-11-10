package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"simulator.iot.integrator.6th/src/domain/models"
)

type SQLiteManager struct {
	conn 	*sql.DB
	mutex 	sync.Mutex
}

func NewSQLiteManager()(*SQLiteManager, error){
	conn, err := sql.Open("sqlite3", "sensors.db")
	if err != nil {
		return nil, err
	}

	db := &SQLiteManager{conn: conn}

	fmt.Println("Conexión a la base de datos establecida")

	if err = db.initTables(); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *SQLiteManager) initTables() error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	query := `CREATE TABLE IF NOT EXISTS sensor_readings (
		id_sensor_reading INTEGER PRIMARY KEY AUTOINCREMENT,
		value REAL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		id_sensor TEXT,
		backed BOOLEAN,
		FOREIGN KEY(id_sensor) REFERENCES sensors(id_sensor) ON DELETE CASCADE
	);`

	_, err := db.conn.Exec(query)
	if err != nil {
		return err
	}

	println("Tablas creadas")
	return nil
}

func (db *SQLiteManager) CreateSensorReading(value float64, idSensor string, backed bool) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	_, err := db.conn.Exec(`
		INSERT INTO sensor_readings (value, id_sensor, backed) VALUES (?, ?, ?)
	`, value, idSensor, backed)

	return err
}

func (db *SQLiteManager) GetSensorReadingsNotSent() ([]models.SensorReadings, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	rows, err := db.conn.Query(`SELECT * FROM sensor_readings WHERE backed = false`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.SensorReadings

	for rows.Next() {
		var id int
		var value float32
		var timestamp, idSensor string
		var backed bool

		if err := rows.Scan(&id, &value, &timestamp, &idSensor, &backed); err != nil {
			return nil, err
		}

		results = append(results, models.SensorReadings{
			IdSensorReading: id,
			Value: value,
			Timestamp: timestamp,
			IdSensor: idSensor,
			Backed: backed,
		})
	}

	return results, nil
}

func (db *SQLiteManager) MarkSensorReadingSent(id int) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	_, err := db.conn.Exec(`
		UPDATE sensor_readings SET backed = true WHERE id_sensor_reading = ?
	`, id)

	return err
}

// Cerrar DB
func (db *SQLiteManager) CloseDB() {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	fmt.Println("Cerrando conexión con la db")
	db.conn.Close()
}