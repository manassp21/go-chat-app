package database

import(
	"go-chat-app/pkg/config"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error{
	var err error
	DB, err := sql.Open("sqlite3", config.AppConfig.DBPath)
	if err!=nil{
		return err 
	}

	err = DB.Ping()
	if err!=nil{
		return err 
	}

	log.Println("database connection secured")

	err = createTables(DB)
	if err != nil{
		return err
	}
	return nil
}

func createTables(db *sql.DB) error{
	usersSQL := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,	
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP 
	);
	`

	_, err := db.Exec(usersSQL)
	if err!=nil{
		return err 
	}

	messagesSQL := `
	CREATE TABLE IF NOT EXISTS messages(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		username TEXT UNIQUE NOT NULL,
		content TEXT NOT NULL,
		room_id TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`

	_, err = db.Exec(messagesSQL)
	if err!=nil{
		return err 
	}

	log.Println("tables created/verified")
	return nil
}

func CloseDB(db *sql.DB) error{
	return db.Close()
}