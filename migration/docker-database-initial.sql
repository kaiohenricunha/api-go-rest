create table wikis(
    id serial primary key,
    name varchar,
    bio varchar
);

INSERT INTO wikis(name, bio) VALUES
('Maynard Keynes', 'John Maynard Keynes, 1st Baron Keynes, CB, FBA, was an English economist whose ideas fundamentally changed the theory and practice of macroeconomics and the economic policies of governments. Originally trained in mathematics, he built on and greatly refined earlier work on the causes of business cycles.'),
('Flannery OConnor', 'Mary Flannery OConnor was an American novelist, short story writer and essayist. She wrote two novels and 31 short stories, as well as a number of reviews and commentaries.');
