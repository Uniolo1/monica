// SPDX-FileCopyrightText: NONE
// SPDX-License-Identifier: Unlicense

/*
 * This is free and unencumbered software released into the public domain.
 *
 * Anyone is free to copy, modify, publish, use, compile, sell, or distribute
 * this software, either in source code form or as a compiled binary, for any
 * purpose, commercial or non-commercial, and by any means.
 *
 * In jurisdictions that recognize copyright laws, the author or authors of this
 * software dedicate any and all copyright interest in the software to the
 * public domain. We make this dedication for the benefit of the public at large
 * and to the detriment of our heirs and successors. We intend this dedication
 * to be an overt act of relinquishment in perpetuity of all present and future
 * rights to this software under copyright law.
 *
 * THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 * WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

pub struct Data {} // user data stuff
type Error = Box<dyn std::error::Error + Send + Sync>;
type Context<'a> = poise::Context<'a, Data, Error>;

#[poise::command(slash_command, prefix_command)]
pub async fn ping(ctx: Context<'_>) -> Result<(), Error> {
    let response = "pong!";
    ctx.send(
        poise::CreateReply::default()
            .content(response)
            .ephemeral(true),
    )
    .await?;
    Ok(())
}

#[poise::command(slash_command, prefix_command)]
pub async fn define(
    ctx: Context<'_>,
    #[description = "Word to search"] query: String,
) -> Result<(), Error> {
    let url = format!(
        "https://api.datamuse.com/words?sp={}&md=d",
        urlencoding::encode(&query)
    );

    match reqwest::get(url).await {
        Ok(response) => {
            let response_raw_json = response.text().await?;
            let response_parsed_json: serde_json::Value = serde_json::from_str(&response_raw_json)?;

            let response_final = response_parsed_json
                .as_array()
                .and_then(|results| results.first())
                .and_then(|first| first.get("defs"))
                .and_then(|defs| defs.as_array())
                .map(|defs| {
                    defs.iter()
                        .filter_map(|item| item.as_str())
                        .collect::<Vec<&str>>()
                        .join("\n")
                })
                .filter(|defs| !defs.is_empty())
                .unwrap_or_else(|| "No definitions found.".to_string());

            ctx.send(
                poise::CreateReply::default()
                    .content(response_final)
                    .ephemeral(true),
            )
            .await?;
        }
        Err(err) => {
            ctx.send(
                poise::CreateReply::default()
                    .content(format!("Request failed: {}", err))
                    .ephemeral(true),
            )
            .await?;
        }
    }

    Ok(())
}
