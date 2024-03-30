-- initalize the database
DROP DATABASE IF EXISTS irb_tracker;
CREATE DATABASE irb_tracker;
USE irb_tracker;

-- department table
CREATE TABLE department (
	id int AUTO_INCREMENT,
	name varchar(100) NOT NULL,
	PRIMARY KEY (id)
);

-- users table
CREATE TABLE users (
	id int AUTO_INCREMENT,
	fname varchar(50) NOT NULL,
	lname varchar(50) NOT NULL,
	email varchar(85) NOT NULL UNIQUE,
	passwd varchar(85) NOT NULL,
	dept int,
	PRIMARY KEY (id),
  foreign key (dept) references department(id)
);

-- majors table
CREATE TABLE majors (
	id int AUTO_INCREMENT NOT NULL UNIQUE,
	name varchar(50) NOT NULL,
	PRIMARY KEY (id)
);

-- student table
CREATE TABLE student (
	id int,
	major int,
	year_group int NOT NULL,
	student_id varchar(8) NOT NULL UNIQUE,
	PRIMARY KEY (id),
  FOREIGN KEY (id) REFERENCES users(id),
  FOREIGN KEY (major) REFERENCES majors(id)
);

-- faculty table
CREATE TABLE faculty (
	id int,
	PRIMARY KEY (id),
  FOREIGN KEY (id) REFERENCES users(id)
);

-- workflow table
CREATE TABLE workflow (
	id int AUTO_INCREMENT,
	name varchar(50) NOT NULL UNIQUE,
	stage_1 varchar(50) NOT NULL,
	stage_2 varchar(50) NOT NULL,
	stage_3 varchar(50) NOT NULL,
	stage_4 varchar(50),
	stage_5 varchar(50),
	PRIMARY KEY (id)
);

-- status table
CREATE TABLE status (
	id int AUTO_INCREMENT,
	sname varchar(30) NOT NULL UNIQUE,
	PRIMARY KEY (id)
);

-- projects table
CREATE TABLE projects (
	id int AUTO_INCREMENT,
	brief varchar(500) NOT NULL,
	title varchar(100) NOT NULL UNIQUE,
	associated_documents varchar(200) NOT NULL,
	principal_investigator int,
	dept int,
	supervisor int,
	start_date date NOT NULL,
	end_date date NOT NULL,
	other_investigators mediumtext,
	purpose mediumtext NOT NULL,
	results_dissemination mediumtext NOT NULL,
	participants_count int NOT NULL,
	participants_type mediumtext NOT NULL,
	recuitment_method mediumtext NOT NULL,
	date_submitted date NOT NULL,
	date_approved date,
	workflow_id int,
	status_id int,
	PRIMARY KEY (id),
  FOREIGN KEY (principal_investigator) REFERENCES users(id),
  FOREIGN KEY (dept) REFERENCES department(id),
  FOREIGN KEY (supervisor) REFERENCES faculty(id),
  FOREIGN KEY (workflow_id) REFERENCES workflow(id),
  FOREIGN KEY (status_id) REFERENCES status(id)
);

-- populate the department table
INSERT INTO department(name)
VALUES
  ("Humanities & Social Sciences"),
  ("Economics & Business Administration"),
  ("Computer Science & Information Systems"),
  ("Engineering");

-- populate the majors table
INSERT INTO majors(name)
VALUES
  ("Computer Science"),
  ("Computer Engineering"),
  ("Electrical and Electronic Engineering"),
  ("Mechanical Engineering"),
  ("Mechatronics Engineering"),
  ("Management Information Systems"),
  ("Business Administration"),
  ("Economics");

-- add the default workflow
INSERT INTO workflow (name, stage_1, stage_2, stage_3, stage_4, stage_5)
VALUES
  ("default", "stage 1", "stage 2", "stage 3", "stage 4", "stage 5");

-- populate the status
INSERT INTO status(sname)
VALUES
  ("Pending"),
  ("Incomplete"),
  ("Changes requested"),
  ("Rejected"),
  ("Approved");
