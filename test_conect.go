
		connStr := "postgres://postgres:anbu@localhost/db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	println("open")
	defer db.Close()
	println("open")
	err = db.Ping()

	if err != nil {
		panic(err)
	}
