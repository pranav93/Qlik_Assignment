CREATE TABLE messages (
    id serial PRIMARY KEY,
    message VARCHAR(255)
);

create extension if not exists "uuid-ossp";

CREATE TABLE latency_information (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    message VARCHAR(255)
);