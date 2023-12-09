-- migrate:up
CREATE TABLE patient_nurses (
    patient_id INT,
    nurse_id INT,
    PRIMARY KEY (patient_id, nurse_id),
    FOREIGN KEY (patient_id) REFERENCES Patients(id),
    FOREIGN KEY (nurse_id) REFERENCES Nurses(id)
);

-- migrate:down
DROP TABLE IF EXISTS patient_nurses;

