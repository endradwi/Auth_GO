CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255),
    password VARCHAR(255),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

CREATE TABLE profile{
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(60),
    last_name VARCHAR(60),
    image VARCHAR(255),
    user_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
}

DROP TABLE users, movies

SELECT * From users ORDER BY id ASC;

CREATE Table cinemas{
    id serial PRIMARY KEY,
    name VARCHAR(50),
    profile_id int REFERENCES profile(id),
    movie_id int REFERENCES movies(id),
    -- cinema_location VARCHAR(50),
    -- cinema_date DATE,
    -- cinema_time TIME,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
}

CREATE TABLE cinema_date{
    id SERIAL PRIMARY KEY
    name_date DATE,
    cinema_id int REFERENCES cinemas(id)
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
}

CREATE TABLE cinema_time{
    id SERIAL PRIMARY KEY,
    name_time TIME,
    cinema_id int REFERENCES cinemas(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP 
}

CREATE TABLE cinema_location{
    id SERIAL PRIMARY KEY,
    name_location VARCHAR(50),
    cinema_id int REFERENCES cinemas(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
}

CREATE TABLE seat{
    id SERIAL PRIMARY KEY,
    name VARCHAR(2),
    price INT(10),
    cinema_id INT REFERENCES cinemas(id),
    create_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
}

CREATE TABLE payment_method{
    id SERIAL PRIMARY KEY,
    name VARCHAR(20),
    image VARCHAR(255),
    user_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
}


CREATE TABLE movies(
    id SERIAL PRIMARY KEY,
    tittle VARCHAR(60),
    genre VARCHAR(100),
    images VARCHAR(255),
    synopsis VARCHAR(255),
    author VARCHAR(60),
    actors VARCHAR(255),
    release_date DATE,
    duration TIME,
    tag VARCHAR(30),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

INSERT INTO users (email, password) VALUES
('admin@mail.com', '1234'),
('nauval@mail.com', '1234'),
('rangga@mail.com', '1234'),
('yuni@mail.com', '1234'),
('yanti@mail.com', '1234'),
('tono@mail.com', '1234')


INSERT INTO movies (title, genre, images, synopsis, author, actors, release_date, duration)
VALUES
('The Great Adventure', 'Action', 'img1.jpg', 'An epic journey through dangerous lands.', 'John Doe', 'Jane Smith, Robert Brown', '2023-06-15', '02:30:00'),
('Lost in Time', 'Sci-Fi', 'img2.jpg', 'A time traveler gets stuck in an alternate dimension.', 'Alice White', 'Chris Evans, Emma Stone', '2024-01-22', '01:45:00'),
('Love in Paris', 'Romance', 'img3.jpg', 'Two strangers meet and fall in love in Paris.', 'Marie Adams', 'Olivia Wilde, Tom Hiddleston', '2023-11-10', '01:50:00'),
('Space Odyssey', 'Sci-Fi', 'img4.jpg', 'Exploring the farthest reaches of the universe.', 'David Clark', 'Scarlett Johansson, Idris Elba', '2023-07-04', '02:40:00'),
('The Last Kingdom', 'Drama', 'img5.jpg', 'A king struggles to maintain his reign in troubled times.', 'Michael Lee', 'Henry Cavill, Natalie Dormer', '2023-03-01', '02:20:00'),
('Horror Night', 'Horror', 'img6.jpg', 'A group of friends encounter terror in an abandoned house.', 'Hannah Green', 'Keanu Reeves, Emma Roberts', '2024-05-14', '01:35:00'),
('Race to the Top', 'Adventure', 'img7.jpg', 'A team of explorers races to reach the top of the world.', 'James Brown', 'Brad Pitt, Angelina Jolie', '2023-10-20', '02:10:00'),
('The Silent Witness', 'Thriller', 'img8.jpg', 'A detective must solve a crime with no clues.', 'Nina Patel', 'Daniel Craig, Cate Blanchett', '2024-02-08', '02:00:00'),
('Pirate Treasure', 'Action', 'img9.jpg', 'A group of pirates seek a legendary treasure.', 'Samuel Clark', 'Johnny Depp, Emma Watson', '2023-09-15', '02:25:00'),
('The Mystery of the Forest', 'Mystery', 'img10.jpg', 'A young girl uncovers the secrets of a haunted forest.', 'Sarah King', 'Millie Bobby Brown, Tom Holland', '2024-03-17', '01:55:00'),
('A Night in Venice', 'Romance', 'img11.jpg', 'A couple reunites under the stars in Venice.', 'Lily Adams', 'Ryan Gosling, Emma Stone', '2024-06-01', '02:00:00'),
('The Dark Past', 'Thriller', 'img12.jpg', 'A man confronts his troubled past to save his future.', 'Edward Miller', 'Matt Damon, Charlize Theron', '2023-12-05', '01:45:00'),
('Underwater Escape', 'Action', 'img13.jpg', 'A submarine crew fights to survive underwater.', 'Rachel Johnson', 'Chris Hemsworth, Zoe Saldana', '2023-08-25', '02:15:00'),
('City of Lights', 'Drama', 'img14.jpg', 'A love story unfolds amidst the lights of the city.', 'Jennifer Moore', 'Ryan Reynolds, Emily Blunt', '2024-04-20', '02:10:00'),
('The Edge of Tomorrow', 'Sci-Fi', 'img15.jpg', 'Soldiers fight in a time loop against alien invaders.', 'Sophie Black', 'Tom Cruise, Emily Blunt', '2023-12-12', '02:00:00'),
('The Forgotten Princess', 'Fantasy', 'img16.jpg', 'A princess discovers her hidden magical powers.', 'Rachel Turner', 'Emma Watson, Henry Cavill', '2023-10-05', '02:30:00'),
('Beneath the Surface', 'Horror', 'img17.jpg', 'A group of divers discover a terrifying underwater secret.', 'John Walker', 'Hugh Jackman, Jennifer Lawrence', '2024-01-10', '01:40:00'),
('The Great Heist', 'Action', 'img18.jpg', 'A group of thieves plan the ultimate heist.', 'George Turner', 'Ryan Gosling, Margot Robbie', '2024-05-22', '02:25:00'),
('Ghost Ship', 'Horror', 'img19.jpg', 'A ghostly ship haunts the seas, seeking revenge.', 'Isabella Harris', 'Chris Pratt, Zoe Kravitz', '2023-09-30', '01:55:00'),
('The Legend of the Dragon', 'Fantasy', 'img20.jpg', 'A young warrior must face an ancient dragon to save the kingdom.', 'Michael Smith', 'Chris Hemsworth, Zendaya', '2023-11-05', '02:40:00'),
('Tales of the Wild', 'Adventure', 'img21.jpg', 'An adventurer sets off to discover a lost civilization.', 'Samuel Green', 'Dwayne Johnson, Emily Blunt', '2023-07-10', '02:00:00'),
('Silent City', 'Drama', 'img22.jpg', 'A city grieves after a devastating tragedy.', 'Olivia White', 'Tom Hardy, Natalie Portman', '2024-03-22', '02:05:00'),
('The King’s Secret', 'Thriller', 'img23.jpg', 'A royal secret threatens the kingdom’s future.', 'Daniel Black', 'Matt Damon, Nicole Kidman', '2024-04-13', '02:15:00'),
('Beyond the Horizon', 'Sci-Fi', 'img24.jpg', 'Explorers journey beyond known space to find a new world.', 'Helen Moore', 'Will Smith, Gal Gadot', '2023-08-18', '02:30:00'),
('Dark Side of the Moon', 'Sci-Fi', 'img25.jpg', 'A team of astronauts investigate strange occurrences on the moon.', 'Peter Allen', 'Chris Hemsworth, Scarlett Johansson', '2023-12-20', '02:00:00'),
('Kingdom of Shadows', 'Fantasy', 'img26.jpg', 'A prince battles against an army of shadows to reclaim his throne.', 'Victoria Green', 'Chris Pine, Emilia Clarke', '2023-10-11', '02:40:00'),
('Fallen Star', 'Drama', 'img27.jpg', 'A former star attempts to make a comeback in the world of entertainment.', 'Alexandra White', 'Meryl Streep, Leonardo DiCaprio', '2024-01-12', '01:55:00'),
('Under the Moonlight', 'Romance', 'img28.jpg', 'A couple finds love under the full moon in the countryside.', 'Sophie Brown', 'Emma Watson, Timothée Chalamet', '2024-02-10', '02:00:00'),
('The Secret Agent', 'Thriller', 'img29.jpg', 'A secret agent races against time to stop a global threat.', 'Robert Black', 'James Bond, Naomi Harris', '2024-05-05', '02:15:00'),
('The Vampire’s Curse', 'Horror', 'img30.jpg', 'A family discovers they are cursed by an ancient vampire.', 'Laura King', 'Zac Efron, Lily Collins', '2024-06-10', '02:05:00'),
('The Last Train', 'Adventure', 'img31.jpg', 'A group of passengers must survive a mysterious train ride.', 'Benjamin Davis', 'Tom Cruise, Michelle Williams', '2024-02-20', '01:45:00'),
('Rise of the Phoenix', 'Fantasy', 'img32.jpg', 'A young hero awakens the power of the Phoenix to save the world.', 'David Johnson', 'Henry Cavill, Emma Watson', '2023-11-30', '02:30:00'),
('City of Shadows', 'Action', 'img33.jpg', 'A vigilante fights crime in the dark underbelly of the city.', 'Sarah White', 'Chris Evans, Gal Gadot', '2024-07-01', '02:20:00'),
('The Secret Door', 'Mystery', 'img34.jpg', 'A detective uncovers hidden doors that lead to a parallel world.', 'James Turner', 'Robert Downey Jr., Cate Blanchett', '2023-09-19', '01:50:00'),
('Beyond the Edge', 'Thriller', 'img35.jpg', 'A man discovers the terrifying truth about his own identity.', 'Rachel Scott', 'Leonardo DiCaprio, Jennifer Lawrence', '2024-05-10', '02:10:00'),
('The Final Hour', 'Action', 'img36.jpg', 'A race against time to stop a catastrophic event from occurring.', 'James White', 'Brad Pitt, Anne Hathaway', '2023-10-01', '02:35:00'),
('Echoes of the Past', 'Drama', 'img37.jpg', 'A family returns to their ancestral home and uncovers dark secrets.', 'Victoria Harris', 'Emma Thompson, Hugh Grant', '2024-04-18', '02:00:00'),
('The Lighthouse Keeper', 'Horror', 'img38.jpg', 'A lighthouse keeper faces an otherworldly threat at sea.', 'Jason Brown', 'Jake Gyllenhaal, Rachel McAdams', '2023-11-25', '01:45:00'),
('Wings of Freedom', 'Action', 'img39.jpg', 'A group of resistance fighters challenges a tyrannical government.', 'Oliver Green', 'Tom Hardy, Emily Blunt', '2023-08-23', '02:00:00'),
('The Enchanted Forest', 'Fantasy', 'img40.jpg', 'A young girl ventures into an enchanted forest to find her destiny.', 'Lily Johnson', 'Saoirse Ronan, Emma Watson', '2024-06-08', '02:20:00'),
('Shadow of the Beast', 'Action', 'img41.jpg', 'A warrior hunts down a dangerous beast that threatens the kingdom.', 'Edward Clark', 'Chris Hemsworth, Gal Gadot', '2023-10-22', '02:15:00'),
('Beyond the Stars', 'Sci-Fi', 'img42.jpg', 'A journey into deep space to discover a new life form.', 'Sarah Davis', 'Chris Pratt, Zoe Saldana', '2024-01-05', '02:30:00'),
('The Night Stalker', 'Thriller', 'img43.jpg', 'A detective faces a serial killer who stalks the night.', 'Jonathan Lee', 'Matthew McConaughey, Julia Roberts', '2024-03-14', '02:10:00'),
('The Lost Treasure', 'Adventure', 'img44.jpg', 'A treasure hunter embarks on a dangerous journey to find a lost artifact.', 'Lucas Brown', 'Johnny Depp, Emma Watson', '2024-04-05', '02:25:00'),
('Dawn of the Dead', 'Horror', 'img45.jpg', 'A zombie apocalypse ravages the world as survivors fight to stay alive.', 'Laura White', 'Brad Pitt, Keira Knightley', '2023-12-01', '02:10:00'),
('The Beast Within', 'Horror', 'img46.jpg', 'A man is haunted by his own monstrous alter ego.', 'Zara Scott', 'Oscar Isaac, Alicia Vikander', '2024-01-25', '01:55:00'),
('The Final Countdown', 'Sci-Fi', 'img47.jpg', 'A group of astronauts must stop a cosmic event that will destroy Earth.', 'Eric Harris', 'Ryan Gosling, Brie Larson', '2024-02-18', '02:30:00'),
('The Last Starfighter', 'Sci-Fi', 'img48.jpg', 'A young man is recruited by an alien race to fight in an intergalactic war.', 'George Lee', 'Tom Hanks, Zoe Kravitz', '2023-10-15', '02:00:00'),
('The Spy Who Loved Me', 'Thriller', 'img49.jpg', 'A spy must navigate a world of deceit and danger.', 'Maria White', 'Daniel Craig, Penélope Cruz', '2024-05-30', '02:10:00'),
('Into the Abyss', 'Horror', 'img50.jpg', 'A group of explorers venture into a cave that harbors dark secrets.', 'Ethan Brown', 'Christian Bale, Anne Hathaway', '2024-07-10', '02:05:00');

