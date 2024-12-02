CREATE TABLE public.m_influence (
	id serial4 NOT NULL,
	"name" varchar(50) NOT NULL,
	CONSTRAINT m_influence_pk PRIMARY KEY (id)
);

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