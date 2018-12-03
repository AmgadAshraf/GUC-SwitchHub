/*CREATE TABLE signedup (
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
    didswitch BOOLEAN DEFAULT false
);

INSERT INTO signedup(fname, lname, id, email, userpassword)
VALUES
('Amgad','Ashraf','37-2058','amgadramses96@gmail.com','a'),
('Akram','Ashraf','37-2076','akramashraf96@gmail.com','a'),
('Mina','Rafik','34-2048','mina.r.mofeed@gmail.com','a'),
('Amgad','Yahoo','34-678','amgadramses@yahoo.com','a'),
('Howaida','Roman','37-1234','howaidaroman@yahoo.com','a'),
('Mar','Gerges','34-876','margergesyouth@gmail.com','a'),
('Veronica','Rafik','34-999','konka_rafik@hotmail.com','a'),
('Youssam','Joseph','34-5559','youssamjoseph@gmail.com','a'),
('Ebram','Nagy','37-5559','ebramnagyy@gmail.com','a'),
('Mina','Ishak','34-65459','minaishak10@gmail.com','a');

*/


/*DROP SCHEMA public CASCADE;
CREATE SCHEMA public;*/