-- Create database
CREATE DATABASE IF NOT EXISTS example_database;

-- Use the database
USE example_database;

-- Create table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);
