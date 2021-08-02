CREATE TABLE books (
	id SERIAL PRIMARY KEY,
	name varchar(50) NOT NULL, 
	description varchar(100) NOT NULL,
	medium_price float NOT NULL,
	author varchar (50) NOT NULL,
	ImageURL varchar NOT NULL,
	CreatedAt TIMESTAMPTZ DEFAULT Now(),
	UpdatedAt TIMESTAMPTZ DEFAULT Now()
);

INSERT INTO books (name, description, medium_price, author, ImageURL) VALUES ('Livro 1', 'Um livro legal', 199.99, 'Um autor ai', 'link da URL');