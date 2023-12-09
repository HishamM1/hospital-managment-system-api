-- migrate:up
CREATE TABLE numbers (
    id INT PRIMARY KEY AUTO_INCREMENT,
	patient_id INT,
	number VARCHAR(20) NOT NULL,
	description VARCHAR(50) NOT NULL,
	FOREIGN KEY (patient_id) REFERENCES Patients(id)
);

-- migrate:down
DROP TABLE IF EXISTS numbers;

