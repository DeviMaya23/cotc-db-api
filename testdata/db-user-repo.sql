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

INSERT INTO public.m_user
(username, "password", "token", created_at, created_by, updated_at, updated_by, deleted_at, deleted_by)
VALUES('isla', '$2a$14$D0JCfJ3BGUiLVfb0NkOdqudJFe64umW6nmcOl4nSloWtJkIb0M8QW', '1', NULL, NULL, NULL, NULL, NULL, NULL);
