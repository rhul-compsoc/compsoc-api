# **`compsoc api`**

api for storing compsoc data

## *`usage`*

Run API with \
`$ cargo run`

For debug mode \
`$ RUST_LOG=debug cargo run`

Composing with Docker \
`$ docker up -d`

### `ADD data`

Address: `addr:port/person` \
Body:

```json
{
    "id" : 1,
    "name" : "string"
}
```

Address: `addr:port/member` \
Body:

```json
{
  "id": 123,
  "student_id": "0123",
  "student_email": "member@member.ac.uk",
  "first_name": "firt_name",
  "last_name": "last_name",
  "active_member": false
}
```

Header:

```json
{
    "authorization" : "Bearer admin",
}
```

### `GET data`

Address: `addr:port/people`
Optional parameters: `offset & limit`

Address: `addr:port/people/:id`

Address: `addr:port/member`
Optional parameters: `offset & limit`
Header:

```json
{
    "authorization" : "Bearer admin",
}
```

Address: `addr:port/member/:id`
Header:

```json
{
    "authorization" : "Bearer admin",
}
```

Address: `addr:port/guild/:id`

### `PATCH Data`

Address: `addr:port/people/:id` \
Body:

```json
{
    "id" : 2,
    "name" : "string"
}
```

Address: `addr:port/member/:id` \
Body:

```json
{
  "id": 123,
  "student_id": "0123",
  "student_email": "member@member.ac.uk",
  "first_name": "firt_name",
  "last_name": "last_name",
  "active_member": false
}
```

Header:

```json
{
    "authorization" : "Bearer admin",
}
```

### `PUT data`

Address: `addr:port/people/:id` \
Body:

```json
{
    "id" : 3,
    "name" : "string"
}
```

Address: `addr:port/member/:id` \
Body:

```json
{
  "id": 123,
  "student_id": "0123",
  "student_email": "member@member.ac.uk",
  "first_name": "firt_name",
  "last_name": "last_name",
  "active_member": false
}
```

Header:

```json
{
    "authorization" : "Bearer admin",
}
```

### `DELETE data`

Address: `addr:port/people/:id` \
Header:

```json
{
    "authorization" : "Bearer admin",
}
```

Address: `addr:port/member/:id` \
Header:

```json
{
    "authorization" : "Bearer admin",
}
```

## *`models`*

### `person`

```rs
pub struct Person {
    pub id: i32,
    pub name: String,
}
```

### `members`

```rs
pub struct Member {
    pub id: i32,
    pub student_id: String,
    pub student_email: String,
    pub first_name: String,
    pub last_name: String,
    pub active_member: bool,
}
```

### `guilds`

```rs
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
```

## *`dependencies`*

- warp
- tokio
- serde
- dotenvy
- env logger & logs
- diesel
- reqwest
