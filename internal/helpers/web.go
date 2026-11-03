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

package helpers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

// wrapper that respects 'ALLOW_HTTP_GET'
func Get(requestUrl string) (int, string) {
	if os.Getenv("ALLOW_HTTP") != "true" && os.Getenv("ALLOW_HTTP") != "True" {
		return 0, "HTTP is disabled!"
	}

	// the proxy
	var transport http.Transport

	if os.Getenv("HTTP_OVER_TOR") != "false" && os.Getenv("HTTP_OVER_TOR") != "False" {
		// I gave up trying to get tor to work so I just used AI for this part. Ensure 'tor' is started
		proxyURL, err := url.Parse("socks5h://127.0.0.1:9050")
		if err != nil {
			return 0, fmt.Sprintf("Proxy configuration error: %v", err)
		}

		transport = http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	} else {
		transport = http.Transport{}
	}

	client := &http.Client{
		Transport: &transport,
		Timeout:   30 * time.Second,
	}

	// send GET request
	resp, err := client.Get(requestUrl)
	if err != nil {
		return 0, fmt.Sprintf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	// read request
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, fmt.Sprintf("Response read failed: %v", err)
	}

	return resp.StatusCode, string(body)
}
