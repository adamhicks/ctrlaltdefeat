create table player_cursors (
  `id` varchar(255) not null,
  `last_event_id` bigint(20) not null default '0',
  `updated_at` datetime not null,

   primary key (id)
);

create table player_rounds (
  id bigint not null auto_increment primary key,
  status int default 0,
  created_at datetime not null,
  updated_at datetime
);

create table player_rounds_events (
  id bigint not null auto_increment primary key,
  timestamp datetime not null,
  foreign_id bigint not null,
  type int not null
);
