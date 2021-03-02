drop table if exists people;

CREATE TABLE people
(
    id SERIAL PRIMARY KEY,
    cpf VARCHAR(20) NOT NULL,
    private                         BOOLEAN,
    incompleto                      BOOLEAN,
    data_ultima_compra              DATE,
    ticket_medio                    DOUBLE PRECISION,
    ticket_ultima_compra            DOUBLE PRECISION,
    cnpj_loja_mais_frequente        VARCHAR(20),
    cnpj_loja_ultima_compra         VARCHAR(20)
);

drop table if exists error;

CREATE TABLE error
(
    id SERIAL PRIMARY KEY,
    chave           VARCHAR(1000),
    message         VARCHAR(1000)
);