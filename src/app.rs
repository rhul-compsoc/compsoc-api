use super::filters;
use warp::Filter;
use dotenvy::dotenv;

#[tokio::main]
pub async fn run() {
    dotenv().ok();
    env_logger::init();

    let api = filters::make_filters();
    let routes = api.with(warp::log("compsoc api"));

    warp::serve(routes).run(([127, 0, 0, 1], 3030)).await;
}
