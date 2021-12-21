-- creates
CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    passwd TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    id_user INTEGER NOT NULL,
    text TEXT NOT NULL,

    FOREIGN KEY (id_user) 
        REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS friends (
    id_user_first INTEGER NOT NULL,
	id_user_second INTEGER NOT NULL,
	date TEXT NOT NULL,

    FOREIGN KEY (id_user_first) 
        REFERENCES users (id), 
    FOREIGN KEY (id_user_second) 
        REFERENCES users (id)  
);

CREATE TABLE IF NOT EXISTS requests (
    id_user_first INTEGER NOT NULL,
	id_user_second INTEGER NOT NULL,

    FOREIGN KEY (id_user_first) 
        REFERENCES users (id),
    FOREIGN KEY (id_user_second) 
        REFERENCES users (id)   
);

-- inserts
INSERT INTO 
    users (id, username, passwd) 
    VALUES 
        (1, "admin", "123"),
        (2, "user1", "pass1"),
        (3, "user2", "pass2"),
        (4, "user3", "pass3")
    ON CONFLICT (id) DO UPDATE SET
        username = excluded.username,
        passwd = excluded.passwd;
