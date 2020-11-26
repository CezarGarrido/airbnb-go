-- insert users
INSERT INTO users
(first_name, last_name, email, "password", created_at, updated_at)
VALUES('Jonh', 'Smith', 'jonh@email.com', 'arbnb123', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO users
(first_name, last_name, email, "password", created_at, updated_at)
VALUES('Amanda', 'Nunes', 'amanda@email.com', 'arbnbamanda', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- insert apartments
INSERT INTO apartments
("name","description" city, adress, description, number_rooms, price, user_id, created_at, updated_at)
VALUES('Wooden Cottage', 'Best apt' 'New York', 'Staten Island, 234', 'A simple house, a cozy wooden chalet, great to enjoy and have fun.', 1, 30.0, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.apartments
("name", description, street, city, state, country, guests, bedrooms, beds, baths, likes, price, status, user_id, created_at, updated_at)
VALUES('Wooden Cottage', '', 'Staten Island, 234', 'New York', 'NY', 'USA', 0, 0, 0, 0, 0, 0, 'available'::character varying, 0, CURRENT_TIMESTAMP, '');


-- insert bookings
INSERT INTO bookings
(apartment_id, owner_id, customer_id, start_date, end_date)
VALUES(1, 1, 2, '2020-09-16 10:37:29', '2020-09-17 10:37:29');