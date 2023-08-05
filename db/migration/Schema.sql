-- users
CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_at BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  updated_at BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
  username TEXT NOT NULL UNIQUE,
  display_name TEXT NOT NULL DEFAULT '',
  -- role TEXT NOT NULL CHECK (role IN ('ADMIN', 'USER'))
  email TEXT NOT NULL DEFAULT '',
  password_hash TEXT NOT NULL,
  avatar_url TEXT NOT NULL DEFAULT ''
);

CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);

-- moods
CREATE TABLE IF NOT EXISTS moods (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	mood VARCHAR(64) NOT NULL,
  user_id INTEGER NOT NULL,
  description TEXT NOT NULL,
  date DATE NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE INDEX IF NOT EXISTS idx_moods_user_id ON moods (user_id);