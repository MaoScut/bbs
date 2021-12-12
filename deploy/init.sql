CREATE DATABASE bbs_db;
USE bbs_db;
CREATE TABLE user_tab (
  id int(64) NOT NULL AUTO_INCREMENT,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  create_time int(64) NOT NULL,
  update_time int(64) NOT NULL,
  PRIMARY KEY ( id )
);