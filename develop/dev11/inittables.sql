DROP TABLE IF EXISTS events;

CREATE TABLE events (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INTEGER,
    title TEXT,
    descr TEXT,
    e_date DATE
);