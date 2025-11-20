-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pull_requests (
    pull_request_id   VARCHAR(120) PRIMARY KEY,
    pull_request_name VARCHAR(255) NOT NULL,
    author_id         VARCHAR(120) NOT NULL,
    status            VARCHAR(16)  NOT NULL,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    merged_at         TIMESTAMPTZ,
    CONSTRAINT fk_pr_author FOREIGN KEY (author_id)
        REFERENCES users(user_id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_pr_author ON pull_requests(author_id);
CREATE INDEX IF NOT EXISTS idx_pr_status ON pull_requests(status);
CREATE INDEX IF NOT EXISTS idx_pr_created_at ON pull_requests(created_at);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
