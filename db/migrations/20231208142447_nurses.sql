-- migrate:up
CREATE TABLE nurses (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS nurses;

