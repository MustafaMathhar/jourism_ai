ALTER TABLE attractions
  ADD COLUMN profile_id bigint default null;

ALTER TABLE categories
  ADD COLUMN profile_id bigint default null;
