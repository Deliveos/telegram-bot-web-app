CREATE SCHEMA schedules;

CREATE TABLE schedules.groups(
  id SERIAL PRIMARY KEY,
  "name" VARCHAR(256),
  chat_id INTEGER
);

CREATE TABLE schedules.subjects(
  id SERIAL PRIMARY KEY,
  "name" VARCHAR(256),
  teacher VARCHAR(256),
  classrom VARCHAR(20)
);

CREATE TABLE schedules.schedules (
  id SERIAL PRIMARY KEY,
  group_id INTEGER REFERENCES schedules.groups(id),
  subject_id INTEGER REFERENCES schedules.subjects(id),
  "weekday" INTEGER,
  "time" TIME
);

CREATE TABLE schedules.replacements (
  id SERIAL PRIMARY KEY,
  schedule_id INTEGER REFERENCES schedules.schedules(id),
  subject_id INTEGER REFERENCES schedules.subjects(id),
  "datetime" TIMESTAMP
);

CREATE TABLE schedules.notifications (
  id SERIAL PRIMARY KEY,
  chat_id INTEGER,
  "text" INTEGER,
  "datetime" TIMESTAMP
);

