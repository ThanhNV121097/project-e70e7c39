CREATE TABLE greetings (
  id boolean PRIMARY KEY CHECK (id),
  text text NOT NULL CHECK (btrim(text) <> ''),
  updated_at timestamptz NOT NULL DEFAULT now()
);

INSERT INTO greetings (id, text) VALUES (TRUE, 'Hello, World!') ON CONFLICT (id) DO NOTHING;
