CREATE DATABASE IF NOT EXISTS dojo_db;

USE dojo_db;

CREATE TABLE users(
  id              INTEGER PRIMARY KEY AUTO_INCREMENT,
  name            VARCHAR(40) NOT NULL,
  token           VARCHAR(255) UNIQUE,
  created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE characters(
  id   INTEGER PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(40) NOT NULL
);

CREATE TABLE collections(
  id              INTEGER PRIMARY KEY AUTO_INCREMENT,
  user_id         INTEGER,
  character_id    INTEGER,

  CONSTRAINT fk_user_id
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE SET NULL ON UPDATE CASCADE,

  CONSTRAINT fk_character_id_collections
    FOREIGN KEY (character_id)
    REFERENCES characters (id)
    ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE probabilities(
  id           INTEGER PRIMARY KEY AUTO_INCREMENT,
  weight       INTEGER,
  character_id INTEGER,

  CONSTRAINT fk_character_id_probabilities
    FOREIGN KEY (character_id)
    REFERENCES characters (id)
    ON DELETE SET NULL ON UPDATE CASCADE
);

INSERT INTO users (name, token) VALUES ("林 遼太朗", "11111");
INSERT INTO users (name, token) VALUES ("John Titor", "22222");

INSERT INTO characters (name) VALUES ("フシギダネ");
INSERT INTO characters (name) VALUES ("ヒトカゲ");
INSERT INTO characters (name) VALUES ("ゼニガメ");

INSERT INTO collections (user_id, character_id) VALUES (1, 1);
INSERT INTO collections (user_id, character_id) VALUES (1, 2);
INSERT INTO collections (user_id, character_id) VALUES (2, 3);

INSERT INTO probabilities (weight, character_id) VALUES (10, 1);
INSERT INTO probabilities (weight, character_id) VALUES (20, 2);
INSERT INTO probabilities (weight, character_id) VALUES (70, 3);