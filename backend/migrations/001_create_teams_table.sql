CREATE TABLE teams (
    team_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users (user_id),
    team_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)