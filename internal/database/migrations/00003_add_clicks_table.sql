-- +goose Up
CREATE TABLE clicks (
    id SERIAL PRIMARY KEY,

    -- Foreign key
    link_id INT NOT NULL,

    -- Click information
    clicked_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ip_address TEXT,
    user_agent TEXT,
    referer TEXT,
    country_code VARCHAR(2) NOT NULL DEFAULT NULL,

    -- Metadata
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Deletion information
    deleted_at TIMESTAMP NULL,

    FOREIGN KEY (link_id) REFERENCES links(id) ON DELETE CASCADE
);


-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS clicks;

-- +goose StatementBegin
-- +goose StatementEnd
