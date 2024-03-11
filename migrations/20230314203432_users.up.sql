CREATE SCHEMA users;

CREATE TABLE users.users(
  id SERIAL PRIMARY KEY,
  firstname VARCHAR(64),
  lastname VARCHAR(64),
  chatId INTEGER
);

CREATE TABLE users.roles(
  id SERIAL PRIMARY KEY,
  name VARCHAR(256)
);

CREATE TABLE users.users_roles(
  role_id INTEGER,
  user_id INTEGER,
  CONSTRAINT fk_users_roles
    FOREIGN KEY(role_id) REFERENCES users.roles(id) ON DELETE SET NULL,
    FOREIGN KEY(user_id) REFERENCES users.users(id) ON DELETE CASCADE
);

CREATE TABLE users.permissions(
  id SERIAL PRIMARY KEY,
  name VARCHAR(256)
);

CREATE TABLE users.permissions_roles(
  role_id INTEGER,
  permission_id INTEGER,
  CONSTRAINT fk_permissions_roles
    FOREIGN KEY(role_id) REFERENCES users.roles(id) ON DELETE SET NULL,
    FOREIGN KEY(permission_id) REFERENCES users.permissions(id) ON DELETE SET NULL
);