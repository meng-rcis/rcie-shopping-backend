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

-- Create Role Table
CREATE TABLE "role" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    description VARCHAR(255)
);

-- Create User Table
CREATE TABLE "user" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(62) NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    password_salt VARCHAR(32) NOT NULL,
    role_id INT NOT NULL REFERENCES "role" (id),
    mobile VARCHAR(15),
    address VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create Shop Status Table
CREATE TABLE "shop_status" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    description VARCHAR(255)
);

-- Create Shop Table
CREATE TABLE "shop" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50) NOT NULL,
    description VARCHAR(255),
    owner_id UUID NOT NULL REFERENCES "user" (id),
    status_id INT NOT NULL REFERENCES "shop_status" (id) DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create Product Status Table
CREATE TABLE "product_status" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    description VARCHAR(255)
);

-- Create Product Table
CREATE TABLE "product" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50) NOT NULL,
    description VARCHAR(255),
    price DECIMAL(10,2) NOT NULL,
    quantity INT NOT NULL,
    shop_id UUID NOT NULL REFERENCES "shop" (id),
    status_id INT NOT NULL REFERENCES "product_status" (id) DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create Order Status Table
CREATE TABLE "order_status" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    description VARCHAR(255)
);

-- Create Order Table
CREATE TABLE "order" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    buyer_id UUID NOT NULL REFERENCES "user" (id),
    product_id UUID NOT NULL REFERENCES "product" (id),
    status_id INT NOT NULL REFERENCES "order_status" (id),
    quantity INT NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create Cart Table
CREATE TABLE "cart" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    buyer_id UUID NOT NULL REFERENCES "user" (id),
    product_id UUID NOT NULL REFERENCES "product" (id),
    quantity INT NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create Wishlist Table
CREATE TABLE "wishlist" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    buyer_id UUID NOT NULL REFERENCES "user" (id),
    product_id UUID NOT NULL REFERENCES "product" (id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create a Function to Update Timestamp at "updated_at" Column
CREATE OR REPLACE FUNCTION update_timestamp_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create a Trigger to Update the Given Tables Timestamp when their Rows are Updated
CREATE TRIGGER update_user_task_updated_at
    BEFORE UPDATE ON "user"
    FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_updated_at();

CREATE TRIGGER update_shop_task_updated_at
    BEFORE UPDATE ON "shop"
    FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_updated_at();

CREATE TRIGGER update_product_task_updated_at
    BEFORE UPDATE ON "product"
    FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_updated_at();

CREATE TRIGGER update_order_task_updated_at
    BEFORE UPDATE ON "order"
    FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_updated_at();

CREATE TRIGGER update_cart_task_updated_at
    BEFORE UPDATE ON "cart"
    FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_updated_at();

-- Insert Role Data
INSERT INTO "role" (name, description) VALUES 
    ('Buyer', 'Client Logged-In as Buyer'),
    ('Seller', 'Client Logged-In as Seller'),
    ('Admin', 'System Administrator');

-- Insert Shop Status Data
INSERT INTO "shop_status" (name, description) VALUES 
    ('Pending', 'Shop is Pending'),
    ('Approved', 'Shop is Approved'),
    ('Rejected', 'Shop is Rejected'),
    ('Suspended', 'Shop is Suspended'),
    ('Banned', 'Shop is Banned');

-- Insert Product Status Data
INSERT INTO "product_status" (name, description) VALUES 
    ('Shown', 'Product is Shown'),
    ('Hidden', 'Product is Hidden');

-- Insert Order Status Data
INSERT INTO "order_status" (name, description) VALUES 
    ('Pending', 'Order is Pending'),
    ('Processing', 'Order is Processing'),
    ('Shipped', 'Order is Shipped'),
    ('Delivered', 'Order is Delivered'),
    ('Cancelled', 'Order is Cancelled');

-- Insert User
INSERT INTO "user" (first_name, last_name, email, username, password_hash, password_salt, role_id) VALUES 
    ('Admin', 'Admin', '', 'admin', 'd5503b08cca52c56cfb12db044a4891a44a5929a52f5ea6c4acf7d1c9c792b83', '1234567890', 3),
    ('John', 'Doe', 'john@outlook.com', 'johndoe', '5a66d38defd48f6c1ccf3229a1ae37b1e096cb67bab14ebfceb832b5d71115fa', '1234567890', 1),
    ('Mary', 'Jane', 'mary@mail.com', 'mary001', '22c3ac3de5dbc7c1b22b4b1f773abd8c56829218e571398aab676135098bd7d2', '1234567890', 2);
