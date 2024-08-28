create table med.login_activity_log (
	id uuid primary key default gen_random_uuid(),
	client_id uuid not null,
	user_id uuid not null,
	refresh_token_hash varchar(65) not null,
	expire_time bigint not null,
	created_at bigint not null default extract(epoch from now())
)
