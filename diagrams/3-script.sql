
-- Users table
CREATE TABLE IF NOT EXISTS users (
	id bigserial NOT NULL,
	uuid uuid NOT NULL,
	first_name varchar NOT NULL,
	last_name varchar NOT NULL,
	email varchar NOT NULL,
	phone varchar,
	password varchar NOT NULL,
	created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP,
	CONSTRAINT pk_user_id PRIMARY KEY (id),
	UNIQUE (email)
);


-- public.apartments definition

-- Drop table

-- DROP TABLE public.apartments;

CREATE TABLE public.apartments (
	id bigserial NOT NULL,
	uuid uuid NOT NULL,
	user_id int8 NOT NULL,
	name varchar NOT NULL,
	description text NOT NULL,
	room_type varchar NOT NULL DEFAULT 'whole_place',
	property_type varchar NOT NULL  DEFAULT 'house',
	street varchar NOT NULL,
	city varchar NOT NULL,
	state varchar NOT NULL,
	country varchar NOT NULL,
	guests int4 NOT NULL DEFAULT 0,
	bedrooms int4 NOT NULL DEFAULT 0,
	beds int4 NOT NULL DEFAULT 0,
	baths int4 NOT NULL DEFAULT 0,
	likes int4 NOT NULL DEFAULT 0,
	price numeric NOT NULL,
	status varchar(20) NOT NULL DEFAULT 'available'::character varying,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	CONSTRAINT pk_apartment_id PRIMARY KEY (id)
);


-- public.apartments foreign keys

ALTER TABLE public.apartments ADD CONSTRAINT apartments_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

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

-- Bookings table
CREATE TABLE IF NOT EXISTS apartment_images (
	id bigserial NOT NULL,
	uuid uuid NOT NULL,
	apartment_id int8 NOT NULL,
	file_name varchar NOT NULL,
	file_size int8 NOT NULL,
	url varchar NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	CONSTRAINT fk_apartment_id FOREIGN KEY (apartment_id) REFERENCES apartments(id)
);