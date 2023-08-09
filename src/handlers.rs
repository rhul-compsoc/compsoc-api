use super::{
    models::{Person, ListOptions, Member},
    database,
    fetch
};
use serde_json::Value;
use std::convert::Infallible;
use warp::http::StatusCode;

pub async fn list_people(
    opts: ListOptions,
) -> Result<impl warp::Reply, Infallible> {
    log::debug!("list_people, {:?}", opts);
    let offset: u64 = opts.offset.unwrap_or_default() as u64;
    let limit: u64 = opts.limit.unwrap_or_default() as u64;

    let peoples: Vec<Person> = database::read_people(offset + limit)
        .clone()
        .into_iter()
        .skip(opts.offset.unwrap_or(0))
        .take(opts.limit.unwrap_or(std::usize::MAX))
        .collect();

    Ok(warp::reply::json(&peoples))
}

pub async fn get_person(
    id: i32,
) -> Result<impl warp::Reply, Infallible> {
    log::debug!("get_person: {:?}", id);
    let exists = database::check_user(id);

    if !exists.unwrap() {
        log::debug!("    -> person not found");
        let null_value: Value = serde_json::Value::Null;
        return Ok(warp::reply::json(&null_value))
    }

    let person: Person = database::read_person(id);
    Ok(warp::reply::json(&person))
}

pub async fn create_person(
    create: Person,
) -> Result<impl warp::Reply, Infallible> {
    log::debug!("create_person: {:?}", create);
    let exists = database::check_user(create.id);

    if exists.unwrap() {
        return Ok(StatusCode::BAD_REQUEST)
    }

    let _ = database::add_person(create);
    Ok(StatusCode::CREATED)
}

pub async fn patch_person(
    id: i32,
    update: Person,
) -> Result<impl warp::Reply, Infallible> {
    log::debug!("update_person: id={}, person={:?}", id, update);
    let exists = database::check_user(id);

    if !exists.unwrap() {
        log::debug!("    -> person id not found");
        return Ok(StatusCode::NOT_FOUND)
    }

    let _ = database::update_person(update);
    Ok(StatusCode::OK)
}

pub async fn put_person(
    id: i32,
    update: Person,
) -> Result<impl warp::Reply, Infallible> {
    log::debug!("update_person: id={}, person={:?}", id, update);
    let exists = database::check_user(id);

    if !exists.unwrap() {
        let _ = database::add_person(update);
        return Ok(StatusCode::CREATED)
    }

    let _ = database::update_person(update);
    Ok(StatusCode::OK)
}

pub async fn delete_person(
    id: i32,
) -> Result<impl warp::Reply, Infallible> {
    log::debug!("delete_person: id={}", id);

    let mut exists = database::check_user(id);

    if !exists.unwrap() {
        log::debug!("    -> person id not found");
        return Ok(StatusCode::NOT_FOUND)
    }

    let _ = database::delete_person(id);
    exists = database::check_user(id);

    if exists.unwrap() {
        log::debug!("    -> failed to delete");
        Ok(StatusCode::BAD_REQUEST)
    } else {
        Ok(StatusCode::NO_CONTENT)
    }
}

pub async fn list_member(
    opts: ListOptions,
) -> Result<impl warp::Reply, Infallible> {
    log::debug!("list_member, {:?}", opts);
    let offset: u64 = opts.offset.unwrap_or_default() as u64;
    let limit: u64 = opts.limit.unwrap_or_default() as u64;

    let members: Vec<Member> = database::read_members(offset + limit)
        .clone()
        .into_iter()
        .skip(opts.offset.unwrap_or(0))
        .take(opts.limit.unwrap_or(std::usize::MAX))
        .collect();

    Ok(warp::reply::json(&members))
}

pub async fn get_member(
    id: i32,
) -> Result<impl warp::Reply, Infallible> {
    log::debug!("get_member: {:?}", id);
    let exists = database::check_member(id);

    if !exists.unwrap() {
        log::debug!("    -> member not found");
        let null_value: Value = serde_json::Value::Null;
        return Ok(warp::reply::json(&null_value))
    }

    let member: Member = database::read_member(id);
    Ok(warp::reply::json(&member))
}

pub async fn create_member(
    create: Member,
) -> Result<impl warp::Reply, Infallible> {
    log::debug!("create_member: {:?}", create);
    let exists = database::check_member(create.id);

    if exists.unwrap() {
        return Ok(StatusCode::BAD_REQUEST)
    }

    let _ = database::add_member(create);
    Ok(StatusCode::CREATED)
}

pub async fn patch_member(
    id: i32,
    update: Member,
) -> Result<impl warp::Reply, Infallible> {
    log::debug!("update_memmber: id={}, member={:?}", id, update);
    let exists = database::check_member(id);

    if !exists.unwrap() {
        log::debug!("    -> member id not found");
        return Ok(StatusCode::NOT_FOUND)
    }

    let _ = database::update_member(update);
    Ok(StatusCode::OK)
}

pub async fn put_member(
    id: i32,
    update: Member,
) -> Result<impl warp::Reply, Infallible> {
    log::debug!("update_member: id={}, member={:?}", id, update);
    let exists = database::check_member(id);

    if !exists.unwrap() {
        let _ = database::add_member(update);
        return Ok(StatusCode::CREATED)
    }

    let _ = database::update_member(update);
    Ok(StatusCode::OK)
}

pub async fn delete_member(
    id: i32,
) -> Result<impl warp::Reply, Infallible> {
    log::debug!("delete_member: id={}", id);

    let mut exists = database::check_member(id);

    if !exists.unwrap() {
        log::debug!("    -> member id not found");
        return Ok(StatusCode::NOT_FOUND)
    }

    let _ = database::delete_member(id);
    exists = database::check_member(id);

    if exists.unwrap() {
        log::debug!("    -> failed to delete");
        Ok(StatusCode::BAD_REQUEST)
    } else {
        Ok(StatusCode::NO_CONTENT)
    }
}

pub async fn get_guild(
    id: String,
) -> Result<impl warp::Reply, Infallible> {
    log::debug!("get_guild: id={}", id);

    let url = format!("https://discord.com/api/guilds/{}?with_counts=true", id);

    let guild = fetch::fetch_guild(url);

    match guild.await {
        Ok(guild) => {
            return Ok(warp::reply::json(&guild));
        },
        Err(err) => {
            log::debug!("    -> person not found, err: {}", err);
            let null_value: Value = serde_json::Value::Null;
            return Ok(warp::reply::json(&null_value))
        },
    }
}
