-- migrate:up
CREATE TABLE doctors (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    number VARCHAR(20) NOT NULL,
    email VARCHAR(255) NOT NULL,
    specialization VARCHAR(255) NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS doctors;


