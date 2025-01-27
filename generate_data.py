import argparse
import random
import uuid
from faker import Faker

fake = Faker()

insert_user          = 'insert into public.user (id, email, password, phone, role) values'
insert_property      = 'insert into public.property (id, owner_id, type, offer, price, p_status, m_status) values'
insert_flat          = 'insert into public.flat (id, property_id, house_id, flat_number, rooms, square)'
insert_private_house = 'insert into public.private_house (id, property_id, address, rooms, square)'
insert_building      = 'insert into public.apartment_building (id, year_built, address, developer)'

user_roles        = ['user', 'moderator']
offer_type        = ['rent', 'sale']
property_type     = ['flat', 'private_house']
moderation_status = ['created', 'approved', 'declined', 'on moderation']
property_status   = ['sold', 'for rent']

users = []
properties = []
houses = []

def construct_user_data():
    return dict(
            id=uuid.uuid4(),
            email=fake.email(),
            password=fake.password(),
            phone=fake.phone_number(),
            role=random.choice(user_roles),
            )

def add_user_insert_values(user):
    return f"('{user['id']}', '{user['email']}', '{user['password']}', '{user['phone']}', '{user['role']}'),\n"

def generate_user(user_num):
    stmt = insert_user
    for _ in range(user_num):
        user = construct_user_data()
        users.append(user)
        stmt += add_user_insert_values(user)

    return stmt.strip(',\n') + ';\n\n'


def construct_property_data():
    return dict(
            id=uuid.uuid4(),
            owner_id=random.choice(users)['id'],
            type=random.choice(property_type),
            offer=random.choice(offer_type),
            price=random.randint(10000, 10000000),
            p_status=random.choice(property_type),
            m_status=random.choice(moderation_status)
            )


def construct_flat_data(property_id):
    return dict(
            id=uuid.uuid4(),
            property_id=property_id,
            house_id=random.choice(houses)['id'],
            flat_number = random.randint(1, 50),
            rooms = random.randint(1, 5),
            square = random.randint(10, 100),
            )


def construct_private_house_data(property_id):
    return dict(
            id=uuid.uuid4(),
            property_id=property_id,
            address=fake.address(),
            rooms = random.randint(5, 10),
            square = random.randint(50, 100),
            )


def add_property_insert_values(property):
    return f"('{property['id']}', '{property['owner_id']}', '{property['type']}', '{property['offer']}', "\
                      f"'{property['price']}', '{property['p_status']}', '{property['m_status']}'),\n"


def add_flat_insert_values(flat):
    return f"('{flat['id']}', '{flat['property_id']}', '{flat['house_id']}', '{flat['flat_number']}', "\
                      f"'{flat['rooms']}', '{flat['square']}'),\n"


def add_private_house_insert_values(private_house):
    return f"('{private_house['id']}', '{private_house['property_id']}', '{private_house['address']}', "\
                      f"'{private_house['rooms']}', '{private_house['square']}'),\n"


def generate_property(property_num):
    property_stmt = insert_property
    for _ in range(property_num):
        property = construct_property_data()
        properties.append(property)
        property_stmt += add_property_insert_values(property)
    property_stmt = property_stmt.strip(',\n') + ';\n\n'

    flat_stmt = ''
    private_house_stmt = ''
    for p in properties:
        if p['type'] == 'flat':
            flat = construct_flat_data(p['id'])
            flat_stmt += add_flat_insert_values(flat)
        else:
            private_house = construct_private_house_data(p['id'])
            private_house_stmt += add_private_house_insert_values(private_house)

    if flat_stmt:
        flat_stmt = flat_stmt.strip(',\n') + ';\n\n'
    else:
        private_house_stmt = private_house_stmt.strip(',\n') + ';\n\n'

    return property_stmt + private_house_stmt + flat_stmt 


def construct_building_data():
    return dict(
            id=uuid.uuid4(),
            year_built=fake.year(),
            address=fake.address(),
            developer=fake.name(),
            )


def add_building_insert_values(building):
    return f"('{building['id']}', '{building['year_built']}', '{building['address']}', "\
                      f"'{building['developer']}'),\n"


def generate_building(building_num):
    stmt = insert_building
    for _ in range(building_num):
        building = construct_building_data()
        stmt += add_building_insert_values(building)

    return stmt.strip(',\n') + ';\n\n'


def get_random_timestamp():
    random_timestamp = fake.date_time_betwee(
    start_date="-1y",
    end_date="now",  
    tzinfo=None      
    )

    return random_timestamp.strftime("%Y-%m-%d %H:%M:%S")



def construct_listing_data():
    return dict(
            id=uuid.uuid4(),
            user_id=random.choice(users)['id'],
            property_id=random.choice(properties)['id'],
            title=fake.text(max_nb_chars=10),
            description=fake.text(max_nb_chars=50),
            status=random.choice(moderation_status),
            created_at=get_random_timestamp()
            )


def add_listing_insert_values(listing):
    return f"('{listing['id']}', '{listing['user_id']}', '{listing['property_id']}', "\
            f"'{listing['title']}', '{listing['description']}', '{listing['status']}', '{listing['created_at']}'),\n"


def generate_listing(listing_num):
    pass


def generate_data(user_num, property_num, building_num, listing_num, dest):
    if user_num:
        dest.write(generate_user(user_num))
    if property_num:
        dest.write(generate_property(property_num))
    if building_num:
        dest.write(generate_building(building_num))
    if listing_num:
        dest.wrtie(generate_listing(listing_num))
    

if __name__ == "main":
    parser = argparse.ArgumentParser(description='Parse info to construct sql statements')

    parser.add_argument('--users', type=int, default=5, help='generate users')
    parser.add_argument('--properties', type=int, default=5, help='generate properties')
    parser.add_argument('--buildings', type=int, default=5, help='genereate buildings')
    parser.add_argument('--listings', type=int, default=5, help='generate listings')
    parser.add_argument('--file', type=str, default='insert.sql', help='destination file')

    args = parser.parse_args()

    with open(args.file, 'w') as file:
        generate_data(args.users, args.properties, args.buildings, args.listings, file)
