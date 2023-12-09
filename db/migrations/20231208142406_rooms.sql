-- migrate:up
CREATE TABLE rooms (
    id INT PRIMARY KEY AUTO_INCREMENT,
    wardboy_id INT,
    type VARCHAR(255) NOT NULL,
    number INT NOT NULL,
    FOREIGN KEY (wardboy_id) REFERENCES WardBoys(id)
);

-- migrate:down
DROP TABLE IF EXISTS rooms;

