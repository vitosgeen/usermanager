package repository

const (
	addUser = `INSERT INTO users (user_id, nickname, first_name, last_name, email, password, is_public, user_role, created_at, updated_at, deleted_at, login_date, created_by)
    			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	updateUser = `UPDATE users
					SET nickname = $1, first_name = $2, last_name = $3, email = $4, password = $5, is_public = $6, updated_at = $7, login_date = $8
					WHERE user_id = $9`

	updateDeletedAt = `UPDATE users
					SET deleted_at = $1
					WHERE user_id = $2`

	deleteUserFromDb = `DELETE FROM users WHERE user_id = $1`
	getUserByID      = `SELECT user_id, nickname, first_name, last_name, email, password, is_public, user_role, created_at, updated_at, deleted_at, login_date, created_by
							FROM users WHERE user_id=$1`

	getUserByNickname = `SELECT user_id, nickname, first_name, last_name, email, password, is_public, user_role, created_at, updated_at, deleted_at, login_date, created_by
							FROM users
							WHERE nickname = $1`

	getUsers = `SELECT user_id, nickname, first_name, last_name, email, password, is_public, user_role, created_at, updated_at, login_date
  				FROM users
 				ORDER BY created_at, updated_at OFFSET $1 LIMIT $2`
)
