CREATE TABLE public.tr_traveller (
	id int4 DEFAULT nextval('traveller_id_seq'::regclass) NOT NULL,
	"name" varchar(50) NOT NULL,
	rarity int2 NOT NULL,
	influence_id int4 NULL
);


-- public.tr_traveller foreign keys

ALTER TABLE public.tr_traveller ADD CONSTRAINT tr_traveller_m_influence_fk FOREIGN KEY (influence_id) REFERENCES public.m_influence(id);