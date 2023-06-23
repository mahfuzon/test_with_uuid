CREATE TABLE IF NOT EXISTS contacts (
    id varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    gender enum("male", "female") NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    phone varchar(255) NOT NULL UNIQUE,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
    ) ENGINE=InnoDB;