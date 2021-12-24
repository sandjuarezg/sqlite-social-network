-- creates
CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL,   
    passwd TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    id_user INTEGER NOT NULL,
    text TEXT NOT NULL,
	date TEXT NOT NULL,

    FOREIGN KEY (id_user) 
        REFERENCES users (id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS friends (
    id_user_first INTEGER NOT NULL,
	id_user_second INTEGER NOT NULL,
	date TEXT NOT NULL,

    FOREIGN KEY (id_user_first) 
        REFERENCES users (id)
        ON DELETE CASCADE, 
    FOREIGN KEY (id_user_second) 
        REFERENCES users (id)  
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS requests (
    id_user_first INTEGER NOT NULL,
	id_user_second INTEGER NOT NULL, 

    FOREIGN KEY (id_user_first) 
        REFERENCES users (id)
        ON DELETE CASCADE,
    FOREIGN KEY (id_user_second) 
        REFERENCES users (id)
        ON DELETE CASCADE   
);

-- inserts
INSERT OR IGNORE INTO 
    users (email, username, passwd) 
    VALUES 
        ("admin@gmail.com", "admin", "123"),
        ("user1@gmail.com", "user1", "pass1"),
        ("user2@gmail.com", "user2", "pass2"),
        ("user3@gmail.com", "user3", "pass3");
