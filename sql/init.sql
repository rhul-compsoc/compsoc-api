create database compsoc;
\c compsoc;

create table public.users
(
    id   int                        not null
        primary key
        unique,
    name varchar(50) default 'Name' not null
);

alter table public.users
    owner to "user";

create table public.members
(
    id            integer                                                not null
        primary key
        constraint members_users_id_fk
            references public.users,
    student_id    varchar(50) default 'student_id'::character varying    not null
        unique,
    first_name    varchar(25) default 'first_name'::character varying    not null,
    last_name     varchar(25) default 'last_name'::character varying     not null,
    student_email varchar(50) default 'student_email'::character varying not null
        unique,
    active_member boolean     default false                              not null
);

alter table public.members
    owner to "user";

create table public.admins
(
    token varchar(64) not null
        constraint admin_pkey
            primary key
);

alter table public.admins
    owner to "user";
