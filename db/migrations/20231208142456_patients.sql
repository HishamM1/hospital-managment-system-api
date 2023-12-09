-- migrate:up
CREATE TABLE patients (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    address VARCHAR(255),
    disease VARCHAR(255),
    start_date DATE,
    doctor_id INT,
    room_id INT,
    treatment_id INT,
    FOREIGN KEY (doctor_id) REFERENCES doctors(id),
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (treatment_id) REFERENCES treatments(id)
);

-- migrate:down
DROP TABLE IF EXISTS patients;

