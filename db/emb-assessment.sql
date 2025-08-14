-- emb lab golang assessment

CREATE DATABASE IF NOT EXISTS emb;

USE emb;

DROP TABLE books;
CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(100) NOT NULL,
    genre VARCHAR(50),
    description TEXT,
    isbn VARCHAR(20),
    image VARCHAR(255),
    published DATE,
    publisher VARCHAR(100),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    created_by VARCHAR(50) DEFAULT 'system',
    updated_by VARCHAR(50) DEFAULT 'system'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SELECT * FROM books ORDER BY id DESC;



