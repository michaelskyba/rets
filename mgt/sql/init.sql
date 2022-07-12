CREATE TABLE records (
	id			INT AUTO_INCREMENT NOT NULL,

	# The "rets_" prefix represents metadata that we generated - i.e. not data
	# coming directly from the RETS-provided records

	# Which class of property we're pulling from
	rets_class	VARCHAR(128) NOT NULL,

	# Date of record entry
	rets_date	DATE NOT NULL,

	PRIMARY KEY (`id`)
);

INSERT INTO accounts
	(rets_class, rets_date)
VALUES
	('CondoProperty', '2020-01-01'),
	('ResidentialProperty', '2022-05-07');
