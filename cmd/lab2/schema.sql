CREATE TABLE student
(
    id         serial PRIMARY KEY,
    name       varchar(50)  NOT NULL,
    surname    varchar(50)  NOT NULL,
    index      varchar(20)  NOT NULL UNIQUE,
    email      varchar(100) NOT NULL UNIQUE,
    stationary boolean      NOT NULL
);

CREATE TABLE project
(
    id           serial PRIMARY KEY,
    name         varchar(100) NOT NULL,
    description  varchar(1000),
    start_date   timestamp    NOT NULL,
    turnout_date date
);

CREATE TABLE task
(
    id            serial PRIMARY KEY,
    project_id    integer REFERENCES project (id) NOT NULL,
    name          varchar(50)                     NOT NULL,
    description   varchar(1000),
    ord           integer,
    creation_date timestamp                       NOT NULL
);

CREATE TABLE student_project
(
    id         serial PRIMARY KEY,
    project_id integer REFERENCES project (id) NOT NULL,
    student_id integer REFERENCES student (id) NOT NULL
);