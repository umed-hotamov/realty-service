insert into public.user (id, email, password, phone, role) values
('55edd9dc-a142-44cc-9080-ff601ddb537f', 'wendy88@example.net', '#5IVnU0m!U', '+1-265-480-3340x1271', 'seller'),
('dbb219bd-2cec-49e6-859d-3c8c8f285fd0', 'deanna72@example.com', '7VODXXvn$2', '001-775-918-9112x0227', 'moderator');

insert into public.apartment_building (id, year_built, address, developer) values
('8dd1e33a-ef4b-41a2-a6a1-b5d758653faa', '2016', '36334 Michele Loaf Suite 327
Lake Mario, DC 14360', 'John Berry'),
('6cfb11ae-2ed6-4a23-8d09-716d13b13ea7', '2003', '5080 Marc Mount
Lake Steven, KY 16826', 'Paul Alexander'),
('5b7abe34-9b4d-4c30-9793-bcc6c42db1eb', '2022', '132 Robert Road Suite 239
East Tiffanyfurt, MP 10835', 'Desiree Pena'),
('d37ff292-ff41-489a-bb23-7b93c83825a3', '1973', 'PSC 2976, Box 1479
APO AP 00894', 'Patricia Mitchell'),
('7ccc5352-233b-40eb-8d57-145fcd601ef1', '1972', '25004 Jennifer Crest
Bennettborough, SC 58605', 'Christopher Cummings'),
('a4570772-c4d1-4136-90ae-70e708583abd', '1978', '7028 Lisa Village
New Melissaside, MT 17249', 'Glenn Cook'),
('ba086b89-3137-4577-b81c-d66abe729e8e', '1995', '3733 Miller Highway Apt. 877
North Saratown, TN 39734', 'Tina Knapp'),
('d84df525-d38d-45f1-a850-d990a3c48c45', '1995', 'Unit 0429 Box 9546
DPO AE 53251', 'Susan Owen'),
('c1ec07d1-fe6e-40b1-9140-bfc6aeb273ab', '2002', '157 Hogan Street
South Abigail, MT 60704', 'Christine Johnson'),
('8442520a-772a-4fed-b078-dee35194b643', '2016', '3662 Natalie Via Suite 763
West Jaredport, NE 87010', 'Jesus Thomas');

insert into public.property (id, owner_id, type, offer, price, p_status, m_status) values
('951984ba-bba4-4e83-a4f7-2625a6d288e0', '55edd9dc-a142-44cc-9080-ff601ddb537f', 'private_house', 'sale', '5834483', 'none', 'on moderation'),
('a9723afc-3e5c-4d3c-9807-b0e7766c5ec4', 'dbb219bd-2cec-49e6-859d-3c8c8f285fd0', 'flat', 'sale', '8155403', 'none', 'on moderation');

insert into public.private_house (id, property_id, address, rooms, square) values
('86edaca6-4089-4e0f-bbef-c5ae76cf8376', '951984ba-bba4-4e83-a4f7-2625a6d288e0', 'Unit 7960 Box 4697
DPO AP 53952', '7', '82');

insert into public.flat (id, property_id, house_id, flat_number, rooms, square) values
('b3e03168-d1bb-426c-abb6-07c987a6fed9', 'a9723afc-3e5c-4d3c-9807-b0e7766c5ec4', 'c1ec07d1-fe6e-40b1-9140-bfc6aeb273ab', '44', '1', '80');

insert into public.listing (id, user_id, property_id, title, description, status, created_at) values
('27215e91-0673-4728-bee0-1314f9c692f6', '55edd9dc-a142-44cc-9080-ff601ddb537f', 'a9723afc-3e5c-4d3c-9807-b0e7766c5ec4', 'Example.', 'Purpose instead must goal throw itself.', 'on moderation', '2024-06-11 08:16:52'),
('4bfd8cf0-4e9f-4742-9177-5ed2cd09c211', '55edd9dc-a142-44cc-9080-ff601ddb537f', '951984ba-bba4-4e83-a4f7-2625a6d288e0', 'Night.', 'Off tend list range accept front.', 'on moderation', '2024-04-07 20:42:44');

