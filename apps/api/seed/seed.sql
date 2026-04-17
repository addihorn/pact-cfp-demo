INSERT INTO todos (title, description, completed)
VALUES
  ('Set up architecture review', 'Walk through 3-tier boundaries with teammates', FALSE),
  ('Create sample API calls', 'Prepare curl examples for demo', FALSE),
  ('Verify dashboard metrics', 'Confirm /metrics exposes request counters', TRUE)
ON CONFLICT DO NOTHING;