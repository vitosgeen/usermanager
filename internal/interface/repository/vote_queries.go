package repository

const (
	addVote = `INSERT INTO vote (vote, created_user_id, created_at) VALUES ($1, $2, $3) RETURNING vote_id`

	updateVote = `UPDATE vote SET vote = $1 WHERE vote_id = $2`

	getVoteByID = `SELECT vote_id, vote, created_user_id, created_at FROM vote WHERE vote_id = $1`

	getVotesByUserID = `SELECT vote_id, vote, created_user_id, created_at FROM vote WHERE created_user_id = $1 ORDER BY created_at DESC`

	getVotesByUserIDs = `SELECT vote_id, vote, created_user_id, created_at FROM vote WHERE created_user_id = ANY ($1) ORDER BY created_at DESC`

	deleteVote = `DELETE FROM vote WHERE vote_id = $1`

	addUserVote = `INSERT INTO user_votes (user_id, vote_id) VALUES ($1, $2) RETURNING id`

	getUserVoteByID = `SELECT id, user_id, vote_id FROM user_votes WHERE id = $1`

	getUserVotesByUserID = `SELECT id, user_id, vote_id FROM user_votes WHERE user_id = $1`

	deleteUserVote = `DELETE FROM user_votes WHERE id = $1`
)
