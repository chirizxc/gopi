CREATE TABLE IF NOT EXISTS gifs (
                                    id INT AUTO_INCREMENT PRIMARY KEY,
                                    uuid VARCHAR(255) NOT NULL UNIQUE,
    path VARCHAR(255) NOT NULL,
    alias VARCHAR(255) NOT NULL UNIQUE
    );
INSERT INTO gifs (uuid, path, alias)
VALUES
    ('123e4567-e89b-12d3-a456-426614174000', '/path/to/gif1.gif', 'gif1'),
    ('123e4567-e89b-12d3-a456-426614174001', '/path/to/gif2.gif', 'gif2'),
    ('123e4567-e89b-12d3-a456-426614174002', '/path/to/gif3.gif', 'gif3');

