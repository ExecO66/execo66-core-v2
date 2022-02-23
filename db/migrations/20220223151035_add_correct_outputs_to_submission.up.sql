ALTER TABLE submission 
    ADD COLUMN tests_run integer NOT NULL DEFAULT 0,
    ADD COLUMN correct_outputs integer NOT NULL DEFAULT 0;