package migrations

import "gorm.io/gorm"

type Migration struct {
	Name string
	Up   func(*gorm.DB) error
	Down func(*gorm.DB) error
}

var AllMigrations = []Migration{
	{
		Name: "create_table_kriteria",
		Up:   UpKriteria,
		Down: DownKriteria,
	},
	{
		Name: "create_table_depot",
		Up:   UpDepot,
		Down: DownDepot,
	},

	{
		Name: "create_table_users",
		Up:   UpUsers,
		Down: DownUsers,
	},
	{
		Name: "create_table_user_location",
		Up:   UpUserLocation,
		Down: DownUserLocation,
	},
	{
		Name: "create_table_kecamatan",
		Up:   UpKecamatan,
		Down: DownKecamatan,
	},
	{
		Name: "create_table_rating",
		Up:   UpRating,
		Down: DownRating,
	},
	{
		Name: "create_table_login_logs",
		Up:   UpLoginLogs,
		Down: DownLoginLogs,
	},
	// Tambahkan migration lainnya di sini
}
