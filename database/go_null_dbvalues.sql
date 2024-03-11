CREATE DATABASE IF NOT EXISTS go_null_dbvalues;
USE go_null_dbvalues;


CREATE TABLE customer (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT,
    gender VARCHAR(10),
    height FLOAT
);

