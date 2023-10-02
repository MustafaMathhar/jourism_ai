CREATE TABLE profiles_to_attractions (
profile_id bigint,
  attraction_id bigint,
  PRIMARY KEY(profile_id,attraction_id)
);
ALTER TABLE attractions
DROP COLUMN profile_id,
DROP COLUMN category_id;
