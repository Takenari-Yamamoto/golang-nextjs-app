CREATE TABLE tasks (
  ID varchar(255) NOT NULL,
  Title varchar(255) NOT NULL,
  Content text NOT NULL,
  Done boolean NOT NULL,
  CreatedBy varchar(255) NOT NULL,
  CreatedAt timestamp NOT NULL,
  PRIMARY KEY (ID)
);
