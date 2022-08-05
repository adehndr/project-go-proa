CREATE TABLE anime_list (id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, title VARCHAR(512) NOT NULL, description VARCHAR(4096),episodes INT(6), aired DATETIME default null, finished DATETIME default null, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
CREATE TABLE task_table (id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, task_detail VARCHAR(512) NOT NULL,  ,deadline DATETIME default null, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);

-- PostgreSQL 
CREATE TABLE task_table (id serial PRIMARY KEY, task_detail VARCHAR(4096), assignee VARCHAR(512), deadline DATE, is_finished BOOLEAN, created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP );
insert into task_table(task_detail,assignee,deadline,is_finished) values('ini tes','ade',now(),false);
