CREATE TABLE inventory
(
    product_id varchar,
    quantity integer,
    CONSTRAINT product_master_ivenory FOREIGN KEY (product_id)  REFERENCES product_master (product_id) 
);

CREATE TABLE category_master
(
    category_id varchar,
    category_name varchar,
    CONSTRAINT category_master_pkey PRIMARY KEY (category_id)
);

CREATE TABLE product_master
(
    product_id varchar,
    name varchar,
    sku varchar,
    category_id varchar,
    price float64,
    specification json ,
    CONSTRAINT product_master_pkey PRIMARY KEY (product_id),
    CONSTRAINT category_id FOREIGN KEY (category_id)
        REFERENCES category_master (category_id)
   );
   



CREATE TABLE cart_reference
(
    reference_id VARCHAR,
    create_date DATE,
    username VARCHAR,
    CONSTRAINT cart_reference_pkey PRIMARY KEY (reference_id)
);


  CREATE TABLE cart (
  reference_id VARCHAR,
   product_id VARCHAR,
   quantity INT,
CONSTRAINT cart_account PRIMARY KEY (  reference_id,product_id),
    CONSTRAINT reference_id FOREIGN KEY (reference_id)
        REFERENCES cart_reference (reference_id),
        CONSTRAINT product_id FOREIGN KEY (product_id)
        REFERENCES product_master (product_id)
);