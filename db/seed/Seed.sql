INSERT INTO
  users (
    `id`,
    `username`,
    `display_name`,
    `email`,
    `password_hash`,
    `avatar_url`
  )
VALUES (
  101,
  'demo',
  'Demo User',
  'demo@moodjournal.com',
  -- password: secret
  '$2a$10$9jRGUM08I.wV9mHNhw7vJuaUUeb799zlYoq9Mk8FIvjvdg62vBi/a',
  'https://gravatar.com/avatar/111c287a4e4b61c57dc91ebe43a4ea1c?s=400&d=robohash&r=x'
);

INSERT INTO
  moods (
		`mood`,
    `user_id`,
		`description`,
		`date`
	)
VALUES 
  (
    'Happy',
    101,
    'Spent the day at the beach with friends, enjoying the sun and waves!',
    '2023-07-25'
  );

INSERT INTO
  moods (
		`mood`,
    `user_id`,
		`description`,
		`date`
	)
VALUES 
  (
    'Happy',
    101,
    'Received a promotion at work today! Celebrated with colleagues in the evening.',
    '2023-07-01'
  );

INSERT INTO
  moods (
		`mood`,
    `user_id`,
		`description`,
		`date`
	)
VALUES 
  (
    'Happy',
    101,
    'Attended a family reunion, had a fantastic time catching up with relatives and making new memories.',
    '2022-08-01'
  );

INSERT INTO
  moods (
		`mood`,
    `user_id`,
		`description`,
		`date`
	)
VALUES 
  (
    'Happy',
    101,
    'Completed a challenging hike and enjoyed breathtaking views from the mountaintop.',
    '2021-08-01'
  );

INSERT INTO
  moods (
		`mood`,
    `user_id`,
		`description`,
		`date`
	)
VALUES 
  (
    'Happy',
    101,
    'Surprised my partner with a thoughtful gift and saw their face light up with joy.',
    '2021-07-31'
  );

INSERT INTO
  moods (
		`mood`,
    `user_id`,
		`description`,
		`date`
	)
VALUES 
  (
    'Sad',
    101,
    'Lost my beloved pet today, and feeling heartbroken.',
    '2023-07-25'
  );

INSERT INTO
  moods (
		`mood`,
    `user_id`,
		`description`,
		`date`
	)
VALUES 
  (
    'Sad',
    101,
    'Anniversary of a loved one passing; missing them dearly.',
    '2022-08-01'
  );

INSERT INTO
  moods (
		`mood`,
    `user_id`,
		`description`,
		`date`
	)
VALUES 
  (
    'Sad',
    101,
    'Had a falling out with a close friend, and its been tough to cope.',
    '2021-07-31'
  );

INSERT INTO
  moods (
		`mood`,
    `user_id`,
		`description`,
		`date`
	)
VALUES 
  (
    'Normal',
    101,
    'Spent a relaxing day at home, reading a good book and enjoying some quiet time.',
    '2023-07-25'
  );

INSERT INTO
  moods (
		`mood`,
    `user_id`,
		`description`,
		`date`
	)
VALUES 
  (
    'Normal',
    101,
    'Met up with friends for a casual dinner and had a great time catching up.',
    '2022-08-01'
  );

INSERT INTO
  moods (
		`mood`,
    `user_id`,
		`description`,
		`date`
	)
VALUES 
  (
    'Normal',
    101,
    'Achieved a personal milestone, feeling content and accomplished.',
    '2021-08-01'
  );
