package main

import (
	"Social/cmd/api/internal/db"
	"Social/cmd/api/internal/env"
	"Social/cmd/api/internal/store"
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15min"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	os.LookupEnv("PATH")

	mux := app.mount()
	log.Fatal(app.run(mux))
}
