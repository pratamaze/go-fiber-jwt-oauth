package migrations

import "go-fiber-auth/config"

func AddIsAdminColumn() {
	config.DB.Exec("ALTER TABLE users ADD COLUMN IF NOT EXISTS is_admin BOOLEAN DEFAULT FALSE;")
}
