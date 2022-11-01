-- Set Environment Variables
\set userdb `echo "$APP_DB_USER"`
\set passdb `echo "$APP_DB_PASS"`
\set dbname `echo "$APP_DB_NAME"`

-- Create Database
CREATE DATABASE :"dbname";

-- Create User and Grant Privileges
CREATE USER :"userdb" WITH ENCRYPTED PASSWORD :'passdb';
GRANT ALL PRIVILEGES ON DATABASE :"dbname" TO :"userdb";
\connect :"dbname" :"userdb"

-- Create Extension and Install into Database
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create User Table
CREATE TABLE "user" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(62) NOT NULL,
    username VARCHAR(50) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    password_salt VARCHAR(32) NOT NULL,
    mobile VARCHAR(15),
    address VARCHAR(255),
    role_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create the Function to Update Timestamp at "updated_at" Column
CREATE OR REPLACE FUNCTION update_user_timestamp_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create the Trigger to Update Timestamp at "updated_at" Column
CREATE TRIGGER update_user_task_updated_at 
    BEFORE UPDATE
    ON "user"
    FOR EACH ROW 
EXECUTE PROCEDURE update_user_timestamp_updated_at();

-- Create Role Table
CREATE TABLE "role" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    description VARCHAR(255) NOT NULL
);

-- Insert Role Data
INSERT INTO "role" (name, description) VALUES 
    ('Admin', 'System Administrator'),
    ('Buyer', 'Client Logged-In as Buyer'),
    ('Seller', 'Client Logged-In as Seller');

-- Insert Admin User
INSERT INTO "user" (first_name, last_name, email, username, password_hash, password_salt, role_id) VALUES 
    ('Admin', 'Admin', '', 'admin', 'd3b6c2a2b1c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0u1v2w3x4y5z6', '1234567890', 1);