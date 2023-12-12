CREATE TABLE IF NOT EXISTS vote (
    vote_id SERIAL PRIMARY KEY,
    vote    INTEGER NOT NULL,
    created_user_id UUID NOT NULL,
    created_at TIMESTAMP
);
CREATE INDEX idx_created_user_id ON vote (created_user_id);
CREATE INDEX idx_created_at ON vote (created_at);

CREATE TABLE IF NOT EXISTS user_votes (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    vote_id BIGINT
);
CREATE INDEX idx_user_id ON user_votes (user_id);