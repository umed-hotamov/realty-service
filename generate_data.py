import argparse
import random
import uuid
from faker import Faker

fake = Faker()

insert_user = """
INSERT INTO public.user (id, role, name, surname, email, password, phone)
VALUES
"""

insert_property = """
INSERT INTO public.property (id, owner_id, building_id, property_type, house_type)
VALUES
"""

insert_property_details = """
INSERT INTO public.property_details (property_id, address, apartment_number, floors, rooms, area)
VALUES
"""

insert_building = """
INSERT INTO public.apartment_building (id, type)
VALUES
"""

insert_building_details = """
INSERT INTO public.building_details (building_id, address, developer, floors, construction_year, parking_place)
VALUES
"""

insert_listing = """
INSERT INTO public.listing (id, user_id, property_id, title, city, offer, price, description, status)
VALUES
"""

user_roles = ['realtor', 'agency', 'private_person']
offer_type = ['rent', 'sale']
property_types = ['apartment', 'house']
house_types = ['cottage', 'country_house', 'townhouse', 'tiny_home']
listing_status = ['sold', 'for_rent', 'available', 'draft']
building_types = ['new', 'secondary']

users = []
properties = []
buildings = []

def construct_user_data():
    return dict(
            id=uuid.uuid4(),
            role=random.choice(user_roles),
            name=fake.name(),
            surname=fake.name(),
            email=fake.email(),
            password=fake.password(),
            phone=fake.phone_number(),
            )


def add_user_insert_values(user):
    return f"('{user['id']}', '{user['role']}', '{user['name']}', '{user['surname']}', "\
           f"'{user['email']}', '{user['password']}', '{user['phone']}'),\n"


def generate_user(user_num):
    stmt = insert_user
    for _ in range(user_num):
        user = construct_user_data()
        users.append(user)
        stmt += add_user_insert_values(user)

    return stmt.strip(',\n') + ';\n\n'


def construct_property_data():
    property_type = random.choice(property_types)
    return dict(
            id=uuid.uuid4(),
            owner_id=random.choice(users)['id'],
            building_id=f"'{random.choice(buildings)['id']}'" if property_type == 'apartment' else 'NULL',
            property_type=property_type,
            house_type=f"'{random.choice(house_types)}'" if property_type == 'house' else 'NULL',
            )


def construct_property_details(property):
    return dict(
        property_id=property['id'],
        address=fake.address(),
        apartment_number=random.randint(1, 100) if property['property_type'] == 'apartment' else 0,
        floors=random.randint(1, 5) if property['property_type'] == 'house' else 0,
        rooms=random.randint(1, 6),
        area=random.randint(30, 200),
    )


def add_property_insert_values(property):
    return f"('{property['id']}', '{property['owner_id']}', {property['building_id']}, '{property['property_type']}', "\
                      f"{property['house_type']}),\n"


def add_property_details_insert_values(property_details):
    return f"('{property_details['property_id']}', '{property_details['address']}', {property_details['apartment_number']}, {property_details['floors']}, "\
                     f"{property_details['rooms']}, {property_details['area']}),\n"


def generate_property(property_num):
    property_stmt = insert_property
    for _ in range(property_num):
        property = construct_property_data()
        properties.append(property)
        property_stmt += add_property_insert_values(property)
    property_stmt = property_stmt.strip(',\n') + ';\n\n'

    details_stmt = insert_property_details
    for p in properties:
        details = construct_property_details(p)
        details_stmt += add_property_details_insert_values(details)
    details_stmt = details_stmt.strip(',\n') + ';\n\n'

    return property_stmt + details_stmt  


def construct_building_data():
    return dict(
        id=uuid.uuid4(),
        type=random.choice(building_types),
    )


def construct_building_details(building_id):
    return dict(
        building_id=building_id,
        address=fake.address(),
        developer=fake.name(),
        floors=random.randint(1, 20),
        construction_year=fake.year(),
        parking_place=fake.boolean()
    )


def add_building_insert_values(building):
   return f"('{building['id']}', '{building['type']}'),\n"


def add_building_details_insert_values(building_details):
    return f"('{building_details['building_id']}', '{building_details['address']}', '{building_details['developer']}', " \
                     f"{building_details['floors']}, {building_details['construction_year']}, '{building_details['parking_place']}'),\n"


def generate_building(building_num):
    building_stmt = insert_building
    for _ in range(building_num):
        building = construct_building_data()
        buildings.append(building)
        building_stmt += add_building_insert_values(building)
    building_stmt = building_stmt.strip(',\n') + ';\n\n'

    details_stmt = insert_building_details
    for b in buildings:
        details = construct_building_details(b['id'])
        details_stmt += add_building_details_insert_values(details)
    details_stmt = details_stmt.strip(',\n') + ';\n\n'
        

    return building_stmt + details_stmt


def construct_listing_data():
    return dict(
            id=uuid.uuid4(),
            user_id=random.choice(users)['id'],
            property_id=random.choice(properties)['id'],
            title=fake.text(max_nb_chars=10),
            city=fake.city(),
            offer=random.choice(offer_type),
            price=random.randint(10000, 1000000),
            description=fake.text(max_nb_chars=50),
            status=random.choice(listing_status),
            )


def add_listing_insert_values(listing):
    return f"('{listing['id']}', '{listing['user_id']}', '{listing['property_id']}', "\
            f"'{listing['title']}', '{listing['city']}', '{listing['offer']}', "\
            f"{listing['price']}, '{listing['description']}', '{listing['status']}'),\n"


def generate_listing(listing_num):
    stmt = insert_listing
    for _ in range(listing_num):
        listing = construct_listing_data()
        stmt += add_listing_insert_values(listing)

    return stmt.strip(',\n') + ';\n\n'


def generate_data(user_num, property_num, building_num, listing_num, dest):
    if user_num:
        dest.write(generate_user(user_num))
    if building_num:
        dest.write(generate_building(building_num))
    if property_num:
        dest.write(generate_property(property_num))
    if listing_num:
        dest.write(generate_listing(listing_num))
    

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Parse info to construct sql statements')

    parser.add_argument('--users', type=int, default=2, help='generate users')
    parser.add_argument('--buildings', type=int, default=10, help='genereate buildings')
    parser.add_argument('--properties', type=int, default=5, help='generate properties')
    parser.add_argument('--listings', type=int, default=2, help='generate listings')
    parser.add_argument('--file', type=str, default='insert.sql', help='destination file')

    args = parser.parse_args()

    with open(args.file, 'w') as file:
        generate_data(args.users, args.properties, args.buildings, args.listings, file)
