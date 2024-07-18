create table med.users (
	id uuid primary key default gen_random_uuid(),
	username varchar(50) not null,
	password varchar(200) not null,
	created_at bigint not null
);
