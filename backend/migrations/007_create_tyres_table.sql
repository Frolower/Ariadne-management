CREATE TABLE parts (
    tire_id VARCHAR(255) NOT NULL PRIMARY KEY,
    tread_remaining INT NOT NULL,
    compound VARCHAR(255) NOT NULL,
    chassis_number VARCHAR(255) REFERENCES cars (chassis_number),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);