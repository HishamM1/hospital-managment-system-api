-- migrate:up
CREATE TABLE doctors (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS doctors;


