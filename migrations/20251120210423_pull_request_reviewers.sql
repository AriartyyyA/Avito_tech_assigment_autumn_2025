-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pull_request_reviewers (
    pull_request_id VARCHAR(120) NOT NULL,
    user_id         VARCHAR(120) NOT NULL,
    PRIMARY KEY (pull_request_id, user_id),
    CONSTRAINT fk_prr_pr FOREIGN KEY (pull_request_id)
        REFERENCES pull_requests(pull_request_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_prr_user FOREIGN KEY (user_id)
        REFERENCES users(user_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_pr_reviewers_user ON pull_request_reviewers(user_id);
CREATE INDEX IF NOT EXISTS idx_pr_reviewers_pr ON pull_request_reviewers(pull_request_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
