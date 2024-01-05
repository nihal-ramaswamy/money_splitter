CREATE TABLE user (
  ID VARCHAR(255) NOT NULL PRIMARY KEY,
  EMAIL VARCHAR(255) NOT NULL UNIQUE,
  PASSWORD VARCHAR(255) NOT NULL,
);

CREATE TABLE group (
  ID VARCHAR(255) NOT NULL PRIMARY KEY,
  GROUP_NAME VARCHAR(255) NOT NULL UNIQUE,
  SIMPLIFY_TXN BOOLEAN NOT NULL DEFAULT FALSE,
);

CREATE TABLE user_in_groups (
  user_id VARCHAR(255) NOT NULL foreign key references user(ID),
  group_id VARCHAR(255) NOT NULL foreign key references group(ID),
);



CREATE TABLE group_txns (
  user_id VARCHAR(255) NOT NULL foreign key references user(ID),
  group_id VARCHAR(255) NOT NULL foreign key references group(ID),
  TXN_ID VARCHAR(255) NOT NULL UNIQUE REFERENCES txn(ID)
);

CREATE TABLE txn (
  ID VARCHAR(255) NOT NULL PRIMARY KEY,
  AMOUNT DECIMAL(10,2) NOT NULL,
  DESCRIPTION VARCHAR(255) NOT NULL,
  DATE DATE NOT NULL,
  GROUP_ID VARCHAR(255) NOT NULL foreign key references group(ID),
  PAYER_ID VARCHAR(255) NOT NULL foreign key references user(ID),
);