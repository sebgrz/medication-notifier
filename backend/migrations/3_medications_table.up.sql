create table med.medications (
	id uuid primary key default gen_random_uuid(),
	user_id uuid not null,
	name varchar(1000) not null,
	day varchar(2) not null,
	time_of_day varchar(3) not null,
	created_at bigint not null default extract(epoch from now())
);
