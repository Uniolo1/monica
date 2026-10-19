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

package commands

var DEFAULT_INSULTS = []string{
	"You oxygen thief!",
	"I hope both sides of your pillow are hot!",

	// found these from: https://bestlifeonline.com/best-insults/
	"You’re the reason God created the middle finger.",
	"I’d rather change a baby’s diaper than have lunch with you.",
	"You’re so ugly, even Stephen King would have nightmares about you.",
	"You are proof God has a sense of humor.",
	"Calling you an idiot would be an insult to all stupid people.",
	"You fear success, but you really have nothing to worry about.",
	"You have an entire life to be an idiot. Why not take today off?",

	// everything below this comment is from the internet - https://psycatgames.com/magazine/conversation-starters/creative-insults/
	"There’s some shit on your clothes. Oh, nope. That’s just you.",
	"You are living proof that even ugly people have sex.", // refers to either them, or there parents
	"If you ate trash, it would be cannibalism.",
	"If I throw a stick, will you leave?",
	"If ignorance is bliss, then you must be one happy idiot.",
}

var DEFAULT_COMPLIMENTS = []string{
	"Nice message!",

	// AI suggestions (not code so it's *fine*)
	"You give people room to be human.",
	"You’re quietly extraordinary",
	"You make the little things count.",
	"You make people feel like they don't have to pretend around you.",
	"You have a way of making the world feel a little less ordinary.",

	// Aleis's suggestion
	"You have a soul that reminds the lands of the falling rain again.",

	// everything below this comment is from wikihow - https://www.wikihow.com/Unique-Compliments
	// (wikihow sucks it had soo little)
	"You could talk to a brick wall and make it feel heard.",
	"Your voice is so calming, you should start an ASMR channel.",
	"Your heart must be 10 times the average size.",
}

var DEFAULT_FLIRTS = []string{
	"Heyy",
	"Are you open source? Cuz' I want to fork your left atrium!", // this one is high-risk I would say... meh... they all are

	// The AI gave me some suggestions for flirting when I asked for compliments... guess that's what people use AI for
	"You make eye contact feel like a dangerous game.",
	"If distraction were a person, it would look exactly like you.",
	"You make kindness look like a superpower.",
	"You’re not everybody’s type. You’re everybody’s plot twist.",
	"You’re the reason my screen time is so high.",
	"Y'know I was having a productive day until you showed up.", // I edited this one slightly
	"You’re a rare kind of beautiful.",
	"Your energy is addictive.",
	"You’re art with a pulse.",

	// An acquaintance of mine (Zachary) recommended I add this... keeping it as a comment
	// "Stawp! I am not a tsundere!",

	// everything below ths comment were flirts I found online - https://bestlifeonline.com/flirty-texts/
	"Hey, stranger. Stop being a stranger.",
	"Wanna trade secrets?",
	"I want to eat everything, including…", // A little inappropriate... might remove.
	"My love language is physical touch. Want to see why?",
	"Want to be my little spoon?",
	"Why are you so cute?",
	"Nobody gets me like you do.",
	"I have no idea what I would do without you.",
}

var DEFAULT_ANIMEGIRL_GIFS = []string{
	// From my discord gif gallery - Uniolo
	"https://tenor.com/view/alya-alisa-mikhailovna-kujou-alya-sometimes-hides-her-feelings-in-russian-anime-anime-dance-gif-1225141651010373452",
	"https://tenor.com/view/rumia-fish-gif-rumia-fish-touhou-gif-23149573",
	"https://tenor.com/view/bunny-gif-22123810",
	"https://klipy.com/gifs/femtanyl-2",
	"https://klipy.com/gifs/marin-kitagawa-33",
	"https://klipy.com/gifs/anime-girl-711",
	"https://klipy.com/gifs/osaka-spin-3",
	"https://klipy.com/gifs/happy-blue-hair",
	"https://klipy.com/gifs/fate-grand-10",
	"https://klipy.com/gifs/silent-witch-monica-everett-30", // _monica_ norton
	"https://klipy.com/gifs/silent-witch-4",                 // ^^^^^^^^^^^^^^^

	// From searching (via gif selector)
	"https://klipy.com/gifs/embarrassed-shy-59",
	"https://klipy.com/gifs/anime-girl-light-blue-hair-anime",
	"https://klipy.com/gifs/anime-lily-1",
	"https://klipy.com/gifs/reze-ok",
	"https://klipy.com/gifs/anime-anime-girl-129",
	"https://klipy.com/gifs/yuki-13",
	"https://klipy.com/gifs/chou-kaguya-hime-kaguya-hime-5",
	"https://klipy.com/gifs/angry-anime-36",
	"https://klipy.com/gifs/wink-anime-2",
	"https://klipy.com/gifs/anim-11",
	"https://klipy.com/gifs/alya-1", // good manga BTW
	"https://klipy.com/gifs/kushima-kamome-3",
	"https://klipy.com/gifs/wink-anime-wink",
	"https://klipy.com/gifs/tokoyami-towa-novice",
	"https://klipy.com/gifs/wink-smile-75",
	"https://klipy.com/gifs/hyakkano-100-girlfriends-1", // good manga BTW (x2)
	"https://klipy.com/gifs/hakari-hanazono-4",
	"https://klipy.com/gifs/shizuka-yoshimoto-kimi-no-koto-ga-daidaidaidaidaisuki-na-100-nin-no-kanojo",
	"https://klipy.com/gifs/chewzuka-cheeks",
	"https://klipy.com/gifs/wink-neuro", // neuro-sama is not anime... but close enough
}
