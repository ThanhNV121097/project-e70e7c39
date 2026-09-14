CREATE TABLE greetings (
  id SMALLINT PRIMARY KEY CHECK (id = 1),
  text TEXT NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO greetings (id, text) VALUES (1, 'Hello, World!');
