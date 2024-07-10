-- Active: 1714057451806@@127.0.0.1@5432@restapi_peopleinfo@public
CREATE TABLE Users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    surname VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    address VARCHAR(255) NOT NULL,
    passportSerie INT NOT NULL,
    passportNumber INT NOT NULL
);
CREATE INDEX idx_passport ON Users(passportSerie, passportNumber);
-- CREATE INDEX idx_user ON TimeLogs(user_id);

CREATE FUNCTION filter_users(
    -- p_id BIGINT,
    p_surname TEXT,
    p_name TEXT,
    p_patronymic TEXT,
    p_address TEXT,
    p_passport_serie INTEGER,
    p_passport_number INTEGER
)
RETURNS TABLE (
    -- id BIGINT,
    user_surname TEXT,
    user_name TEXT,
    user_patronymic TEXT,
    user_address TEXT,
    user_passport_serie INTEGER,
    user_passport_number INTEGER
)
AS $$
BEGIN
    RETURN QUERY
    SELECT
        -- id,
        surname AS user_surname,
        name AS user_name,
        patronymic AS user_patronymic,
        address AS user_address,
        passport_serie AS user_passport_serie,
        passport_number AS user_passport_number
    FROM users
    WHERE
        -- (p_id IS NULL OR id = p_id) AND
        (p_surname IS NULL OR surname ILIKE p_surname || '%') AND
        (p_name IS NULL OR name ILIKE p_name || '%') AND
        (p_patronymic IS NULL OR patronymic ILIKE p_patronymic || '%') AND
        (p_address IS NULL OR address ILIKE p_address || '%') AND
        (p_passport_serie IS NULL OR passport_serie = p_passport_serie) AND
        (p_passport_number IS NULL OR passport_number = p_passport_number);
END;
$$
 LANGUAGE plpgsql;
 


--    Перенести в тесты
INSERT INTO Users (surname, name, patronymic, address, passportSerie, passportNumber) VALUES
    ('Иванов', 'Иван', 'Иванович', 'ул. Ленина, д. 5, кв. 10', 1234, 567890),
    ('Петрова', 'Мария', 'Сергеевна', 'пр. Победы, д. 7, кв. 15', 4567, 123456),
    ('Сидоров', 'Алексей', 'Михайлович', 'ул. Гагарина, д. 3, кв. 20', 7890, 456123),
    ('Кузнецова', 'Анна', 'Викторовна', 'ул. Мира, д. 9, кв. 25', 2345, 678901),
    ('Смирнов', 'Дмитрий', 'Андреевич', 'пр. Ленинский, д. 11, кв. 30', 6789, 234567),
    ('Орлова', 'Екатерина', 'Сергеевна', 'ул. Пушкина, д. 13, кв. 35', 3456, 789012),
    ('Волков', 'Максим', 'Олегович', 'ул. Гоголя, д. 15, кв. 40', 8901, 345678),
    ('Соколова', 'Ольга', 'Николаевна', 'пр. Октябрьский, д. 17, кв. 45', 4567, 890123),
    ('Романов', 'Артем', 'Александрович', 'ул. Чехова, д. 19, кв. 50', 7890, 456789),
    ('Лебедева', 'Татьяна', 'Владимировна', 'ул. Толстого, д. 21, кв. 55', 2345, 678012);


    SELECT  surname, name, patronymic, address, passport_serie, passport_number 
    FROM filter_users('Иванов', 'Иван', 'Иванова', 'ул. Ленина, д. 5, кв. 10', 1234, 567890);