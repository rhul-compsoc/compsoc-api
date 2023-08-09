use std::env;
use super::models::Guild;
use reqwest::{
    header::{HeaderMap, HeaderValue, AUTHORIZATION},
    Client
};

pub async fn fetch_guild(
    url: String,
) -> Result<Guild, Box<dyn std::error::Error>> {
    let client = Client::new();

    let auth_token = env::var("DISCORD_TOK").expect("DISCORD_TOK must be set");
    let mut headers = HeaderMap::new();
    headers.insert(AUTHORIZATION, HeaderValue::from_str(&auth_token)?);

    let res = client
        .get(url)
        .headers(headers)
        .send()
        .await?;

    if res.status().is_success() {
        let guild: Guild = res.json().await?;
        Ok(guild)
    } else {
        Err(format!("Request failed with status code: {}", res.status()).into())
    }
}
