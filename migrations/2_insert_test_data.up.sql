INSERT INTO public.user (id, role, name, surname, email, password, phone)
VALUES
('46bf6f7c-4f64-4134-81a4-0d9b5cb447c5', 'private_person', 'Brenda Mann', 'Melissa Perez', 'nsanchez@example.com', '_7PBgH5uZ7', '824-552-6702x52146'),
('ed15d801-c7dd-47b3-b6e5-7dc0097020fe', 'private_person', 'Lynn Lopez', 'Julie Tate', 'michelleewing@example.org', 'F62&2TJfbG', '+1-530-928-7150x412');

INSERT INTO public.apartment_building (id, type)
VALUES
('012b36fd-472c-49a5-9c45-ca46bfcef337', 'secondary'),
('99bf794d-94b9-4c33-88cf-d39670279e1c', 'new'),
('36c1dd19-570c-4c41-af9b-6c9cab1f4c44', 'secondary'),
('cdaf2027-32e8-4567-b5d1-d2220d7864a1', 'secondary'),
('26012eb0-7a4d-425b-9b13-5e2d37494eaf', 'secondary'),
('c5434720-ac72-453b-b28f-6e8b31b6b89f', 'secondary'),
('2093519c-7e20-421d-a5a8-79770ab6a243', 'secondary'),
('932364c7-dedb-44d2-8d55-408a8957df52', 'new'),
('3f54aaa5-e6c1-44c5-ab9e-4796a5ab6986', 'secondary'),
('2dba779f-22ba-4c35-a59b-147bcf4d0159', 'new');

INSERT INTO public.building_details (building_id, address, developer, floors, construction_year, parking_place)
VALUES
('012b36fd-472c-49a5-9c45-ca46bfcef337', 'PSC 5197, Box 8806
APO AA 20964', 'Elizabeth Ramirez', 11, 1977, 'True'),
('99bf794d-94b9-4c33-88cf-d39670279e1c', '7714 Hoover Forge
New Jeffrey, GU 38628', 'Kelly Wilkinson', 9, 1975, 'True'),
('36c1dd19-570c-4c41-af9b-6c9cab1f4c44', '5122 Phillips Landing
South Christopher, TX 71227', 'Bradley Figueroa', 7, 1971, 'True'),
('cdaf2027-32e8-4567-b5d1-d2220d7864a1', '218 Alison Common
East Heatherside, SC 89415', 'Kenneth Jones', 9, 1992, 'True'),
('26012eb0-7a4d-425b-9b13-5e2d37494eaf', '7972 Shane Knolls
North Travisland, NY 41149', 'Julie Walker', 1, 2016, 'True'),
('c5434720-ac72-453b-b28f-6e8b31b6b89f', '816 Rodney Streets
Wilsonport, GU 20122', 'James Perry', 15, 1991, 'True'),
('2093519c-7e20-421d-a5a8-79770ab6a243', '6460 Nicole Squares Apt. 606
Ashleyhaven, SD 32591', 'Spencer Stewart', 15, 1981, 'True'),
('932364c7-dedb-44d2-8d55-408a8957df52', '03369 Rhonda Row Suite 923
West Samantha, KY 11248', 'Megan Arellano', 20, 1978, 'False'),
('3f54aaa5-e6c1-44c5-ab9e-4796a5ab6986', '81820 Rios Expressway
South Joshuatown, AZ 43365', 'John Fernandez', 10, 2014, 'True'),
('2dba779f-22ba-4c35-a59b-147bcf4d0159', '065 Brown Highway
East Chelsea, PA 59386', 'Michelle Garcia MD', 13, 1972, 'False');

INSERT INTO public.property (id, owner_id, property_type, building_id, house_type)
VALUES
('04d88736-785c-409b-9962-95c7a0703fea', 'ed15d801-c7dd-47b3-b6e5-7dc0097020fe', 'apartment', '2dba779f-22ba-4c35-a59b-147bcf4d0159', NULL),
('b83f81eb-d4c5-4f2c-86b6-cee101c9bcbc', 'ed15d801-c7dd-47b3-b6e5-7dc0097020fe', 'apartment', '932364c7-dedb-44d2-8d55-408a8957df52', NULL),
('099499e3-9349-4c8c-b072-1e5efface117', '46bf6f7c-4f64-4134-81a4-0d9b5cb447c5', 'apartment', 'c5434720-ac72-453b-b28f-6e8b31b6b89f', NULL),
('b4487938-1c31-4855-8f3b-9f52c936c0f9', '46bf6f7c-4f64-4134-81a4-0d9b5cb447c5', 'apartment', '2093519c-7e20-421d-a5a8-79770ab6a243', NULL),
('07a4117c-148e-4f13-a964-f92ed6f7e7da', 'ed15d801-c7dd-47b3-b6e5-7dc0097020fe', 'house', NULL, 'cottage');

INSERT INTO public.property_details (property_id, address, apartment_number, floors, rooms, area)
VALUES
('04d88736-785c-409b-9962-95c7a0703fea', '653 Michael Lodge
New Emily, DC 89455', 94, 0, 4, 36),
('b83f81eb-d4c5-4f2c-86b6-cee101c9bcbc', '0632 Lynch Manors Suite 801
North Brendaport, AL 35059', 93, 0, 2, 44),
('099499e3-9349-4c8c-b072-1e5efface117', '563 Mendez Spring
Timothyside, AS 11953', 90, 0, 3, 67),
('b4487938-1c31-4855-8f3b-9f52c936c0f9', '1616 Brandon Road
West Danielland, CT 80189', 48, 0, 4, 71),
('07a4117c-148e-4f13-a964-f92ed6f7e7da', '76756 Christian Ford Suite 528
North Tracy, SC 79304', 0, 3, 5, 105);

INSERT INTO public.listing (id, user_id, property_id, title, city, offer, price, description, status)
VALUES
('e18765bb-eb52-4be8-ac4b-dfb402f057a4', 'ed15d801-c7dd-47b3-b6e5-7dc0097020fe', '07a4117c-148e-4f13-a964-f92ed6f7e7da', 'Modern.', 'West Jason', 'rent', 117955, 'Get how argue section middle.', 'for_rent'),
('d79dede0-21e4-4930-8aa9-529de4ef3f2d', '46bf6f7c-4f64-4134-81a4-0d9b5cb447c5', '04d88736-785c-409b-9962-95c7a0703fea', 'Economy.', 'Rojasbury', 'rent', 902607, 'Machine us cell real certainly.', 'for_rent');

