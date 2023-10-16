-- Drop tables if they exist
DROP TABLE IF EXISTS Customers CASCADE;
DROP TABLE IF EXISTS Addresses CASCADE;
DROP TABLE IF EXISTS KYCDetails CASCADE;
DROP TABLE IF EXISTS Nominee CASCADE;
DROP TABLE IF EXISTS Transactions CASCADE;
DROP TABLE IF EXISTS Contact CASCADE;
DROP TABLE IF EXISTS Account CASCADE;-- Create the Customers table


CREATE TABLE Customers (
    customer_id UUID PRIMARY KEY,
    account_id UUID REFERENCES Accounts(account_id) UNIQUE,
    kyc_details_id UUID REFERENCES KYCDetails(kyc_details_id) UNIQUE,
    contact_id UUID REFERENCES Contact(contact_id) UNIQUE
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
    nominee_id UUID REFERENCES Nominee(nominee_id)
);

-- Create the Nominee table
CREATE TABLE Nominee (
    nominee_id UUID PRIMARY KEY
    -- Other nominee fields
);

-- Create the Transactions table
CREATE TABLE Transactions (
    transaction_id UUID PRIMARY KEY,
    customer_id UUID REFERENCES Customers(customer_id)
    -- Other transaction fields
);

-- Create the Contact table
CREATE TABLE Contact (
    contact_id UUID PRIMARY KEY,
    customer_id UUID REFERENCES Customers(customer_id)
    -- Other contact fields
);

-- Create the Account table
CREATE TABLE Account (
    account_id UUID PRIMARY KEY
    -- Other account fields
);
