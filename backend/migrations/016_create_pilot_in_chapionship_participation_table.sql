CREATE TABLE pilotInChampionship (
    document_number INT NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    championship_id INT NOT NULL REFERENCES championships (championship_id),
    chassis_number VARCHAR(255) NOT NULL REFERENCES cars (chassis_number),
    FOREIGN KEY (document_number, first_name, last_name) REFERENCES pilots (document_number, first_name, last_name),
    PRIMARY KEY (document_number, first_name, last_name, chassis_number, championship_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);