
func DatabaseMigration() {
	db, err := config.GormDatabaseConn()

	if err != nil {
		fmt.Errorf("Error", err.Error())
	}

	db.AutoMigrate(
		&model.Person{},
		&model.User{},
		&model.Employee{},
		&model.CreditCard{})
}