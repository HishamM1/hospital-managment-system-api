-- migrate:up
CREATE TABLE numbers (
    id INT PRIMARY KEY AUTO_INCREMENT,
	patient_id INT NOT NULL,
	patient_number VARCHAR(20) NOT NULL,
	family_number VARCHAR(20) NOT NULL,
	FOREIGN KEY (patient_id) REFERENCES patients(id)
);

-- migrate:down
DROP TABLE IF EXISTS numbers;

