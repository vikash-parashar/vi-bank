DROP DATABASE IF EXISTS vibank;
CREATE DATABASE IF NOT EXISTS vibank;
DROP TABLE IF EXISTS users, kyc_details, addresses, contacts, customers, accounts, transactions, nominees;


CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    user_type VARCHAR(20) NOT NULL,
    customer_id UUID,
    user_email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS kyc_details (
    id UUID PRIMARY KEY,
    customer_id UUID,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    date_of_birth DATE NOT NULL,
    gender VARCHAR(10) NOT NULL,
    aadhaar_number VARCHAR(16) NOT NULL,
    pan_card_number VARCHAR(10) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS addresses (
    id UUID PRIMARY KEY,
    customer_id UUID,
    street VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    pin_code VARCHAR(10) NOT NULL,
    country VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS contacts (
    id UUID PRIMARY KEY,
    customer_id UUID,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY,
    kyc_id UUID,
    address_id UUID,
    contact_id UUID,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY,
    customer_id UUID,
    balance NUMERIC(15, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    account_type VARCHAR(20) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY,
    source_account_id UUID,
    destination_account_id UUID,
    transaction_amount NUMERIC(15, 2) NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    transaction_status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS nominees (
    id UUID PRIMARY KEY,
    nominee_first_name VARCHAR(255) NOT NULL,
    nominee_last_name VARCHAR(255) NOT NULL,
    nominee_date_of_birth DATE NOT NULL,
    relation_with_nominee VARCHAR(20) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
