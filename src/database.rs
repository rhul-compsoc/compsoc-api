use super::{
    models::*, schema::{
        admins::dsl::*, users::dsl::*, members::dsl::*,
    }
};
use diesel::pg::PgConnection;
use diesel::prelude::*;
use std::env;

pub fn establish_connection() -> PgConnection {

    let database_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");

    PgConnection::establish(&database_url)
        .unwrap_or_else(|_| panic!("Error connecting to {}", database_url))
}

pub fn read_people(
    limit: u64,
) -> Vec<Person> {
    let connection = &mut establish_connection();
    let mut limit: i64 = limit as i64;

    if limit == 0 {
        limit = std::i64::MAX;
    }

    users
        .limit(limit)
        .select(Person::as_select())
        .load(connection)
        .expect("Error loading people")
}

pub fn read_person(
    person_id: i32,
) -> Person {
    let connection = &mut establish_connection();

    users
        .find(person_id)
        .first::<Person>(connection)
        .expect("Error loading person")
}

pub fn add_person(
    person: Person,
) -> Result<(), diesel::result::Error> {
    let connection = &mut establish_connection();

    diesel::insert_into(users)
        .values(person)
        .execute(connection)?;

    Ok(())
}

pub fn update_person(
    person: Person,
) -> Result<(), diesel::result::Error> {
    let connection = &mut establish_connection();

    diesel::update(users.find(person.id))
        .set(&person)
        .execute(connection)?;

    Ok(())
}

pub fn delete_person(
    person_id: i32,
) -> Result<(), diesel::result::Error> {
    let connection = &mut establish_connection();

    diesel::delete(users.find(person_id))
        .execute(connection)?;

    Ok(())
}

pub fn check_user(
    person_id: i32,
) -> Result<bool, diesel::result::Error> {
    let connection = &mut establish_connection();

    let user_exists = users
        .find(person_id)
        .first::<Person>(connection)
        .is_ok();

    Ok(user_exists)
}

pub fn read_members(
    limit: u64,
) -> Vec<Member> {
    let connection = &mut establish_connection();
    let mut limit: i64 = limit as i64;

    if limit == 0 {
        limit = std::i64::MAX;
    }

    members
        .limit(limit)
        .select(Member::as_select())
        .load(connection)
        .expect("Error loading members")
}

pub fn read_member(
    member_id: i32,
) -> Member {
    let connection = &mut establish_connection();

    members
        .find(member_id)
        .first::<Member>(connection)
        .expect("Error loading member")
}

pub fn add_member(
    member: Member,
) -> Result<(), diesel::result::Error> {
    let connection = &mut establish_connection();

    diesel::insert_into(members)
        .values(member)
        .execute(connection)?;

    Ok(())
}

pub fn update_member(
    member: Member,
) -> Result<(), diesel::result::Error> {
    let connection = &mut establish_connection();

    diesel::update(members.find(member.id))
        .set(&member)
        .execute(connection)?;

    Ok(())
}

pub fn delete_member(
    member_id: i32,
) -> Result<(), diesel::result::Error> {
    let connection = &mut establish_connection();

    diesel::delete(members.find(member_id))
        .execute(connection)?;

    Ok(())
}

pub fn check_member(
    member_id: i32,
) -> Result<bool, diesel::result::Error> {
    let connection = &mut establish_connection();

    let member_exists = members
        .find(member_id)
        .first::<Member>(connection)
        .is_ok();

    Ok(member_exists)
}

pub fn check_admin(
    tok: String,
) -> Result<bool, diesel::result::Error> {
    let connection = &mut establish_connection();

    let user_exists = admins
        .find(tok)
        .first::<Admin>(connection)
        .is_ok();

    Ok(user_exists)
}
