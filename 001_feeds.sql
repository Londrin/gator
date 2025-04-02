-- +goose Up
CREATE TABLE feeds (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	name TEXT UNIQUE NOT NULL,
	url TEXT UNIQUE NOT NULL,
	user_id INTEGER REFERENCES 
);

-- +goose Down
DROP TABLE feeds;
