-- Dummy data for SQL lite
DROP TABLE IF EXISTS customers;
CREATE TABLE customers (
    customer_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    date_of_birth DATE NOT NULL,
    city TEXT NOT NULL,
    zipcode TEXT NOT NULL,
    status INTEGER NOT NULL DEFAULT '1'
);

INSERT INTO customers values
(1000, "Scuba Steave", "1999-06-25", "New York", "20020", 1),
(1001, "Lola Bunny", "1998-11-15", "Toonland", "00001", 1),
(1002, "Cloud Strife", "1997-09-15", "Playstation", "99999", 1),
(1003, "Tifa Lockheart", "1997-09-15", "Playstation", "99998", 0),
(1004, "RedXIII", "1997-09-15", "Playstation", "99997", 0);

-- print to show data exist
select * from customers;