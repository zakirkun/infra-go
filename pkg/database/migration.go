package database

func NewMigration(models ...interface{}) error {
	db, _ := DB.OpenDB()

	err := db.AutoMigrate(models...)

	return err
}
