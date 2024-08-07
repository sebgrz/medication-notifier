create table med.push_token (
	id uuid primary key default gen_random_uuid(),
	user_id uuid not null,
	token varchar(500) not null,
	created_at bigint not null default extract(epoch from now())
);
