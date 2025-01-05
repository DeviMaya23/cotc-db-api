CREATE TABLE public.m_user (
	id serial NOT NULL,
	username varchar(20) NOT NULL,
	"password" varchar(250) NOT NULL,
	token varchar(250) NOT NULL,
	created_at timestamptz NULL,
	created_by varchar(50) NULL,
	updated_at timestamptz NULL,
	updated_by varchar(50) NULL,
	deleted_at timestamptz NULL,
	deleted_by varchar(50) NULL,
	CONSTRAINT m_user_pkey PRIMARY KEY (id)
);

