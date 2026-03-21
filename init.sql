-- Users table
CREATE TABLE
    IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL
    );

-- Todos table (linked to users)
CREATE TABLE
    IF NOT EXISTS todos (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL,
        title TEXT NOT NULL,
        completed BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        -- foreign key relationship
        CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );

-- Insert users
INSERT INTO
    users (name, email)
VALUES
    ('Alice', 'alice@test.com'),
    ('Bob', 'bob@test.com'),
    ('Charlie', 'charlie@test.com'),
    ('Dave', 'dave@test.com'),
    ('Eve', 'eve@test.com');

-- Insert todos (each user gets some)
INSERT INTO
    todos (user_id, title, completed)
VALUES
    -- Alice (id 1)
    (1, 'Buy groceries', false),
    (1, 'Go to gym', true),
    -- Bob (id 2)
    (2, 'Finish report', false),
    (2, 'Call mom', false),
    -- Charlie (id 3)
    (3, 'Read book', true),
    (3, 'Write blog post', false),
    -- Dave (id 4)
    (4, 'Clean room', false),
    (4, 'Fix bug in code', true),
    -- Eve (id 5)
    (5, 'Plan trip', false),
    (5, 'Book hotel', false);