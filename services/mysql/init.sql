ALTER DATABASE sklad CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

use sklad

CREATE TABLE product (
    id int NOT NULL AUTO_INCREMENT,
    short_name varchar(255),
    description varchar(255),
    product_part_number varchar(255),
    PRIMARY KEY (id)
);

INSERT INTO product (short_name, description, product_part_number) VALUE ('Logitech MX Master 3S Performance','Мышь Logitech MX Master 3S Performance Wireless Mouse Bluetooth Graphite (910-006559)', '910-006559'),
('Logitech Lift for Mac Vertical Ergonomic Mouse Off','Мышь Logitech Lift for Mac Vertical Ergonomic Mouse Off White (910-006477)', '910-006477'),
('Logitech Signature M650 L Wireless Mouse LEFT Off','Мышь Logitech Signature M650 L Wireless Mouse LEFT Off-White (910-006240)', '910-006240');



CREATE TABLE customers (
    id int NOT NULL AUTO_INCREMENT,
    first_name varchar(255),
    last_name varchar(255),
    PRIMARY KEY (id)
);

INSERT INTO customers (first_name, last_name) VALUE ('Ivan','Ma'),('Sergiy','Ja'), ('Max','Ha');


CREATE TABLE skald (
    id int NOT NULL AUTO_INCREMENT,
    sklad_id varchar(255),
    short_name varchar(255),
    description varchar(255),
    address varchar(255),
    PRIMARY KEY (id)
);


INSERT INTO skald (sklad_id, short_name, description, address) VALUE ('ukraine-odesa-00001', 'Відд 1','The Head branch outside of city','Одеса, Київське шосе (ран. Ленінградське шосе), 27' ),
('ukraine-odesa-00003','Відд 3','The second largest branch within the city','Одеса, вул. Дальницька, 23/4' ),
('ukraine-odesa-00067', 'Відд 67','Одеса відділення №67','Одеса, вул. Марсельська, 35 б' );


CREATE TABLE remaining (
    id int NOT NULL AUTO_INCREMENT,
    product_part_number  varchar(255),
    sklad_id  varchar(255),
    qty int,
    PRIMARY KEY (id)
);
INSERT INTO remaining (product_part_number, sklad_id, qty) VALUE 
('910-006559','ukraine-odesa-00003', 100 ),
('910-006559','ukraine-odesa-00001', 40 ),
('910-006240','ukraine-odesa-00067', 1 ),
('910-006240','ukraine-odesa-00001', 101 );




CREATE TABLE ping_log (
    id int NOT NULL AUTO_INCREMENT,
    ping_time  DATETIME,
    PRIMARY KEY (id)
);