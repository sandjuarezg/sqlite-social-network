-- creates
CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL,   
    passwd TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    text TEXT NOT NULL,
	date TEXT NOT NULL,

    FOREIGN KEY (user_id) 
        REFERENCES users (id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS friends (
    user_id_first INTEGER NOT NULL,
	user_id_second INTEGER NOT NULL,
	date TEXT NOT NULL,

    FOREIGN KEY (user_id_first) 
        REFERENCES users (id)
        ON DELETE CASCADE, 
    FOREIGN KEY (user_id_second) 
        REFERENCES users (id)  
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS requests (
    user_id_first INTEGER NOT NULL,
	user_id_second INTEGER NOT NULL, 

    FOREIGN KEY (user_id_first) 
        REFERENCES users (id)
        ON DELETE CASCADE,
    FOREIGN KEY (user_id_second) 
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
