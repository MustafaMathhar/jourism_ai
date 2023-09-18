CREATE TABLE profiles (
id bigint not null auto_increment,
  first_name varchar(100) not null,
  last_name varchar(100) not null,
  email varchar(320) not null,
  category_id bigint,
  attraction_id bigint,
  country_id bigint,
  PRIMARY KEY(id)
)
