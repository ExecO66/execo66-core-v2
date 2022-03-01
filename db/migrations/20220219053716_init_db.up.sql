CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TYPE IF EXISTS user_status_enum;
DROP TYPE IF EXISTS login_provider_enum;
DROP TYPE IF EXISTS exec_lang_enum;

CREATE TYPE user_status_enum as ENUM ('student', 'teacher');
CREATE TYPE login_provider_enum as ENUM ('google');
CREATE TYPE exec_lang_enum as ENUM ('java');

CREATE TABLE IF NOT EXISTS "user"(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    username text NOT NULL,
    email text NOT NULL,
    user_status user_status_enum NOT NULL,
    provider login_provider_enum NOT NULL,
    provider_id text NOT NULL,
    profile_picture text NOT NULL
);

CREATE TABLE IF NOT EXISTS assignment(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id uuid REFERENCES "user"(id) NOT NULL,
    due_date timestamptz NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    code_locations text[] NOT NULl,
    test_inputs text NOT NULL,
    lang exec_lang_enum NOT NULL,
    available_until timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS assignment_exec(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    assignment_id uuid REFERENCES assignment(id) NOT NULL,
    output text NOT NULL,
    test_input text NOT NULL
);

CREATE TABLE IF NOT EXISTS student_assignment(
    student_id uuid REFERENCES "user"(id) NOT NULL,
    assignment_id uuid REFERENCES assignment(id) NOT NULL,
    PRIMARY KEY (student_id, assignment_id)
);

CREATE TABLE IF NOT EXISTS submission(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    submitter_id uuid REFERENCES "user"(id) NOT NULL,
    assignment_id uuid REFERENCES assignment(id) NOT NULL,
    submit_date timestamptz NOT NULL,
    code_locations text[] NOT NULL
);

CREATE TABLE IF NOT EXISTS submission_exec(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    submission_id uuid REFERENCES submission(id) NOT NULL,
    output text NOT NULL,
    test_input text NOT NULL
);