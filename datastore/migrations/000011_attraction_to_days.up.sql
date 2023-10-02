CREATE TABLE attractions_to_days (
attraction_id bigint not null,
  day_id varchar(15) not null,
  PRIMARY KEY(attraction_id,day_id)
);

ALTER TABLE days
  DROP COLUMN attraction_id;
