CREATE TABLE signedup (
    fname VARCHAR NOT NULL,
    lname VARCHAR NOT NULL,
    id VARCHAR UNIQUE NOT NULL,
    email VARCHAR PRIMARY KEY ,
    userpassword VARCHAR NOT NULL
);

CREATE TABLE switching (
    id VARCHAR UNIQUE NOT NULL,
    major VARCHAR  NOT NULL,
    tutorialfrom VARCHAR NOT NULL,
    tutorialto VARCHAR NOT NULL,
    germanlevel VARCHAR NOT NULL,
    englishlevel VARCHAR NOT NULL,
    email VARCHAR PRIMARY KEY NOT NULL,
    didswitch BOOLEAN DEFAULT false,
    sentconfirmationemail BOOLEAN DEFAULT false
);

/*INSERT INTO signedup(fname, lname, id, email, userpassword)
VALUES
('Akram','Ashraf','37-2076','akramashraf96@gmail.com','pass123'),
('Amgad','Ashraf','37-2058','amgadramses96@gmail.com','pass12');
*/

/*DROP SCHEMA public CASCADE;
CREATE SCHEMA public;*/