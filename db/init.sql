CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS signatures (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    questions_answers JSON NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);