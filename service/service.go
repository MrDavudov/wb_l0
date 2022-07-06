package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MrDavudov/LessonPosgresql/models"
	_ "github.com/lib/pq"
)

var Db *sql.DB

var InitQueryCreate = `
CREATE TABLE IF NOT EXISTS orders
	(order_uid VARCHAR(100) UNIQUE,
	track_number VARCHAR(100),
	entry VARCHAR(100),
	locale VARCHAR(100),
	internal_signature VARCHAR(100),
	customer_id VARCHAR(100),
	delivery_service VARCHAR(100),
	shardkey VARCHAR(100),
	sm_id INT,
	date_created VARCHAR(100),
	oof_shard VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS deliveries
	(order_uid VARCHAR(100) UNIQUE,
	name VARCHAR(100),
	phone VARCHAR(100),
	zip VARCHAR(100),
	city VARCHAR(100),
	address VARCHAR(100),
	region VARCHAR(100),
	email VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS payments
	(order_uid VARCHAR(100) UNIQUE,
	request_id VARCHAR(100),
	currency VARCHAR(100),
	provider VARCHAR(100),
	amount INT,
	payment_dt INT,
	bank VARCHAR(100),
	delivery_cost INT,
	goods_total INT,
	custom_fee INT
);

CREATE TABLE IF NOT EXISTS items
	(order_uid VARCHAR(100),
	chrt_id INT,
	track_number VARCHAR(100),
	price VARCHAR(100),
	rid VARCHAR(100),
	name VARCHAR(100),
	sale INT,
	size VARCHAR(100),
	total_price INT,
	nm_id INT,
	brand VARCHAR(100),
	status INT
);

CREATE TABLE IF NOT EXISTS nats
	(
	    host VARCHAR(100) UNIQUE,
	    last VARCHAR(100)
);`

func DBInit(cfg models.Config) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connect db: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error ping db: %s", err)
	}

	_, err = db.Exec(InitQueryCreate)

	return db
}

func DBCreate(models.Model) {
	var id int
	q := `INSERT INTO orders 
			(track_number, entry, locale, internal_signature,
			customer_id, delivery_service, shardkey, sm_id, date_created,
			oof_shard) 
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id`

	r.db.QueryRow(

}