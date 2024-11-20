CREATE TABLE users (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(50) NOT NULL,
  password CHAR(32) NOT NULL,
  email VARCHAR(100),
  created_at DATETIME,
  updated_at DATETIME
);
ALTER TABLE users ADD UNIQUE (username);

CREATE TABLE words (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    uid INT UNSIGNED NOT NULL,
    word VARCHAR ( 255 ) NOT NULL,
    definition TEXT,
    example_sentence TEXT,
    chinese_translation VARCHAR ( 255 ),
    pronunciation VARCHAR ( 255 ),
    proficiency_level SMALLINT UNSIGNED,
    created_at DATETIME,
    updated_at DATETIME
);

ALTER TABLE words ADD UNIQUE (uid, word);