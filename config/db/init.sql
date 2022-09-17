CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title character varying(255) NOT NULL UNIQUE,
    author character varying(255) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO books (id, title, author) VALUES
(1, 'Fairy Tale', 'Stephen King'),
(2, 'Desperation In Death', 'J.D. Robb'),
(3, 'Verity', 'John Johnson'),
(4, 'It Ends With Us', 'Colleen Hoover'),
(5, 'Where The Crawdads Sing', 'Delia Owens'),
(6, 'A Court Of Silver Flames', 'Sarah J. Maas'),
(7, 'Ugly Love', 'Colleen Hoover'),
(8, 'The Seven Husbands Of Evelyn Hugo', 'Taylor Jenkins Reid'),
(9, 'November 9', 'Colleen Hoover'),
(10, 'The American Roommate Experiment', 'Elena Armas'),
(11, 'Clive Cusslers Hellburner', 'Mike Maden'),
(12, 'Hell And Back', 'Craig Johnson'),
(13, 'Fire And Blood', 'George R.R. Martin'),
(14, 'Love On The Brain', 'Ali Hazelwood'),
(15, 'Carrie Soto Is Back', 'Taylor Jenkins Reid');