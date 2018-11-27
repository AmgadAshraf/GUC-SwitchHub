/*CREATE TABLE signedup (
    fname VARCHAR(15) NOT NULL,
    lname VARCHAR(15) NOT NULL,
    id VARCHAR(10) UNIQUE NOT NULL,
    email VARCHAR(50) PRIMARY KEY ,
    userpassword VARCHAR(50) NOT NULL
);

CREATE TABLE met (
    id VARCHAR(10) UNIQUE NOT NULL,
    tutorialfrom VARCHAR(5) NOT NULL,
    tutorialto VARCHAR(5) NOT NULL,
    germanlevel INTEGER,
    englishlevel VARCHAR(10),
    email VARCHAR(50) PRIMARY KEY,
    remainingswitches int DEFAULT 2
);*/

INSERT INTO signedup VALUES
('Amgad','Ashraf','37-2058','amgadramses96@gmail.com','password123'),
('Akram','Ashraf','37-2076','akramashraf96@gmail.com','password12345');


/*DROP SCHEMA public CASCADE;
CREATE SCHEMA public;

*/