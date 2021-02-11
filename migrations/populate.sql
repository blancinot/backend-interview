INSERT INTO users (id, name) VALUES (
    'testuid',
    'testname'
);

INSERT INTO users (id, name) VALUES (
    'testuid2',
    'testname2'
);

INSERT INTO account (id, user_id, total) VALUES (
    'testaid0',
    'testuid',
    0
);

INSERT INTO account (id, user_id, total) VALUES (
    'testaid1',
    'testuid',
    0
);

INSERT INTO account (id, user_id, total) VALUES (
    'testaid2',
    'testuid',
    2300
);

INSERT INTO account (id, user_id, total) VALUES (
    'testaid3',
    'testuid2',
    1800
);

INSERT INTO transaction (id, amount, account_id, created_at) VALUES (
    'testtx',
    700,
    'testaid0',
    10
);

INSERT INTO transaction (id, amount, account_id, created_at) VALUES (
    'testtx2',
    100,
    'testaid1',
    50
);

INSERT INTO transaction (id, amount, account_id, created_at) VALUES (
    'testtx3',
    300,
    'testaid1',
    60
);
