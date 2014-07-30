CREATE TABLE persons (
  id    serial PRIMARY KEY,
  name  varchar(160),
  age   integer,
  job   varchar(160),
  email varchar(160)
);

INSERT INTO persons VALUES (DEFAULT, 'Jeaneth Farmer', 23, 'TangoSource Engineer', 'xaid27@gmail.com');
