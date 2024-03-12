CREATE TABLE IF NOT EXISTS "commands" (
    "id" serial PRIMARY KEY,
    "script" varchar NOT NULL,
    "description" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "commands_output" (
    "id" serial PRIMARY KEY,
    "id_command" int NOT NULL,
    "output" varchar,
    "time" timestamp NOT NULL,
    FOREIGN KEY (id_command)  REFERENCES commands (id)
);