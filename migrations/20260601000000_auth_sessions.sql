-- +goose Up
-- +goose StatementBegin
CREATE TYPE auth_session_state AS ENUM (
    'active',
    'logouted',
    'expired'
);

CREATE TABLE auth_sessions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    valid_after_at TIMESTAMPTZ NOT NULL,
    valid_before_at TIMESTAMPTZ NOT NULL,
    state auth_session_state NOT NULL,
    refresh_id UUID NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE auth_sessions;
DROP TYPE auth_session_state;
-- +goose StatementEnd
