CREATE TABLE users (
    user_id serial PRIMARY KEY,
    steam_id varchar (32),
    name varchar (32),
    last_connected timestamp
);

CREATE INDEX steam_id_idx ON users USING btree(steam_id);

CREATE TABLE ip_addresses (
    ipaddress_id serial PRIMARY KEY,
    ipaddress int
);

CREATE INDEX ipaddress_idx ON ip_addresses USING btree(ipaddress);

CREATE TABLE ip_address_users (
    ref_id serial PRIMARY KEY,
    ipaddress_id integer,
    user_id integer,
    connection_time timestamp
);

ALTER TABLE ip_address_users ADD FOREIGN KEY (ipaddress_id) REFERENCES ip_addresses;
ALTER TABLE ip_address_users ADD FOREIGN KEY (user_id) REFERENCES users;
