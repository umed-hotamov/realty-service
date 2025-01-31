insert into public.user (id, email, password, phone, role) values
('a6057299-6c3d-4be1-b77b-d707f5eec16c', 'audreyweaver@example.org', 'X^X3qQLx+b', '691.211.2462x28030', 'moderator'),
('399118c9-9a81-47f8-9fc4-4621d6290fc0', 'elizabeth09@example.net', '%JGSzGIYl6', '558.640.4862x733', 'seller');

insert into public.apartment_building (id, year_built, address, developer) values
('c08c8a41-9ec8-49e0-a06b-ce45802a406e', '1977', '281 Tina Squares
Markfort, MA 64529', 'Christopher Torres'),
('103216f8-1b84-4340-971c-62be158b9821', '2009', '7039 Mills Vista Suite 977
Cruzhaven, IA 33297', 'Billy Sawyer'),
('7e770104-a5a7-4bc6-98a9-c9d943261f72', '1980', '54019 Andrew Road
Lake Brian, CA 75301', 'Martin Davis'),
('e4459525-ef4e-47e3-a2ab-86f7c5d0d159', '1986', '0617 Amy Lodge
East Andrew, OR 69217', 'Nicole Beard'),
('d4db8dad-6e00-4fa5-8d44-c3c37a0f6e4f', '2004', '796 Anthony Manors Suite 634
Jenniferton, MT 76701', 'James Adams'),
('9bf63377-e904-4d3c-a5c3-b54093b98609', '2022', '673 Martinez Dam
Dominiqueshire, DE 12267', 'Jennifer Nelson'),
('9af8e4cf-efb0-4631-94d4-fbc4293d4d95', '2011', 'PSC 7711, Box 9145
APO AE 22208', 'George Adams'),
('b6ac9f28-9626-4842-aa24-3d4f20b231fc', '1970', '9575 Hoffman Parks
South Michelle, HI 00536', 'Warren Gardner'),
('791bccf9-3869-4b62-b38f-431287ae6087', '2006', '683 Regina Lodge Apt. 757
Nicholasside, NE 65687', 'Nicole Stuart'),
('9c960ef6-0f7f-4519-baeb-bf9ca2c06be9', '1999', '6193 Jose Forest Apt. 183
Paigefurt, CA 29936', 'Stephanie Johnson');

insert into public.property (id, owner_id, type, offer, price, p_status, m_status) values
('b514ae21-a72b-47d0-a375-b648d1cff671', '399118c9-9a81-47f8-9fc4-4621d6290fc0', 'flat', 'sale', '4797050', 'alive', 'on moderation'),
('c808c54c-dfc2-4381-93e1-e43175dcbe46', 'a6057299-6c3d-4be1-b77b-d707f5eec16c', 'private_house', 'sale', '1946183', 'alive', 'on moderation'),
('e3f6e5c7-ca32-4cb1-ba1f-5e3cc6fbeb7c', 'a6057299-6c3d-4be1-b77b-d707f5eec16c', 'private_house', 'sale', '1870403', 'alive', 'on moderation'),
('8541d2f9-0003-4658-99a9-09cf79d28d3f', 'a6057299-6c3d-4be1-b77b-d707f5eec16c', 'private_house', 'sale', '3934343', 'alive', 'on moderation'),
('16a2dea1-6e6c-43ce-b352-32625c3de323', '399118c9-9a81-47f8-9fc4-4621d6290fc0', 'flat', 'rent', '4744902', 'alive', 'on moderation');

insert into public.private_house (id, property_id, address, rooms, square) values
('f9b1af81-c2a8-4414-bd0d-03d5f95ebbbe', 'c808c54c-dfc2-4381-93e1-e43175dcbe46', '193 Rebecca Extensions
New Kristine, VA 12016', '10', '84'),
('0add2441-67f6-4521-a7c5-0a36f3cebf0e', 'e3f6e5c7-ca32-4cb1-ba1f-5e3cc6fbeb7c', '42894 Sutton Court
Amyfort, IL 28134', '8', '87'),
('cb5c34b2-dcef-4492-bf79-82c179177fdd', '8541d2f9-0003-4658-99a9-09cf79d28d3f', '78966 Leach Circles
Port Mark, WI 27421', '9', '95');

insert into public.flat (id, property_id, house_id, flat_number, rooms, square) values
('694e8f5f-25f3-4d58-ae6c-3b0c465639b0', 'b514ae21-a72b-47d0-a375-b648d1cff671', '9c960ef6-0f7f-4519-baeb-bf9ca2c06be9', '15', '1', '92'),
('c899ba6b-b40e-4495-9c46-c3d5c6f0ba90', '16a2dea1-6e6c-43ce-b352-32625c3de323', '791bccf9-3869-4b62-b38f-431287ae6087', '31', '5', '53');

insert into public.listing (id, user_id, property_id, title, description, status, created_at) values
('ef545f09-df5a-4f75-be46-9ef1f98f3e63', 'a6057299-6c3d-4be1-b77b-d707f5eec16c', 'e3f6e5c7-ca32-4cb1-ba1f-5e3cc6fbeb7c', 'Amount.', 'Player site community. Lot sign deal discover.', 'on moderation', '2024-06-25 00:11:52'),
('06c45bb1-ba8e-4acc-bc1b-03a22bc29936', 'a6057299-6c3d-4be1-b77b-d707f5eec16c', 'c808c54c-dfc2-4381-93e1-e43175dcbe46', 'Design.', 'Most cause get high bank large. Half must open.', 'on moderation', '2025-01-31 14:51:06');

