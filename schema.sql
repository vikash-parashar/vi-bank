-- -- Drop tables if they exist
-- DROP TABLE IF EXISTS Customers CASCADE;
-- DROP TABLE IF EXISTS Addresses CASCADE;
-- DROP TABLE IF EXISTS KYCDetails CASCADE;
-- DROP TABLE IF EXISTS Nominee CASCADE;
-- DROP TABLE IF EXISTS Transactions CASCADE;
-- DROP TABLE IF EXISTS Contact CASCADE;
-- DROP TABLE IF EXISTS Account CASCADE;-- Create the Customers table


-- CREATE TABLE Customers (
--     customer_id UUID PRIMARY KEY,
--     account_id UUID REFERENCES Accounts(account_id) UNIQUE,
--     kyc_details_id UUID REFERENCES KYCDetails(kyc_details_id) UNIQUE,
--     contact_id UUID REFERENCES Contact(contact_id) UNIQUE
-- );

-- -- Create the Address table
-- CREATE TABLE Address (
--     address_id UUID PRIMARY KEY,
--     customer_id UUID REFERENCES Customers(customer_id),
--     address_type VARCHAR(50),  -- Adjust the field size as needed
--     street VARCHAR(100),  -- Adjust the field size as needed
--     city VARCHAR(50),  -- Adjust the field size as needed
--     state VARCHAR(50),  -- Adjust the field size as needed
--     pin_code VARCHAR(20),  -- Adjust the field size as needed
--     country VARCHAR(50),  -- Adjust the field size as needed
--     created_at TIMESTAMP,
--     updated_at TIMESTAMP
-- );


-- -- Create the KYCDetails table
-- CREATE TABLE KYCDetails (
--     kyc_id UUID PRIMARY KEY,
--     customer_id UUID REFERENCES Customers(customer_id),
--     first_name VARCHAR(50),  -- Adjust the field size as needed
--     last_name VARCHAR(50),  -- Adjust the field size as needed
--     date_of_birth DATE,
--     gender VARCHAR(10),  -- Adjust the field size as needed
--     created_at TIMESTAMP,
--     updated_at TIMESTAMP
-- );

-- -- Create the Document table
-- CREATE TABLE Document (
--     document_name VARCHAR(100),  -- Adjust the field size as needed
--     document_number VARCHAR(100),  -- Adjust the field size as needed
--     expire_on DATE
-- );

-- -- Create the Nominee table
-- CREATE TABLE Nominee (
--     nominee_id UUID PRIMARY KEY,
--     nominee_first_name VARCHAR(50),  -- Adjust the field size as needed
--     nominee_last_name VARCHAR(50),  -- Adjust the field size as needed
--     nominee_date_of_birth DATE,
--     relation_with_nominee VARCHAR(50),  -- Adjust the field size as needed
--     created_at TIMESTAMP,
--     updated_at TIMESTAMP
-- );


-- -- Create the Transaction table
-- CREATE TABLE Transaction (
--     transaction_id UUID PRIMARY KEY,
--     customer_id UUID REFERENCES Customers(customer_id),
--     source_account_id UUID REFERENCES Account(account_id),
--     destination_account_id UUID REFERENCES Account(account_id),
--     transaction_amount NUMERIC(15, 2),  -- Assuming you want a fixed-point decimal for transaction amount
--     timestamp TIMESTAMP,
--     transaction_status VARCHAR(20)  -- Assuming you want to store the status as a string
-- );


-- -- Create the Contact table
-- CREATE TABLE Contact (
--     contact_id UUID PRIMARY KEY,
--     customer_id UUID REFERENCES Customers(customer_id),
--     phone VARCHAR(20),  -- Assuming you want to store phone numbers as strings
--     email VARCHAR(100),  -- Assuming you want to store email addresses
--     created_at TIMESTAMP,
--     updated_at TIMESTAMP,
--     deleted_at TIMESTAMP  -- Use this field if you want to implement soft deletion
-- );


-- -- Create the Account table
-- CREATE TABLE Account (
--     account_id UUID PRIMARY KEY,
--     customer_id UUID REFERENCES Customers(customer_id) UNIQUE,
--     balance NUMERIC(15, 2), -- Assuming you want a fixed-point decimal for balance
--     currency VARCHAR(3), -- Currency code, e.g., "USD", "EUR", etc.
--     account_type VARCHAR(50), -- Type of account, e.g., "Savings", "Checking", etc.
--     nominee_id UUID REFERENCES Nominee(nominee_id),
--     created_at TIMESTAMP,
--     updated_at TIMESTAMP
-- );

-- --


-- Drop tables if they exist
DROP TABLE IF EXISTS Transactions CASCADE;
DROP TABLE IF EXISTS Contact CASCADE;
DROP TABLE IF EXISTS Nominee CASCADE;
DROP TABLE IF EXISTS KYCDetails CASCADE;
DROP TABLE IF EXISTS Addresses CASCADE;
DROP TABLE IF EXISTS Account CASCADE;
DROP TABLE IF EXISTS Customers CASCADE;

-- Create the Customers table
CREATE TABLE Customers (
    customer_id UUID PRIMARY KEY,
    account_id UUID UNIQUE,
    kyc_details_id UUID UNIQUE,
    contact_id UUID UNIQUE
);

-- Create the Addresses table
CREATE TABLE Addresses (
    address_id UUID PRIMARY KEY,
    customer_id UUID REFERENCES Customers(customer_id)
);

-- Create the KYCDetails table
CREATE TABLE KYCDetails (
    kyc_details_id UUID PRIMARY KEY,
    customer_id UUID REFERENCES Customers(customer_id),
    nominee_id UUID
);

-- Create the Nominee table
CREATE TABLE Nominee (
    nominee_id UUID PRIMARY KEY,
    nominee_first_name VARCHAR(50),
    nominee_last_name VARCHAR(50),
    nominee_date_of_birth DATE,
    relation_with_nominee VARCHAR(50)
);

-- Create the Transactions table
CREATE TABLE Transactions (
    transaction_id UUID PRIMARY KEY,
    customer_id UUID REFERENCES Customers(customer_id),
    source_account_id UUID,
    destination_account_id UUID,
    transaction_amount NUMERIC(15, 2),
    timestamp TIMESTAMP,
    transaction_status VARCHAR(20)
);

-- Create the Contact table
CREATE TABLE Contact (
    contact_id UUID PRIMARY KEY,
    customer_id UUID REFERENCES Customers(customer_id),
    phone VARCHAR(20),
    email VARCHAR(100),
    deleted_at TIMESTAMP
);

-- Create the Account table
CREATE TABLE Account (
    account_id UUID PRIMARY KEY,
    customer_id UUID,
    balance NUMERIC(15, 2),
    currency VARCHAR(3),
    account_type VARCHAR(50),
    nominee_id UUID,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
