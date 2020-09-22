CREATE TABLE messages (
    id serial PRIMARY KEY,
    message VARCHAR(255)
);

create extension if not exists "uuid-ossp";

CREATE TABLE latency_information (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    latency_ns VARCHAR(255),
    path VARCHAR(255)
);