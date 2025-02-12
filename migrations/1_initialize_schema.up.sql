create type user_role as enum('realtor', 'agency', 'private_person');

create table public.user (
  id uuid primary key,
  role user_role not null,
  name varchar(255) not null,
  surname varchar(255) not null,
  email varchar(255) unique not null,
  password varchar(255) not null,
  phone varchar(32) unique not null
);

create type building_type as enum('new', 'secondary');

create table public.apartment_building (
  id uuid primary key,
  type building_type not null
);

create type house_type as enum('cottage', 'country_house', 'townhouse', 'tiny_home');
create type property_type as enum('apartment', 'house');

create table public.property (
  id uuid primary key,
  owner_id uuid not null,
  property_type property_type not null,
  building_id uuid,
  house_type house_type,
  foreign key (owner_id) references public.user(id) on delete cascade,
  foreign key (building_id) references public.apartment_building(id) on delete cascade
);

create table public.property_details (
  property_id uuid primary key,
  address varchar(255) not null,
  apartment_number integer,
  floors int,
  rooms int,
  area numeric,
  foreign key(property_id) references public.property(id) on delete cascade
);

create table public.building_details (
  building_id uuid primary key,
  address varchar(255) not null,
  developer varchar(255),
  floors int,
  construction_year int,
  parking_place boolean,
  foreign key(building_id) references public.apartment_building(id) on delete cascade
);

create type offer_type as enum('rent', 'sale');
create type listing_status as enum ('sold', 'for_rent', 'available', 'draft');

create table public.listing (
  id uuid primary key,
  user_id uuid not null,
  property_id uuid not null,
  title varchar(255),
  city varchar(255),
  offer offer_type not null,
  price bigint,
  description text,
  status listing_status not null,
  created_at timestamp DEFAULT NOW(),
  foreign key (user_id) references public.user(id) on delete cascade, 
  foreign key (property_id) references public.property(id) on delete cascade
);
