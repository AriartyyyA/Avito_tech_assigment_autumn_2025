-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    user_id VARCHAR(120) PRIMARY KEY,
    username VARCHAR(120) NOT NULL,
    team_name VARCHAR(120) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,

    CONSTRAINT fk_users_team FOREIGN KEY (team_name) 
        REFERENCES teams(team_name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_users_team_name ON users(team_name);
CREATE INDEX IF NOT EXISTS idx_users_is_active ON users(is_active);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
