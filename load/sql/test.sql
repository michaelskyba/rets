DROP TABLE IF EXISTS records;
DROP TABLE IF EXISTS residential_records;
DROP TABLE IF EXISTS condo_records;

CREATE TABLE residential_records (
	id			INT AUTO_INCREMENT NOT NULL,
	rets_date	DATE NOT NULL,
	A_c VARCHAR(128),
	Ml_num VARCHAR(128),
	PRIMARY KEY (`id`)
);

CREATE TABLE condo_records (
	id			INT AUTO_INCREMENT NOT NULL,
	rets_date	DATE NOT NULL,
	A_c VARCHAR(128),
	Ml_num VARCHAR(128),
	PRIMARY KEY (`id`)
);

INSERT INTO residential_records
	(rets_date, A_c, Ml_num)
VALUES
	("2020-01-01", "Central Air", "S9999999"),
	("2022-05-07", "Central Air", "S8888888");

INSERT INTO condo_records
	(rets_date, A_c, Ml_num)
VALUES
	("2021-05-11", "None", "S6666666"),
	("2022-02-17", "Central Air", "S7777777");
