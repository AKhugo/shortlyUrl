-- +goose Up

CREATE TABLE IF NOT EXISTS links (

    -- Primary key
    id SERIAL PRIMARY KEY,

    -- Foreign key
    user_id INT NOT NULL,

    -- Link information
    short_code TEXT NOT NULL,
    original_url TEXT NOT NULL,
    access_limit INT NOT NULL DEFAULT 0,
    domain VARCHAR(255) NOT NULL,


    -- Link status
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_public BOOLEAN NOT NULL DEFAULT FALSE,

    -- Metadata
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Deletion information
    deleted_at TIMESTAMP NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose StatementBegin

-- +goose StatementEnd

-- +goose Down

DROP TABLE IF EXISTS links;

-- +goose StatementBegin
-- +goose StatementEnd
