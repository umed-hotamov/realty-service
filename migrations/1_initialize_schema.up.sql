create type user_role as enum ('user', 'moderator');

create table public.user (
  id uuid primary key,
  email varchar(255) unique not null,
  password varchar(255) not null,
  phone varchar(32) not null,
  role user_role not null
);

create type offer_type as enum('rent', 'sale');
create type property_type as enum ('flat', 'private_house');
create type property_status as enum ('sold', 'for rent');
create type moderation_status as enum ('created', 'approved', 'declined', 'on moderation');

create table public.property (
  id uuid primary key,
  owner_id uuid not null,
  type property_type not null,
  offer offer_type not null,
  price numeric not null,
  p_status property_status,
  m_status moderation_status not null,
  foreign key (owner_id) references public.user(id) on delete cascade
);

create table public.apartment_building (
  id uuid primary key,
  year_built integer not null,
  address varchar(255) not null,
  developer varchar(255) not null
);

create table public.flat (
  id uuid primary key,
  property_id uuid unique not null,
  house_id uuid not null,
  flat_number integer not null,
  rooms integer not null,
  square numeric not null,
  foreign key (property_id) references public.property(id) on delete cascade,
  foreign key (house_id) references public.apartment_building(id) on delete cascade
);

create table public.private_house (
  id uuid primary key,
  property_id uuid unique not null,
  address varchar(255) not null,
  rooms integer not null,
  square numeric not null,
  foreign key (property_id) references public.property(id) on delete cascade
);

create table public.listing (
  id uuid primary key,
  user_id uuid not null,
  property_id uuid unique not null,
  title varchar(255),
  description text,
  status moderation_status not null,
  created_at timestamp not null,
  foreign key (property_id) references public.property(id) on delete restrict
);
