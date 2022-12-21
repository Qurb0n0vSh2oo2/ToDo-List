CREATE TABLE tasks(
        id serial PRIMARY KEY,
        name varchar(30) NOT NULL,
        description TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        status VARCHAR(15) NOT NULL,
        user_id serial,
        FOREIGN KEY (user_id) REFERENCES users (id)
    );


-- drop table token