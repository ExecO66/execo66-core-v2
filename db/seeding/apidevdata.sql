GRANT USAGE ON SCHEMA public TO postgres;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO postgres;

SET search_path TO public;

SELECT table_name
FROM information_schema.tables
WHERE table_schema = 'public'
ORDER BY table_name;

BEGIN;

INSERT INTO "user" (id, username, email, user_status, provider, provider_id, profile_picture)
VALUES 
('00000000-0000-0000-0000-000000000001','Student1 Student', 'student1@gmail.com', 'student', 'google', '198a4d', 'https://picsum.photos/200/200'),
('00000000-0000-0000-0000-000000000002','Student2 Student', 'student2@gmail,com', 'student', 'google', '214vaa', 'https://picsum.photos/200/200'),
('00000000-0000-0000-0000-000000000003','Teacher1 Teacher', 'teacher1@gmail,com', 'teacher', 'google', 'g119j1', 'https://picsum.photos/200/200'),
('00000000-0000-0000-0000-000000000004','Teacher2 Teacher', 'teacher2@gmail,com', 'teacher', 'google', 'oiagq1', 'https://picsum.photos/200/200');

INSERT INTO assignment (id, owner_id, due_date, title, description, code_locations, test_inputs) 
VALUES 
('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000003', now() + INTERVAL '7 DAY', 'CS A Lab 1', 'Some description 1', '{"https://example.com/1"}', 'str::1input1;1input2;1input3;', 'java'),
('00000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000003', now() - INTERVAL '1 DAY', 'CS A Lab 2', 'Some description 2', '{"https://example.com/2"}', 'str::2input1;2input2;3input3;', 'java'),
('00000000-0000-0000-0000-000000000003', '00000000-0000-0000-0000-000000000003', now() + INTERVAL '14 DAY', 'CS A Lab 3', 'Some description 3', '{"https://example.com/3"}', 'str::3input1;3input2;3input3;', 'java'),
('00000000-0000-0000-0000-000000000004', '00000000-0000-0000-0000-000000000004', now() - INTERVAL '2 DAY', 'CS A Lab 4', 'Some description 4', '{"https://example.com/4"}', 'str::4input1;4input2;4input3;', 'java');

INSERT INTO assignment_exec (id, assignment_id, output, test_input)
VALUES
('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000001', '1output1', '1input1'),
('00000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000001', '1output2', '1input2'),
('00000000-0000-0000-0000-000000000003', '00000000-0000-0000-0000-000000000001', '1output3', '1input3'),

('00000000-0000-0000-0000-000000000004', '00000000-0000-0000-0000-000000000002', '2output1', '2input1'),
('00000000-0000-0000-0000-000000000005', '00000000-0000-0000-0000-000000000002', '2output2', '2input2'),
('00000000-0000-0000-0000-000000000006', '00000000-0000-0000-0000-000000000002', '2output3', '2input3'),

('00000000-0000-0000-0000-000000000007', '00000000-0000-0000-0000-000000000003', '3output1', '3input1'),
('00000000-0000-0000-0000-000000000008', '00000000-0000-0000-0000-000000000003', '3output2', '3input2'),
('00000000-0000-0000-0000-000000000009', '00000000-0000-0000-0000-000000000003', '3output3', '3input3'),

('00000000-0000-0000-0000-000000000010', '00000000-0000-0000-0000-000000000004', '4output1', '4input1'),
('00000000-0000-0000-0000-000000000011', '00000000-0000-0000-0000-000000000004', '4output2', '4input2'),
('00000000-0000-0000-0000-000000000012', '00000000-0000-0000-0000-000000000004', '4output3', '4input3');

INSERT INTO student_assignment (student_id, assignment_id)
VALUES
('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000002'),
('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000003'),
('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000004'),
('00000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000003');

INSERT INTO submission (id, submitter_id, assignment_id, submit_date, code_locations, tests_run, correct_outputs)
VALUES 
('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000002', now() - INTERVAL '2 DAY', '{"http://example.com/1"}', 3, 3),
('00000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000002', now() - INTERVAL '3 DAY', '{"http://example.com/3"}', 3, 1),
('00000000-0000-0000-0000-000000000003', '00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000003', now(), '{"http://example.com/1"}', 3, 3);

INSERT INTO submission_exec (submission_id, output, test_input)
VALUES
('00000000-0000-0000-0000-000000000001', '2output1', '2input1'),
('00000000-0000-0000-0000-000000000001', '2output2', '2input2'),
('00000000-0000-0000-0000-000000000001', '2output3', '2input3'),

('00000000-0000-0000-0000-000000000002', '2output1', '2input1'),
('00000000-0000-0000-0000-000000000002', '2incorrect2', '2input2'),
('00000000-0000-0000-0000-000000000002', '2incorrect3', '2input3'),

('00000000-0000-0000-0000-000000000003', '3output1', '3input1'),
('00000000-0000-0000-0000-000000000003', '3output2', '3input2'),
('00000000-0000-0000-0000-000000000003', '3output3', '3input3');

COMMIT;