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


DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts (
    account_id INTEGER PRIMARY KEY AUTOINCREMENT,
    customer_id INTEGER,
    opening_date DATE NOT NULL,
    account_type TEXT NOT NULL,
    amount INTEGER NOT NULL,
    status INTEGER NOT NULL DEFAULT '1'
);

INSERT INTO accounts VALUES
 (1, 1000, '2024-09-09 13:00:00',"savings", 5126.25, "1" );

DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
    transaction_id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER,
    amount INTEGER NOT NULL,
    transaction_type TEXT NOT NULL,
    transaction_date DATE NOT NULL
);

INSERT INTO accounts VALUES
    (1, 1000, '2024-09-09 13:00:00',"savings", 5126.25, "1" );


-- print to show data exist
select * from customers;
select * from accounts;
select * from transactions;


