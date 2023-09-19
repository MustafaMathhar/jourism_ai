CREATE TABLE if not exists plans (
id bigint not null auto_increment,
  name varchar(100) not null,
  profile_id bigint,
  start_date date,
  end_date date,
  PRIMARY KEY(id),
  KEY profile_id_idx (profile_id)
);
CREATE TABLE days (
id varchar(15) not null,
  plan_id bigint not null,
  attraction_id bigint,
  PRIMARY KEY(id),
  KEY plan_id_idx (plan_id),
  key attraction_id_idx (attraction_id)
);
