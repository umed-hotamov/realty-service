insert into public.user (id, email, password, phone, role) values
('47a2e3bd-44dd-4a14-914f-24ebf762947f', 'whitetodd@example.com', 'x7Y1SrFR#R', '(460)939-9766x12628', 'moderator'),
('bb9f8cc5-13de-4431-b8a9-63bfc241de9a', 'psimmons@example.com', '^_F3sVGxQf', '+1-641-282-4974', 'moderator');

insert into public.apartment_building (id, year_built, address, developer) values
('5400658c-db40-40cf-8179-16ef68b52f95', '1977', '9128 Thompson Brooks
Rosariostad, NH 64647', 'Misty Castillo'),
('98006a67-6f46-4377-abc7-d3abbea4a89c', '1999', '9693 Annette Manors Apt. 380
Andrewshire, CT 75178', 'Victoria Gonzales'),
('d6517438-ed6d-4147-a0a3-e4d6d059431b', '2013', '816 Chapman Port Suite 655
New Catherineborough, MH 88061', 'Craig Jackson'),
('f11685b7-5d82-4cde-bbc4-33aed587b195', '2023', 'PSC 6245, Box 7025
APO AE 26440', 'Ana Booth'),
('d2bb04df-91ed-4e03-9fe7-d3787bb9665f', '1981', '864 Ashley Knolls
North Alyssaside, NH 44578', 'Carla Wright'),
('e30ed284-8cff-418f-acdd-6fe97ca76a54', '2018', '3733 Morris Hill
South Timothy, KY 93683', 'Anthony Smith'),
('9ab10ca4-321a-4a92-b22f-b827323ddeb0', '1999', '31625 Tran Points
Wrightburgh, MI 14988', 'Jesse Jones'),
('43888cc3-7eee-4de7-bfc3-aa42d5ef4020', '2024', '21836 Keith Mount
Krystalfort, ID 11823', 'Jaime Powers DVM'),
('8b5012c4-5b4c-4f8c-a30b-4e1c118e849f', '2021', '42182 Ortiz Cliff
South Darrenside, MA 30881', 'Nancy Henderson'),
('09cc7b89-ff19-4749-b710-188c7abe0ca6', '1987', '40852 Michael Forges
Thompsonburgh, TN 06807', 'Matthew Lewis');

insert into public.property (id, owner_id, type, offer, price, p_status, m_status) values
('1f20845d-8f62-48c8-88eb-ee34325df6f3', 'bb9f8cc5-13de-4431-b8a9-63bfc241de9a', 'private_house', 'rent', '7696548', 'none', 'on moderation'),
('55ca50d5-f742-48d0-a0c6-2b6dd1df030f', '47a2e3bd-44dd-4a14-914f-24ebf762947f', 'flat', 'sale', '3524653', 'none', 'on moderation');

insert into public.private_house (id, property_id, address, rooms, square) values
('939a41b4-13e9-4683-83eb-16694de337e0', '1f20845d-8f62-48c8-88eb-ee34325df6f3', '5037 Madison Fork Apt. 160
Victoriaton, OR 98143', '8', '78');

insert into public.flat (id, property_id, house_id, flat_number, rooms, square) values
('3eef206d-4525-4943-8bfa-32c3172116da', '55ca50d5-f742-48d0-a0c6-2b6dd1df030f', '43888cc3-7eee-4de7-bfc3-aa42d5ef4020', '34', '1', '79');

insert into public.listing (id, user_id, property_id, title, description, status, created_at) values
('5ea63a6a-5bcf-45a3-b01c-48a2d932f6e7', 'bb9f8cc5-13de-4431-b8a9-63bfc241de9a', '1f20845d-8f62-48c8-88eb-ee34325df6f3', 'Yard who.', 'Their probably change drug each marriage.', 'on moderation', '2024-04-12 20:28:31'),
('53c29e27-3bd8-4c56-ae40-ff063300b48f', '47a2e3bd-44dd-4a14-914f-24ebf762947f', '55ca50d5-f742-48d0-a0c6-2b6dd1df030f', 'Deep.', 'Wrong ground campaign laugh.', 'on moderation', '2024-12-01 04:38:13');

