INSERT INTO users (id, name, username, bio, image_url, password, email) VALUES
(gen_random_uuid(), 'Sardorbek Axadjonov', 'sardorbek', 'I love coding', 'https://example.com/img1.jpg', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'axadjonovsardorbeck@gmail.com'),
(gen_random_uuid(), 'John Doe', 'johndoe', 'I love biking', 'https://example.com/img1.jpg', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'johndoe@example.com'),
(gen_random_uuid(), 'Jane Smith', 'janesmith', 'Photography enthusiast', 'https://example.com/img2.jpg', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'janesmith@example.com'),
(gen_random_uuid(), 'Alice Johnson', 'alicej', 'Travel blogger', 'https://example.com/img3.jpg', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'alicej@example.com'),
(gen_random_uuid(), 'Bob Brown', 'bobbrown', 'Fitness trainer', 'https://example.com/img4.jpg', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'bobbrown@example.com'),
(gen_random_uuid(), 'Charlie Black', 'charlieb', 'Gamer and streamer', 'https://example.com/img5.jpg', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'charlieb@example.com'),
(gen_random_uuid(), 'David White', 'davidwhite', 'Music producer', 'https://example.com/img6.jpg', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'davidwhite@example.com'),
(gen_random_uuid(), 'Eva Green', 'evagreen', 'Food lover', 'https://example.com/img7.jpg', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'evagreen@example.com'),
(gen_random_uuid(), 'Frank Blue', 'frankblue', 'Tech geek', 'https://example.com/img8.jpg', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'frankblue@example.com'),
(gen_random_uuid(), 'Grace Pink', 'gracepink', 'Fashion designer', 'https://example.com/img9.jpg', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'gracepink@example.com'),
(gen_random_uuid(), 'Henry Gray', 'henrygray', 'Fitness coach', 'https://example.com/img10.jpg', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'henrygray@example.com');

INSERT INTO twits (id, user_id, twit_id, content, image_url) VALUES
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'sardorbek'), NULL, 'Hello Twitter!', ''),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'janesmith'), NULL, 'Just took an amazing photo!', 'https://example.com/photo.jpg'),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'alicej'), NULL, 'Traveling to Italy soon!', ''),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'bobbrown'), NULL, 'Morning workout complete!', ''),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'sardorbek'), NULL, 'New game release today!', 'https://example.com/game.jpg'),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'sardorbek'), NULL, 'Produced a new song!', ''),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'evagreen'), NULL, 'Just tried an amazing new recipe.', 'https://example.com/food.jpg'),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'frankblue'), NULL, 'New tech review is up!', ''),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'gracepink'), NULL, 'Designing a new dress.', 'https://example.com/design.jpg'),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'henrygray'), NULL, 'Fitness tips for the day.', '');


INSERT INTO follows (id, follower_id, followed_id) VALUES
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'johndoe'), (SELECT id FROM users WHERE username = 'sardorbek')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'alicej'), (SELECT id FROM users WHERE username = 'sardorbek')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'sardorbek'), (SELECT id FROM users WHERE username = 'davidwhite')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'evagreen'), (SELECT id FROM users WHERE username = 'sardorbek')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'gracepink'), (SELECT id FROM users WHERE username = 'henrygray')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'bobbrown'), (SELECT id FROM users WHERE username = 'sardorbek')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'davidwhite'), (SELECT id FROM users WHERE username = 'alicej')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'sardorbek'), (SELECT id FROM users WHERE username = 'charlieb')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'sardorbek'), (SELECT id FROM users WHERE username = 'gracepink')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'sardorbek'), (SELECT id FROM users WHERE username = 'evagreen'));

INSERT INTO likes (id, user_id, twit_id) VALUES
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'sardorbek'), (SELECT id FROM twits WHERE content = 'Morning workout complete!')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'sardorbek'), (SELECT id FROM twits WHERE content = 'Just tried an amazing new recipe.')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'sardorbek'), (SELECT id FROM twits WHERE content = 'Fitness tips for the day.')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'bobbrown'), (SELECT id FROM twits WHERE content = 'Produced a new song!')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'charlieb'), (SELECT id FROM twits WHERE content = 'New game release today!')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'davidwhite'), (SELECT id FROM twits WHERE content = 'Hello Twitter!')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'evagreen'), (SELECT id FROM twits WHERE content = 'New tech review is up!')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'frankblue'), (SELECT id FROM twits WHERE content = 'Designing a new dress.')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'sardorbek'), (SELECT id FROM twits WHERE content = 'Just took an amazing photo!')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'henrygray'), (SELECT id FROM twits WHERE content = 'Traveling to Italy soon!'));
