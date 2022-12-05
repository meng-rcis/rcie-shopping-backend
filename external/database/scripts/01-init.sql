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
    mobile VARCHAR(20),
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
    price DECIMAL(10,2) NOT NULL CHECK (price > 0),
    quantity INT NOT NULL CHECK (quantity >= 0),
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
    owner_id UUID NOT NULL REFERENCES "user" (id),
    product_id UUID NOT NULL REFERENCES "product" (id),
    status_id INT NOT NULL REFERENCES "order_status" (id) DEFAULT 1,
    quantity INT NOT NULL CHECK (quantity >= 0),
    total_price DECIMAL(10,2) NOT NULL CHECK (total_price > 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create Cart Table
CREATE TABLE "cart" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id UUID NOT NULL REFERENCES "user" (id),
    product_id UUID NOT NULL REFERENCES "product" (id),
    quantity INT NOT NULL CHECK (quantity >= 0),
    total_price DECIMAL(10,2) NOT NULL CHECK (total_price > 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create Wishlist Table
CREATE TABLE "wishlist" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id UUID NOT NULL REFERENCES "user" (id),
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
INSERT INTO "user" (first_name, last_name, email, username, password_hash, password_salt, role_id, mobile, address) VALUES 
    ('Admin', 'Admin', '', 'admin', 'd5503b08cca52c56cfb12db044a4891a44a5929a52f5ea6c4acf7d1c9c792b83', 'tOvyVv6VNs', 3, NULL, NULL),
    ('John', 'Doe', 'john@outlook.com', 'johndoe', 'bdf281208b843d12f1b322674536f5b75533bd6b52a60f969752597284e96267', 'mT8kgmRfep', 1, '+66811111111', 'Bangkok'),
    ('Jane', 'Catterin', 'jane@outlook.com', 'janecat', 'd22ccae83a060c6379c7d9d0f07dd3a05e9b96bc9ec1e7ae1a6a9d946df739a3', 'p5U1fGte3y', 1, NULL, 'Nonthaburi'),
    ('Mary', 'Jane', 'mary@mail.com', 'mary001', 'c26507b2360af67bb3e4c1158ba108e76ed665d11da5a00b35ad737c9d562ba8', 'xNjc22n5kY', 2, '+66822222222', 'Bangkok'),
    ('Kate', 'Smith', 'kate@hotmail.com', 'katecha', '00f3deb28fcc9dfc2fd02725503b0a5d6dcee323705294784b22b591eb5a925b', 'ljUy3RtWer', 2, '+66833333333', NULL);

-- Insert Shop
INSERT INTO "shop" (name, description, owner_id, status_id) VALUES 
    ('Mary Shop', 
    'Lorem Ipsum is simply dummy text of the printing and typesetting industry.',
    (SELECT id FROM "user" WHERE username = 'mary001'),
    2),
    ('SmithPlay', 
    'Contrary to popular belief, Lorem Ipsum is not simply random text.',
    (SELECT id FROM "user" WHERE username = 'katecha'),
    2);


-- Insert Product
INSERT INTO "product" (name, description, price, quantity, shop_id, status_id) VALUES 
    ('Mary Product 1', 'Mary Product 1 Description', 10.00, 10, (SELECT id FROM "shop" WHERE name = 'Mary Shop'), 1),
    ('Mary Product 2', 'Mary Product 2 Description', 200.00, 20, (SELECT id FROM "shop" WHERE name = 'Mary Shop'), 1),
    ('Mary Product 3', 'Mary Product 3 Description', 320.00, 50, (SELECT id FROM "shop" WHERE name = 'Mary Shop'), 1),
    ('Mary Product 4', 'Mary Product 4 Description', 550.00, 50, (SELECT id FROM "shop" WHERE name = 'Mary Shop'), 2),
    ('Mary Product 5', 'Mary Product 5 Description', 600.00, 50, (SELECT id FROM "shop" WHERE name = 'Mary Shop'), 2),
    ('Smith Product 1', 'Smith Product 1 Description', 300.00, 30, (SELECT id FROM "shop" WHERE name = 'SmithPlay'), 1),
    ('Smith Product 2', 'Smith Product 2 Description', 350.00, 10, (SELECT id FROM "shop" WHERE name = 'SmithPlay'), 1),
    ('Smith Product 3', 'Smith Product 3 Description', 400.00, 40, (SELECT id FROM "shop" WHERE name = 'SmithPlay'), 2),
    ('Smith Product 4', 'Smith Product 4 Description', 300.00, 30, (SELECT id FROM "shop" WHERE name = 'SmithPlay'), 2);

-- Insert Order
INSERT INTO "order" (owner_id, product_id, quantity, total_price, status_id) VALUES 
    ((SELECT id FROM "user" WHERE username = 'johndoe'), (SELECT id FROM "product" WHERE name = 'Mary Product 1'), 4, 40.00, 4),
    ((SELECT id FROM "user" WHERE username = 'johndoe'), (SELECT id FROM "product" WHERE name = 'Smith Product 1'), 1, 300.00, 1),
    ((SELECT id FROM "user" WHERE username = 'janecat'), (SELECT id FROM "product" WHERE name = 'Mary Product 5'), 2, 1200.00, 2);

-- Insert Cart
INSERT INTO "cart" (owner_id, product_id, quantity, total_price) VALUES 
    ((SELECT id FROM "user" WHERE username = 'johndoe'), (SELECT id FROM "product" WHERE name = 'Mary Product 2'), 4, 800.00),
    ((SELECT id FROM "user" WHERE username = 'johndoe'), (SELECT id FROM "product" WHERE name = 'Mary Product 3'), 1, 320.00),
    ((SELECT id FROM "user" WHERE username = 'janecat'), (SELECT id FROM "product" WHERE name = 'Smith Product 3'), 4, 1600.00);

-- Insert Wishlist
INSERT INTO "wishlist" (owner_id, product_id) VALUES 
    ((SELECT id FROM "user" WHERE username = 'johndoe'), (SELECT id FROM "product" WHERE name = 'Mary Product 2')),
    ((SELECT id FROM "user" WHERE username = 'johndoe'), (SELECT id FROM "product" WHERE name = 'Mary Product 5')),
    ((SELECT id FROM "user" WHERE username = 'janecat'), (SELECT id FROM "product" WHERE name = 'Mary Product 2'));
