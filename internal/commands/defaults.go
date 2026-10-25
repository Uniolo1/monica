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

const DEFAULT_DEFAULT_AI_PROMPT string = "You are MonicaAI, a helpful AI assistant!"

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

	// Friend suggestion:
	"Are you a triangle, because your acute",

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
	"https://tenor.com/view/bocchi-the-rock-gif-15650315287216926532",
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
	"https://klipy.com/gifs/neko-cat-9",
	"https://klipy.com/gifs/cat-girl-45",
	"https://klipy.com/gifs/cat-girl-anime-cat-girl",
	"https://klipy.com/gifs/neko-neko-niiiii-chibi",
	"https://klipy.com/gifs/anime-ok-11",
	"https://klipy.com/gifs/princess-connect-cat-girl",
	"https://klipy.com/gifs/monika-15",
	"https://klipy.com/gifs/anime-fran-22",
	"https://klipy.com/gifs/catgirl-5",
	"https://klipy.com/gifs/anime-sleepy-30",
	"https://klipy.com/gifs/mikasa-pudding-chibi-mikasa",
	"https://klipy.com/gifs/konata-konata-happy",  // The Arch Cat posted it immediately after I searched it up lol
	"https://klipy.com/gifs/lucky-star-thumbs-up", // TODO: watch lucky star
	"https://klipy.com/gifs/lucky-star-intro",
	"https://klipy.com/gifs/anime-tired-lucky-star",
	"https://klipy.com/gifs/lucky-star-27",
	"https://klipy.com/gifs/dancinha-charmy", // good anime
	"https://klipy.com/gifs/black-clover-nero-3",
	"https://klipy.com/gifs/black-clover-black-clover-movie-44",
	"https://klipy.com/gifs/noelle-silva-black-clover-4",
	"https://klipy.com/gifs/black-clover-noelle-4",
	//"https://klipy.com/gifs/noelle-silva-black-clover", // not technically visible... ehh.... commenting it out
	"https://klipy.com/gifs/black-clover-53",
	"https://klipy.com/gifs/asuka-26",
	"https://klipy.com/gifs/asuka-langley-84",
	"https://klipy.com/gifs/asuka-langley-33",
	"https://klipy.com/gifs/asuka-cry-asuka-langley",
	"https://klipy.com/gifs/asuka-asuka-langley-13",
	"https://klipy.com/gifs/no-24335",
	"https://klipy.com/gifs/asuka-evangelion-3",
	"https://klipy.com/gifs/rei-rei-ayanami-18",
	"https://klipy.com/gifs/rei-neon-genesis-evangelion-1",
	"https://klipy.com/gifs/neon-genesis-evangelion-rei-ayanami-1",
	"https://klipy.com/gifs/hitagi-senjougahara-1", // good manga (Bakemonogatari), my favorite manga as of writing this
	"https://klipy.com/gifs/tsubasa-hanekawa-neko",
	"https://klipy.com/gifs/monogatari-series-oshino-shinobu",
	"https://klipy.com/gifs/anime-bakemonogatari-7",
	"https://klipy.com/gifs/monogatari-yotsugi-ononoki",
	"https://klipy.com/gifs/yashiro-nene-hanako-kun", // good anime (and manga)
	"https://klipy.com/gifs/tbhk-toilet-bound-hanako-kun-37",
	"https://klipy.com/gifs/anime-happy-anime-smile-2",
	"https://klipy.com/gifs/toilet-bound-hanako-kun-tbhk-5",
	"https://klipy.com/gifs/nene-yashiro-after-school-hanako-kun",
	"https://klipy.com/gifs/tbhk-toilet-bound-hanako-kun-32",
	"https://klipy.com/gifs/tokyo-ghoul-108",
	"https://klipy.com/gifs/mikasa-ackerman-shingeki-no-kyojin-1",
	"https://klipy.com/gifs/nezuko-demon-slayer-6",
	"https://klipy.com/gifs/nezuko-29",
	"https://klipy.com/gifs/nezuko-saying-hi",
	"https://klipy.com/gifs/nezuko-9",
	"https://klipy.com/gifs/nezuko-anime-9",
	"https://klipy.com/gifs/nezuko-31",
	"https://klipy.com/gifs/nezuko-blinking",
	"https://klipy.com/gifs/nezuko-dancing-nezuko",
	"https://klipy.com/gifs/nezuko-nezuko-kamado-14",
	"https://klipy.com/gifs/nezuko-33",
	"https://klipy.com/gifs/shiina-mashiro-sakurasou-2",
	"https://klipy.com/gifs/anime-winter-2026-17",
	"https://klipy.com/gifs/anime-eating-1",
	"https://klipy.com/gifs/shiboyugi-anime-9",
	"https://klipy.com/gifs/citrus-9UF",
	"https://klipy.com/gifs/yuri-doki-doki-literature-club-2",
	"https://klipy.com/gifs/yuri-anime-4",
	"https://klipy.com/gifs/mitsuri-cry-1",
	"https://klipy.com/gifs/mitsuri-mitsuri-kanroji-5",
	"https://klipy.com/gifs/asufuturechamp-mitsuri",
	"https://klipy.com/gifs/mitsuri-mitsuri-love",
	"https://klipy.com/gifs/mitsuri-13",
	"https://klipy.com/gifs/demon-slayer-mitsuri-kanroji",
	"https://klipy.com/gifs/mitsuri-anime-2",
	"https://klipy.com/gifs/mitsuri-21",
	"https://klipy.com/gifs/demon-slayer-kny-4",
	"https://klipy.com/gifs/code-geass-cc-6", // great anime!
	"https://klipy.com/gifs/pizza-codegeass-1",
	"https://klipy.com/gifs/cc-code-geass-3",
	"https://klipy.com/gifs/code-geass-anime-3",
	"https://klipy.com/gifs/cc-stare-code-geass",
	"https://klipy.com/gifs/cc-smile",
	"https://klipy.com/gifs/cc-254",
	"https://klipy.com/gifs/shirley-cry", // why is this the only Shirley gif? smh
	"https://klipy.com/gifs/laugh-laughing-63",
	"https://klipy.com/gifs/misato-evangelion-11",
	"https://klipy.com/gifs/misato-9",
	"https://klipy.com/gifs/misato-evangelion-26",
	"https://klipy.com/gifs/misato-type",
	"https://klipy.com/gifs/misato-scratch-head",
	"https://klipy.com/gifs/misato-2",
	"https://klipy.com/gifs/shy-cute-32", // I was looking for Yui Ikari and found this
	"https://klipy.com/gifs/i-wanna-cute",
	"https://klipy.com/gifs/hirasawa-yui-2",
	"https://klipy.com/gifs/yui-yui-hirasawa-24",
	"https://klipy.com/gifs/clap-clap-castanet",
	"https://klipy.com/gifs/huh-what-377",
	"https://klipy.com/gifs/im-sorry-U8C",
	"https://klipy.com/gifs/k-on-yui-hirasawa-130",
	"https://klipy.com/gifs/mitsuba-mitsuba-sousuke-4",
	"https://klipy.com/gifs/mitsuba-mitsuba-sousuke-2",
	"https://klipy.com/gifs/toilet-bound-hanako-kun-mitsuba-7",
	"https://klipy.com/gifs/toilet-bound-hanako-kun-mitsuba-sousuke",
	"https://klipy.com/gifs/mitsuba-sousuke-mitsuba-8",
	"https://klipy.com/gifs/mitsuba",
	"https://klipy.com/gifs/cznburak-demon-slayer", // wrong aoi but okay
	"https://klipy.com/gifs/tbhk-toilet-bound-hanako-kun-26",
	"https://klipy.com/gifs/blink-blink-9",
	"https://klipy.com/gifs/aoi-teehee",
	"https://klipy.com/gifs/ue-ue-1",
	"https://klipy.com/gifs/good-luck-525",
	"https://klipy.com/gifs/hehe-294",
	"https://klipy.com/gifs/nya-10",
	"https://klipy.com/gifs/hm-118",
	"https://klipy.com/gifs/yes-23296",
	"https://klipy.com/gifs/nanamine-sakura-hanako-kun",
	"https://klipy.com/gifs/nanamine-sakura-tbhk-1",
	"https://klipy.com/gifs/tbhk-toiletboundhanako",
	"https://klipy.com/gifs/yashiro-toilet-bound-hanako-kun",
	"https://klipy.com/gifs/sakura-kinomoto-sakura-card-captor-1",
	"https://klipy.com/gifs/rae-rae-taylor", // no clue who this or where its from
	"https://klipy.com/gifs/rae-rae-taylor-1",
	"https://klipy.com/gifs/alya-kujou", // good manga (probs already said that)
	"https://klipy.com/gifs/firefly-alya",
	"https://klipy.com/gifs/alya-5",
	"https://klipy.com/gifs/alya-anime-1",
	"https://klipy.com/gifs/-36094",
	"https://klipy.com/gifs/alya-alisa",
	"https://klipy.com/gifs/alya-alya-sometimes-hides-her-feelings-in-russian-2",
	"https://klipy.com/gifs/holo-wisewolf", // great manga
	"https://klipy.com/gifs/spice-and-wolf-spice-and-wolf-remake",
	"https://klipy.com/gifs/spice-and-wolf-holo-7",
	"https://klipy.com/gifs/holo-spice-and-wolf-9",
	"https://klipy.com/gifs/holo-wolf",
	"https://klipy.com/gifs/holo-spice-and-wolf-8",
	"https://klipy.com/gifs/tatsumaki-phone", // great manga (and S1 anime)
	"https://klipy.com/gifs/tatsumaki-1",
	"https://klipy.com/gifs/tatsumaki-one-punch-man-8",
	"https://klipy.com/gifs/tatsumaki-7",
	"https://klipy.com/gifs/sakamoto-days-osaragi",
	"https://klipy.com/gifs/mahiro-mahiro-oyama-11",
	"https://klipy.com/gifs/emma-the-promised-neverland",
	"https://klipy.com/gifs/the-promised-neverland-emma-3",
	"https://klipy.com/gifs/the-promised-neverland-emma-4",
	"https://klipy.com/gifs/angel-beats-yui-7", // good anime!
	"https://klipy.com/gifs/kanade-angel-beats-2",
	"https://klipy.com/gifs/kanade-angel-beats-3",
	"https://klipy.com/gifs/roxnook-yui-2",
	"https://klipy.com/gifs/angel-beats-25",
	"https://klipy.com/gifs/bungo-stray-dogs-lucy-1", // good anime!
	"https://klipy.com/gifs/kei-cote-1",              // good anime (ignoring memes)
	"https://klipy.com/gifs/airi-sakura-classroom-of-the-elite",
	"https://klipy.com/gifs/classroom-of-the-elite-youkoso-jitsuryoku-shijou-shugi-no-kyoushitsu-e-13",
	"https://klipy.com/gifs/sakayanagi-classroom-of-the-elite-1",
	"https://klipy.com/gifs/classroom-of-the-elite-youkoso-jitsuryoku-shijou-shugi-no-kyoushitsu-e-25",
	"https://klipy.com/gifs/kei-karuizawa-kei-classroom-of-the-elite-1",
	"https://klipy.com/gifs/classroom-of-the-elite-youkoso-jitsuryoku-shijou-shugi-no-kyoushitsu-10",
	"https://klipy.com/gifs/irina-irina-jelavic-1", // good anime
	"https://klipy.com/gifs/kaede-kaede-kayano-3",
	"https://klipy.com/gifs/kaede-kaede-kayano-1",
	"https://klipy.com/gifs/assassination-classroom-79", // actually a boy BTW
	"https://klipy.com/gifs/bye-later-4",                // great anime (an anime movie BTW)
	"https://klipy.com/gifs/shouko-nishimiya-a-silent-voice-1",
	"https://klipy.com/gifs/a-silent-voice-shouko",
	"https://klipy.com/gifs/shouko-a-silent-voice",
	"https://klipy.com/gifs/a-silent-voice-anime-3",
	"https://klipy.com/gifs/a-silent-voice-nishimiya-koenokatachi",
	"https://klipy.com/gifs/miki-kawai-a-silent-voice",
	"https://klipy.com/gifs/girl-anime-27",
	"https://klipy.com/gifs/anime-tyan",
	"https://klipy.com/gifs/hi-oomf-cute-anime-girl-vtuber",
	"https://klipy.com/gifs/beast-tamer-kanade",
	"https://klipy.com/gifs/baby-zlS",
	"https://klipy.com/gifs/cosmic-princess-kaguya-pleading",
	"https://klipy.com/gifs/azumanga-daioh-cute-anime",
	"https://klipy.com/gifs/sewayaki-senko-san-kitsune",
	"https://klipy.com/gifs/anime-girl-smile-cute-anime-girl",
	"https://klipy.com/gifs/pout-red",
	"https://klipy.com/gifs/chihaya-anon-anon-chihaya",
	"https://klipy.com/gifs/cute-anime-girl-phone-yuri-crowsbo",
	"https://klipy.com/gifs/kasumi-laugh-anime-laugh",
	"https://klipy.com/gifs/anime-kanna-12",
	"https://klipy.com/gifs/horimiya-ayasaki-remi",
	"https://klipy.com/gifs/lycoris-recoil-lecoreco-2",
	"https://klipy.com/gifs/momo-the-gamer-momo-vtuber-1",
	"https://klipy.com/gifs/mira-mirajane",
	"https://klipy.com/gifs/familiar-of-zero-zero-no-tsukaima-4",
	"https://klipy.com/gifs/honda-tohru-tohru-1",
	"https://klipy.com/gifs/tohsaka-rin-1",
	"https://klipy.com/gifs/nichijou-ona2000-1",
	"https://klipy.com/gifs/clannad-fujibayashi-kyou",
	"https://klipy.com/gifs/anime-tsundere",
	"https://klipy.com/gifs/spin-happy-5",
	"https://klipy.com/gifs/haru-aomi-chitose-is-in-the-ramune-bottle-6",
	"https://klipy.com/gifs/ao-sorakado-summer-pockets-5",
	"https://klipy.com/gifs/taiga-toradora-16",

	// gif's others send - The Arch Cat(tm):
	"https://klipy.com/gifs/throw-rocks-cartoons",
	"https://klipy.com/gifs/yama-no-susume-yukimura-aoi",
	"https://klipy.com/gifs/yippie-yay-3",
	"https://klipy.com/gifs/rikka-takanashi-love-chunibyo-other-delusions",
	"https://klipy.com/gifs/chou-kaguya-hime-kaguya-hime-3",
	"https://klipy.com/gifs/smile-cute-97",
	"https://klipy.com/gifs/cute-anime-142",
	"https://klipy.com/gifs/anime-flex",
	"https://klipy.com/gifs/bocchi-the-rock-ryo-yamada-15",
	"https://klipy.com/gifs/kushima-kamome-14",
	"https://klipy.com/gifs/onimai-oniichan-wa-oshimai-58",
	"https://klipy.com/gifs/akira-kogami-1",
	"https://klipy.com/gifs/osaka-looking-up",
	"https://klipy.com/gifs/maka-soul-eater-11",
	"https://klipy.com/gifs/witch-hat-atelier-richeh",
	"https://klipy.com/gifs/rem-re-zero-23",
	"https://klipy.com/gifs/marina-takopis-original-sin",
	"https://klipy.com/gifs/himouto-umaru-chan-transform",
	"https://klipy.com/gifs/bocchi-the-rock-crying-1",
	"https://klipy.com/gifs/mahiro-oyama-3",
	"https://klipy.com/gifs/onimai-oniichan-wa-oshimai-13",
	// how the hell does the arch in the cat have so many gifs?
	// isn't even making a Discord bot that requires a large amount of anime gifs smh

	// gif's others send - IRL friends
	"https://klipy.com/gifs/koneko-koneko-toujou-1",
	"https://klipy.com/gifs/kopa22-anime-63",
	"https://klipy.com/gifs/anime-planes12265-2",
	"https://klipy.com/gifs/momo-momo-belia-deviluke",
	"https://klipy.com/gifs/yuuki-mikan-mikan",
	"https://klipy.com/gifs/to-love-ru-to-love-ru-darkness",
	"https://klipy.com/gifs/to-love-ru-golden-darkness-1",
	"https://klipy.com/gifs/marty1234-1",

	// gif's others send - others:
	"https://klipy.com/gifs/silent-witch-chinmoku-no-majo-1",
	"https://klipy.com/gifs/everett-monica-silent-witch-18",
	"https://klipy.com/gifs/silent-witch-norton",
	"https://klipy.com/gifs/throw-rocks-cartoons", // this gif is used very commonly in Neon Genesis Linux
}
