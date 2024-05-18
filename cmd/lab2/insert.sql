INSERT INTO student (name, surname, index, email, stationary)
VALUES
    ('Alice', 'Smith', '12345', 'alice@example.com', true),
    ('Bob', 'Johnson', '67890', 'bob@example.com', false),
    ('Charlie', 'Brown', '24680', 'charlie@example.com', true),
    ('David', 'Lee', '13579', 'david@example.com', false),
    ('Eva', 'Garcia', '98765', 'eva@example.com', true),
    ('Frank', 'Miller', '54321', 'frank@example.com', false),
    ('Grace', 'Wang', '76543', 'grace@example.com', true),
    ('Henry', 'Chen', '98767', 'henry@example.com', false),
    ('Ivy', 'Nguyen', '23456', 'ivy@example.com', true),
    ('Jack', 'Taylor', '34567', 'jack@example.com', false);

INSERT INTO project (name, description, start_date, turnout_date)
VALUES
    ('Project A', 'Description for Project A', '2024-05-01', '2024-06-15'),
    ('Project B', 'Description for Project B', '2024-05-10', NULL),
    ('Project C', 'Description for Project C', '2024-05-15', NULL),
    ('Project D', 'Description for Project D', '2024-05-20', NULL),
    ('Project E', 'Description for Project E', '2024-05-25', NULL),
    ('Project F', 'Description for Project F', '2024-06-01', NULL),
    ('Project G', 'Description for Project G', '2024-06-05', NULL),
    ('Project H', 'Description for Project H', '2024-06-10', NULL),
    ('Project I', 'Description for Project I', '2024-06-15', NULL),
    ('Project J', 'Description for Project J', '2024-06-20', NULL);

INSERT INTO task (project_id, name, description, ord, creation_date)
VALUES
    (1, 'Task 1 for Project A', 'Description for Task 1', 1, '2024-05-01'),
    (1, 'Task 2 for Project A', 'Description for Task 2', 2, '2024-05-02'),
    (1, 'Task 3 for Project A', 'Description for Task 3', 3, '2024-05-03'),
    (2, 'Task 1 for Project B', 'Description for Task 1', 1, '2024-05-10'),
    (2, 'Task 2 for Project B', 'Description for Task 2', 2, '2024-05-11'),
    (3, 'Task 1 for Project C', 'Description for Task 1', 1, '2024-05-15'),
    (3, 'Task 2 for Project C', 'Description for Task 2', 2, '2024-05-16'),
    (4, 'Task 1 for Project D', 'Description for Task 1', 1, '2024-05-20'),
    (4, 'Task 2 for Project D', 'Description for Task 2', 2, '2024-05-21'),
    (5, 'Task 1 for Project E', 'Description for Task 1', 1, '2024-05-25');

INSERT INTO student_project (project_id, student_id)
VALUES
    (1, 1), -- Alice is working on Project A
    (1, 2), -- Bob is also working on Project A
    (1, 3), -- Charlie is working on Project A
    (2, 4), -- David is working on Project B
    (2, 5), -- Eva is also working on Project B
    (3, 6), -- Frank is working on Project C
    (3, 7), -- Grace is also working on Project C
    (4, 8), -- Henry is working on Project D
    (4, 9), -- Ivy is also working on Project D
    (5, 10); -- Jack is working on Project E
