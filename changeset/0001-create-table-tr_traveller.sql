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
