CREATE DATABASE IF NOT EXISTS olive_db;
USE olive_db;

CREATE TABLE IF NOT EXISTS olive_db.reservations (
  id                INT(11) NOT NULL AUTO_INCREMENT,
  title             CHAR(64) NOT NULL,
  description       BLOB(256),
  status            TINYINT NOT NULL,
  fee               INT,
  departure_date    DATETIME,
  arrival_date      DATETIME,
  created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  updated_at        TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
