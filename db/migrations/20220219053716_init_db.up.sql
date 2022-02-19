CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TYPE IF EXISTS user_status_enum;
DROP TYPE IF EXISTS login_provider_enum;

CREATE TYPE user_status_enum as ENUM ('student', 'teacher');
CREATE TYPE login_provider_enum as ENUM ('google');

CREATE TABLE IF NOT EXISTS public.user(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    username text NOT NULL,
    email text NOT NULL,
    user_status user_status_enum NOT NULL,
    provider login_provider_enum NOT NULL,
    provider_id text NOT NULL
);

CREATE TABLE IF NOT EXISTS public.assignment(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id uuid REFERENCES public.user(id) NOT NULL,
    due_date timestamptz NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    code_locations text[] NOT NULl,
    test_inputs text NOT NULL
);

CREATE TABLE IF NOT EXISTS public.assignment_exec(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    assignment_id uuid REFERENCES public.assignment(id) NOT NULL,
    output text NOT NULL,
    test_input text NOT NULL
);

CREATE TABLE IF NOT EXISTS public.student_assignment(
    student_id uuid REFERENCES public.user(id) NOT NULL,
    assignment_id uuid REFERENCES public.assignment(id) NOT NULL,
    PRIMARY KEY (student_id, assignment_id)
);

CREATE TABLE IF NOT EXISTS public.submission(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    submitter_id uuid REFERENCES public.user(id) NOT NULL,
    assignment_id uuid REFERENCES public.assignment(id) NOT NULL,
    submit_date timestamptz NOT NULL,
    code_locations text[] NOT NULL
);

CREATE TABLE IF NOT EXISTS public.submission_exec(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    submission_id uuid REFERENCES public.submission(id) NOT NULL,
    output text NOT NULL,
    test_input text NOT NULL
);