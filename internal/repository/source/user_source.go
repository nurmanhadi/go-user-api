package source

const (
	USER_ADD             = "INSERT INTO users(id, name, email, password) VALUES(?,?,?,?)"
	USER_FIND_BY_ID      = "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?"
	USER_FIND_BY_EMAIL   = "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?"
	USER_COUNT_BY_EMAIL  = "SELECT COUNT(*) FROM users WHERE email = ?"
	USER_DELETE_BY_ID    = "DELETE FROM users WHERE id = ?"
	USER_UPDATE_PASSWORD = "UPDATE users SET password = ? WHERE id = ?"
	USER_UPDATE_NAME     = "UPDATE users SET name = ? WHERE id = ?"
)
