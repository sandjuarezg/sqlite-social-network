-- creates
CREATE TABLE user (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    passwd TEXT NOT NULL
);

CREATE TABLE post (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    id_user INTEGER NOT NULL,
    text TEXT NOT NULL,

    FOREIGN KEY (id_user) 
        REFERENCES user (id)
);

CREATE TABLE friend (
    id_user_first INTEGER NOT NULL,
	id_user_second INTEGER NOT NULL,
	date TEXT NOT NULL,

    FOREIGN KEY (id_user_first) 
        REFERENCES user (id), 
    FOREIGN KEY (id_user_second) 
        REFERENCES user (id)  
);

CREATE TABLE request (
    id_user_first INTEGER NOT NULL,
	id_user_second INTEGER NOT NULL,

    FOREIGN KEY (id_user_first) 
        REFERENCES user (id),
    FOREIGN KEY (id_user_second) 
        REFERENCES user (id)   
);

-- inserts
INSERT INTO 
    user (username, passwd) 
    VALUES 
        ("sand", "123"),
        ("aaa", "passaaa"),
        ("bbb", "passbbb"),
        ("ccc", "passccc");

INSERT INTO 
    post (id_user, text) 
    VALUES 
        (1, "Hi, I'm Sand"),
        (1, "How are u"),
        (2, "insert text here");

INSERT INTO 
    friend (id_user_first, id_user_second, date) 
    VALUES 
        (1, 2, date('now')),
        (1, 3, date('now'));

INSERT INTO 
    request (id_user_first, id_user_second) 
    VALUES 
        (2, 3);
