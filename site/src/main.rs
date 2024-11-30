use reqwest::{Client, Error};
use serde::Deserialize;
use warp::Filter;
use base64::{engine::general_purpose::STANDARD, Engine};

const BASE_URL: &str = "http://localhost:1111";
const PORT: u16 = 4444;

const HTML_TEMPLATE: &str = r#"
<html>
<head>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            flex-wrap: wrap;
            align-items: center;
            justify-content: center;
            padding: 10px;
        }
        .gif-container {
            display: flex;
            flex-direction: column;
            align-items: center;
            margin: 10px;
        }
        img {
            width: 200px;
            height: auto;
            margin-bottom: 2px;
        }
        p {
            margin-top: 2px;
            font-size: 14px;
        }
    </style>
</head>
<body>
    {GIFS}
</body>
</html>
"#;

#[derive(Debug, Deserialize)]
struct Alias {
    aliases: Vec<String>,
}

#[tokio::main]
async fn main() -> Result<(), Error> {
    let client = Client::new();
    let url = format!("{}/gifs", BASE_URL);
    let response = client.get(url).send().await?;

    if !response.status().is_success() {
        eprintln!("алиасы не получены");
        return Ok(());
    }

    let aliases: Alias = response.json().await?;
    let html = generate_html(&aliases.aliases).await;

    println!("Server start: http://localhost:{}/", PORT);

    warp::serve(
        warp::path::end()
            .map(move || warp::reply::html(html.clone()))
    )
        .run(([127, 0, 0, 1], PORT))
        .await;

    Ok(())
}

async fn generate_html(aliases: &[String]) -> String {
    let client = Client::new();
    let mut html = String::new();

    for alias in aliases {
        let url = format!("{}/gif/{}", BASE_URL, alias);
        let response = client.get(&url).send().await;

        if let Ok(response) = response {
            if response.status().is_success() {
                let data: Vec<u8> = response.bytes().await.unwrap_or_default().to_vec();
                let base64 = STANDARD.encode(&data);
                html += &format!(
                    r#"
                    <div class="gif-container">
                        <img src="data:image/gif;base64,{}" alt="gif" />
                        <p>{}</p>
                    </div>
                    "#,
                    base64, alias
                );
            }
        }
    }

    HTML_TEMPLATE.replace("{GIFS}", &html)
}
