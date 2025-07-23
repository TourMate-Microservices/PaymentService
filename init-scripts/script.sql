-- Create database
CREATE DATABASE payment-service;

-- Connect to the new database
\c payment-service

-- Create Account table
CREATE TABLE feedbacks (
    feedbackId SERIAL PRIMARY KEY,
    customerId INT NOT NULL,
    tourGuideId INT NOT NULL,
    createdDate TIMESTAMP NOT NULL,
    content TEXT NOT NULL,
    rating INT NOT NULL,
    isDeleted BOOLEAN NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    invoiceId INT NOT NULL
);

CREATE TABLE payments (
    paymentId SERIAL PRIMARY KEY,
    price REAL NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    customer_id INT NOT NULL,
    invoice_id INT
);