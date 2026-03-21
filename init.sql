CREATE TABLE
    IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email TEXT NOT NULL,
        name TEXT NOT NULL
    )
INSERT INTO
    users
VALUES
    ('cameron', 'cameron@gmail.com'),
    ('rachel', 'rachel@gmail.com'),
    ('joh', 'joh@gmail.com'),
    ('clav', 'clav@gmail.com'),
    ('fratleader', 'fratleader@gmail.com'),
    ('rizzler', 'rizzler@gmail.com');