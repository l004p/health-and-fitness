-- INSERT INTO users (username, user_password, first_name, last_name, user_email) VALUES
--     ('Baker12', '123', 'Hannah', 'Baker', 'hannahbaker@email'),
--     ('FB_Bryce', '123', 'Bryce', 'Walker', 'whodunnit@email');

-- INSERT INTO user_roles (user_id, role_id) VALUES
--     (1, 1),
--     (1, 2),
--     (1, 3),
--     (2, 1);

--  INSERT INTO user_roles (user_id, role_id) VALUES
--      (8, 1),
--      (8, 2),
--      (8, 3);

--  INSERT INTO member_trained (member_id, trainer_id) VALUES
--      (2, 6),
--      (5, 6);

--  INSERT INTO trainer_interests (trainer_id, interest) VALUES
--      (8, 'CARDIO'),
--      (8, 'RUNNING'),
--      (1, 'RUNNING'),
--      (1, 'ENDURANCE'),
--      (8, 'MOBILITY');

-- INSERT INTO memberships (member_id, title, membership_status) VALUES
--     (1, 'REGULAR MEMBERSHIP', 'ACTIVE'),
--     (2, 'REGULAR MEMBERSHIP', 'ACTIVE');

INSERT INTO rooms (capacity, room_description) VALUES
    (30, 'YOGA STUDIO'),
    (10, 'WEIGHT ROOM'),
    (40, 'BIKE ROOM'),
    (12, 'MULTIPURPOSE ROOM');

INSERT INTO equipment (equipment_type, equipment_status) VALUES
    ('BIKE', 'ACTIVE'),
    ('BIKE', 'ACTIVE'),
    ('BIKE', 'ACTIVE'),
    ('BIKE', 'ACTIVE'),
    ('BIKE', 'ACTIVE'),
    ('BIKE', 'ACTIVE'),
    ('TREADMILL', 'ACTIVE'),
    ('TREADMILL', 'ACTIVE'),
    ('TREADMILL', 'ACTIVE'),
    ('TREADMILL', 'ACTIVE'),
    ('TREADMILL', 'ACTIVE'),
    ('TREADMILL', 'ACTIVE'),
    ('STAIRMASTER', 'ACTIVE'),
    ('STAIRMASTER', 'ACTIVE'),
    ('STAIRMASTER', 'ACTIVE'),
    ('ROWING MACHINE', 'MAINTENANCE'),
    ('ROWING MACHINE', 'MAINTENANCE'),
    ('ROWING MACHINE', 'MAINTENANCE'),
    ('STAIRMASTER', 'MAINTENANCE'),
    ('STAIRMASTER', 'MAINTENANCE'),
    ('TREADMILL', 'MAINTENANCE'),
    ('TREADMILL', 'MAINTENANCE'),
    ('TREADMILL', 'MAINTENANCE'),
    ('TREADMILL', 'MAINTENANCE');

INSERT INTO equipment_rooms (room_id, equipment_id) VALUES
    (3, 1),
    (3, 2),
    (3, 3),
    (3, 4),
    (3, 5),
    (3, 6),
    (4, 12),
    (4, 13),
    (4, 15);

INSERT INTO goals (member_id, goal_date, title, unit, goal_value, goal_type) VALUES 
    (1, '2024-06-23', 'LOSE WEIGHT', 'kg', 60, 'BODY WEIGHT'),
    (1, '2024-06-23', 'RUNNING DISTANCE', 'km', 3, 'DISTANCE'),
    (2, '2024-06-23', 'LIFTING GOAL', 'kg', 10, 'LIFTING WEIGHT');

INSERT INTO metrics (member_id, unit, metric_value, metric_type) VALUES 
    (1, 'kg', 79, 'BODY WEIGHT'),
    (1, 'kg', 72, 'BODY WEIGHT'),
    (1, 'kg', 3, 'LIFTING WEIGHT'),
    (1, 'kg', 5, 'LIFTING WEIGHT'),
    (1, 'kg', 7, 'LIFTING WEIGHT'),
    (1, 'km', 1, 'DISTANCE'),
    (2, 'kg', 1, 'LIFTING WEIGHT'),
    (2, 'kg', 2, 'LIFTING WEIGHT'),
    (2, 'kg', 4, 'LIFTING WEIGHT');

INSERT INTO calender_events (room_id, event_type, max_attendees, event_status, event_title) VALUES
    (1, 'CLASS', 20, 'UPCOMING', 'YOGA CLASS'),
    (2, 'CLASS', 5, 'UPCOMING', 'WEIGHT CLASS');


INSERT INTO attending_event (member_id, event_id) VALUES
    (1, 1),
    (2, 1),
    (6, 1);

INSERT INTO equipment_rooms (room_id, equipment_id) VALUES
    (3, 7),
    (3, 8),
    (3, 9),
    (3, 10),
    (3, 11),
    (3, 14);

INSERT INTO bills (user_id, bill_description, bill_status) VALUES
    (8, 'CLASS', 'PENDING'),
    (8, 'REGULAR MEMBERSHIP', 'CLEARED'),
    (8, 'CLASS', 'CANCELLED');

INSERT INTO bills (user_id, bill_description, bill_status) VALUES
    (5, 'CLASS', 'PENDING'),
    (5, 'REGULAR MEMBERSHIP', 'CLEARED'),
    (5, 'CLASS', 'CANCELLED');

INSERT INTO bills (user_id, bill_description, bill_status) VALUES
    (6, 'CLASS', 'PENDING'),
    (6, 'EXTENDED MEMBERSHIP', 'CANCELLED'),
    (6, 'CLASS', 'CANCELLED');

INSERT INTO bills (user_id, bill_description, bill_status) VALUES
    (5, 'CLASS', 'PENDING'),
    (5, 'EXTENDED MEMBERSHIP', 'PENDING'),
    (5, 'CLASS', 'PENDING');