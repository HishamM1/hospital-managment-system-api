-- migrate:up
CREATE TABLE treatments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    type VARCHAR(255) NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS treatments;

