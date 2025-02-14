INSERT INTO public.user (id, role, name, surname, email, password, phone)
VALUES
('66a38dbc-fc24-49a5-93c6-084965b7f948', 'private_person', 'Elizabeth Wilson', 'William Stanley', 'tammygarrison@example.com', '&1*hlSyu5M', '7315767298'),
('42950f9d-7f73-4bae-af7a-e6d4c01caffa', 'private_person', 'Danielle Pierce', 'Alexander Lopez', 'mpark@example.org', '&5%JZCT^hC', '235-788-7668');

INSERT INTO public.apartment_building (id, type)
VALUES
('d14f285c-a3ec-46e2-a82c-1a4f780499be', 'secondary'),
('af91a583-767b-44dc-9569-d58b2d4689c2', 'new'),
('1b76437e-73b1-4fea-8d36-32fbbe5d8573', 'secondary'),
('0e05b4fa-8afa-44a8-9068-713368a76477', 'secondary'),
('475540c9-7358-48a8-aec0-a465fc8cc453', 'secondary'),
('a823d9d1-17a9-40f0-b171-78f4b3d66298', 'secondary'),
('76ee1a41-bb13-4065-93d6-c7c7f72e3761', 'secondary'),
('71d57dbe-df2d-41d1-989d-ca8eef7c046a', 'secondary'),
('4faffd45-0f0f-454f-9676-be9ff1ceddd5', 'new'),
('74f1466c-184d-4bff-b328-e8cfe57bab19', 'new');

INSERT INTO public.building_details (building_id, address, developer, floors, construction_year, parking_place)
VALUES
('d14f285c-a3ec-46e2-a82c-1a4f780499be', '297 Cynthia Flat
Baldwinhaven, MT 31585', 'Joshua Mason', 13, 1989, 'True'),
('af91a583-767b-44dc-9569-d58b2d4689c2', '46864 Arthur Grove
Jonesborough, SD 32578', 'Victoria Evans', 8, 1988, 'False'),
('1b76437e-73b1-4fea-8d36-32fbbe5d8573', '7594 Hardin Vista
East Williamtown, SD 33339', 'Jacob Stevens', 14, 1978, 'True'),
('0e05b4fa-8afa-44a8-9068-713368a76477', '77790 Amanda Falls
Sellersbury, AS 06148', 'Danielle Webb', 4, 2001, 'True'),
('475540c9-7358-48a8-aec0-a465fc8cc453', 'USS Wilson
FPO AA 18218', 'Robert Combs', 15, 1987, 'False'),
('a823d9d1-17a9-40f0-b171-78f4b3d66298', '854 Macdonald Alley
Markport, AS 23642', 'Linda Wilson', 2, 1974, 'False'),
('76ee1a41-bb13-4065-93d6-c7c7f72e3761', '6295 John Shoal
Salazarmouth, NE 16017', 'Allison Salazar', 7, 1978, 'False'),
('71d57dbe-df2d-41d1-989d-ca8eef7c046a', '331 Paul Roads
Port Michaelview, MD 59975', 'Albert Guerra', 8, 2021, 'False'),
('4faffd45-0f0f-454f-9676-be9ff1ceddd5', '9385 Bell Route Suite 817
North Ronald, HI 39070', 'Maria Brewer', 19, 1979, 'False'),
('74f1466c-184d-4bff-b328-e8cfe57bab19', '65068 Dickson Crescent
Lisaport, KY 93392', 'Amanda Powell', 10, 1989, 'True');

INSERT INTO public.property (id, owner_id, building_id, property_type, house_type)
VALUES
('5a2cb689-8648-4ce3-adfa-ccff938fdfad', '66a38dbc-fc24-49a5-93c6-084965b7f948', '1b76437e-73b1-4fea-8d36-32fbbe5d8573', 'apartment', NULL),
('1de52c9e-354e-4a4b-91b5-4d77317d3766', '42950f9d-7f73-4bae-af7a-e6d4c01caffa', NULL, 'house', 'cottage'),
('02e20c1a-4341-44ff-84d9-72352a7d55e3', '42950f9d-7f73-4bae-af7a-e6d4c01caffa', '4faffd45-0f0f-454f-9676-be9ff1ceddd5', 'apartment', NULL),
('9c06aac9-a9f6-4d50-8f20-715180b85dc3', '66a38dbc-fc24-49a5-93c6-084965b7f948', '74f1466c-184d-4bff-b328-e8cfe57bab19', 'apartment', NULL),
('96140c48-ff32-4485-9320-ee459bdef820', '66a38dbc-fc24-49a5-93c6-084965b7f948', '76ee1a41-bb13-4065-93d6-c7c7f72e3761', 'apartment', NULL);

INSERT INTO public.property_details (property_id, address, apartment_number, floors, rooms, area)
VALUES
('5a2cb689-8648-4ce3-adfa-ccff938fdfad', '7131 Torres Crest
Rodriguezhaven, FM 88887', 53, 0, 5, 161),
('1de52c9e-354e-4a4b-91b5-4d77317d3766', '9255 Ronald Hollow
North Michael, MH 98763', 0, 3, 2, 191),
('02e20c1a-4341-44ff-84d9-72352a7d55e3', '299 Cathy Glen
Port Amber, CA 87723', 74, 0, 1, 139),
('9c06aac9-a9f6-4d50-8f20-715180b85dc3', '245 Greg Haven Suite 599
Williamport, GU 45888', 39, 0, 1, 107),
('96140c48-ff32-4485-9320-ee459bdef820', '66403 Winters Fork Apt. 004
Randallhaven, WI 64264', 6, 0, 1, 93);

INSERT INTO public.listing (id, user_id, property_id, title, city, offer, price, description, status)
VALUES
('f6f4f067-2404-4b03-914f-d09267b67253', '66a38dbc-fc24-49a5-93c6-084965b7f948', '9c06aac9-a9f6-4d50-8f20-715180b85dc3', 'Manager.', 'Jenniferberg', 'sale', 914303, 'Face size local later wear whom occur.', 'draft'),
('29ad4da8-a8df-48da-b328-c660a626e198', '42950f9d-7f73-4bae-af7a-e6d4c01caffa', '9c06aac9-a9f6-4d50-8f20-715180b85dc3', 'We a list.', 'West Kimberly', 'sale', 262650, 'Protect economic leg full safe.', 'for_rent');

