-- migrate:up
CREATE TABLE bills (
    id INT PRIMARY KEY AUTO_INCREMENT,
    patient_id INT,
    amount DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES patients(id)
);

-- migrate:down
DROP TABLE IF EXISTS bills;

