-- migrate:up
CREATE TABLE patient_nurses (
    patient_id INT,
    nurse_id INT,
    PRIMARY KEY (patient_id, nurse_id),
    FOREIGN KEY (patient_id) REFERENCES patients(id),
    FOREIGN KEY (nurse_id) REFERENCES nurses(id)
);

-- migrate:down
DROP TABLE IF EXISTS patient_nurses;

