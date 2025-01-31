insert into public.user (id, email, password, phone, role) values
('368cd392-489d-4601-b4d8-3cd2f2403272', 'tammie90@example.org', 'B_H(1OflO@', '(433)776-2116x147', 'moderator'),
('d3908aa0-ac19-479c-ad89-dd9b942e8b59', 'hubbardbrenda@example.org', '894KnWXp_2', '+1-958-591-9238x376', 'moderator');

insert into public.apartment_building (id, year_built, address, developer) values
('691f65a9-97b9-422f-9679-09f7dd43d827', '2015', '30499 Matthew Locks
Lake Zachary, FL 81163', 'Daniel Gomez'),
('ecb0bec3-7d02-44ff-923d-4fc76c84d738', '1976', '66880 Johnson Stravenue
New Jennifer, OH 69383', 'Meagan Garrett'),
('87d0c6bd-69f9-4655-b97e-dba421b78012', '2015', 'USCGC Cook
FPO AA 66985', 'Nicole Rodriguez'),
('922efb0e-fd07-4657-a7a4-4426dbce3207', '2020', '0525 Kayla Dale Apt. 196
Samanthaville, MO 54561', 'David Pruitt'),
('38c80a59-6854-4d03-b740-59c821b0b5ba', '2011', '3098 Choi Union Suite 599
Rushborough, MN 08288', 'Kirk Perry'),
('0ba7b176-32ae-4716-a7e1-6d00b481afb2', '2018', '69811 Kyle Streets
New Clintonhaven, VI 27695', 'William Freeman'),
('30563eab-5205-41eb-a22c-dfb9809ef238', '1991', '0684 Vargas Station Apt. 005
Tamaramouth, AS 07376', 'Dustin Morgan'),
('cc67dfa6-9dd6-4760-b13f-7a6557c41fc4', '1975', '14852 Lee Crossing
Lake Toddport, DC 89257', 'Nathaniel Smith'),
('b3d435af-a4a3-41df-b44d-78d51cec08f0', '2008', '362 Martinez Mall Apt. 745
East Dennisland, ID 11902', 'Amanda King'),
('9ac8f55a-a9ff-49e6-9702-7b2befe70f96', '2004', '238 Fry Spurs
North Angela, PW 56164', 'Christopher Thompson');

insert into public.property (id, owner_id, type, offer, p_status, m_status, price, rooms, square) values
('aa5bcafb-59eb-4996-baae-0bb86c75021d', '368cd392-489d-4601-b4d8-3cd2f2403272', 'private_house', 'sale', 'alive', 'on moderation', '3387611', '10', '60'),
('30c623b2-0d76-4cc7-bdc0-969918b82540', '368cd392-489d-4601-b4d8-3cd2f2403272', 'private_house', 'sale', 'alive', 'on moderation', '1827562', '8', '68'),
('2966e123-bdfb-4421-844d-74475b4fa872', 'd3908aa0-ac19-479c-ad89-dd9b942e8b59', 'flat', 'sale', 'alive', 'on moderation', '2029864', '5', '89'),
('05b85c7e-a00c-4a4b-88d6-8bbfd0515c95', '368cd392-489d-4601-b4d8-3cd2f2403272', 'private_house', 'sale', 'alive', 'on moderation', '4219774', '10', '68'),
('fe83b323-719c-4fff-bff9-f9f316058555', '368cd392-489d-4601-b4d8-3cd2f2403272', 'flat', 'sale', 'alive', 'on moderation', '2246826', '6', '89');

insert into public.private_house (id, property_id, address) values
('0545cd31-c7a7-4761-b541-2d7e2a3fd61a', 'aa5bcafb-59eb-4996-baae-0bb86c75021d', 'USNS Miranda
FPO AE 50435'),
('7e6b0382-1e11-418b-878f-6cd22ab94ef2', '30c623b2-0d76-4cc7-bdc0-969918b82540', '6986 Thomas Drive
Parkview, PR 72373'),
('7f42fc27-cf46-4de7-8575-10b4eca592ce', '05b85c7e-a00c-4a4b-88d6-8bbfd0515c95', '04809 Contreras Ways Suite 402
Churchmouth, PR 30125');

insert into public.flat (id, property_id, house_id, flat_number) values
('19e8794f-ef5a-492c-b406-87b751d35b31', '2966e123-bdfb-4421-844d-74475b4fa872', '38c80a59-6854-4d03-b740-59c821b0b5ba', '2'),
('dfdc750f-66c1-461a-a26b-71d88a062cc9', 'fe83b323-719c-4fff-bff9-f9f316058555', '87d0c6bd-69f9-4655-b97e-dba421b78012', '40');

insert into public.listing (id, user_id, property_id, title, description, status, created_at) values
('b94761a4-dd4b-494e-90e6-1a58c77d66e0', 'd3908aa0-ac19-479c-ad89-dd9b942e8b59', '05b85c7e-a00c-4a4b-88d6-8bbfd0515c95', 'Total.', 'Many prove day better I wind.', 'on moderation', '2024-12-08 05:13:25'),
('bcd2062c-7820-4800-adf9-a72a11bb2355', '368cd392-489d-4601-b4d8-3cd2f2403272', 'aa5bcafb-59eb-4996-baae-0bb86c75021d', 'Stuff.', 'Race poor sea set everyone always contain.', 'on moderation', '2024-08-08 11:54:20');

