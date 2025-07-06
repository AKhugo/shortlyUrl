-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (

    -- Primary key
    id SERIAL PRIMARY KEY,

    -- Identity & Authentication information
    username VARCHAR(255) NOT NULL UNIQUE, -- username is the primary identifier for the user
    email VARCHAR(255) NOT NULL UNIQUE, -- email is the secondary identifier for the user
    password_hash VARCHAR(255) NOT NULL, -- password is the password for the user

    -- Account information
    first_name VARCHAR(255) NOT NULL, -- first name of the user
    last_name VARCHAR(255) NOT NULL, -- last name of the user
    full_name VARCHAR(255) NOT NULL, -- full name of the user
    avatar_url VARCHAR(255) DEFAULT NULL, -- avatar url of the user
    phone_number VARCHAR(255) DEFAULT NULL, -- phone number of the user


    -- Confirmation d'email
    email_verified BOOLEAN DEFAULT FALSE, -- email is verified
    verification_token TEXT DEFAULT NULL, -- token for email verification
    verification_expires TIMESTAMP DEFAULT NULL, -- expiration date for email verification

    -- Account status
    is_active BOOLEAN NOT NULL DEFAULT TRUE, -- account is active
    is_verified BOOLEAN NOT NULL DEFAULT FALSE, -- account is verified

    -- Forgot password
    forgot_password_token TEXT DEFAULT NULL, -- token for forgot password
    forgot_password_expires TIMESTAMP DEFAULT NULL, -- expiration date for forgot password

    -- remember me
    remember_me_token TEXT DEFAULT NULL, -- token for remember me
    remember_me_expires TIMESTAMP DEFAULT NULL, -- expiration date for remember me

    -- Connection information
    last_login_at TIMESTAMP DEFAULT NULL, -- last login date
    failed_login_attempts INT NOT NULL DEFAULT 0, -- number of failed login attempts
    failed_login_at TIMESTAMP DEFAULT NULL, -- last failed login date
    last_failed_login_ip VARCHAR(255) DEFAULT NULL, -- last failed login ip

    -- Deletion information
    deleted_at TIMESTAMP NULL, -- deletion date


    -- metadata
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- creation date
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- last update date
    last_ip TEXT DEFAULT NULL, -- last ip address
    user_agent TEXT DEFAULT NULL -- user agent
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
