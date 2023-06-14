CREATE TABLE IF NOT EXISTS contratos (
  id SERIAL NOT NULL,
  tipo VARCHAR(50) NOT NULL UNIQUE,
  CONSTRAINT CONTRATO_PK PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS professor (
  id SERIAL NOT NULL,
  nome VARCHAR(100) NOT NULL,
  id_detalhes_contrato INTEGER NOT NULL,
  CONSTRAINT PROFESSOR_PK PRIMARY KEY (id),
  FOREIGN KEY (id_detalhes_contrato) REFERENCES contratos (id)
);
CREATE TABLE IF NOT EXISTS professor_clt (
  id SERIAL NOT NULL,
  salario FLOAT NOT NULL,
  CONSTRAINT CLT_PK PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS professor_horista (
  id SERIAL NOT NULL,
  valor_hora FLOAT NOT NULL,
  horas_trabalhadas INTEGER NOT NULL,
  CONSTRAINT HORISTA_PK PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS professor_pj (
  id SERIAL NOT NULL,
  contrato FLOAT NOT NULL,
  CONSTRAINT PJ_PK PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS contrato_professor (
  id SERIAL NOT NULL,
  id_contrato INTEGER NOT NULL,
  id_professor INTEGER NOT NULL,
  id_detalhes_contrato INTEGER NOT NULL,
  inativo BOOL NOT NULL,
  CONSTRAINT CONTRATO_PROFESSOR_PK PRIMARY KEY (id),
  FOREIGN KEY (id_contrato) REFERENCES contratos (id),
  FOREIGN KEY (id_professor) REFERENCES professor (id)
);

--- INSERTS

INSERT INTO contratos (tipo)
VALUES ('CLT'),
  ('HORISTA'),
  ('PJ');

INSERT INTO professor (nome, id_detalhes_contrato)
VALUES ('Montanha', 3),
  ('Calado', 3),
  ('Jairo', 2),
  ('Luciano', 1);

INSERT INTO professor_pj (contrato)
VALUES (17200.00),
  (22200.00);

INSERT INTO professor_horista (valor_hora, horas_trabalhadas)
VALUES (10.00, 150);

INSERT INTO professor_clt (salario)
VALUES (15000.00);

INSERT INTO contrato_professor (
    id_contrato,
    id_professor,
    id_detalhes_contrato,
    inativo
  )
VALUES
  ((SELECT id_detalhes_contrato FROM professor WHERE id=1), 1, 2, false),
  ((SELECT id_detalhes_contrato FROM professor WHERE id=2), 2, 1, false),
  ((SELECT id_detalhes_contrato FROM professor WHERE id=4), 4, 1, false),
  ((SELECT id_detalhes_contrato FROM professor WHERE id=3), 3, 1, false);