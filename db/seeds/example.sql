INSERT INTO users (email, password)
VALUES ('mail@mail.com', 'password');

INSERT INTO
    profiles (
        user_id,
        full_name,
        nickname,
        occupation,
        location,
        short_summary,
        summary
    )
VALUES (
        1,
        'Ellion Blessan',
        'Leon',
        'Software Engineer and Machine Learning enthusiast',
        'Indonesia',
        'I make web apps and mobile apps coupled with Machine Learning capabilities',
        'Dedicated Software Engineer with a strong foundation in back-end development and experience in machine learning. Proficient in Go and Python. Skilled in database management, API development, and version control. Eager to apply my theoretical knowledge and practical skills to contribute to innovative projects and further develop my expertise in back-end technologies.'
    );