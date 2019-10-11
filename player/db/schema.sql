create table player_cursors (
  `id` varchar(255) not null,
  `last_event_id` bigint(20) not null default '0',
  `updated_at` datetime not null,

   primary key (id)
);