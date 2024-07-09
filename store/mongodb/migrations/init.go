package migrations

func init() {
	// Register the migration with the migration manager
	_ = RegisterMigration("0001_init", up_0001, down_0001)
	_ = RegisterMigration("0002_add_operators_field", up_0002, down_0002)
}