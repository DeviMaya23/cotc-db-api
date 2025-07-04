CREATE TABLE public.m_influence (
	id serial4 NOT NULL,
	"name" varchar(50) NOT NULL,
	created_at timestamptz NULL,
	created_by varchar(50) NULL,
	updated_at timestamptz NULL,
	updated_by varchar(50) NULL,
	deleted_at timestamptz NULL,
	deleted_by varchar(50) NULL,
	CONSTRAINT m_influence_pk PRIMARY KEY (id)
);

CREATE TABLE public.tr_traveller (
	id serial NOT NULL,
	"name" varchar(50) NOT NULL,
	rarity int2 NOT NULL,
	influence_id int4 NULL,
	created_at timestamptz NULL,
	created_by varchar(50) NULL,
	updated_at timestamptz NULL,
	updated_by varchar(50) NULL,
	deleted_at timestamptz NULL,
	deleted_by varchar(50) NULL,
	CONSTRAINT tr_traveller_pk PRIMARY KEY (id)
);


-- public.tr_traveller foreign keys

ALTER TABLE public.tr_traveller ADD CONSTRAINT tr_traveller_m_influence_fk FOREIGN KEY (influence_id) REFERENCES public.m_influence(id);

INSERT INTO public.m_influence
("name")
VALUES('Wealth');
INSERT INTO public.m_influence
("name")
VALUES('Power');
INSERT INTO public.m_influence
("name")
VALUES('Fame');
INSERT INTO public.m_influence
("name")
VALUES('Opulence');
INSERT INTO public.m_influence
("name")
VALUES('Dominance');
INSERT INTO public.m_influence
("name")
VALUES('Prestige');

INSERT INTO public.tr_traveller
("name", rarity, influence_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by)
VALUES('Fiore', 5, 3, NULL, NULL, NULL, NULL, NULL, NULL);