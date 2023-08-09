use super::handlers;
use super::models::{ListOptions, Person, Member};
use warp::Filter;

pub fn make_filters() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    people()
        .or(member())
        .or(guild())
}

/// All people filters combined.
pub fn people() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    people_list()
        .or(people_get())
        .or(people_create())
        .or(people_patch())
        .or(people_put())
        .or(people_delete())
}

/// GET /people?offset=3&limit=5
pub fn people_list() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    warp::path!("people")
        .and(warp::get())
        .and(warp::query::<ListOptions>())
        .and_then(handlers::list_people)
}

/// GET /people/:id
pub fn people_get() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    warp::path!("people" / i32)
        .and(warp::get())
        .and_then(handlers::get_person)
}

/// POST /people (with JSON body).
pub fn people_create() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    warp::path!("people")
        .and(warp::post())
        .and(person_json_body())
        .and_then(handlers::create_person)
}

/// PATCH /people/:id (with JSON body).
pub fn people_patch() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    warp::path!("people" / i32)
        .and(warp::patch())
        .and(person_json_body())
        .and_then(handlers::patch_person)
}

/// PUT /people/:id (with JSON body).
pub fn people_put() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    warp::path!("people" / i32)
        .and(warp::put())
        .and(person_json_body())
        .and_then(handlers::put_person)
}

/// DELETE /people/:id (with JSON body).
pub fn people_delete() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    let admin_only = warp::header::exact("authorization", "Bearer admin");

    warp::path!("people" / i32)
        .and(admin_only)
        .and(warp::delete())
        .and_then(handlers::delete_person)
}

pub fn member() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    member_list()
        .or(member_get())
        .or(member_create())
        .or(member_patch())
        .or(member_put())
        .or(member_delete())
}

/// GET /member?offset=3&limit=5
pub fn member_list() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    let admin_only = warp::header::exact("authorization", "Bearer admin");

    warp::path!("member")
        .and(admin_only)
        .and(warp::get())
        .and(warp::query::<ListOptions>())
        .and_then(handlers::list_member)
}

/// GET /member/:id
pub fn member_get() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    let admin_only = warp::header::exact("authorization", "Bearer admin");

    warp::path!("member" / i32)
        .and(admin_only)
        .and(warp::get())
        .and_then(handlers::get_member)
}

/// POST /member (with JSON body).
pub fn member_create() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    let admin_only = warp::header::exact("authorization", "Bearer admin");

    warp::path!("member")
        .and(admin_only)
        .and(warp::post())
        .and(member_json_body())
        .and_then(handlers::create_member)
}

/// PATCH /member/:id (with JSON body).
pub fn member_patch() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    let admin_only = warp::header::exact("authorization", "Bearer admin");

    warp::path!("member" / i32)
        .and(admin_only)
        .and(warp::patch())
        .and(member_json_body())
        .and_then(handlers::patch_member)
}

/// PUT /member/:id (with JSON body).
pub fn member_put() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    let admin_only = warp::header::exact("authorization", "Bearer admin");

    warp::path!("member" / i32)
        .and(admin_only)
        .and(warp::put())
        .and(member_json_body())
        .and_then(handlers::put_member)
}

/// DELETE /member/:id (with JSON body).
pub fn member_delete() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    let admin_only = warp::header::exact("authorization", "Bearer admin");

    warp::path!("member" / i32)
        .and(admin_only)
        .and(warp::delete())
        .and_then(handlers::delete_member)
}

pub fn guild() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    guild_get()
}

/// GET /guild/:id
pub fn guild_get() -> impl Filter<Extract = (impl warp::Reply,), Error = warp::Rejection> + Clone {
    warp::path!("guild" / String)
        .and(warp::get())
        .and_then(handlers::get_guild)
}

fn person_json_body() -> impl Filter<Extract = (Person,), Error = warp::Rejection> + Clone {
    // When accepting a body, we want a JSON body
    // (and to reject huge payloads)...
    warp::body::content_length_limit(1024 * 16).and(warp::body::json())
}

fn member_json_body() -> impl Filter<Extract = (Member,), Error = warp::Rejection> + Clone {
    // When accepting a body, we want a JSON body
    // (and to reject huge payloads)...
    warp::body::content_length_limit(1024 * 16).and(warp::body::json())
}
