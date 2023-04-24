CREATE TABLE pedido
(
    id    bigint unsigned not null primary key auto_increment,
    producto  VARCHAR(255)    NOT NULL,
    categoria VARCHAR(255)    NOT NULL,
    cantidad  INTEGER         NOT NULL
);
