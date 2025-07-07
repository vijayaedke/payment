
CREATE TABLE accounts (
    account_id INT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    document_number VARCHAR(50)
);


CREATE TABLE operation_types (
    operation_type_id INT NOT NULL PRIMARY KEY,
    description VARCHAR(50)
);

INSERT INTO operation_types (operation_type_id, description) VALUES 
(1, "PURCHASE"),
(2, "INSTALLMENT PURCHASE"),
(3, "WITHDRAWAL"),
(4, "PAYMENT");

CREATE TABLE transactions (
    transaction_id INT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    account_id INT not NULL,
    operation_type_id INT not NULL,
    amount FLOAT,
    event_date DATETIME,
    FOREIGN KEY (account_id) REFERENCES accounts(account_id),
    FOREIGN KEY (operation_type_id) REFERENCES operation_types(operation_type_id)
);
