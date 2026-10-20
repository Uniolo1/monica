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
	"os"
	"time"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	// load configuration from "/etc" first
	_ = godotenv.Load("/etc/monica.env")

	// Not *that* big of a deal if os.Getenv returns "" here
	_ = godotenv.Load(fmt.Sprintf("%s/.config/monica/monica.env", os.Getenv("HOME")))
	_ = godotenv.Load(fmt.Sprintf("%s/monica/monica.env", os.Getenv("XDG_CONFIG_HOME")))

	// relative to the current directory (NOTE: overwrites env above)
	_ = godotenv.Load(".env")

	// Happy Pride Month!
	if time.Now().Month() == time.June {
		fmt.Println("Happy Pride Month!")
	}
}
