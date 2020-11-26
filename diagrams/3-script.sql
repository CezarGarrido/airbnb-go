
-- Users table
CREATE TABLE IF NOT EXISTS users (
	id bigserial NOT NULL,
	first_name varchar NOT NULL,
	last_name varchar NOT NULL,
	email varchar NOT NULL,
	password varchar NOT NULL,
	created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP,
	CONSTRAINT pk_user_id PRIMARY KEY (id)
);

-- Apartments table
CREATE TABLE IF NOT EXISTS apartments (
	id bigserial NOT NULL,
    user_id int8 NOT NULL,
	name varchar NOT NULL,
	description text NOT NULL,
	street varchar NOT NULL,
	city varchar NOT NULL,
	state varchar NOT NULL,
	country varchar NOT NULL,
	guests integer NOT NULL DEFAULT 0,
	bedrooms integer NOT NULL DEFAULT 0,
	beds integer NOT NULL DEFAULT 0,
	baths integer NOT NULL DEFAULT 0,
	likes integer NOT NULL DEFAULT 0,
	price numeric NOT NULL,
	status varchar(20) NOT NULL DEFAULT 'available',
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	CONSTRAINT pk_apartment_id PRIMARY KEY (id)
);


-- Bookings table
CREATE TABLE IF NOT EXISTS bookings (
	id bigserial NOT NULL,
	apartment_id int8 NOT NULL,
	owner_id int8 NOT NULL,
	customer_id int8 NOT NULL,
	start_date date NOT NULL,
	end_date date NOT NULL,
	CONSTRAINT pk_booking_id PRIMARY KEY (id),
	CONSTRAINT fk_apartment_id FOREIGN KEY (apartment_id) REFERENCES apartments(id),
	CONSTRAINT fk_customer_id FOREIGN KEY (customer_id) REFERENCES users(id),
	CONSTRAINT fk_owner_id FOREIGN KEY (owner_id) REFERENCES users(id)
);
