CREATE TABLE IF NOT EXISTS stocks (
    stockid SERIAL PRIMARY KEY,
    stockname VARCHAR(100) NOT NULL,
    stockprice NUMERIC(10,2) NOT NULL,
    stockcompany VARCHAR(100) NOT NULL
);
