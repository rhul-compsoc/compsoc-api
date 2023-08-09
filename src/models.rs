use serde_derive::{Deserialize, Serialize};
use diesel::prelude::*;

#[derive(Debug, Deserialize, Serialize, Clone, Queryable, Selectable, Insertable, AsChangeset)]
#[diesel(table_name = super::schema::users)]
#[diesel(check_for_backend(diesel::pg::Pg))]
pub struct Person {
    pub id: i32,
    pub name: String,
}

#[derive(Debug, Deserialize, Serialize, Clone, Queryable, Selectable, Insertable, AsChangeset)]
#[diesel(table_name = super::schema::members)]
#[diesel(check_for_backend(diesel::pg::Pg))]
pub struct Member {
    pub id: i32,
    pub student_id: String,
    pub student_email: String,
    pub first_name: String,
    pub last_name: String,
    pub active_member: bool,
}

#[derive(Debug, Deserialize)]
pub struct ListOptions {
    pub offset: Option<usize>,
    pub limit: Option<usize>,
}

#[derive(Debug, Deserialize, Serialize, Clone)]
pub struct Guild {
    id: String,
    name: String,
    icon: Option<String>,
    description: Option<String>,
    home_header: Option<String>,
    splash: Option<String>,
    discovery_splash: Option<String>,
    features: Vec<String>,
    banner: Option<String>,
    owner_id: String,
    application_id: Option<String>,
    region: String,
    afk_channel_id: Option<String>,
    afk_timeout: i32,
    system_channel_id: String,
    system_channel_flags: i32,
    widget_enabled: bool,
    widget_channel_id: Option<String>,
    verification_level: i32,
    roles: Vec<Role>,
    default_message_notifications: i32,
    mfa_level: i32,
    explicit_content_filter: i32,
    max_presences: Option<String>,
    max_members: i32,
    max_stage_video_channel_users: i32,
    max_video_channel_users: i32,
    vanity_url_code: Option<String>,
    premium_tier: i32,
    premium_subscription_count: i32,
    preferred_locale: String,
    rules_channel_id: Option<String>,
    safety_alerts_channel_id: Option<String>,
    public_updates_channel_id: Option<String>,
    hub_type: Option<String>,
    premium_progress_bar_enabled: bool,
    latest_onboarding_question_id: Option<String>,
    nsfw: bool,
    nsfw_level: i32,
    emojis: Vec<String>,
    stickers: Vec<String>,
    incidents_data: Option<String>,
    inventory_settings: Option<String>,
    embed_enabled: bool,
    embed_channel_id: Option<String>,
    approximate_member_count: i32,
    approximate_presence_count: i32,
}

#[derive(Debug, Deserialize, Serialize, Clone)]
struct Role {
    id: String,
    name: String,
    description: Option<String>,
    permissions: i64,
    permissions_new: Option<String>,
    position: i32,
    color: i32,
    hoist: bool,
    managed: bool,
    mentionable: bool,
    icon: Option<String>,
    unicode_emoji: Option<String>,
    flags: i32,
}

#[derive(Debug, Deserialize, Serialize, Clone, Queryable, Selectable)]
#[diesel(table_name = super::schema::admins)]
#[diesel(check_for_backend(diesel::pg::Pg))]
pub struct Admin {
    pub token: String,
}