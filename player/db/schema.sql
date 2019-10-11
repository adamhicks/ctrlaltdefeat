create table player_cursors (
  `id` varchar(255) not null,
  `last_event_id` bigint(20) not null default '0',
  `updated_at` datetime not null,

   primary key (id)
);

create table round_parts (
    `id` bigint not null auto_increment,

    `match_id` int not null,
    `round_id` int not null,
    `player_id` varchar(255) not null,
    `rank` int not null,
    `p1_part` int not null,
    `p2_part` int not null,
    `p3_part` int not null,
    `p4_part` int not null,

    primary key (id),
    unique by_match_round_player(`match_id`, `round_id`, `player_id`)
);
