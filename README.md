# Awesome Go Extra
***All data are from [awesome-go](https://github.com/avelino/awesome-go) and [GitHub API](https://docs.github.com/en/rest/reference/repos#get-a-repository).***

Records are sorted by ***Star*** | [CreatedAt](./README-created.md) | [PushedAt](./README-pushed.md)

## Audio and Music
*Libraries for manipulating audio.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[oto](https://github.com/hajimehoshi/oto)|♪ A low-level library to play sound on multiple platforms ♪|907|77|18|2017-05-04T12:16:30Z|2021-11-06T16:37:30Z|
[portaudio](https://github.com/gordonklaus/portaudio)|Go bindings for the PortAudio audio I/O library|484|73|11|2015-09-16T07:59:23Z|2020-09-11T16:11:49Z|
[music-theory](https://github.com/go-music-theory/music-theory)|Go models of Note, Scale, Chord and Key|373|37|8|2016-03-17T03:50:17Z|2020-07-22T22:26:52Z|
[portmidi](https://github.com/rakyll/portmidi)|Go bindings for libportmidi|262|52|12|2013-11-10T14:24:56Z|2020-10-20T18:07:03Z|
[id3v2](https://github.com/bogem/id3v2)|🎵 ID3 decoding and encoding library for Go|211|32|12|2016-05-15T18:36:53Z|2021-12-07T13:00:41Z|
[GoAudio](https://github.com/DylanMeeus/GoAudio)|Go tools for audio processing &amp; creation 🎶|182|16|5|2020-07-05T01:35:41Z|2021-11-24T02:03:43Z|
[flac](https://github.com/mewkiz/flac)|Package flac provides access to FLAC (Free Lossless Audio Codec) streams.|178|31|10|2012-11-01T20:13:58Z|2021-01-27T23:39:45Z|
[malgo](https://github.com/gen2brain/malgo)|Mini audio library|158|27|4|2017-11-09T18:27:52Z|2021-09-25T18:19:22Z|
[mix](https://github.com/go-mix/mix)|Sequence-based Go-native audio mixer for music apps|144|22|11|2016-01-03T15:53:57Z|2020-05-14T17:14:34Z|
[gaad](https://github.com/Comcast/gaad)|GAAD (Go Advanced Audio Decoder)|92|16|3|2016-07-11T14:19:16Z|2021-10-01T22:31:53Z|
[minimp3](https://github.com/tosone/minimp3)|Decode mp3 base on https://github.com/lieff/minimp3|64|11|0|2018-01-26T09:10:31Z|2021-03-29T11:39:59Z|
[vorbis](https://github.com/mccoyst/vorbis)|A &#34;native&#34; ogg vorbis decoder for Go (uses inline stb_vorbis)|28|5|2|2013-07-12T02:45:39Z|2019-03-30T17:50:34Z|
[gosamplerate](https://github.com/dh1tw/gosamplerate)|Go Bindings for libsamplerate|13|8|0|2016-11-20T21:19:31Z|2020-07-10T23:35:33Z|


## Authentication and OAuth
*Libraries for implementing authentications schemes.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[casbin](https://github.com/casbin/casbin)|An authorization library that supports access control models like ACL, RBAC, ABAC in Golang|10816|1200|34|2017-04-08T07:51:23Z|2021-12-11T00:47:18Z|
[oauth2](https://github.com/golang/oauth2)|Go OAuth2|3930|801|150|2014-04-14T15:07:35Z|2021-11-22T19:55:26Z|
[goth](https://github.com/markbates/goth)|Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications.|3442|419|72|2014-10-14T20:38:12Z|2021-11-25T22:04:30Z|
[authboss](https://github.com/volatiletech/authboss)|The boss of http auth.|2942|180|30|2015-01-03T05:12:02Z|2021-08-12T03:22:16Z|
[go-jose](https://github.com/square/go-jose)|An implementation of JOSE standards (JWE, JWS, JWT) in Go|1845|329|56|2014-11-14T18:27:31Z|2021-10-25T14:41:13Z|
[go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server)|A standalone, specification-compliant,  OAuth2 server written in Golang.|1838|283|27|2015-11-01T13:30:09Z|2021-01-04T11:45:40Z|
[loginsrv](https://github.com/tarent/loginsrv)|JWT login microservice with plugable backends such as OAuth2, Google, Github, htpasswd, osiam, ..|1834|159|29|2016-11-11T12:11:21Z|2021-02-27T06:20:06Z|
[osin](https://github.com/openshift/osin)|Golang OAuth2 server library|1708|382|4|2013-09-10T19:52:00Z|2021-11-18T23:12:04Z|
[gologin](https://github.com/dghubble/gologin)|Go login handlers for authentication providers (OAuth1, OAuth2)|1460|115|0|2015-06-23T04:40:52Z|2021-12-09T17:37:04Z|
[gorbac](https://github.com/mikespook/gorbac)|goRBAC provides a lightweight role-based access control (RBAC) implementation in Golang.|1227|151|2|2013-12-26T10:00:41Z|2021-05-05T22:46:44Z|
[scs](https://github.com/alexedwards/scs)|HTTP Session Management for Go|981|96|4|2016-08-08T16:42:05Z|2021-12-06T17:23:40Z|
[paseto](https://github.com/o1egl/paseto)|Platform-Agnostic Security Tokens implementation in GO (Golang)|558|27|3|2018-01-23T05:27:39Z|2021-10-04T23:07:30Z|
[permissions2](https://github.com/xyproto/permissions2)|  :closed_lock_with_key: Middleware for keeping track of users, login states and permissions|440|36|0|2014-11-19T12:23:37Z|2021-07-28T11:29:31Z|
[go-guardian](https://github.com/shaj13/go-guardian)|Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication.|322|28|8|2020-05-14T12:15:56Z|2021-12-10T01:08:57Z|
[jwt](https://github.com/cristalhq/jwt)|Safe, simple and fast JSON Web Tokens for Go|307|26|0|2019-07-20T18:14:58Z|2021-12-11T18:36:51Z|
[jwt](https://github.com/pascaldekloe/jwt)|JSON Web Token library|275|21|0|2018-03-21T11:59:53Z|2021-11-01T15:10:40Z|
[jeff](https://github.com/abraithwaite/jeff)|🍍Jeff provides the simplest way to manage web sessions in Go.|234|14|1|2018-08-02T19:31:23Z|2021-06-23T12:10:59Z|
[httpauth](https://github.com/goji/httpauth)|HTTP Authentication middlewares|212|27|5|2014-05-26T22:53:57Z|2021-09-28T09:25:46Z|
[jwt-auth](https://github.com/adam-hanna/jwt-auth)|This package provides json web token (jwt) middleware for goLang http servers|210|39|3|2016-07-05T23:31:43Z|2021-08-01T23:03:58Z|
[branca](https://github.com/hako/branca)|:key: Secure alternative to JWT. Authenticated Encrypted API Tokens for Go.|164|18|1|2018-01-09T15:27:31Z|2020-08-07T06:24:03Z|
[sessionup](https://github.com/swithek/sessionup)|Straightforward HTTP session management|113|6|4|2019-07-23T18:55:21Z|2021-08-27T10:31:54Z|
[otpgen](https://github.com/RijulGulati/otpgen)|Library to generate TOTP/HOTP codes|110|7|0|2021-06-01T18:02:59Z|2021-08-06T16:31:11Z|
[session](https://github.com/icza/session)|Go session management for web servers (including support for Google App Engine - GAE).|105|16|4|2016-02-08T09:07:07Z|2021-07-26T20:29:04Z|
[sjwt](https://github.com/brianvoe/sjwt)|Simple JWT Golang|95|8|0|2019-06-20T04:06:21Z|2019-09-19T20:43:18Z|
[jwt](https://github.com/robbert229/jwt)|This is an implementation of JWT in golang!|94|25|9|2016-06-05T22:01:37Z|2020-11-24T16:31:55Z|
[rbac](https://github.com/zpatrick/rbac)|Minimalistic RBAC package for Go applications|88|15|0|2018-08-02T00:11:04Z|2018-08-29T19:03:47Z|
[sessions](https://github.com/adam-hanna/sessions)|A dead simple, highly performant, highly customizable sessions middleware for go http servers.|61|8|2|2017-04-29T01:09:28Z|2020-04-15T18:04:28Z|
[securecookie](https://github.com/chmike/securecookie)|Fast, secure and efficient secure cookie encoder/decoder |54|8|0|2017-09-03T14:40:29Z|2021-09-24T18:14:03Z|
[go-email-normalizer](https://github.com/dimuska139/go-email-normalizer)|Golang library for providing a canonical representation of email address.|33|2|0|2020-08-21T23:13:04Z|2021-09-18T11:31:53Z|
[otpgo](https://github.com/jltorresm/otpgo)|Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go.|26|3|1|2020-08-19T13:20:23Z|2021-02-27T19:43:16Z|
[scope](https://github.com/SonicRoshan/scope)|Easily Manage OAuth2 Scopes In Go|17|5|0|2019-09-23T10:48:14Z|2021-05-25T13:48:05Z|
[sessiongate-go](https://github.com/f0rmiga/sessiongate-go)|A driver for the SessionGate Redis module - easy session management using the Go language.|9|2|0|2017-10-20T03:39:11Z|2018-11-09T19:25:29Z|
[cookiestxt](https://github.com/mengzhuo/cookiestxt)|cookiestxt implement parser of cookies txt format|8|3|1|2017-10-09T11:27:19Z|2021-03-08T11:45:56Z|


## Blockchain
*Tools for building blockchains.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[go-ethereum](https://github.com/ethereum/go-ethereum)|Official Go implementation of the Ethereum protocol|34019|12574|294|2013-12-26T13:05:46Z|2021-12-11T15:51:05Z|
[tendermint](https://github.com/tendermint/tendermint)|⟁ Tendermint Core (BFT Consensus) in Go|4475|1446|337|2014-05-14T23:21:35Z|2021-12-11T14:30:02Z|
[cosmos-sdk](https://github.com/cosmos/cosmos-sdk)|:chains: A Framework for Building High Value Public Blockchains :sparkles:|3122|1342|575|2016-02-06T07:15:53Z|2021-12-11T19:38:45Z|
[gossamer](https://github.com/ChainSafe/gossamer)|🕸️ Gossamer: A Go implementation of the Polkadot Host|292|73|153|2019-01-28T17:40:01Z|2021-12-11T02:06:12Z|


## Bot Building
*Libraries for building and working with bots.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api)|Golang bindings for the Telegram Bot API|3245|532|11|2015-06-25T05:33:57Z|2021-12-08T04:30:23Z|
[olivia](https://github.com/olivia-ai/olivia)|💁‍♀️Your new best friend powered by an artificial neural network|3111|314|20|2018-06-05T18:19:31Z|2021-12-10T17:36:52Z|
[telebot](https://github.com/tucnak/telebot)|Telebot is a Telegram bot framework in Go.|2223|301|42|2015-06-25T19:27:50Z|2021-12-08T18:16:25Z|
[kelp](https://github.com/stellar/kelp)|Kelp is a free and open-source trading bot for the Stellar DEX and 100&#43; centralized exchanges|811|186|160|2018-08-08T23:31:18Z|2021-11-26T23:52:55Z|
[bot](https://github.com/go-chat-bot/bot)|IRC, Slack, Telegram and RocketChat bot written in go|711|172|11|2015-09-22T16:41:13Z|2020-10-04T14:12:27Z|
[golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot)|A golang implementation of a console-based trading bot for cryptocurrency exchanges|653|166|10|2017-05-14T22:11:41Z|2021-06-04T00:21:58Z|
[slacker](https://github.com/shomali11/slacker)|Slack Bot Framework|573|83|8|2017-05-20T01:41:20Z|2021-12-05T13:11:51Z|
[joe](https://github.com/go-joe/joe)|A general-purpose bot library inspired by Hubot but written in Go.   :robot:|435|24|5|2019-03-03T11:19:18Z|2020-07-26T11:24:21Z|
[tbot](https://github.com/yanzay/tbot)|Go library for Telegram Bot API|320|48|0|2015-09-11T16:19:25Z|2021-03-22T20:26:16Z|
[go-sarah](https://github.com/oklahomer/go-sarah)|Simple yet customizable bot framework written in Go.|212|14|0|2016-11-06T10:04:43Z|2021-10-30T04:08:50Z|
[go-twitch-irc](https://github.com/gempir/go-twitch-irc)|go irc client for twitch.tv|208|42|15|2017-03-23T21:31:35Z|2021-12-11T20:04:42Z|
[tenyks](https://github.com/kyleterry/tenyks)|The Tenyks IRC bot.|171|18|12|2012-08-26T02:02:24Z|2019-09-11T01:43:50Z|
[hanu](https://github.com/sbstjn/hanu)|Golang Framework for writing Slack bots|134|23|2|2016-09-16T07:10:42Z|2021-06-16T04:18:00Z|
[go-tgbot](https://github.com/olebedev/go-tgbot)|Golang  telegram bot API wrapper, session-based router and middleware|111|4|2|2016-12-11T06:06:32Z|2018-06-25T04:50:26Z|
[slack-bot](https://github.com/innogames/slack-bot)|Ready to use Slack bot for lazy developers: start Jenkins jobs, watch Jira tickets, watch pull requests...|66|20|6|2019-07-19T07:49:06Z|2021-12-09T21:18:47Z|
[margelet](https://github.com/zhulik/margelet)|Telegram Bot Framework for Go|66|11|0|2015-11-21T13:02:17Z|2016-09-18T11:47:01Z|
[ephemeral-roles](https://github.com/ewohltman/ephemeral-roles)|A Discord bot for managing ephemeral roles based upon voice channel member presence.|51|8|7|2017-12-19T15:20:30Z|2021-12-08T12:00:26Z|
[slackscot](https://github.com/alexandre-normand/slackscot)|Slack bot core/framework written in Go with support for reactions to message updates/deletes|48|10|1|2015-10-22T04:54:55Z|2021-11-22T00:27:28Z|
[echotron](https://github.com/NicoNex/echotron)|Concurrent library for Telegram Bots written in pure go.|38|6|1|2019-07-22T17:31:49Z|2021-12-10T19:27:54Z|
[govkbot](https://github.com/nikepan/govkbot)|VK bot package for Go|37|4|1|2016-07-11T22:09:54Z|2021-08-06T18:46:29Z|
[larry](https://github.com/ezeoleaf/larry)|Larry 🐦 is a really simple Twitter bot generator that tweets random repositories from Github built in Go|29|9|4|2020-11-16T23:25:12Z|2021-11-15T09:49:55Z|
[micha](https://github.com/onrik/micha)|Client lib for Telegram bot api|17|4|0|2016-04-14T12:09:44Z|2021-05-30T07:10:13Z|
[teleterm](https://github.com/alfiankan/teleterm)|Telegram Bot Exec Terminal Command|5|2|0|2020-12-31T22:34:18Z|2021-10-22T02:53:06Z|


## Build Automation
*Libraries and tools helping with build automation.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[task](https://github.com/go-task/task)|A task runner / simpler Make alternative written in Go|4250|264|100|2017-02-27T00:46:04Z|2021-12-04T15:51:11Z|
[realize](https://github.com/oxequa/realize)|Realize is the #1 Golang Task Runner which enhance your workflow by automating the most common tasks and using the best performing Golang live reloading.|4152|216|68|2016-07-12T08:07:25Z|2021-05-14T21:47:38Z|
[mmake](https://github.com/tj/mmake)|Modern Make |1593|43|11|2017-02-15T22:01:21Z|2020-03-02T16:01:44Z|
[goyek](https://github.com/goyek/goyek)|Create build pipelines in Go |269|20|3|2020-10-11T13:20:55Z|2021-11-29T06:03:55Z|
[taskctl](https://github.com/taskctl/taskctl)|Concurrent task runner, developer&#39;s routine tasks automation toolkit. Simple modern alternative to GNU Make 🧰|145|14|9|2019-11-12T13:19:09Z|2021-05-21T20:14:40Z|
[1build](https://github.com/gopinath-langote/1build)|Frictionless way of managing project-specific commands|116|28|32|2019-04-23T17:05:38Z|2021-11-24T04:26:15Z|
[gilbert](https://github.com/go-gilbert/gilbert)|Build system and task runner for Go projects|93|6|0|2019-01-30T09:02:31Z|2020-04-25T14:24:42Z|
[gaper](https://github.com/maxcnunes/gaper)|Builds and restarts a Go project when it crashes or some watched file changes|50|4|7|2018-06-16T02:46:38Z|2019-12-04T21:20:45Z|
[anko](https://github.com/GuilhermeCaruso/anko)|:crystal_ball: Simple application watcher|19|0|0|2021-03-02T14:08:42Z|2021-03-28T15:09:08Z|


### Standard CLI
*Libraries for building standard or basic Command Line applications.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[cobra](https://github.com/spf13/cobra)|A Commander for modern Go CLI interactions|24337|2110|340|2013-09-03T20:40:26Z|2021-12-11T19:41:44Z|
[cli](https://github.com/urfave/cli)|A simple, fast, and fun package for building command line apps in Go|16959|1473|71|2013-07-13T19:32:06Z|2021-12-05T19:56:34Z|
[elvish](https://github.com/elves/elvish)|Elvish = Expressive Programming Language &#43; Versatile Interactive Shell|4557|262|256|2013-06-16T08:43:32Z|2021-12-11T12:15:05Z|
[kingpin](https://github.com/alecthomas/kingpin)|CONTRIBUTIONS ONLY: A Go (golang) command line and flag parser|3176|245|25|2014-05-14T20:09:04Z|2021-10-26T19:12:45Z|
[dnote](https://github.com/dnote/dnote)|A simple command line notebook for programmers|2190|90|59|2017-03-30T23:07:25Z|2021-12-08T11:26:19Z|
[go-flags](https://github.com/jessevdk/go-flags)|go command line option parser|2090|255|41|2012-08-31T13:57:58Z|2021-12-02T13:05:36Z|
[pflag](https://github.com/spf13/pflag)|Drop-in replacement for Go&#39;s flag package, implementing POSIX/GNU-style --flags.|1638|269|118|2013-08-30T14:53:31Z|2021-12-11T08:20:16Z|
[cli](https://github.com/mitchellh/cli)|A Go library for implementing command-line interfaces.|1429|109|9|2013-11-03T06:47:54Z|2021-08-13T13:07:39Z|
[go-arg](https://github.com/alexflint/go-arg)|Struct-based argument parsing in Go|1315|77|7|2015-11-01T01:30:06Z|2021-11-10T16:40:28Z|
[liner](https://github.com/peterh/liner)|Pure Go line editor with history, inspired by linenoise|841|110|13|2012-08-15T16:34:55Z|2021-11-19T16:09:21Z|
[complete](https://github.com/posener/complete)|bash completion written in go &#43; bash completion for go command|802|63|21|2017-05-05T21:34:07Z|2021-04-07T21:04:50Z|
[ops](https://github.com/nanovms/ops)|ops - build and run nanos unikernels|788|77|140|2018-09-10T17:57:47Z|2021-12-08T15:51:33Z|
[flaggy](https://github.com/integrii/flaggy)|Idiomatic Go input parsing with subcommands, positional values, and flags at any position. No required project or package layout and no external dependencies.|769|29|18|2018-03-05T05:55:05Z|2021-09-01T03:22:41Z|
[mow.cli](https://github.com/jawher/mow.cli)|A versatile library for building CLI applications in Go|768|52|26|2014-12-18T19:34:20Z|2021-10-23T15:12:39Z|
[kong](https://github.com/alecthomas/kong)|Kong is a command-line parser for Go|737|74|17|2018-04-10T06:50:32Z|2021-12-04T20:40:21Z|
[cli](https://github.com/mkideal/cli)|CLI - A package for building command line app with go|625|42|3|2016-02-26T16:45:29Z|2021-09-15T13:56:17Z|
[argparse](https://github.com/akamensky/argparse)|Argparse for golang. Just because `flag` sucks|383|44|8|2017-11-24T06:42:20Z|2021-08-13T04:27:10Z|
[climax](https://github.com/tucnak/climax)|Climax is an alternative CLI with the human face|191|18|7|2015-11-03T21:04:57Z|2020-09-05T07:02:16Z|
[hiboot](https://github.com/hidevopsio/hiboot)|hiboot is a high performance web and cli application framework with dependency injection support|162|28|6|2018-03-16T11:21:46Z|2021-06-29T05:34:35Z|
[commandeer](https://github.com/jaffee/commandeer)|Automatically sets up command line flags based on struct fields and tags.|146|15|3|2017-10-12T02:51:05Z|2021-06-16T20:17:08Z|
[wmenu](https://github.com/dixonwille/wmenu)|An easy to use menu structure for cli applications that prompts users to make choices.|145|21|1|2016-04-20T13:09:44Z|2021-08-31T17:22:54Z|
[sflags](https://github.com/octago/sflags)|Generate flags by parsing structures|129|28|9|2016-12-04T14:49:27Z|2021-07-26T01:27:06Z|
[flag](https://github.com/cosiner/flag)|Flag is a simple but powerful command line option parsing library for Go support infinite level subcommand|117|7|1|2016-10-05T16:49:41Z|2020-12-27T11:14:27Z|
[clif](https://github.com/ukautz/clif)|Another CLI framework for Go. It works on my machine.|111|14|3|2015-05-30T18:30:08Z|2019-02-18T14:43:25Z|
[job](https://github.com/liujianping/job)|JOB, make your short-term command as a long-term job. 将命令行规划成任务的工具|110|10|1|2019-04-09T11:14:51Z|2020-06-30T10:17:38Z|
[cli](https://github.com/teris-io/cli)|Simple and complete API for building command line applications in Go|96|8|2|2017-05-24T23:07:07Z|2021-05-09T19:28:00Z|
[cmdr](https://github.com/hedzr/cmdr)|Golang library with POSIX-compliant command-line UI (CLI) and Hierarchical-configuration. Better substitute for &#39;flag&#39;.|89|8|1|2019-05-15T09:58:02Z|2021-12-11T10:24:56Z|
[env](https://github.com/codingconcepts/env)|Tag-based environment configuration for structs|88|8|1|2017-06-14T20:01:55Z|2020-08-21T22:01:19Z|
[clir](https://github.com/leaanthony/clir)|A Simple and Clear CLI library. Dependency free.|78|11|4|2019-11-18T19:52:00Z|2021-10-11T08:16:34Z|
[gocmd](https://github.com/devfacet/gocmd)|A Go library for building command line applications.|56|6|1|2018-01-08T04:52:02Z|2021-05-08T04:04:02Z|
[wlog](https://github.com/dixonwille/wlog)|A simple logging interface that supports cross-platform color and concurrency.|52|7|0|2016-04-13T16:47:40Z|2021-08-31T17:23:26Z|
[strumt](https://github.com/antham/strumt)|Strumt is a library to create prompt chain|44|5|0|2017-06-19T19:33:16Z|2021-04-28T21:56:59Z|
[go-getoptions](https://github.com/DavidGamba/go-getoptions)|Fully featured Go (golang) command line option parser with built-in auto-completion support.|38|8|1|2015-12-18T02:21:14Z|2021-12-10T07:03:50Z|
[flagvar](https://github.com/sgreben/flagvar)|A collection of CLI argument types for the Go `flag` package.|37|2|1|2018-05-18T18:45:16Z|2020-07-11T12:26:29Z|
[cmd](https://github.com/posener/cmd)|The standard library flag package with its missing features|33|2|0|2019-10-29T00:32:11Z|2020-09-27T14:26:26Z|
[argv](https://github.com/cosiner/argv)||31|8|1|2017-01-22T10:37:21Z|2020-04-16T04:13:15Z|
[acmd](https://github.com/cristalhq/acmd)|Simple, useful and opinionated CLI package in Go.|27|0|1|2021-10-27T15:13:31Z|2021-11-17T19:03:44Z|
[go-commander](https://github.com/yitsushi/go-commander)|Go library to simplify CLI workflow|24|5|1|2016-10-10T10:09:41Z|2020-05-24T20:27:55Z|
[sand](https://github.com/Zaba505/sand)|Package for creating interpreters|16|2|0|2018-11-18T22:44:41Z|2018-11-21T19:13:47Z|
[go-command-chain](https://github.com/rainu/go-command-chain)|A go library for easy configure and run command chains. Such like pipelining in unix shells.|16|1|1|2021-05-12T17:47:41Z|2021-06-02T14:25:06Z|
[ts](https://github.com/liujianping/ts)|timestamp convert &amp; compare tool. 时间戳转换与对比工具|13|2|0|2019-06-25T10:21:13Z|2019-07-02T02:41:06Z|
[go-andotp](https://github.com/RijulGulati/go-andotp)|CLI program to encrypt/decrypt andOTP files|10|0|0|2021-05-09T16:58:51Z|2021-06-03T19:08:16Z|
[subcmd](https://github.com/bobg/subcmd)||1|0|0|2020-07-29T15:04:00Z|2021-09-03T15:39:52Z|


### Advanced Console UIs
*Libraries for building Console Applications and Console User Interfaces.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[termui](https://github.com/gizak/termui)|Golang terminal dashboard|11449|735|87|2015-02-03T14:09:27Z|2021-11-21T12:56:32Z|
[gocui](https://github.com/jroimartin/gocui)|Minimalist Go package aimed at creating Console User Interfaces.|7730|514|69|2014-01-04T02:50:20Z|2021-11-08T23:12:38Z|
[go-prompt](https://github.com/c-bata/go-prompt)|Building powerful interactive prompts in Go, inspired by python-prompt-toolkit.|4243|268|90|2017-08-14T16:02:09Z|2021-10-06T15:02:46Z|
[termbox-go](https://github.com/nsf/termbox-go)|Pure Go termbox implementation|4194|362|41|2012-01-12T21:03:03Z|2021-05-09T15:36:41Z|
[pterm](https://github.com/pterm/pterm)|✨ #PTerm is a modern Go module to beautify console output. Featuring charts, progressbars, tables, trees, and much more 🚀 It&#39;s completely configurable and 100% cross-platform compatible.|2223|71|24|2020-09-17T15:52:59Z|2021-11-17T14:20:36Z|
[progressbar](https://github.com/schollz/progressbar)|A really basic thread-safe progress bar for Golang applications|2203|127|21|2017-10-26T18:28:10Z|2021-11-17T18:32:39Z|
[uiprogress](https://github.com/gosuri/uiprogress)|A go library to render progress bars in terminal applications|1869|123|27|2015-11-17T00:59:24Z|2021-08-30T09:11:08Z|
[termdash](https://github.com/mum4k/termdash)|Terminal based dashboard.|1848|96|38|2018-03-24T12:01:49Z|2021-11-24T07:29:42Z|
[asciigraph](https://github.com/guptarohit/asciigraph)|Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies.|1801|71|7|2018-06-17T10:37:16Z|2021-10-08T15:29:11Z|
[spinner](https://github.com/briandowns/spinner)|Go (golang) package with 80 configurable terminal spinner/progress indicators.|1624|105|13|2014-12-13T00:36:19Z|2021-12-06T16:29:44Z|
[mpb](https://github.com/vbauerster/mpb)|multi progress bar for Go cli applications|1495|88|1|2016-12-14T11:56:29Z|2021-12-04T08:14:54Z|
[uilive](https://github.com/gosuri/uilive)|uilive is a go library for updating terminal output in realtime|1391|68|11|2015-11-16T06:13:10Z|2021-09-10T12:25:23Z|
[aurora](https://github.com/logrusorgru/aurora)|Golang ultimate ANSI-colors that supports Printf/Sprintf methods|1148|52|4|2016-11-06T22:37:12Z|2021-02-09T22:00:44Z|
[color](https://github.com/gookit/color)|🎨 Terminal color rendering library, support 8/16 colors, 256 colors, RGB color rendering output, support Print/Sprintf methods, compatible with Windows. GO CLI 控制台颜色渲染工具库，支持16色，256色，RGB色彩渲染输出，使用类似于 Print/Sprintf，兼容并支持 Windows 环境的色彩渲染|984|66|1|2018-07-01T07:28:17Z|2021-11-13T13:01:04Z|
[termenv](https://github.com/muesli/termenv)|Advanced ANSI style &amp; color support for your terminal applications|842|34|14|2019-12-07T06:35:57Z|2021-08-30T23:22:25Z|
[uitable](https://github.com/gosuri/uitable)|A go library to improve readability in terminal apps using tabular data|621|27|3|2015-11-13T21:59:21Z|2020-10-21T11:54:50Z|
[go-isatty](https://github.com/mattn/go-isatty)||579|80|6|2014-04-01T01:53:09Z|2021-11-21T12:24:09Z|
[go-colorable](https://github.com/mattn/go-colorable)||571|76|8|2014-07-30T02:38:06Z|2021-11-23T14:53:07Z|
[gommon](https://github.com/labstack/gommon)|Common packages for Go|420|94|16|2015-03-12T22:35:57Z|2021-11-29T15:58:08Z|
[chalk](https://github.com/ttacon/chalk)|Intuitive package for prettifying terminal/console output. http://godoc.org/github.com/ttacon/chalk|385|19|4|2014-07-18T19:38:58Z|2019-08-28T23:55:36Z|
[simpletable](https://github.com/alexeyco/simpletable)|Simple tables in terminal with Go|333|24|2|2017-03-29T07:27:23Z|2021-04-23T14:55:10Z|
[tabby](https://github.com/cheynewallace/tabby)|A tiny library for super simple Golang tables|301|13|2|2018-12-17T23:35:39Z|2020-12-23T01:20:33Z|
[yacspin](https://github.com/theckman/yacspin)|Yet Another CLi Spinner; providing over 70 easy to use and customizable terminal spinners for multiple OSes|213|7|1|2019-12-29T07:41:23Z|2021-12-11T21:03:43Z|
[go-colortext](https://github.com/daviddengcn/go-colortext)|Change the color of console text.|209|20|4|2013-01-23T03:38:54Z|2020-03-29T21:12:20Z|
[box-cli-maker](https://github.com/Delta456/box-cli-maker)|Make Highly Customized Boxes for your CLI|158|5|0|2020-05-01T07:23:56Z|2021-10-04T17:09:19Z|
[cfmt](https://github.com/mingrammer/cfmt)|:art: Contextual fmt inspired by bootstrap color classes|81|6|1|2018-03-15T19:04:27Z|2018-12-07T17:31:52Z|
[tabular](https://github.com/InVisionApp/tabular)|Tabular simplifies printing ASCII tables from command line utilities|55|5|0|2018-04-23T21:17:03Z|2018-05-14T19:04:57Z|
[ctc](https://github.com/wzshiming/ctc)|Console Text Colors - The non-invasive cross-platform terminal color library does not need to modify the Print method|36|3|0|2018-04-27T18:07:42Z|2020-07-15T08:09:32Z|
[cfmt](https://github.com/i582/cfmt)|Small library for simple and convenient formatted stylized output to the console.|32|2|0|2020-11-13T20:29:45Z|2021-07-01T14:07:37Z|
[colourize](https://github.com/TreyBastian/colourize)|An ANSI colour terminal package for Go|24|4|0|2015-05-11T11:49:39Z|2016-05-10T09:50:02Z|
[marker](https://github.com/cyucelen/marker)| 🖍️ Marker is the easiest way to match and mark strings for colorful terminal outputs!|22|13|4|2019-08-28T15:44:08Z|2021-09-08T01:22:25Z|
[table](https://github.com/tomlazar/table)|pretty colorfull tables in go with less effort|13|1|0|2020-09-22T05:42:34Z|2021-03-04T17:40:38Z|
[go-ataman](https://github.com/workanator/go-ataman)|Another Text Attribute Manupulator|11|2|0|2017-05-17T19:04:57Z|2020-12-23T05:36:05Z|


## Configuration
*Libraries for configuration parsing.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[viper](https://github.com/spf13/viper)|Go configuration with fangs|17587|1522|439|2014-04-02T14:33:33Z|2021-12-09T17:58:28Z|
[godotenv](https://github.com/joho/godotenv)|A Go port of Ruby&#39;s dotenv library (Loads environment variables from `.env`.)|4369|264|55|2013-07-30T07:45:19Z|2021-10-05T15:44:45Z|
[envconfig](https://github.com/kelseyhightower/envconfig)|Golang library for managing configuration data from environment variables|3887|322|51|2013-11-06T17:01:55Z|2021-12-09T08:11:00Z|
[ini](https://github.com/go-ini/ini)|Package ini provides INI file read and write functionality in Go|2764|331|27|2014-12-18T07:36:37Z|2021-12-02T14:07:41Z|
[env](https://github.com/caarlos0/env)|Simple lib to parse environment variables to structs|2151|151|2|2015-07-28T02:14:37Z|2021-12-01T20:03:41Z|
[kong](https://github.com/alecthomas/kong)|Kong is a command-line parser for Go|737|74|17|2018-04-10T06:50:32Z|2021-12-04T20:40:21Z|
[koanf](https://github.com/knadh/koanf)|Simple, lightweight, extensible, configuration management library for Go. Support for JSON, TOML, YAML, env, command line, file, S3 etc. Alternative to viper.|666|63|1|2019-06-18T06:34:05Z|2021-11-30T09:45:23Z|
[konfig](https://github.com/lalamove/konfig)|Composable, observable and performant config handling for Go for the distributed processing era|614|45|4|2019-01-18T17:03:03Z|2020-10-28T08:24:08Z|
[confita](https://github.com/heetch/confita)|Load configuration in cascade from multiple backends into a struct|419|45|22|2017-12-21T10:49:18Z|2021-07-24T10:21:20Z|
[cleanenv](https://github.com/ilyakaznacheev/cleanenv)|✨Clean and minimalistic environment configuration reader for Golang|409|43|17|2019-07-12T15:28:52Z|2021-12-09T11:41:57Z|
[aconfig](https://github.com/cristalhq/aconfig)|Simple, useful and opinionated config loader.|320|22|10|2020-06-26T19:43:20Z|2021-11-04T08:20:09Z|
[config](https://github.com/gookit/config)|📝 Go config manage(load,get,set). support JSON, YAML, TOML, INI, HCL, ENV and Flags. Multi file load, data override merge, parse ENV var. Go应用配置加载管理，支持多种格式，多文件加载，远程文件加载，支持数据合并，解析环境变量名|305|38|0|2018-07-07T08:11:39Z|2021-11-19T02:42:31Z|
[config](https://github.com/JeremyLoy/config)|12 factor configuration as a typesafe struct in as little as two function calls|284|13|2|2019-04-02T13:41:22Z|2021-11-18T16:50:16Z|
[store](https://github.com/tucnak/store)|A dead simple configuration manager for Go applications|256|18|2|2015-10-03T19:17:28Z|2017-09-05T11:38:35Z|
[hjson-go](https://github.com/hjson/hjson-go)|Hjson for Go|246|32|4|2016-08-05T22:59:18Z|2021-11-22T22:48:37Z|
[config](https://github.com/olebedev/config)|JSON or YAML configuration wrapper with convenient access methods.|241|41|4|2014-04-21T15:09:39Z|2021-12-09T09:15:05Z|
[envconfig](https://github.com/vrischmann/envconfig)|Small library to read your configuration from environment variables|216|25|1|2015-04-21T23:37:17Z|2021-10-24T13:21:10Z|
[config](https://github.com/golobby/config)|A lightweight yet powerful configuration manager for Go projects|213|18|0|2019-10-15T22:51:19Z|2021-12-03T21:26:57Z|
[config](https://github.com/joshbetz/config)|🛠 A configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP|209|13|0|2017-04-02T18:37:05Z|2021-11-12T16:58:10Z|
[fig](https://github.com/kkyr/fig)|A minimalist Go configuration library|182|15|4|2020-01-16T18:43:19Z|2021-12-10T21:47:42Z|
[gcfg](https://github.com/go-gcfg/gcfg)|read INI-style configuration files into Go structs; supports user-defined types and subsections|155|51|9|2015-08-17T14:40:55Z|2021-07-02T06:41:18Z|
[xdg](https://github.com/adrg/xdg)|Go implementation of the XDG Base Directory Specification and XDG user directories|154|9|0|2014-08-22T08:23:40Z|2021-10-27T12:47:13Z|
[goconfig](https://github.com/gosidekick/goconfig)|goconfig uses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file.|150|22|6|2016-12-18T11:22:41Z|2021-10-21T20:30:46Z|
[envh](https://github.com/antham/envh)|Go helpers to manage environment variables|95|2|0|2017-01-12T11:25:48Z|2021-04-28T21:51:57Z|
[envcfg](https://github.com/tomazk/envcfg)|Un-marshaling environment variables to Go structs|94|7|0|2014-11-29T11:43:53Z|2017-06-19T15:53:22Z|
[harvester](https://github.com/beatlabs/harvester)|Harvest configuration, watch and notify subscriber|93|22|4|2019-04-09T07:37:19Z|2021-11-29T07:58:10Z|
[onion](https://github.com/goraz/onion)|Layer based configuration for golang|92|10|7|2015-07-22T14:28:21Z|2021-08-22T16:51:14Z|
[configuro](https://github.com/sherifabdlnaby/configuro)|An opinionated configuration loading framework for Containerized and Cloud-Native applications.|79|9|0|2020-04-09T22:10:34Z|2021-03-09T04:21:18Z|
[xdg](https://github.com/OpenPeeDeeP/xdg)|A cross platform package that follows the XDG Standard|65|6|1|2017-07-20T15:58:29Z|2020-10-19T13:34:26Z|
[gofigure](https://github.com/ian-kent/gofigure)|Go configuration made easy!|60|9|1|2014-11-25T00:12:40Z|2019-09-15T00:17:39Z|
[configure](https://github.com/paked/configure)|Configure is a Go package that gives you easy configuration of your project through redundancy|54|9|2|2015-06-14T07:46:56Z|2019-02-18T14:01:49Z|
[go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm)|Go package that interfaces with AWS System Manager|43|10|0|2019-01-24T09:01:19Z|2021-03-17T17:47:08Z|
[configuration](https://github.com/BoRuDar/configuration)|Library for setting values to structs&#39; fields from env, flags, files or default tag|42|8|1|2019-11-27T17:58:49Z|2021-11-04T23:19:10Z|
[uconfig](https://github.com/omeid/uconfig)|Lightweight, zero-dependency, and extendable configuration management library for Go|40|4|0|2017-05-11T01:21:44Z|2021-12-01T21:03:29Z|
[hocon](https://github.com/gurkankaymak/hocon)|go implementation of lightbend&#39;s HOCON configuration library https://github.com/lightbend/config|37|8|2|2020-03-01T18:20:12Z|2021-11-23T12:14:46Z|
[gone](https://github.com/One-com/gone)|Golang packages for writing small daemons and servers.|36|6|0|2016-09-05T09:39:11Z|2021-05-24T14:23:37Z|
[go-up](https://github.com/ufoscout/go-up)|go-up! A simple configuration library with recursive placeholders resolution and no magic.|35|6|1|2018-02-18T09:50:00Z|2020-01-14T07:21:58Z|
[ingo](https://github.com/schachmat/ingo)|persistent storage for flags in go|34|8|0|2016-02-07T22:57:40Z|2017-04-03T01:15:10Z|
[mini](https://github.com/sasbury/mini)|A golang package for parsing ini-style configuration files|28|6|1|2015-04-29T23:52:36Z|2018-12-26T23:28:05Z|
[genv](https://github.com/sakirsensoy/genv)|Genv is a library for Go (golang) that makes it easy to read and use environment variables in your projects. It also allows environment variables to be loaded from the .env file.|25|2|0|2019-07-15T10:25:57Z|2019-07-27T11:56:32Z|
[conflate](https://github.com/the4thamigo-uk/conflate)|Library providing routines to merge and validate JSON, YAML and/or TOML files|20|3|0|2018-02-01T19:06:15Z|2020-09-21T09:50:49Z|
[go-ssm-config](https://github.com/ianlopshire/go-ssm-config)|Go utility for loading configuration parameters from AWS SSM (Parameter Store)|10|8|4|2019-12-02T18:47:38Z|2020-12-15T16:19:27Z|
[envconf](https://github.com/ian-kent/envconf)|Configure Go applications from the environment|9|4|0|2014-10-26T12:12:26Z|2014-10-26T12:12:40Z|
[ini](https://github.com/wlevene/ini)|ini parser for golang|8|1|0|2021-08-13T12:13:44Z|2021-12-02T09:11:37Z|
[env](https://github.com/nasermirzaei89/env)|Golang Get Environment Variables Package|7|1|0|2019-07-24T06:37:13Z|2021-11-21T23:55:20Z|
[go-ini](https://github.com/subpop/go-ini)|automatic mirror of https://git.sr.ht/~spc/go-ini|5|2|1|2019-09-11T18:38:20Z|2021-04-06T17:32:24Z|
[typenv](https://github.com/diegomarangoni/typenv)|Go minimalist typed environment variables library|4|1|0|2020-06-30T18:26:09Z|2020-07-22T16:23:05Z|
[swap](https://github.com/oblq/swap)|Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env).|4|1|0|2020-04-12T23:28:19Z|2021-11-07T11:00:53Z|
[go-conf](https://github.com/ThomasObenaus/go-conf)|Library for easy configuration of a golang service|2|1|1|2021-01-27T21:41:47Z|2021-10-19T12:43:09Z|
[gonfig](https://github.com/miladabc/gonfig)|Tag based configuration loader from different providers|2|0|0|2021-01-21T13:44:44Z|2021-08-02T20:37:02Z|


## Continuous Integration
*Tools for help with continuous integration.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[drone](https://github.com/harness/drone)|Drone is a Container-Native, Continuous Delivery Platform|24262|2381|51|2014-02-07T07:54:44Z|2021-12-08T12:02:05Z|
[cds](https://github.com/ovh/cds)|Enterprise-Grade Continuous Delivery &amp; DevOps Automation Open Source Platform|3665|343|170|2016-10-11T08:28:23Z|2021-12-10T13:41:35Z|
[goveralls](https://github.com/mattn/goveralls)||704|130|13|2013-04-17T10:58:40Z|2021-09-21T08:34:22Z|
[gotestfmt](https://github.com/haveyoudebuggedit/gotestfmt)|go test output for humans|169|1|1|2021-04-29T21:17:30Z|2021-12-08T10:43:13Z|
[overalls](https://github.com/go-playground/overalls)|:jeans:Multi-Package go project coverprofile for tools like goveralls|107|27|3|2015-07-30T11:30:11Z|2019-12-30T18:54:48Z|
[duci](https://github.com/duck8823/duci)|The simple ci server |69|4|7|2018-04-01T01:51:02Z|2021-12-09T20:15:42Z|
[gomason](https://github.com/nikogura/gomason)|A tool for testing, building, signing, and publishing binaries.|51|7|2|2017-11-18T00:59:11Z|2021-07-08T22:29:02Z|
[roveralls](https://github.com/lawrencewoodman/roveralls)|A Go recursive coverage testing tool|15|5|0|2016-06-26T07:45:32Z|2017-11-19T19:39:13Z|


## CSS Preprocessors
*Libraries for preprocessing CSS files.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[gcss](https://github.com/yosssi/gcss)|Pure Go CSS Preprocessor|446|36|8|2014-09-04T14:38:20Z|2014-10-12T14:07:10Z|
[go-libsass](https://github.com/wellington/go-libsass)|Go wrapper for libsass, the only Sass 3.5 compiler for Go|181|24|13|2015-04-19T15:09:47Z|2020-10-23T19:07:14Z|


## Data Structures
*Generic datastructures and algorithms in Go.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[gods](https://github.com/emirpasic/gods)|GoDS (Go Data Structures). Containers (Sets, Lists, Stacks, Maps, Trees), Sets (HashSet, TreeSet, LinkedHashSet), Lists (ArrayList, SinglyLinkedList, DoublyLinkedList), Stacks (LinkedListStack, ArrayStack), Maps (HashMap, TreeMap, HashBidiMap, TreeBidiMap, LinkedHashMap), Trees (RedBlackTree, AVLTree, BTree, BinaryHeap), Comparators, Iterators, Enumerables, Sort, JSON|10854|1288|56|2015-03-04T14:19:52Z|2021-11-03T14:33:49Z|
[go-datastructures](https://github.com/Workiva/go-datastructures)|A collection of useful, performant, and threadsafe Go datastructures.|6288|732|25|2014-10-29T13:55:17Z|2021-04-23T19:44:34Z|
[golang-set](https://github.com/deckarep/golang-set)|A simple set type for the Go language. Trusted by Docker, 1Password, Ethereum and Hashicorp.|2149|183|12|2013-07-03T21:52:01Z|2021-03-19T07:05:45Z|
[gota](https://github.com/go-gota/gota)|Gota: DataFrames and data wrangling in Go (Golang)|1892|188|47|2016-02-06T17:23:25Z|2021-11-26T12:35:36Z|
[roaring](https://github.com/RoaringBitmap/roaring)|Roaring bitmaps in Go (golang)|1392|156|60|2014-07-10T20:14:34Z|2021-12-01T15:25:15Z|
[BoomFilters](https://github.com/tylertreat/BoomFilters)|Probabilistic data structures for processing continuous, unbounded streams.|1391|99|10|2015-02-06T02:01:26Z|2021-03-15T20:15:27Z|
[bloom](https://github.com/bits-and-blooms/bloom)|Go package implementing Bloom filters|1383|180|3|2011-05-21T14:18:41Z|2021-09-30T13:35:09Z|
[gocache](https://github.com/eko/gocache)|☔️ A complete Go cache library that brings you multiple ways of managing your caches|1015|98|12|2019-10-05T08:13:54Z|2021-11-13T14:39:44Z|
[cuckoofilter](https://github.com/seiflotfy/cuckoofilter)|Cuckoo Filter: Practically Better Than Bloom|851|76|12|2015-06-28T23:22:09Z|2021-08-04T10:41:37Z|
[bitset](https://github.com/bits-and-blooms/bitset)|Go package implementing bitsets|760|124|3|2011-05-11T03:33:44Z|2021-09-09T18:13:12Z|
[hyperloglog](https://github.com/axiomhq/hyperloglog)|HyperLogLog with lots of sugar (Sparse, LogLog-Beta bias correction and TailCut space reduction)|758|57|3|2017-06-18T11:18:12Z|2021-10-21T16:48:51Z|
[algorithms](https://github.com/shady831213/algorithms)|CLRS study. Codes are written with golang.|591|96|0|2018-01-31T09:27:56Z|2021-03-17T08:01:38Z|
[trie](https://github.com/derekparker/trie)|Data structure and relevant algorithms for extremely fast prefix/fuzzy string searching.|558|97|10|2014-03-06T22:01:49Z|2020-08-06T13:25:33Z|
[gostl](https://github.com/liyue201/gostl)|Data structure and algorithm library for go, designed to provide functions similar to C&#43;&#43; STL|344|69|1|2019-10-12T01:10:24Z|2021-09-28T07:02:04Z|
[go-geoindex](https://github.com/hailocab/go-geoindex)|Go native library for fast point tracking and K-Nearest queries|332|42|3|2015-01-22T12:26:17Z|2018-02-20T21:58:39Z|
[ttlcache](https://github.com/ReneKroon/ttlcache)|An in-memory string-interface{} map with various expiration options for golang|312|60|5|2014-12-13T01:55:40Z|2021-12-09T17:28:23Z|
[merkletree](https://github.com/cbergoon/merkletree)|A Merkle Tree implementation written in Go.|297|76|6|2017-04-12T02:50:11Z|2019-11-21T17:40:51Z|
[go-edlib](https://github.com/hbollon/go-edlib)|📚 String comparison and edit distance algorithms library, featuring : Levenshtein, LCS, Hamming, Damerau levenshtein (OSA and Adjacent transpositions algorithms), Jaro-Winkler, Cosine, etc...|293|16|2|2020-08-18T09:30:59Z|2021-12-10T05:10:00Z|
[deque](https://github.com/gammazero/deque)|Fast ring-buffer deque (double-ended queue)|261|24|0|2018-04-24T02:57:55Z|2021-05-08T23:45:03Z|
[hilbert](https://github.com/google/hilbert)|Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves.|240|36|2|2015-08-06T15:50:00Z|2018-11-22T06:15:33Z|
[goskiplist](https://github.com/ryszard/goskiplist)|A skip list implementation in Go|230|58|6|2012-05-09T05:44:59Z|2019-10-29T10:07:30Z|
[go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree)|Adaptive Radix Trees implemented in Go|216|34|0|2016-04-01T01:40:40Z|2020-08-16T07:15:37Z|
[cuckoo-filter](https://github.com/linvon/cuckoo-filter)|Cuckoo Filter go implement, better than Bloom Filter, configurable and space optimized  布谷鸟过滤器的Go实现，优于布隆过滤器，可以定制化过滤器参数，并进行了空间优化|186|14|0|2021-02-19T12:27:43Z|2021-10-10T08:50:35Z|
[skiplist](https://github.com/MauriceGit/skiplist)|A Go library for an efficient implementation of a skip list: https://godoc.org/github.com/MauriceGit/skiplist|181|28|3|2018-06-23T16:01:51Z|2021-11-05T23:06:23Z|
[binpacker](https://github.com/zhuangsirui/binpacker)|A binary stream packer and unpacker|178|30|2|2016-02-02T10:06:11Z|2021-10-08T04:16:12Z|
[levenshtein](https://github.com/agnivade/levenshtein)|Go implementation to calculate Levenshtein Distance.|166|13|1|2014-07-30T14:03:55Z|2021-05-21T05:40:33Z|
[goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue)|Go concurrent-safe, goroutine-safe, thread-safe queue|153|13|1|2019-01-10T21:21:23Z|2020-05-12T09:16:45Z|
[bloom](https://github.com/zentures/bloom)|Bloom filters implemented in Go.|144|16|1|2013-09-03T02:27:35Z|2018-04-16T07:52:10Z|
[iter](https://github.com/disksing/iter)|Go implementation of C&#43;&#43; STL iterators and algorithms.|138|8|1|2019-10-20T09:29:49Z|2019-12-15T15:29:09Z|
[ring](https://github.com/tannerryan/ring)|Package ring provides a high performance and thread safe Go implementation of a bloom filter.|120|13|2|2019-01-27T04:02:20Z|2020-09-10T16:36:16Z|
[go-rquad](https://github.com/arl/go-rquad)|:pushpin: State of the art point location and neighbour finding algorithms for region quadtrees, in Go|118|4|0|2016-09-12T21:46:37Z|2020-04-19T09:26:33Z|
[encoding](https://github.com/zentures/encoding)|Integer Compression Libraries for Go|110|14|1|2013-09-20T19:29:57Z|2017-12-23T22:15:28Z|
[bit](https://github.com/yourbasic/bit)|Bitset data structure|107|17|0|2017-05-03T19:05:35Z|2018-03-13T07:45:26Z|
[remember-go](https://github.com/rocketlaunchr/remember-go)|Cache Slow Database Queries|105|7|1|2019-04-04T20:24:25Z|2021-04-19T07:43:10Z|
[bitmap](https://github.com/kelindar/bitmap)|Simple dense bitmap index in Go with binary operators|103|8|1|2021-05-28T06:51:29Z|2021-12-07T20:11:43Z|
[conjungo](https://github.com/InVisionApp/conjungo)|A small flexible merge library in go|98|13|10|2016-12-29T23:50:38Z|2020-10-23T10:46:02Z|
[skiplist](https://github.com/gansidui/skiplist)|skiplist for golang|75|19|1|2014-11-18T16:29:53Z|2014-11-21T05:13:52Z|
[go-mcache](https://github.com/OrlovEvgeny/go-mcache)|Fast in-memory key:value store/cache with TTL|74|12|1|2018-04-14T23:31:21Z|2020-01-21T12:43:35Z|
[bloom](https://github.com/yourbasic/bloom)|Probabilistic set data structure|63|9|0|2017-05-06T19:57:47Z|2017-06-19T17:00:50Z|
[levenshtein](https://github.com/agext/levenshtein)|Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.|59|5|0|2016-04-08T00:14:31Z|2020-10-15T13:29:05Z|
[count-min-log](https://github.com/seiflotfy/count-min-log)|Go implementation of Count-Min-Log|53|4|0|2015-08-16T22:31:36Z|2017-02-12T13:09:21Z|
[crunch](https://github.com/superwhiskers/crunch)|take bytes out of things easily ✨🍪|47|7|0|2019-02-27T03:56:52Z|2021-08-07T04:30:16Z|
[goset](https://github.com/zoumo/goset)|Set is a useful collection but there is no built-in implementation in Go lang.|43|13|0|2017-08-25T09:21:30Z|2020-12-11T10:18:54Z|
[nan](https://github.com/kak-tus/nan)|Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers|43|7|0|2020-05-05T20:20:54Z|2021-09-29T16:32:34Z|
[hide](https://github.com/emvi/hide)|ID type with marshalling to/from hash to prevent sending IDs to clients.|40|4|0|2019-01-16T13:54:17Z|2021-11-09T19:21:48Z|
[pipeline](https://github.com/hyfather/pipeline)|Pipelines using goroutines|37|7|1|2018-04-25T00:11:36Z|2021-11-02T22:47:16Z|
[concurrent-writer](https://github.com/free/concurrent-writer)|Highly concurrent drop-in replacement for bufio.Writer|36|7|0|2017-09-18T15:29:59Z|2017-11-17T21:28:32Z|
[timedmap](https://github.com/zekroTJA/timedmap)|A thread safe map which has expiring key-value pairs.|34|5|0|2019-01-30T12:55:37Z|2021-11-20T19:34:03Z|
[deque](https://github.com/edwingeng/deque)|A highly optimized double-ended queue|33|1|0|2019-02-01T03:32:28Z|2021-05-10T08:39:07Z|
[typ](https://github.com/gurukami/typ)|Null Types, Safe primitive type conversion and fetching value from complex structures.|30|2|0|2019-03-03T05:34:23Z|2021-10-15T16:11:56Z|
[null](https://github.com/emvi/null)|Nullable Go types that can be marshalled/unmarshalled to/from JSON.|23|2|1|2018-07-04T21:18:45Z|2021-11-09T16:04:18Z|
[dict](https://github.com/srfrog/dict)|Python-like dictionaries for Go|22|3|0|2019-04-23T02:04:25Z|2020-10-25T20:55:30Z|
[cmap](https://github.com/lrita/cmap)|a thread-safe concurrent map for go|20|1|0|2019-11-26T03:54:59Z|2020-08-18T17:10:05Z|
[go-ef](https://github.com/amallia/go-ef)|A Go implementation of the Elias-Fano encoding|18|6|0|2017-09-22T01:47:16Z|2017-09-25T20:07:11Z|
[set](https://github.com/StudioSol/set)|A simple Set data structure implementation in Go (Golang) using LinkedHashMap.|16|7|0|2018-07-20T21:53:37Z|2021-10-01T13:28:05Z|
[ptrie](https://github.com/viant/ptrie)|A prefix tree implementation in go |16|5|0|2019-05-20T14:13:05Z|2020-09-02T23:51:09Z|
[mspm](https://github.com/BlackRabbitt/mspm)|Multi-String Pattern Matching Algorithm Using TrieHashNode|16|2|0|2018-05-17T18:59:44Z|2018-05-19T06:36:38Z|
[treap](https://github.com/perdata/treap)|golang persistent immutable treap sorted sets|12|4|0|2018-09-16T01:38:03Z|2019-12-18T09:31:05Z|
[ordered-concurrently](https://github.com/tejzpr/ordered-concurrently)|Ordered-concurrently a library for parallel processing with ordered output in Go. Process work concurrently / in parallel and returns output in a channel in the order of input. It is useful in concurrently / parallelly processing items in a queue, and get output in the order provided by the queue.|12|1|2|2021-02-28T17:56:05Z|2021-11-22T21:51:20Z|
[fsm](https://github.com/cocoonspace/fsm)|Finite State Machine package in Go|12|0|0|2021-10-11T10:12:51Z|2021-10-12T20:13:09Z|
[parapipe](https://github.com/nazar256/parapipe)|Paralleling pipeline|12|0|1|2021-04-09T06:49:56Z|2021-06-07T08:11:36Z|
[gofal](https://github.com/xxjwxc/gofal)|fractional api base on golang . golang math tools fractional molecular denominator 分数计算 分子 分母 运算|9|2|0|2019-08-05T07:37:55Z|2019-10-08T03:02:59Z|
[gdcache](https://github.com/ulovecode/gdcache)|gdcache is a pure non-intrusive cache library implemented by golang, you can use it to implement your own cache.|7|1|0|2021-07-20T12:52:02Z|2021-10-14T17:31:29Z|
[goterator](https://github.com/yaa110/goterator)|Lazy iterator implementation for Golang|6|2|0|2020-08-12T19:47:57Z|2020-12-02T04:17:39Z|
[parsefields](https://github.com/MonaxGT/parsefields)|Tools for parse JSON-like logs for collecting unique fields and events|6|1|0|2019-04-12T22:15:10Z|2019-05-05T18:55:53Z|
[slices](https://github.com/srfrog/slices)|Functions that operate on slices. Similar to functions from package strings or package bytes that have been adapted to work with slices.|5|1|0|2020-07-02T23:17:34Z|2020-11-09T08:18:51Z|
[bloomfilter](https://github.com/OldPanda/bloomfilter)|Yet another Bloomfilter implementation in Go, compatible with Java&#39;s Guava library|5|1|0|2021-01-01T01:28:04Z|2021-06-30T00:59:36Z|
[dsu](https://github.com/ihebu/dsu)|Disjoint Set data structure implementation in Go|4|0|0|2021-04-27T16:35:38Z|2021-06-22T23:07:05Z|
[hashsplit](https://github.com/bobg/hashsplit)||3|1|1|2020-04-26T00:30:09Z|2021-08-19T02:46:31Z|
[merkle](https://github.com/bobg/merkle)|Merkle hash trees|1|0|0|2018-10-13T15:25:10Z|2021-08-19T00:28:32Z|


## Database
*Databases implemented in Go.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[prometheus](https://github.com/prometheus/prometheus)|The Prometheus monitoring system and time series database.|40004|6647|566|2012-11-24T11:14:12Z|2021-12-11T10:18:41Z|
[tidb](https://github.com/pingcap/tidb)|TiDB is an open source distributed HTAP database compatible with the MySQL protocol |29790|4806|2693|2015-09-06T04:01:52Z|2021-12-11T17:33:32Z|
[cockroach](https://github.com/cockroachdb/cockroach)|CockroachDB - the open source, cloud-native distributed SQL database.|22659|2950|4449|2014-02-06T00:18:47Z|2021-12-11T20:38:45Z|
[influxdb](https://github.com/influxdata/influxdb)|Scalable datastore for metrics, events, and real-time analytics|22541|3070|1374|2013-09-26T14:31:10Z|2021-12-10T22:38:12Z|
[dgraph](https://github.com/dgraph-io/dgraph)|Native GraphQL Database with graph backend|17085|1248|93|2015-08-25T07:15:56Z|2021-12-07T11:36:18Z|
[groupcache](https://github.com/golang/groupcache)|groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.|10951|1220|36|2013-07-22T21:55:07Z|2021-05-11T15:58:00Z|
[badger](https://github.com/dgraph-io/badger)|Fast key-value DB in Go.|10097|861|8|2017-01-26T05:09:49Z|2021-12-10T09:18:47Z|
[rqlite](https://github.com/rqlite/rqlite)|The lightweight, distributed relational database built on SQLite|9111|486|41|2014-08-23T04:31:18Z|2021-12-02T11:54:49Z|
[milvus](https://github.com/milvus-io/milvus)|An open-source vector database for embedding similarity search and AI applications.|8951|1356|230|2019-09-16T06:43:43Z|2021-12-11T18:01:11Z|
[go-cache](https://github.com/patrickmn/go-cache)|An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.|5647|686|57|2012-01-02T13:07:13Z|2021-11-20T20:02:27Z|
[VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics)|VictoriaMetrics: fast, cost-effective monitoring solution and time series database|5522|482|397|2018-09-30T09:58:01Z|2021-12-10T12:01:43Z|
[bigcache](https://github.com/allegro/bigcache)|Efficient cache for gigabytes of data written in Go.|5283|446|58|2016-03-23T07:18:52Z|2021-12-03T09:48:59Z|
[bbolt](https://github.com/etcd-io/bbolt)|An embedded key/value database for Go.|5003|395|123|2017-06-17T01:42:09Z|2021-11-26T19:54:36Z|
[goleveldb](https://github.com/syndtr/goleveldb)|LevelDB key/value database in Go.|4761|719|80|2013-01-23T04:08:58Z|2021-12-09T12:15:28Z|
[dtm](https://github.com/yedf/dtm)|🔥A cross-language distributed transaction manager. Support xa, tcc, saga, transactional messages.  跨语言分布式事务管理器|3967|391|15|2021-05-16T00:56:28Z|2021-12-11T07:32:39Z|
[ledisdb](https://github.com/ledisdb/ledisdb)|A high performance NoSQL Database Server powered by Go|3766|428|2|2014-04-30T00:43:09Z|2021-03-11T08:55:48Z|
[buntdb](https://github.com/tidwall/buntdb)|BuntDB is an embeddable, in-memory key/value database for Go with custom indexing and geospatial support|3572|249|4|2016-07-19T22:11:40Z|2021-12-08T23:25:40Z|
[immudb](https://github.com/codenotary/immudb)|immudb - world’s fastest immutable database, built on a zero trust model|3363|144|74|2019-11-07T08:22:16Z|2021-12-11T13:28:09Z|
[tiedot](https://github.com/HouzuoGuo/tiedot)|A rudimentary implementation of a basic document (NoSQL) database in Go|2643|263|26|2013-05-26T10:03:49Z|2021-09-05T17:47:27Z|
[rosedb](https://github.com/roseduan/rosedb)|🚀A fast, stable and embedded k-v storage in pure Golang, supports string, list, hash, set, sorted set. 一个 Go 语言实现的快速、稳定、内嵌的 k-v 存储引擎。|2097|324|3|2020-12-06T07:02:48Z|2021-12-11T05:31:37Z|
[nutsdb](https://github.com/xujiajun/nutsdb)|A simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set.|1812|168|26|2018-12-07T07:03:38Z|2021-12-07T13:12:58Z|
[gcache](https://github.com/bluele/gcache)|An in-memory cache library for golang. It supports multiple eviction policies: LRU, LFU, ARC|1783|207|19|2015-01-24T18:17:07Z|2021-08-10T18:49:50Z|
[cache2go](https://github.com/muesli/cache2go)|Concurrency-safe Go caching library with expiration capabilities and access counters|1653|468|24|2013-11-11T03:45:02Z|2021-11-24T06:43:12Z|
[godis](https://github.com/HDT3213/godis)|A Golang implemented Redis Server and Cluster. Go 语言实现的 Redis 服务器和分布式集群|1547|247|0|2019-06-01T07:49:11Z|2021-11-18T06:54:11Z|
[fastcache](https://github.com/VictoriaMetrics/fastcache)|Fast thread-safe inmemory cache for big number of entries in Go. Minimizes GC overhead|1298|105|33|2018-11-22T22:50:13Z|2021-12-02T08:19:21Z|
[CovenantSQL](https://github.com/CovenantSQL/CovenantSQL)|A decentralized, trusted, high performance, SQL database with blockchain features|1247|140|26|2018-04-11T09:52:58Z|2021-08-23T03:35:51Z|
[diskv](https://github.com/peterbourgon/diskv)|A disk-backed key-value store.|1120|96|9|2012-03-21T16:44:32Z|2021-11-10T01:12:08Z|
[databunker](https://github.com/securitybunker/databunker)|Secure SDK/vault for personal records/PII built to comply with GDPR|927|45|3|2019-12-08T21:55:55Z|2021-12-10T13:32:43Z|
[eliasdb](https://github.com/krotik/eliasdb)|EliasDB a graph-based database.|860|45|12|2016-08-13T13:53:28Z|2021-05-22T10:11:10Z|
[moss](https://github.com/couchbase/moss)|moss - a simple, fast, ordered, persistable, key-val storage library for golang|843|54|46|2016-02-06T20:27:22Z|2021-12-09T07:40:08Z|
[pogreb](https://github.com/akrylysov/pogreb)|Embedded key-value store for read-heavy workloads written in Go|833|60|10|2018-01-06T23:16:36Z|2021-08-27T13:45:37Z|
[column](https://github.com/kelindar/column)|High-performance, columnar, in-memory store with bitmap indexing in Go|832|25|5|2021-05-26T21:27:45Z|2021-12-07T21:19:14Z|
[levigo](https://github.com/jmhodges/levigo)|levigo is a Go wrapper for LevelDB|402|83|4|2012-01-17T08:17:54Z|2020-06-09T15:18:05Z|
[pudge](https://github.com/recoilme/pudge)|Fast and simple key/value store written using Go&#39;s standard library|301|22|0|2018-11-20T10:11:53Z|2021-07-04T02:08:38Z|
[vasto](https://github.com/chrislusf/vasto)|A distributed key-value store. On Disk. Able to grow or shrink without service interruption.|225|26|4|2018-01-16T05:16:57Z|2019-03-07T20:29:11Z|
[kivik](https://github.com/go-kivik/kivik)|Kivik provides a common interface to CouchDB or CouchDB-like databases for Go and GopherJS.|216|31|13|2017-02-09T14:14:54Z|2021-06-06T21:05:30Z|
[piladb](https://github.com/fern4lvarez/piladb)|Lightweight RESTful database engine based on stack data structures|192|20|9|2015-09-08T23:12:22Z|2020-10-29T19:19:06Z|
[golang-scribble](https://github.com/nanobox-io/golang-scribble)|A tiny Golang JSON database|141|18|1|2018-06-21T22:13:33Z|2019-03-09T22:57:36Z|
**[ARCHIVED]**  [slowpoke](https://github.com/recoilme/slowpoke)|Low-level key/value store in pure Go. |98|8|0|2018-02-19T09:22:37Z|2019-09-30T09:10:54Z|
[cache](https://github.com/akyoto/cache)|:handbag: Cache arbitrary data with an expiration time.|96|10|0|2019-05-11T12:42:45Z|2020-02-26T05:54:10Z|
[unitdb](https://github.com/unit-io/unitdb)|Fast specialized time-series database for IoT, real-time internet connected devices and AI analytics.|79|10|0|2019-08-29T18:21:27Z|2021-10-28T10:30:09Z|
[bcache](https://github.com/iwanbk/bcache)|Eventually consistent distributed in-memory  cache Go library|74|11|3|2018-12-26T15:45:16Z|2019-05-01T02:01:34Z|
[couchcache](https://github.com/codingsince1985/couchcache)|A RESTful caching micro-service in Go backed by Couchbase|53|5|0|2015-04-05T07:13:05Z|2021-10-02T02:59:37Z|
[hare](https://github.com/jameycribbs/hare)|Hare is a nimble little database management system for Go.|46|5|1|2016-10-05T20:05:45Z|2021-02-25T00:05:34Z|
[clusteredBigCache](https://github.com/oaStuff/clusteredBigCache)|golang bigcache with clustering as a library.|39|4|2|2017-12-18T07:48:07Z|2018-01-22T22:02:54Z|
[coffer](https://github.com/claygod/coffer)|Simply ACID* key-value database. At the medium or even low latency it tries to provide greater throughput without losing the ACID properties of the database. The database provides the ability to create record headers at own discretion and use them as transactions. The maximum size of stored data is limited by the size of the computer&#39;s RAM.|28|2|0|2019-05-13T18:30:23Z|2021-09-28T18:42:32Z|
[tempdb](https://github.com/rafaeljesus/tempdb)|Key-value store for temporary items :memo:|15|2|0|2017-03-17T18:03:42Z|2018-02-14T19:03:13Z|
[ttlcache](https://github.com/cheshir/ttlcache)|Simple in-memory key-value storage with TTL for each record.|3|2|0|2021-01-06T19:24:26Z|2021-03-21T22:19:47Z|
[Bitcask](https://git.mills.io/prologic/bitcask)|Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM&#43;WAL).|-|-|-|-|-|


## Database Drivers
*Libraries for connecting and operating databases.*
	

### Relational Databases



	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[mysql](https://github.com/go-sql-driver/mysql)|Go MySQL Driver is a MySQL driver for Go&#39;s (golang) database/sql package|11714|2043|99|2012-12-09T20:33:55Z|2021-11-09T11:43:03Z|
[pq](https://github.com/lib/pq)|Pure Go Postgres driver for database/sql|6936|832|284|2012-03-12T18:50:22Z|2021-11-26T21:29:09Z|
[go-sqlite3](https://github.com/mattn/go-sqlite3)|sqlite3 driver for go using database/sql|5289|888|93|2011-11-11T12:36:50Z|2021-11-03T07:59:44Z|
[pgx](https://github.com/jackc/pgx)|PostgreSQL driver and toolkit for Go|4817|483|192|2013-03-30T19:06:26Z|2021-12-11T20:55:32Z|
[go-mssqldb](https://github.com/denisenkom/go-mssqldb)|Microsoft SQL server driver written in go language|1484|388|136|2013-12-16T00:10:47Z|2021-12-11T19:28:05Z|
[go-oci8](https://github.com/mattn/go-oci8)|Oracle driver for Go using database/sql|574|206|8|2012-02-29T12:19:16Z|2021-10-25T19:04:43Z|
[sqlhooks](https://github.com/qustavo/sqlhooks)|Attach hooks to any database/sql driver|496|34|7|2016-04-20T18:37:14Z|2021-10-27T07:23:34Z|
[godror](https://github.com/godror/godror)|GO DRiver for ORacle DB|307|57|3|2019-11-21T21:23:17Z|2021-12-04T10:07:21Z|
[firebirdsql](https://github.com/nakagami/firebirdsql)|Firebird RDBMS sql driver for Go (golang)|160|47|13|2013-08-27T13:09:14Z|2021-10-29T06:46:42Z|
[go-adodb](https://github.com/mattn/go-adodb)|Microsoft ActiveX Object DataBase driver for go that using exp/sql|124|31|17|2011-11-14T04:32:50Z|2020-02-19T21:30:52Z|
[gofreetds](https://github.com/minus5/gofreetds)|Go Sql Server database driver.|105|44|18|2012-12-06T17:29:26Z|2020-11-30T22:32:55Z|
[sqinn-go](https://github.com/cvilsmeier/sqinn-go)|SQLite with pure Go|88|9|1|2020-06-06T20:37:12Z|2021-05-27T18:57:09Z|
[calcite-avatica-go](https://github.com/apache/calcite-avatica-go)|Mirror of Apache Calcite - Avatica Go SQL Driver|85|23|1|2017-08-08T07:00:08Z|2020-10-01T10:02:16Z|
[bgc](https://github.com/viant/bgc)|Datastore Connectivity for BigQuery in go|15|5|0|2016-06-13T20:24:26Z|2020-02-13T15:00:33Z|
[pig](https://github.com/alexeyco/pig)|Simple pgx wrapper to execute and scan query results|4|1|0|2021-04-15T15:33:23Z|2021-04-18T16:51:29Z|


### NoSQL Databases



	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[redis](https://github.com/go-redis/redis)|Type-safe Redis client for Golang|13071|1643|120|2012-07-25T13:01:39Z|2021-12-04T15:35:56Z|
[redigo](https://github.com/gomodule/redigo)|Go client for Redis|8773|1218|25|2012-04-14T04:31:58Z|2021-12-03T14:49:38Z|
[mongo-go-driver](https://github.com/mongodb/mongo-go-driver)|The Go driver for MongoDB|6315|726|6|2017-02-08T17:18:02Z|2021-12-09T15:50:21Z|
[mgo](https://github.com/globalsign/mgo)|The MongoDB driver for Go|1924|235|65|2017-04-13T11:14:04Z|2021-10-29T16:04:56Z|
[rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go)|Go language driver for RethinkDB|1580|180|16|2013-09-12T13:56:27Z|2021-09-26T22:12:29Z|
[gomemcache](https://github.com/bradfitz/gomemcache)|Go Memcached client library #golang|1416|397|52|2011-06-28T19:29:12Z|2021-11-05T16:08:02Z|
[qmgo](https://github.com/qiniu/qmgo)|Qmgo - The Go driver for MongoDB. It‘s based on official mongo-go-driver but easier to use like Mgo.|746|87|26|2020-08-04T09:06:00Z|2021-11-29T02:04:08Z|
[mgm](https://github.com/Kamva/mgm)|Mongo Go Models (mgm) is a fast and simple MongoDB ODM for Go (based on official Mongo Go Driver)|421|40|4|2019-12-27T14:40:51Z|2021-12-08T03:47:21Z|
[redeo](https://github.com/bsm/redeo)|High-performance framework for building redis-protocol compatible TCP servers/services|407|31|3|2014-03-06T08:46:18Z|2020-12-11T15:22:37Z|
[neoism](https://github.com/jmcvetta/neoism)|Neo4j client for Golang|381|58|16|2012-07-12T07:42:33Z|2020-02-16T09:28:03Z|
[aerospike-client-go](https://github.com/aerospike/aerospike-client-go)|Aerospike Client Go |370|172|18|2014-07-26T02:56:21Z|2021-12-06T12:50:20Z|
[gocb](https://github.com/couchbase/gocb)|The Couchbase Go SDK|329|95|0|2015-01-15T20:01:32Z|2021-12-09T18:02:46Z|
[go-couchbase](https://github.com/couchbase/go-couchbase)|Couchbase client in Go|311|90|41|2012-01-19T22:52:08Z|2021-11-08T23:15:41Z|
[go-rejson](https://github.com/nitishm/go-rejson)|Golang client for redislabs&#39; ReJSON module with support for multilple redis clients (redigo, go-redis)|197|30|10|2018-04-23T00:51:05Z|2021-04-05T13:41:10Z|
[godis](https://github.com/piaohao/godis)|redis client implement by golang, inspired by jedis.|98|16|0|2019-06-14T03:14:22Z|2020-05-12T07:08:10Z|
[Neo4j-GO](https://github.com/davemeehan/Neo4j-GO)|Neo4j REST Client in golang|76|16|0|2011-06-04T16:08:35Z|2018-06-20T12:15:38Z|
[arangolite](https://github.com/solher/arangolite)|Lightweight Golang driver for ArangoDB|69|19|5|2015-10-04T17:27:34Z|2021-03-10T17:27:51Z|
[go-pilosa](https://github.com/pilosa/go-pilosa)|Go client library for Pilosa|50|22|13|2016-09-30T21:37:10Z|2020-03-08T19:32:12Z|
[goforestdb](https://github.com/couchbase/goforestdb)|Go bindings for ForestDB|32|5|7|2014-05-14T15:36:12Z|2016-12-15T22:01:01Z|
[neo4j](https://github.com/cihangir/neo4j)|Neo4j Rest API Client for Go lang|26|8|8|2013-05-18T08:54:01Z|2015-04-02T17:38:48Z|
[goriak](https://github.com/zegl/goriak)|goriak - Go language driver for Riak KV|25|6|4|2016-10-05T16:48:17Z|2021-09-15T17:43:18Z|
[xredis](https://github.com/shomali11/xredis)|Go Redis Client|15|3|0|2017-06-14T00:19:26Z|2019-06-08T14:36:42Z|
[godscache](https://github.com/defcronyke/godscache)|An unofficial Google Cloud Platform Go Datastore wrapper that adds caching using memcached. For App Engine Flexible, Compute Engine, Kubernetes Engine, and more.|9|2|0|2018-05-08T20:19:39Z|2019-02-08T07:04:54Z|
[asc](https://github.com/viant/asc)|Datastore Connectivity for Aerospike for go|6|2|0|2016-06-13T20:22:31Z|2019-04-20T03:34:22Z|
[gocosmos](https://github.com/btnguyen2k/gocosmos)|Go driver for Azure CosmosDB SQL API|5|3|0|2020-12-06T07:03:43Z|2021-07-14T07:35:10Z|
[gocql](https://gocql.github.io)|Go language driver for Apache Cassandra.|-|-|-|-|-|


### Search and Analytic Databases.



	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[bleve](https://github.com/blevesearch/bleve)|A modern text indexing library for go|8021|609|264|2014-04-17T21:02:18Z|2021-12-10T17:34:37Z|
[elastic](https://github.com/olivere/elastic)|Elasticsearch client for Go.|6415|1054|72|2012-12-06T17:15:33Z|2021-12-01T08:49:58Z|
**[ARCHIVED]**  [riot](https://github.com/go-ego/riot)|Go Open Source, Distributed, Simple and efficient Search Engine; Warning: This is V1 and beta version, because of big memory consume, and the V2 will be rewrite all code.|6024|467|50|2017-06-21T14:17:59Z|2020-10-13T13:31:05Z|
[go-elasticsearch](https://github.com/elastic/go-elasticsearch)|The official Go client for Elasticsearch|3805|423|65|2017-03-27T17:56:15Z|2021-12-07T19:09:42Z|
[elastigo](https://github.com/mattbaird/elastigo)|A Go (golang) based Elasticsearch client library.|946|256|72|2012-10-12T04:19:59Z|2019-02-05T18:17:02Z|
[elasticsql](https://github.com/cch123/elasticsql)|convert sql to elasticsearch DSL in golang(go)|834|153|8|2016-08-24T07:29:43Z|2021-11-02T09:43:07Z|
[skizze](https://github.com/seiflotfy/skizze)|A probabilistic data structure service and storage|80|9|0|2016-01-17T12:10:40Z|2016-05-09T18:15:30Z|
[goes](https://github.com/OwnLocal/goes)|A library to interact with Elasticsearch in Go!|24|14|0|2015-12-28T18:52:03Z|2020-10-19T19:31:25Z|


### Multiple Backends.



	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[cayley](https://github.com/cayleygraph/cayley)|An open-source graph database|14001|1264|92|2014-06-05T18:49:41Z|2021-11-19T13:24:51Z|
[gokv](https://github.com/philippgille/gokv)|Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more)|380|43|32|2018-10-08T18:55:22Z|2021-09-15T07:51:51Z|
[cachego](https://github.com/faabiosr/cachego)|Golang Cache component - Multiple drivers|159|10|1|2016-10-05T18:10:03Z|2021-09-20T18:32:00Z|
[dsc](https://github.com/viant/dsc)|Datastore Connectivity in go|23|7|0|2016-06-13T20:18:10Z|2020-03-20T20:01:40Z|


## Date and Time
*Libraries for working with dates and times.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[now](https://github.com/jinzhu/now)|Now is a time toolkit for golang|3495|208|4|2013-11-18T10:55:30Z|2021-12-03T06:20:55Z|
[dateparse](https://github.com/araddon/dateparse)|GoLang Parse many date strings without knowing format in advance.|1578|117|50|2014-04-21T02:55:48Z|2021-11-11T01:29:06Z|
[carbon](https://github.com/golang-module/carbon)|A simple, semantic and developer-friendly golang package for datetime|1516|91|4|2020-09-07T09:07:35Z|2021-11-17T04:21:41Z|
[carbon](https://github.com/uniplaces/carbon)|Carbon for Golang, an extension for Time|653|51|11|2016-08-03T10:55:52Z|2021-11-04T16:28:39Z|
[durafmt](https://github.com/hako/durafmt)|:clock8:  Better time duration formatting in Go! |408|42|6|2016-05-20T21:49:33Z|2021-06-08T08:57:54Z|
[timeutil](https://github.com/leekchan/timeutil)|timeutil - useful extensions (Timedelta, Strftime, ...) to the golang&#39;s time package|184|13|1|2015-08-02T01:32:06Z|2019-02-03T13:14:43Z|
[gostradamus](https://github.com/bykof/gostradamus)|Gostradamus: Better DateTimes for Go 🕰️|159|3|0|2020-04-07T12:29:21Z|2021-11-21T18:24:23Z|
[go-persian-calendar](https://github.com/yaa110/go-persian-calendar)|The implementation of Persian (Solar Hijri) Calendar in Go|111|15|5|2016-01-31T18:40:23Z|2021-11-20T18:46:10Z|
[iso8601](https://github.com/relvacode/iso8601)|A fast ISO8601 date parser for Go|96|6|1|2017-04-25T15:54:18Z|2021-07-05T09:18:45Z|
[date](https://github.com/rickb777/date)|A Go package for working with dates|83|19|7|2015-11-23T22:58:07Z|2021-09-11T08:34:18Z|
[timespan](https://github.com/SaidinWoT/timespan)|Golang package to manipulate time intervals.|75|10|3|2014-10-07T03:57:02Z|2019-03-19T18:38:15Z|
[feiertage](https://github.com/wlbr/feiertage)|Gesetzliche Feiertage und mehr in Deutschland und Österreich (Bank holidays/public holidays in Austria and Germany)|40|5|1|2015-11-04T14:19:27Z|2021-08-16T20:16:45Z|
[go-sunrise](https://github.com/nathan-osman/go-sunrise)|Go package for calculating the sunrise and sunset times for a given location|38|7|0|2017-06-15T20:49:41Z|2021-06-07T17:58:34Z|
[go-str2duration](https://github.com/xhit/go-str2duration)|Convert string to duration in golang|33|2|1|2020-02-02T06:04:07Z|2020-08-11T00:48:43Z|
[kair](https://github.com/GuilhermeCaruso/kair)|:clock1: Date and Time - Golang Formatting Library|21|6|0|2018-10-03T15:44:07Z|2020-06-18T03:06:36Z|
[cronrange](https://github.com/1set/cronrange)|time range expression in cron style|13|5|3|2019-11-10T01:30:45Z|2021-08-12T00:03:26Z|
[nulltime](https://github.com/kirillDanshin/nulltime)||11|3|0|2016-03-06T01:44:58Z|2017-03-22T04:30:28Z|
[tuesday](https://github.com/osteele/tuesday)|Ruby-compatible strftime for golang|9|2|1|2017-08-10T20:46:26Z|2021-06-19T03:38:18Z|
[go-week](https://github.com/stoewer/go-week)|A Go package to work with ISO 8601 week dates|7|5|2|2018-02-23T07:02:37Z|2021-11-15T17:56:19Z|
[strftime](https://github.com/awoodbeck/strftime)|C99-compatible strftime formatter for use with Go time.Time instances.|7|3|0|2018-02-10T00:35:46Z|2018-02-21T15:59:14Z|


## Distributed Systems
*Packages that help with building Distributed Systems.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[etcd](https://github.com/etcd-io/etcd)|Distributed reliable key-value store for the most critical data of a distributed system|38156|8185|192|2013-07-06T21:57:21Z|2021-12-10T08:21:35Z|
[kit](https://github.com/go-kit/kit)|A standard library for microservices.|21879|2219|46|2015-02-03T00:01:19Z|2021-11-23T18:14:48Z|
[go-micro](https://github.com/asim/go-micro)|A pluggable Go framework for distributed systems development|17309|1944|56|2015-01-13T23:30:18Z|2021-12-10T11:32:21Z|
[kratos](https://github.com/go-kratos/kratos)|A modular-designed and easy-to-use microservices framework in Go.|15743|3138|38|2019-01-10T10:42:31Z|2021-12-11T07:30:05Z|
[grpc-go](https://github.com/grpc/grpc-go)|The Go language implementation of gRPC. HTTP/2 based RPC|15008|3315|129|2014-12-08T18:59:34Z|2021-12-10T19:22:25Z|
[go-zero](https://github.com/zeromicro/go-zero)|go-zero is a web and rpc framework written in Go. It&#39;s born to ensure the stability of the busy sites with resilient design. Builtin goctl greatly improves the development productivity.|13045|1565|36|2020-08-07T15:37:57Z|2021-12-11T10:51:04Z|
[micro](https://github.com/micro/micro)|Micro is a distributed cloud operating system|10751|952|83|2015-01-16T22:35:14Z|2021-12-09T10:06:09Z|
[nats-server](https://github.com/nats-io/nats-server)|High-Performance server for NATS.io, the cloud and edge native messaging system.|10194|997|187|2012-10-29T16:12:24Z|2021-12-11T18:10:51Z|
[rpcx](https://github.com/smallnest/rpcx)|Best microservices framework in Go, like alibaba Dubbo, but with more features, Scale easily. Try it. Test it. If you feel it&#39;s better, use it! 𝐉𝐚𝐯𝐚有𝐝𝐮𝐛𝐛𝐨, 𝐆𝐨𝐥𝐚𝐧𝐠有𝐫𝐩𝐜𝐱!|6487|984|14|2016-05-18T09:34:05Z|2021-12-10T23:48:22Z|
[raft](https://github.com/hashicorp/raft)|Golang implementation of the Raft consensus protocol|5389|721|17|2013-11-05T00:41:20Z|2021-12-07T18:52:18Z|
[lura](https://github.com/luraproject/lura)|Ultra performant API Gateway with middlewares. A project hosted at The Linux Foundation|4732|457|58|2016-11-04T18:37:13Z|2021-11-05T19:29:18Z|
[torrent](https://github.com/anacrolix/torrent)|Full-featured BitTorrent client package and utilities|4121|509|81|2015-01-08T21:10:42Z|2021-12-11T19:44:42Z|
[dragonboat](https://github.com/lni/dragonboat)|A feature complete and high performance multi-group Raft library in Go.  |3975|413|28|2018-12-23T07:02:04Z|2021-11-18T06:04:09Z|
[emitter](https://github.com/emitter-io/emitter)|High performance, distributed and low latency publish-subscribe platform.|3121|290|13|2016-10-29T08:52:21Z|2021-07-03T14:54:46Z|
[glow](https://github.com/chrislusf/glow)|Glow is an easy-to-use distributed computation system written in Go, similar to Hadoop Map Reduce, Spark, Flink, Storm, etc. I am also working on another similar pure Go system, https://github.com/chrislusf/gleam , which is more flexible and more performant.|3015|235|11|2015-06-14T00:33:48Z|2018-11-02T06:09:14Z|
[gleam](https://github.com/chrislusf/gleam)|Fast, efficient, and scalable distributed map/reduce system, DAG execution, in memory or on disk, written in pure Go, runs standalone or distributedly.|2939|264|36|2016-08-26T08:44:48Z|2021-05-13T22:17:25Z|
[liftbridge](https://github.com/liftbridge-io/liftbridge)|Lightweight, fault-tolerant message streams.|2173|94|37|2017-10-13T19:50:26Z|2021-12-07T21:08:43Z|
[hprose-golang](https://github.com/hprose/hprose-golang)|Hprose is a cross-language RPC. This project is Hprose for Golang.|1189|207|21|2014-02-14T03:16:43Z|2021-12-03T07:42:28Z|
[ringpop-go](https://github.com/uber/ringpop-go)|Scalable, fault-tolerant application-layer sharding for Go applications|711|63|25|2015-06-05T22:48:53Z|2021-02-23T00:14:24Z|
[rain](https://github.com/cenkalti/rain)|🌧 BitTorrent client and library in Go|658|43|5|2014-05-21T09:17:24Z|2021-12-10T19:14:48Z|
[gorpc](https://github.com/valyala/gorpc)|Simple, fast and scalable golang rpc library for high load|648|96|15|2014-11-20T17:02:37Z|2019-09-11T11:57:02Z|
[go-health](https://github.com/InVisionApp/go-health)|Library for enabling asynchronous health checks in your service|609|39|9|2017-11-29T21:00:07Z|2020-01-12T09:34:32Z|
[redislock](https://github.com/bsm/redislock)|Simplified distributed locking implementation using Redis|553|82|1|2019-06-24T11:10:10Z|2021-10-05T10:16:30Z|
[resgate](https://github.com/resgateio/resgate)|A Realtime API Gateway used with NATS to build REST, real time, and RPC APIs, where all your clients are synchronized seamlessly.|541|48|7|2018-02-22T12:06:26Z|2021-11-17T08:12:55Z|
[go-sundheit](https://github.com/AppsFlyer/go-sundheit)|A library built to provide support for defining service health for golang services. It allows you to register async health checks for your dependencies and the service itself, provides a health endpoint that exposes their status, and health metrics.|441|23|3|2019-04-08T12:54:01Z|2021-08-15T13:51:05Z|
[consistent](https://github.com/buraksezer/consistent)|Consistent hashing with bounded loads in Golang|431|53|5|2018-03-25T15:38:27Z|2021-06-06T10:45:01Z|
[arpc](https://github.com/lesismal/arpc)|More effective network communication, two-way calling, notify and broadcast supported.|427|37|2|2020-05-19T11:30:05Z|2021-12-11T04:45:08Z|
[digota](https://github.com/digota/digota)|ecommerce microservice|419|66|10|2017-08-14T12:01:37Z|2021-02-14T21:42:48Z|
[sleuth](https://github.com/ursiform/sleuth)|A Go library for master-less peer-to-peer autodiscovery and RPC between HTTP services|338|23|0|2016-04-23T14:21:41Z|2018-03-21T15:59:30Z|
[go-jump](https://github.com/dgryski/go-jump)|go-jump: Jump consistent hashing|335|28|1|2014-06-15T22:12:04Z|2021-10-18T20:05:52Z|
[dht](https://github.com/anacrolix/dht)|dht is used by anacrolix/torrent, and is intended for use as a library in other projects both torrent related and otherwise|214|51|2|2016-12-14T00:34:42Z|2021-12-10T04:30:24Z|
[jsonrpc](https://github.com/ybbus/jsonrpc)|A simple go implementation of json rpc 2.0 client over http|196|68|4|2016-11-10T11:27:55Z|2021-09-08T14:19:53Z|
[jsonrpc](https://github.com/osamingo/jsonrpc)|The jsonrpc package helps implement of JSON-RPC 2.0|157|17|4|2016-10-28T13:36:59Z|2021-10-15T12:47:14Z|
[go-doudou](https://github.com/unionj-cloud/go-doudou)|go-doudou（doudou pronounce /dəudəu/）is a gossip protocol and OpenAPI 3.0 spec based decentralized microservice framework. It supports monolith service application as well. Currently, it supports restful service only.|105|11|0|2021-02-24T07:21:40Z|2021-12-09T05:44:08Z|
[dynamolock](https://github.com/cirello-io/dynamolock)|DynamoDB Lock Client for Go|75|36|2|2018-07-08T11:13:00Z|2021-12-10T17:42:52Z|
[celeriac.v1](https://github.com/svcavallar/celeriac.v1)|Golang client library for adding support for interacting and monitoring Celery workers, tasks and events.|68|9|1|2015-10-10T07:27:33Z|2020-10-16T04:43:47Z|
[doublejump](https://github.com/edwingeng/doublejump)|A revamped Google&#39;s jump consistent hash|67|13|0|2018-06-26T16:04:50Z|2021-07-24T02:05:09Z|
[semaphore](https://github.com/jexia/semaphore)|Take control of your data, connect with anything, and expose it anywhere through protocols such as HTTP, GraphQL, and gRPC.|64|15|9|2020-02-05T16:39:39Z|2021-06-04T08:27:06Z|
[dot](https://github.com/dotchain/dot)|distributed data sync with operational transformation/transforms |61|2|0|2017-12-18T01:08:12Z|2019-09-30T00:29:15Z|
[outboxer](https://github.com/italolelis/outboxer)|A library that implements the outboxer pattern in go|60|9|11|2019-02-01T09:50:13Z|2021-12-06T11:11:01Z|
[flowgraph](https://github.com/vectaport/flowgraph)|Flowgraph package for scalable asynchronous system development|41|5|0|2018-08-29T21:45:26Z|2021-04-24T16:09:30Z|
[pglock](https://github.com/cirello-io/pglock)|PostgreSQL Lock Client for Go|37|7|0|2018-12-17T17:43:41Z|2021-11-17T15:45:47Z|
[drmaa](https://github.com/dgruber/drmaa)|Compute cluster (HPC) job submission library for Go (#golang) based on the open DRMAA standard.|36|17|0|2013-03-17T12:58:02Z|2020-10-06T06:39:46Z|
[go-mysql-lock](https://github.com/sanketplus/go-mysql-lock)|MySQL Backed Locking Primitive|35|6|3|2020-06-06T16:30:07Z|2021-07-25T17:36:16Z|
[go-pdu](https://github.com/pdupub/go-pdu)|Parallel Digital Universe - A decentralized identity-based social network|31|4|0|2018-10-08T08:13:22Z|2021-10-29T10:27:15Z|
[micro](https://github.com/gmsec/micro)|A Go distributed systems development framework|14|6|0|2020-05-03T01:16:16Z|2021-10-22T11:49:10Z|
[dynatomic](https://github.com/tylfin/dynatomic)|Dynatomic is a library for using dynamodb as an atomic counter|13|2|0|2019-02-08T17:45:14Z|2020-11-04T16:28:08Z|
[consistenthash](https://github.com/mbrostami/consistenthash)|A Go library that implements Consistent Hashing|9|2|0|2020-04-22T16:01:25Z|2020-05-02T00:32:37Z|
[failured](https://github.com/andy2046/failured)|Adaptive Accrual Failure Detector|2|0|0|2021-07-26T10:11:01Z|2021-08-02T03:08:02Z|
[pjrpc](https://gitlab.com/pjrpc/pjrpc)|Golang JSON-RPC Server-Client with Protobuf spec.|-|-|-|-|-|


## Dynamic DNS
*Tools for updating dynamic DNS records.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[godns](https://github.com/TimothyYe/godns)|A dynamic DNS client tool supports AliDNS, Cloudflare, Google Domains, DNSPod, HE.net &amp; DuckDNS &amp; DreamHost, etc, written in Go.|917|163|9|2014-05-11T11:49:17Z|2021-11-30T11:35:01Z|
[ddns](https://github.com/skibish/ddns)|Personal DDNS client with Digital Ocean Networking DNS as backend.|198|19|0|2017-03-13T21:02:27Z|2021-10-10T10:23:19Z|
[dyndns](https://gitlab.com/alcastle/dyndns)|Background Go process to regularly and automatically check your IP Address and make updates to (one or many) Dynamic DNS records for Google domains whenever your address changes.|-|-|-|-|-|


## Email
*Libraries and tools that implement email creation and sending.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[MailHog](https://github.com/mailhog/MailHog)|Web and API based SMTP testing|9375|721|194|2014-04-16T22:28:49Z|2021-12-08T13:14:14Z|
[hermes](https://github.com/matcornic/hermes)|Golang package that generates clean, responsive HTML e-mails for sending transactional mail|2377|195|29|2017-03-25T18:25:36Z|2021-12-05T01:25:36Z|
[email](https://github.com/jordan-wright/email)|Robust and flexible email library for Go|1881|254|49|2013-12-12T20:11:59Z|2021-08-16T11:43:35Z|
[go-imap](https://github.com/emersion/go-imap)| :inbox_tray: An IMAP library for clients and servers|1408|194|65|2016-04-26T17:59:18Z|2021-10-18T17:56:08Z|
[sendgrid-go](https://github.com/sendgrid/sendgrid-go)|The Official Twilio SendGrid Led, Community Driven Golang API Library|778|238|15|2013-09-12T03:31:13Z|2021-12-02T18:18:18Z|
[mailgun-go](https://github.com/mailgun/mailgun-go)|Go library for sending mail with the Mailgun API.|566|123|1|2014-02-28T00:28:44Z|2021-12-10T12:11:34Z|
[chasquid](https://github.com/albertito/chasquid)|SMTP (email) server with a focus on simplicity, security, and ease of operation [mirror]|491|34|4|2016-11-03T01:28:05Z|2021-12-06T17:56:47Z|
[email-verifier](https://github.com/AfterShip/email-verifier)|:white_check_mark: A Go library for email verification without sending any emails.|339|39|0|2020-12-18T08:47:28Z|2021-11-29T02:44:29Z|
[go-simple-mail](https://github.com/xhit/go-simple-mail)|Golang package for send email. Support keep alive connection, TLS and SSL. Easy for bulk SMTP.|231|41|8|2019-09-15T05:38:54Z|2021-12-10T14:01:10Z|
[go-message](https://github.com/emersion/go-message)|:envelope: A streaming Go library for the Internet Message Format and mail messages|230|70|21|2016-12-31T09:31:52Z|2021-11-27T10:45:44Z|
[hectane](https://github.com/hectane/hectane)|Lightweight SMTP client written in Go|211|26|16|2015-08-28T01:36:47Z|2020-11-29T20:53:17Z|
[douceur](https://github.com/aymerick/douceur)|A simple CSS parser and inliner in Go|203|35|9|2015-04-09T10:21:26Z|2021-06-05T19:55:34Z|
[mailchain](https://github.com/mailchain/mailchain)|Using Mailchain, blockchain users can now send and receive rich-media HTML messages with attachments via a blockchain address.|95|44|45|2019-04-11T17:37:31Z|2021-12-11T17:47:09Z|
[go-premailer](https://github.com/vanng822/go-premailer)|Inline styling for html mail in golang|77|13|3|2015-02-16T22:19:18Z|2021-03-06T20:26:39Z|
[go-dkim](https://github.com/toorop/go-dkim)|DKIM package for golang|73|31|4|2015-04-29T15:38:27Z|2020-11-03T13:16:31Z|
[smtp](https://github.com/mailhog/smtp)|MailHog SMTP Protocol|66|26|6|2014-12-24T16:13:49Z|2021-10-20T15:16:17Z|
[go-email-validator](https://github.com/go-email-validator/go-email-validator)|📧 Golang Email address validator|17|5|1|2020-12-10T18:27:20Z|2021-12-10T19:07:56Z|
[go-smtp-mock](https://github.com/mocktools/go-smtp-mock)|SMTP mock server written on Golang. Mimic any 📤 SMTP server behaviour for your test environment with fake SMTP server.|16|0|1|2021-08-31T13:54:57Z|2021-12-11T18:17:52Z|


## Embeddable Scripting Languages
*Embedding other languages inside your go code.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[gopher-lua](https://github.com/yuin/gopher-lua)|GopherLua: VM and compiler for Lua in Go|4443|497|96|2015-02-15T13:23:37Z|2021-11-23T09:57:15Z|
[goja](https://github.com/dop251/goja)|ECMAScript/JavaScript engine in pure Go|2537|212|19|2016-11-04T22:04:06Z|2021-12-11T11:25:15Z|
[tengo](https://github.com/d5/tengo)|A fast script language for Go|2518|154|47|2019-01-09T07:17:17Z|2021-11-13T23:15:04Z|
[go-lua](https://github.com/Shopify/go-lua)|A Lua VM in Go|2210|160|40|2013-12-20T17:29:43Z|2021-11-24T14:52:28Z|
[expr](https://github.com/antonmedv/expr)|Expression language for Go|2106|175|37|2018-07-14T15:57:34Z|2021-12-10T02:07:18Z|
[go-python](https://github.com/sbinet/go-python)|naive go bindings to the CPython2 C-API|1330|130|26|2012-07-09T15:43:31Z|2021-04-14T08:55:37Z|
[anko](https://github.com/mattn/anko)|Scriptable interpreter written in golang|1190|115|19|2014-03-28T07:29:40Z|2021-12-04T14:28:09Z|
[cel-go](https://github.com/google/cel-go)|Fast, portable, non-Turing complete expression evaluation with gradual typing (Go)|985|109|29|2018-03-09T22:57:58Z|2021-12-11T01:43:22Z|
[go-php](https://github.com/deuill/go-php)|PHP bindings for the Go programming language (Golang)|825|96|21|2015-09-17T21:23:52Z|2021-11-28T08:15:10Z|
[core](https://github.com/metacall/core)|MetaCall: The ultimate polyglot programming experience.|816|72|39|2018-12-26T22:02:57Z|2021-12-10T09:27:56Z|
**[ARCHIVED]**  [go-duktape](https://github.com/olebedev/go-duktape)|[abandoned] Duktape JavaScript engine bindings for Go|779|92|8|2015-01-08T05:09:05Z|2021-10-14T11:38:32Z|
[golua](https://github.com/aarzilli/golua)|Go bindings for Lua C API - in progress|571|162|4|2010-12-06T21:39:53Z|2021-11-19T15:09:33Z|
[gisp](https://github.com/jcla1/gisp)|Simple LISP in Go|474|36|1|2014-01-11T14:05:43Z|2017-08-25T13:48:45Z|
[gval](https://github.com/PaesslerAG/gval)|Expression evaluation in golang|411|53|5|2017-09-27T08:32:49Z|2021-12-05T13:55:08Z|
[gentee](https://github.com/gentee/gentee)|Gentee - script programming language for automation. It uses VM and compiler written in Go (Golang).|84|9|0|2018-01-14T15:49:05Z|2021-12-11T09:31:45Z|
[binder](https://github.com/alexeyco/binder)|High level go to Lua binder. Write less, do more.|54|8|0|2017-04-02T17:14:52Z|2018-07-29T22:00:27Z|
[purl](https://github.com/ian-kent/purl)|Perl, but fluffy like a cat!|34|3|2|2014-11-29T19:06:01Z|2014-12-07T17:45:34Z|
[ngaro](https://github.com/db47h/ngaro)|An embeddable implementation of the Ngaro Virtual Machine for Go programs|20|2|1|2016-08-09T15:23:50Z|2018-06-03T10:57:43Z|
[ecal](https://github.com/krotik/ecal)|A simple embeddable scripting language which supports concurrent event processing.|16|3|0|2020-11-30T15:58:56Z|2021-05-23T09:52:36Z|


## Error Handling
*Libraries for handling errors.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
**[ARCHIVED]**  [errors](https://github.com/pkg/errors)|Simple error handling primitives|7516|580|42|2015-12-27T12:05:38Z|2021-11-02T20:32:11Z|
[go-multierror](https://github.com/hashicorp/go-multierror)|A Go (golang) package for representing a list of errors as a single error.|1444|94|17|2014-12-15T20:12:26Z|2021-11-12T21:01:59Z|
[eris](https://github.com/rotisserie/eris)|eris provides a better way to handle, trace, and log errors in Go 🎆|887|22|1|2019-09-07T16:50:33Z|2021-07-30T14:40:52Z|
[errorx](https://github.com/joomcode/errorx)|A comprehensive error handling library for Go|788|27|4|2018-08-17T08:02:10Z|2021-04-07T12:03:57Z|
[tracerr](https://github.com/ztrue/tracerr)|Golang errors with stack trace and source fragments.|696|28|1|2019-02-06T18:57:46Z|2019-03-15T03:57:28Z|
[errlog](https://github.com/snwfdhmp/errlog)|Reduce debugging time while programming Go. Use static and stack-trace analysis to determine which func call causes the error.|399|17|0|2019-02-16T23:19:05Z|2020-11-30T18:28:01Z|
[emperror](https://github.com/emperror/emperror)|The Emperor takes care of all errors personally|235|16|5|2017-06-13T00:24:28Z|2020-10-04T16:48:36Z|
[errors](https://github.com/emperror/errors)|Drop-in replacement for the standard library errors package and github.com/pkg/errors|104|6|6|2019-07-09T13:02:52Z|2021-03-21T00:12:22Z|
[errors](https://github.com/bnkamalesh/errors)|A drop-in replacement for Go errors, with some added sugar! Unwrap user-friendly messages, HTTP status code, easy wrapping with multiple error types.|24|4|0|2020-07-17T18:57:04Z|2021-06-02T04:36:15Z|
[falcon](https://github.com/SonicRoshan/falcon)|A Simple Yet Highly Powerful Package For Error Handling|7|1|0|2019-09-09T12:49:43Z|2019-10-10T09:59:47Z|
[errors](https://github.com/PumpkinSeed/errors)|Simple and efficient error package |3|1|0|2020-01-08T21:12:51Z|2020-01-09T21:13:15Z|
[errors](https://github.com/neuronlabs/errors)|Simple golang error handling with classification primitives.|3|1|0|2019-07-26T00:15:36Z|2019-08-02T15:28:00Z|


## File Handling
*Libraries for handling files and file systems.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[afero](https://github.com/spf13/afero)|A FileSystem Abstraction System for Go|4101|372|90|2014-10-28T14:19:05Z|2021-11-14T16:16:41Z|
[pdfcpu](https://github.com/pdfcpu/pdfcpu)|A PDF processor written in Go.|2832|217|39|2017-06-18T17:27:38Z|2021-12-11T19:44:58Z|
[gdu](https://github.com/dundee/gdu)|Fast disk usage analyzer with console interface written in Go|1447|61|5|2018-02-24T15:04:23Z|2021-12-11T10:57:05Z|
[notify](https://github.com/rjeczalik/notify)|File system event notification library on steroids.|701|106|40|2014-09-08T16:09:34Z|2021-08-09T11:31:54Z|
[copy](https://github.com/otiai10/copy)|Go copy directory recursively|361|79|10|2017-09-01T03:18:56Z|2021-12-08T08:53:51Z|
[bigfile](https://github.com/bigfile/bigfile)|Bigfile -- a file transfer system that supports http, rpc and ftp protocol   https://bigfile.site  |212|40|1|2019-07-15T10:43:50Z|2020-02-26T01:29:46Z|
[afs](https://github.com/viant/afs)|Abstract File Storage|163|15|0|2019-08-19T18:43:38Z|2021-10-27T13:48:41Z|
[vfs](https://github.com/C2FO/vfs)|Pluggable, extensible virtual file system for Go|141|12|10|2017-08-01T18:06:14Z|2021-12-06T19:54:52Z|
[go-csv-tag](https://github.com/artonge/go-csv-tag)|Read csv file from go using tags|95|21|1|2017-06-18T15:31:16Z|2021-11-14T17:04:52Z|
[go-exiftool](https://github.com/barasher/go-exiftool)|Golang wrapper for Exiftool : extract as much metadata as possible (EXIF, ...) from files (pictures, pdf, office documents, ...)|91|22|3|2019-05-12T20:34:09Z|2021-08-18T21:08:18Z|
[opc](https://github.com/qmuntal/opc)|Go implementation of the Open Packaging Conventions (OPC)|67|5|0|2018-11-06T14:49:06Z|2021-03-01T20:00:33Z|
[skywalker](https://github.com/dixonwille/skywalker)|A package to allow one to concurrently go through a filesystem with ease|65|6|1|2017-08-01T20:08:25Z|2021-08-31T17:22:09Z|
[tarfs](https://github.com/posener/tarfs)|An implementation of the FileSystem interface for tar files.|48|7|1|2017-03-10T22:13:11Z|2020-03-13T18:47:56Z|
[checksum](https://github.com/codingsince1985/checksum)|Compute message digest for large files in Go|47|13|0|2014-11-05T09:37:00Z|2021-11-29T08:44:34Z|
[parquet](https://github.com/parsyl/parquet)|A library for reading and writing parquet files.|41|6|0|2019-01-29T21:52:30Z|2021-10-10T12:39:19Z|
[baraka](https://github.com/xis/baraka)|a tool for handling file uploads simple|33|6|1|2020-07-12T21:56:50Z|2021-08-29T21:03:18Z|
[go-gtfs](https://github.com/artonge/go-gtfs)|Load GTFS files in golang|29|16|0|2017-07-09T09:30:31Z|2020-10-08T14:23:27Z|
[flop](https://github.com/homedepot/flop)|Go file operations library chasing GNU APIs.|29|7|0|2019-03-01T13:41:39Z|2021-12-07T15:59:35Z|
[gut](https://github.com/1set/gut)|🍱 yet another collection of go utilities &amp; tools|19|5|13|2019-10-05T23:47:24Z|2020-11-17T17:52:05Z|
[go-decent-copy](https://github.com/hugocarreira/go-decent-copy)|copy files for humans|14|7|1|2018-10-16T07:08:24Z|2020-01-03T16:44:55Z|
[todotxt](https://github.com/1set/todotxt)|Parser for todo.txt files in Go ✅|9|1|0|2020-11-06T17:41:59Z|2020-11-30T16:17:27Z|
[pathtype](https://github.com/jonchun/pathtype)|Add a type for paths in Go.|6|2|0|2021-08-03T09:59:44Z|2021-08-12T15:10:37Z|
[higgs](https://github.com/dastoori/higgs)|A tiny cross-platform Go library to hide/unhide files and directories|6|2|2|2020-12-13T18:33:10Z|2021-04-16T00:26:37Z|
[stl](https://gitlab.com/russoj88/stl)|Modules to read and write STL (stereolithography) files.  Concurrent algorithm for reading.|-|-|-|-|-|


## Financial
*Packages for accounting and finance.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[ticker](https://github.com/achannarasappa/ticker)|Terminal stock ticker with live updates and position tracking|3970|210|18|2021-01-24T03:50:46Z|2021-11-27T15:26:30Z|
[decimal](https://github.com/shopspring/decimal)|Arbitrary-precision fixed-point decimal numbers in go|3571|418|74|2015-02-25T20:12:57Z|2021-12-06T03:38:19Z|
[go-money](https://github.com/Rhymond/go-money)|Go implementation of Fowler&#39;s Money pattern|1023|94|23|2017-03-20T16:23:54Z|2021-11-29T05:23:41Z|
[accounting](https://github.com/leekchan/accounting)|money and currency formatting for golang|693|56|8|2015-08-10T13:23:56Z|2021-10-06T19:11:59Z|
[techan](https://github.com/sdcoffey/techan)|Technical Analysis Library for Golang|565|97|19|2017-03-08T03:04:08Z|2021-11-17T16:09:20Z|
[go-finance](https://github.com/FlashBoys/go-finance)|:warning: Deprecrated in favor of https://github.com/piquette/finance-go |536|52|4|2016-02-28T00:37:46Z|2018-03-09T02:50:46Z|
[ach](https://github.com/moov-io/ach)|ACH implements a reader, writer, and validator for Automated Clearing House (ACH) files. The HTTP server is available in a Docker image and the Go package is available.|273|80|23|2016-12-14T21:12:49Z|2021-12-10T14:59:24Z|
[currency](https://github.com/bojanz/currency)|Currency handling for Go.|268|11|3|2020-04-16T15:34:39Z|2021-11-25T23:17:16Z|
[orderbook](https://github.com/i25959341/orderbook)|Matching Engine for Limit Order Book in Golang|228|84|5|2018-04-24T18:05:26Z|2021-05-16T21:28:00Z|
[go-finance](https://github.com/alpeb/go-finance)|Go library containing a collection of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.|118|18|1|2017-06-01T15:58:33Z|2021-12-02T20:16:28Z|
[transaction](https://github.com/claygod/transaction)|Embedded database for accounts transactions.|97|15|0|2017-10-11T13:50:30Z|2021-07-05T20:43:15Z|
[ofxgo](https://github.com/aclindsa/ofxgo)|Golang library for querying and parsing OFX|96|23|0|2015-11-08T13:56:53Z|2021-10-18T01:58:17Z|
[vat](https://github.com/dannyvankooten/vat)|Go package for dealing with EU VAT. Does VAT number validation &amp; rates retrieval.|85|12|4|2016-06-18T16:10:09Z|2021-02-05T12:45:06Z|
[sleet](https://github.com/BoltApp/sleet)|Payment abstraction library - one interface for multiple payment processors ( inspired by Ruby&#39;s ActiveMerchant )|76|10|4|2019-11-13T21:56:58Z|2021-12-09T01:45:48Z|
**[ARCHIVED]**  [go-finnhub](https://github.com/m1/go-finnhub)|Simple and easy to use client for stock market, forex and crypto data from finnhub.io written in Go. Access real-time financial market data from 60&#43; stock exchanges, 10 forex brokers, and 15&#43; crypto exchanges|64|14|0|2020-01-13T20:47:13Z|2020-02-01T14:53:23Z|
[currency](https://github.com/bnkamalesh/currency)|A currency computations package.|44|6|0|2017-05-09T06:06:38Z|2021-11-13T17:10:30Z|
[fastme](https://github.com/newity/fastme)||26|6|0|2020-10-29T13:57:10Z|2021-09-20T15:24:53Z|
[payme](https://github.com/jovandeginste/payme)|QR code generator (ASCII &amp; PNG) for SEPA payments|8|0|0|2021-05-03T21:56:06Z|2021-12-02T08:35:13Z|
[go-finance](https://github.com/pieterclaerhout/go-finance)|Finance related Go functions (e.g. exchange rates, VAT number checking, …)|5|3|0|2019-09-30T06:49:07Z|2019-10-23T13:05:23Z|


## Forms
*Libraries for working with forms.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[nosurf](https://github.com/justinas/nosurf)|CSRF protection middleware for Go.|1231|104|9|2013-08-22T17:47:34Z|2020-10-22T21:11:02Z|
[binding](https://github.com/mholt/binding)|Reflectionless data binding for Go&#39;s net/http (not actively maintained)|784|80|8|2014-05-20T23:35:14Z|2018-03-28T23:47:34Z|
[csrf](https://github.com/gorilla/csrf)|gorilla/csrf provides Cross Site Request Forgery (CSRF) prevention middleware for Go web applications &amp; services 🔒|724|107|5|2015-08-03T00:35:16Z|2021-10-26T19:04:41Z|
[form](https://github.com/go-playground/form)|:steam_locomotive: Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.|503|30|9|2016-05-26T13:26:40Z|2021-07-08T05:00:48Z|
[conform](https://github.com/leebenson/conform)|Trims, sanitizes &amp; scrubs data based on struct tags (go, golang)|241|30|0|2016-01-05T18:00:06Z|2021-09-29T18:12:34Z|
[formam](https://github.com/monoculum/formam)|a package for decode form&#39;s values into struct in Go|163|18|2|2014-10-25T00:23:30Z|2021-10-03T00:24:15Z|
[forms](https://github.com/albrow/forms)|A lightweight go library for parsing form data or json from an http.Request.|124|17|2|2014-08-07T16:11:30Z|2017-07-02T12:22:45Z|
[qs](https://github.com/sonh/qs)|Go module for encoding structs into URL query parameters|60|1|0|2020-10-02T09:50:35Z|2021-06-09T20:01:53Z|
[bind](https://github.com/robfig/bind)||26|4|0|2014-08-06T00:13:10Z|2014-08-16T17:03:51Z|
[queryparam](https://github.com/TomWright/queryparam)|Go package to easily convert a URL&#39;s query parameters/values into usable struct values of the correct types.|10|6|0|2018-06-14T10:23:05Z|2020-09-23T15:23:11Z|


## Functional
*Packages to support functional programming in Go.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[go-underscore](https://github.com/tobyhede/go-underscore)| Helpfully Functional Go -  A useful collection of Go utilities. Designed for programmer happiness. |1220|67|4|2014-07-02T10:27:16Z|2019-02-14T21:27:45Z|
[fpGo](https://github.com/TeaEntityLab/fpGo)|Monad, Functional Programming features for Golang|209|13|0|2018-05-24T09:08:45Z|2021-08-21T10:02:07Z|
[fuego](https://github.com/seborama/fuego)|Functional Experiment in Golang|101|9|0|2018-11-05T22:24:09Z|2020-11-11T22:18:31Z|
[gofp](https://github.com/rbrahul/gofp)|A super simple Lodash like utility library with essential functions that empowers the development in Go|97|4|0|2021-02-19T00:01:39Z|2021-02-23T02:11:36Z|


## Game Development
*Awesome game development libraries.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[ebiten](https://github.com/hajimehoshi/ebiten)|A dead simple 2D game library for Go|5535|358|225|2013-06-16T15:13:01Z|2021-12-09T12:03:54Z|
[leaf](https://github.com/name5566/leaf)|A game server framework in Go (golang)|4245|1130|14|2014-08-04T12:40:08Z|2021-07-11T11:08:50Z|
[pixel](https://github.com/faiface/pixel)|A hand-crafted 2D game library in Go|3749|216|38|2016-11-19T11:15:34Z|2021-10-14T01:17:34Z|
[goworld](https://github.com/xiaonanln/goworld)|Scalable Distributed Game Server Engine with Hot Swapping in Golang|2001|376|18|2017-06-03T15:02:46Z|2021-06-21T13:23:15Z|
[nano](https://github.com/lonng/nano)|Lightweight, facility, high performance golang based game server framework|1882|308|19|2017-08-02T06:05:14Z|2021-07-05T02:45:14Z|
[engine](https://github.com/g3n/engine)|Go 3D Game Engine (http://g3n.rocks)|1745|160|43|2017-03-07T18:25:09Z|2021-12-09T12:59:59Z|
[go-sdl2](https://github.com/veandco/go-sdl2)|SDL2 binding for Go|1661|194|53|2013-06-05T18:30:03Z|2021-11-01T14:51:53Z|
[engo](https://github.com/EngoEngine/engo)|Engo is an open-source 2D game engine written in Go.|1449|120|53|2014-11-12T05:50:03Z|2021-11-30T18:57:10Z|
[pitaya](https://github.com/topfreegames/pitaya)|Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK.|1254|262|23|2018-03-19T19:40:36Z|2021-11-11T17:31:34Z|
[termloop](https://github.com/JoelOtter/termloop)|Terminal-based game engine for Go, built on top of Termbox|1247|74|5|2015-05-23T17:12:34Z|2021-08-06T17:39:44Z|
[gonet](https://github.com/xtaci/gonet)|A Game Server Skeleton in golang.|1163|301|0|2013-04-11T02:18:23Z|2017-05-12T07:31:41Z|
[oak](https://github.com/oakmound/oak)|A pure Go game engine|977|62|11|2017-07-15T16:24:27Z|2021-12-11T19:10:29Z|
[raylib-go](https://github.com/gen2brain/raylib-go)|Go bindings for raylib, a simple and easy-to-use library to enjoy videogames programming.|699|74|9|2017-01-27T08:31:45Z|2021-11-14T11:16:13Z|
[engine](https://github.com/azul3d/engine)|Azul3D - A 3D game engine written in Go!|517|48|83|2016-02-29T04:54:44Z|2021-10-24T04:33:05Z|
[go-astar](https://github.com/beefsack/go-astar)|Go implementation of the A* search algorithm|462|64|1|2014-05-28T02:00:03Z|2020-08-27T23:23:15Z|
[go3d](https://github.com/ungerik/go3d)|A performance oriented 2D/3D math package for Go|215|36|1|2011-06-27T13:02:26Z|2021-10-26T19:35:42Z|
[prototype](https://github.com/gonutz/prototype)|Simple 2D game prototyping framework.|65|5|1|2015-03-04T09:24:39Z|2021-12-10T17:39:44Z|
[tile](https://github.com/kelindar/tile)|Tile is a 2D grid engine, built with data and cache friendly ways, includes pathfinding and observers.|36|2|0|2020-08-19T13:23:18Z|2021-07-08T05:12:10Z|


## Generation and Generics
*Tools to enhance the language with features like generics via code generation.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[go-linq](https://github.com/ahmetb/go-linq)|.NET LINQ capabilities in Go|2778|194|7|2013-12-19T03:05:00Z|2021-11-25T10:42:52Z|
[jennifer](https://github.com/dave/jennifer)|Jennifer is a code generator for Go|2247|112|16|2016-12-04T20:57:38Z|2021-12-08T18:29:53Z|
[gen](https://github.com/clipperhouse/gen)|Type-driven code generation for Go|1336|84|32|2013-10-13T20:26:36Z|2020-01-10T22:44:15Z|
[goderive](https://github.com/awalterschulze/goderive)|Code Generation for Functional Programming, Concurrency and Generics in Golang|939|37|16|2017-02-10T21:46:49Z|2021-12-10T20:55:35Z|
[gowrap](https://github.com/hexdigest/gowrap)|GoWrap is a command line tool for generating decorators for Go interfaces|549|50|5|2018-09-15T09:20:42Z|2021-10-25T17:26:13Z|
[interfaces](https://github.com/rjeczalik/interfaces)|Code generation tools for Go.|311|20|5|2015-12-06T00:04:50Z|2021-04-27T07:31:41Z|
[go-enum](https://github.com/abice/go-enum)|An enum generator for go|256|29|0|2017-08-10T22:07:31Z|2021-12-06T02:20:32Z|
[pkgreflect](https://github.com/ungerik/pkgreflect)|A Go preprocessor for package scoped reflection|99|16|0|2012-09-03T07:53:00Z|2017-09-05T12:27:27Z|
[goverter](https://github.com/jmattheis/goverter)|Generate type-safe Go converters by simply defining an interface|92|2|3|2021-03-09T20:39:27Z|2021-11-05T16:55:53Z|
[efaceconv](https://github.com/t0pep0/efaceconv)||51|9|1|2016-11-18T11:38:54Z|2017-10-12T07:16:32Z|
[gotype](https://github.com/wzshiming/gotype)|Golang source code parsing, usage like reflect package|36|7|0|2017-12-05T04:09:47Z|2021-08-02T13:55:42Z|
[GENERIS](https://github.com/senselogic/GENERIS)|Versatile Go code generator.|29|1|0|2019-03-10T19:33:31Z|2021-08-07T07:09:40Z|
[go-xray](https://github.com/pieterclaerhout/go-xray)|Helpers for making the use of reflection easier|20|2|0|2019-10-01T08:40:51Z|2019-11-20T17:31:59Z|
[typeregistry](https://github.com/xiaoxin01/typeregistry)|create type dynamically in Golang|11|1|0|2020-01-14T15:50:38Z|2020-02-20T13:00:03Z|


## Geographic
*Geographic tools and servers*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[tile38](https://github.com/tidwall/tile38)|Real-time Geospatial and Geofencing|7876|485|117|2016-03-04T23:07:44Z|2021-12-10T01:19:47Z|
[geo](https://github.com/golang/geo)|S2 geometry library in Go|1304|149|11|2014-12-03T23:02:15Z|2021-11-19T07:57:49Z|
[mbtileserver](https://github.com/consbio/mbtileserver)|Basic Go server for mbtiles|289|54|12|2014-11-01T04:12:14Z|2021-12-10T02:35:52Z|
[osm](https://github.com/paulmach/osm)|General purpose library for reading, writing and working with OpenStreetMap data|181|31|2|2016-02-02T00:59:03Z|2021-09-05T17:28:37Z|
[wgs84](https://github.com/wroge/wgs84)|A pure Go package for coordinate transformations.|66|5|0|2019-06-08T17:17:59Z|2021-12-10T18:02:51Z|
[geoserver](https://github.com/hishamkaram/geoserver)|geoserver is a Go library for manipulating a GeoServer instance via the GeoServer REST API.|65|15|1|2018-03-26T21:36:49Z|2021-11-26T17:00:58Z|
[godal](https://github.com/airbusgeo/godal)|golang wrapper for github.com/OSGEO/gdal|55|8|4|2021-02-05T17:27:05Z|2021-12-10T16:45:57Z|
[gismanager](https://github.com/hishamkaram/gismanager)|Publish Your GIS Data(Vector Data) to PostGIS and Geoserver|39|7|0|2018-09-29T12:51:37Z|2018-10-30T08:55:19Z|
[simplefeatures](https://github.com/peterstace/simplefeatures)|Simple Features is a pure Go Implementation of the OpenGIS Simple Feature Access Specification|38|4|44|2019-06-07T07:52:01Z|2021-12-07T18:38:04Z|
[pbf](https://github.com/maguro/pbf)|OpenStreetMap PBF golang parser|28|5|1|2017-09-18T23:13:18Z|2021-04-16T22:36:07Z|
[s2-geojson](https://github.com/pantrif/s2-geojson)|Draw a polygon on the map or paste a geoJSON and explore how the s2.RegionCoverer covers it with S2 cells depending on the min and max levels|15|4|1|2020-03-27T09:47:32Z|2020-04-05T06:44:10Z|


## Go Compilers
*Tools for compiling Go to other languages.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[gopherjs](https://github.com/gopherjs/gopherjs)|A compiler from Go to JavaScript for running Go code in a browser|10701|499|205|2013-08-27T22:23:58Z|2021-11-21T16:54:38Z|
[tardisgo](https://github.com/tardisgo/tardisgo)|Golang-&gt;Haxe-&gt;CPP/CSharp/Java/JavaScript transpiler  |413|28|4|2014-01-08T11:07:33Z|2016-11-19T18:08:43Z|
[c4go](https://github.com/Konstantin8105/c4go)|Transpiling C code to Go code|291|35|23|2018-03-28T06:24:57Z|2021-11-15T11:17:02Z|
[esp32-transpiler](https://github.com/andygeiss/esp32-transpiler)|Transpile Golang into Arduino code to use fully automated testing at your IoT projects.|33|3|0|2018-03-14T14:22:55Z|2021-07-19T11:06:51Z|
[f4go](https://github.com/Konstantin8105/f4go)|Transpiling fortran code to golang code|31|8|5|2018-07-08T17:05:43Z|2021-11-30T13:42:22Z|


## Goroutines
*Tools for managing and working with Goroutines.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[ants](https://github.com/panjf2000/ants)|🐜🐜🐜 ants is a high-performance and low-cost goroutine pool in Go, inspired by fasthttp./ ants 是一个高性能且低损耗的 goroutine 池。|7121|875|20|2018-05-19T01:13:38Z|2021-12-02T13:28:49Z|
[tunny](https://github.com/Jeffail/tunny)|A goroutine pool for Go|2744|237|1|2014-04-02T16:14:58Z|2021-07-12T17:03:25Z|
[goworker](https://github.com/benmanns/goworker)|goworker is a Go-based background worker that runs 10 to 100,000* times faster than Ruby-based workers.|2614|241|32|2013-07-22T17:04:27Z|2021-12-09T16:32:27Z|
[workerpool](https://github.com/gammazero/workerpool)|Concurrency limiting goroutine pool|693|90|5|2016-05-17T14:32:06Z|2021-10-29T00:40:18Z|
[grpool](https://github.com/ivpusic/grpool)|Lightweight Goroutine pool|663|97|5|2015-07-22T00:15:04Z|2019-01-27T23:07:22Z|
[pool](https://github.com/go-playground/pool)|:speedboat: a limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation|663|59|4|2015-10-28T16:36:08Z|2021-06-28T13:01:34Z|
[gowp](https://github.com/xxjwxc/gowp)|golang worker pool , Concurrency limiting goroutine pool|346|55|1|2019-09-14T11:43:50Z|2021-05-20T11:30:11Z|
[pond](https://github.com/alitto/pond)|🔘 Minimalistic and High-performance goroutine worker pool written in Go|346|20|3|2020-03-21T14:56:33Z|2021-11-20T13:55:13Z|
[go-floc](https://github.com/workanator/go-floc)|Floc: Orchestrate goroutines with ease.|215|16|0|2017-07-03T07:34:06Z|2021-08-10T10:33:23Z|
[go-flow](https://github.com/kamildrazkiewicz/go-flow)|Simply way to control goroutines execution order based on dependencies|179|19|1|2016-09-25T14:46:09Z|2019-05-14T12:10:41Z|
[semaphore](https://github.com/marusama/semaphore)|Fast resizable golang semaphore primitive|130|7|0|2017-11-22T14:00:58Z|2021-03-28T09:27:47Z|
[go-workers](https://github.com/catmullet/go-workers)|👷 Library for safely running groups of workers concurrently or consecutively that require input and output through channels|125|8|2|2020-10-06T15:39:43Z|2021-09-24T10:05:25Z|
[artifex](https://github.com/mborders/artifex)|Simple in-memory job queue for Golang using worker-based dispatching|122|8|0|2018-10-31T19:34:31Z|2020-08-18T21:33:48Z|
[errgroup](https://github.com/neilotoole/errgroup)|errgroup with goroutine worker limits|102|7|5|2020-06-26T06:07:39Z|2021-09-19T04:32:47Z|
[async](https://github.com/StudioSol/async)|A safe way to execute functions asynchronously, recovering them in case of panic. It also provides an error stack aiming to facilitate fail causes discovery.|95|10|2|2017-06-30T17:08:33Z|2020-11-19T17:27:17Z|
[semaphore](https://github.com/kamilsk/semaphore)|🚦 Semaphore pattern implementation with timeout of lock/unlock operations.|89|11|6|2016-10-08T11:48:12Z|2020-04-16T19:25:15Z|
[cyclicbarrier](https://github.com/marusama/cyclicbarrier)|CyclicBarrier golang implementation|87|12|0|2018-01-11T10:38:46Z|2020-06-30T10:11:31Z|
[gpool](https://github.com/sherifabdlnaby/gpool)|gpool - a generic context-aware resizable goroutines pool to bound concurrency based on semaphore. |82|3|0|2018-12-03T04:23:35Z|2019-12-16T17:37:15Z|
[worker-pool](https://github.com/vardius/worker-pool)|Go simple async worker pool|80|12|0|2017-10-04T09:18:31Z|2021-01-17T02:27:13Z|
[gollback](https://github.com/vardius/gollback)|Go asynchronous simple function utilities, for managing execution of closures and callbacks|70|5|1|2019-05-11T05:56:37Z|2021-11-30T12:25:12Z|
[Hunch](https://github.com/AaronJan/Hunch)|Hunch provides functions like: All, First, Retry, Waterfall etc., that makes asynchronous flow control more intuitive.|67|5|0|2019-06-05T13:21:04Z|2020-10-13T14:56:47Z|
[threadpool](https://github.com/shettyh/threadpool)|Golang simple thread pool implementation|62|14|0|2017-09-06T18:45:39Z|2020-03-23T11:51:49Z|
[routine](https://github.com/x-mod/routine)|go routine control, abstraction of the Main and some useful Executors.如果你不会管理Goroutine的话，用它|45|6|0|2019-03-04T12:25:23Z|2020-10-08T05:51:14Z|
[nursery](https://github.com/arunsworld/nursery)|Structured Concurrency in Go|38|4|1|2019-11-23T19:26:02Z|2021-07-08T15:59:22Z|
[kyoo](https://github.com/dirkaholic/kyoo)|Unlimited job queue for go, using a pool of concurrent workers processing the job queue entries|36|1|0|2020-01-06T20:35:11Z|2020-03-29T16:11:58Z|
[goccm](https://github.com/zenthangplus/goccm)|Limits the number of goroutines that are allowed to run concurrently|33|4|2|2019-08-16T02:26:53Z|2021-10-05T16:37:09Z|
[async](https://github.com/reugn/async)|Synchronization and asynchronous computation utilities library for Go|33|1|0|2019-12-28T09:48:40Z|2021-11-23T14:24:25Z|
[parallel-fn](https://github.com/rafaeljesus/parallel-fn)|Run functions in parallel :comet:|31|1|0|2017-06-18T09:47:54Z|2018-01-01T20:34:49Z|
[oversight](https://github.com/cirello-io/oversight)|[Mirror] Erlang-like supervisor trees|25|3|0|2018-11-09T14:46:48Z|2021-12-01T16:14:09Z|
[go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup)|A sync.WaitGroup with error handling and concurrency control|24|1|0|2018-08-08T16:12:35Z|2020-02-21T09:12:59Z|
[go-trylock](https://github.com/subchen/go-trylock)|TryLock support on read-write lock for Golang|22|6|0|2018-04-26T06:02:47Z|2021-05-07T03:38:43Z|
[stl](https://github.com/ssgreg/stl)|Software Transactional Locks|20|3|0|2018-06-19T10:50:11Z|2020-07-24T08:20:52Z|
[gohive](https://github.com/loveleshsharma/gohive)|🐝 A Highly Performant and easy to use goroutine pool for Go|15|2|2|2019-05-31T10:44:24Z|2021-11-27T10:45:02Z|
[channelify](https://github.com/ddelizia/channelify)|Make functions return a channel for parallel processing via go routines.|15|1|1|2020-10-05T13:12:48Z|2021-02-25T17:33:41Z|
[gowl](https://github.com/hamed-yousefi/gowl)|Gowl is a process management and process monitoring tool at once. An infinite worker pool gives you the ability to control the pool and processes and monitor their status.|12|2|4|2021-04-12T19:15:53Z|2021-07-22T10:48:03Z|
[conexec](https://github.com/ITcathyh/conexec)|A concurrent toolkit to help execute funcs concurrently in an efficient and safe way. It supports specifying the overall timeout to avoid blocking.|11|1|0|2019-12-24T07:35:11Z|2020-06-28T03:09:55Z|
[queue](https://github.com/AnikHasibul/queue)|package queue gives you a queue group accessibility. Helps you to limit goroutines, wait for the end of the all goroutines and much more.|10|1|0|2018-12-21T07:15:00Z|2019-05-18T11:05:23Z|
[hands](https://github.com/duanckham/hands)|Hands is a process controller used to control the execution and return strategies of multiple goroutines.|7|2|1|2020-04-04T11:04:11Z|2020-04-16T02:34:07Z|
[execpool](https://github.com/hexdigest/execpool)|A pool that spins up a given number of processes in advance and attaches stdin and stdout when needed. Very similar to FastCGI but works for any command.|7|1|0|2021-06-17T18:41:46Z|2021-07-06T20:39:16Z|
[concurrency-limiter](https://github.com/vivek-ng/concurrency-limiter)||5|2|0|2020-11-22T02:35:52Z|2020-12-04T21:15:00Z|
[go-tools](https://github.com/nikhilsaraf/go-tools)|A collection of tools for Golang|5|2|0|2018-11-14T02:53:08Z|2019-03-27T19:18:09Z|
[breaker](https://github.com/kamilsk/breaker)|🚧 Flexible mechanism to make execution flow interruptible.|2|0|0|2021-07-11T10:35:18Z|2021-07-11T10:32:17Z|


## Hardware
*Libraries, tools, and tutorials for interacting with hardware.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[arduino-cli](https://github.com/arduino/arduino-cli)|Arduino command line tool|3166|282|239|2018-08-08T15:57:32Z|2021-12-10T16:30:10Z|
[go-rpio](https://github.com/stianeikeland/go-rpio)|:electric_plug: Raspberry Pi GPIO library for go-lang|1754|200|31|2013-07-30T03:58:33Z|2021-12-05T16:17:03Z|
[ghw](https://github.com/jaypipes/ghw)|Golang hardware discovery/inspection library|1082|121|25|2017-05-26T16:39:02Z|2021-12-10T22:37:07Z|
[emgo](https://github.com/ziutek/emgo)|Emgo: Bare metal Go (language for programming embedded systems)|945|64|13|2014-07-09T10:55:20Z|2021-12-05T21:00:21Z|
[sysinfo](https://github.com/zcalusic/sysinfo)|Sysinfo is a Go library providing Linux OS / kernel / hardware system information.|301|60|13|2016-08-22T01:46:45Z|2021-10-21T14:56:31Z|
[goroslib](https://github.com/aler9/goroslib)|ROS client library for the Go programming language|146|28|3|2020-01-19T20:02:35Z|2021-12-10T04:06:28Z|
[go-osc](https://github.com/hypebeast/go-osc)|Open Sound Control (OSC) library for Golang. Implemented in pure Go.|136|38|15|2013-08-26T14:10:42Z|2021-08-15T01:27:36Z|
[joystick](https://github.com/0xcafed00d/joystick)|Go Joystick API|29|11|1|2015-07-24T14:51:47Z|2020-02-14T23:53:45Z|


## Images
*Libraries for manipulating images.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[gocv](https://github.com/hybridgroup/gocv)|Go package for computer vision using OpenCV 4 and beyond.|4498|658|179|2017-09-18T21:54:17Z|2021-11-10T18:17:51Z|
[imaginary](https://github.com/h2non/imaginary)|Fast, simple, scalable, Docker-ready HTTP microservice for high-level image processing|4078|384|99|2015-03-04T18:51:40Z|2021-12-08T20:51:52Z|
[imaging](https://github.com/disintegration/imaging)|Imaging is a simple image processing package for Go|4015|336|17|2012-12-06T20:21:21Z|2020-12-18T19:30:12Z|
[bild](https://github.com/anthonynsimon/bild)|Image processing algorithms in pure Go|3418|187|13|2016-08-01T15:54:29Z|2021-11-15T13:57:57Z|
[gg](https://github.com/fogleman/gg)|Go Graphics - 2D rendering in Go with a simple API.|3153|234|65|2016-02-18T21:05:08Z|2021-11-16T20:02:59Z|
[ln](https://github.com/fogleman/ln)|3D line art engine.|3005|111|12|2016-01-10T04:28:10Z|2019-07-19T09:00:40Z|
[resize](https://github.com/nfnt/resize)|Pure golang image resizing |2762|271|10|2012-08-02T19:48:26Z|2020-11-20T20:05:09Z|
[pt](https://github.com/fogleman/pt)|A path tracer written in Go.|1977|113|8|2015-01-23T19:39:29Z|2019-03-21T10:07:26Z|
[svgo](https://github.com/ajstarks/svgo)|Go Language Library for SVG generation|1778|148|9|2010-03-05T23:24:10Z|2021-10-24T23:50:47Z|
[bimg](https://github.com/h2non/bimg)|Go package for fast high-level image processing powered by libvips C library|1721|284|126|2015-03-17T14:14:02Z|2021-12-09T16:12:00Z|
[picfit](https://github.com/thoas/picfit)|An image resizing server written in Go|1585|136|17|2014-12-06T17:30:45Z|2021-09-17T10:07:15Z|
[smartcrop](https://github.com/muesli/smartcrop)|smartcrop finds good image crops for arbitrary crop sizes|1566|105|7|2014-04-07T22:40:03Z|2021-05-19T04:35:11Z|
[gift](https://github.com/disintegration/gift)|Go Image Filtering Toolkit|1473|108|2|2014-07-12T18:47:40Z|2020-11-21T15:45:54Z|
[imagick](https://github.com/gographics/imagick)|Go binding to ImageMagick&#39;s MagickWand C API|1376|163|10|2013-04-30T17:31:48Z|2021-09-08T03:48:56Z|
[gowitness](https://github.com/sensepost/gowitness)|🔍 gowitness - a golang, web screenshot utility using Chrome Headless|1313|185|20|2017-10-31T08:36:35Z|2021-09-29T16:29:51Z|
[go-opencv](https://github.com/go-opencv/go-opencv)|Go bindings for OpenCV / 2.x API in gocv / 1.x API in opencv|1254|198|45|2013-12-09T09:43:26Z|2019-05-24T14:30:18Z|
[geopattern](https://github.com/pravj/geopattern)|:triangular_ruler: Create beautiful generative image patterns from a string in golang.|1149|63|3|2014-10-22T17:26:30Z|2019-01-08T20:17:57Z|
[stegify](https://github.com/DimitarPetrov/stegify)|🔍 Go tool for LSB steganography, capable of hiding any file within an image.|968|111|0|2018-11-29T16:45:58Z|2020-07-08T13:43:58Z|
[canvas](https://github.com/tdewolff/canvas)|Cairo in Go: vector to raster, SVG, PDF, EPS, WASM, OpenGL, Gio, etc.|850|52|13|2017-05-20T18:10:51Z|2021-12-09T03:16:45Z|
[image2ascii](https://github.com/qeesung/image2ascii)|:foggy: Convert image to ASCII|607|55|4|2018-10-20T05:06:25Z|2021-07-27T10:56:28Z|
[govips](https://github.com/davidbyttow/govips)|A lightning fast image processing and resizing library for Go|576|128|20|2016-12-25T04:32:56Z|2021-12-11T11:27:23Z|
[draft](https://github.com/lucasepe/draft)|Generate High Level Cloud Architecture diagrams using YAML syntax.|517|23|0|2020-06-05T16:11:40Z|2021-09-08T18:02:56Z|
[govatar](https://github.com/o1egl/govatar)|Avatar generation library for GO language|469|29|1|2016-01-18T12:12:28Z|2021-03-14T12:22:11Z|
[goimagehash](https://github.com/corona10/goimagehash)|Go Perceptual image hashing package|446|52|10|2017-07-28T17:15:58Z|2021-02-18T21:10:22Z|
[mort](https://github.com/aldor007/mort)|Storage and image processing server written in Go|440|18|8|2017-11-19T13:37:58Z|2021-11-17T08:20:42Z|
[go-nude](https://github.com/koyachi/go-nude)|Nudity detection with Go.|336|38|2|2014-05-02T08:32:29Z|2018-11-22T15:22:42Z|
[rez](https://github.com/bamiaux/rez)|Image resizing in pure Go and SIMD|202|18|1|2014-01-16T21:16:15Z|2017-07-31T18:51:31Z|
[darkroom](https://github.com/gojek/darkroom)||177|33|7|2019-07-01T10:17:08Z|2021-06-29T15:19:25Z|
[mergi](https://github.com/noelyahan/mergi)|go library for image programming (merge, crop, resize, watermark, animate, ease, transit)|156|23|2|2018-09-24T03:40:47Z|2020-05-29T19:49:07Z|
[gltf](https://github.com/qmuntal/gltf)|:eyeglasses: Go library for [d]encoding glTF 2.0 files|139|23|4|2019-01-15T17:43:54Z|2021-10-27T13:23:53Z|
[img](https://github.com/hawx/img)|A selection of image manipulation tools|138|9|1|2012-07-28T19:57:47Z|2015-05-01T15:11:26Z|
[steganography](https://github.com/auyer/steganography)|Pure Golang Library that allows simple LSB steganography on images|121|20|0|2018-05-21T17:27:36Z|2021-07-29T15:48:34Z|
[go-cairo](https://github.com/ungerik/go-cairo)|Go binding for the cairo graphics library|116|28|0|2012-08-22T18:27:01Z|2021-03-17T13:39:46Z|
[cameron](https://github.com/aofei/cameron)|An avatar generator for Go.|78|8|1|2018-05-05T22:13:11Z|2021-11-13T09:43:00Z|
[go-gd](https://github.com/bolknote/go-gd)|Go bingings for GD (http://www.boutell.com/gd/)|52|16|1|2011-05-12T06:33:54Z|2018-05-07T19:29:26Z|
[gridder](https://github.com/shomali11/gridder)|A Grid based 2D Graphics library|47|6|0|2020-04-10T00:13:10Z|2021-09-30T17:31:42Z|
[go-webp](https://github.com/kolesa-team/go-webp)|Simple and fast webp library for golang|36|8|0|2020-02-18T09:53:07Z|2021-09-15T04:03:25Z|
[goimghdr](https://github.com/corona10/goimghdr)|The imghdr module determines the type of image contained in a file for go|36|3|0|2018-02-25T09:34:44Z|2019-06-14T10:13:28Z|
[webp-server](https://github.com/mehdipourfar/webp-server)|Simple and minimal image server capable of storing, resizing, converting and caching images.|30|6|0|2020-11-22T12:03:12Z|2021-01-14T20:14:09Z|
**[ARCHIVED]**  [tga](https://github.com/ftrvxmtrx/tga)|Go package for decoding and encoding TARGA image format|28|10|1|2012-10-08T01:09:20Z|2015-05-24T08:11:41Z|
[go-webcolors](https://github.com/jyotiska/go-webcolors)|Port of webcolors library from Python to Go|24|5|0|2014-04-24T14:41:22Z|2015-08-21T04:56:56Z|
[mpo](https://github.com/donatj/mpo)|JPEG-MPO Decoder / Converter Library and CLI Tool|7|3|1|2015-04-14T22:37:59Z|2020-06-18T16:55:56Z|
[scout](https://github.com/jonoton/scout)|Scout is a standalone open source software solution for DIY video security.|2|0|0|2020-09-25T17:28:58Z|2021-10-12T17:16:48Z|


## IoT (Internet of Things)
*Libraries for programming devices of the IoT.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[gobot](https://github.com/hybridgroup/gobot)|Golang framework for robotics, drones, and the Internet of Things (IoT)|7485|933|163|2013-09-21T14:09:19Z|2021-11-20T02:27:07Z|
[flogo](https://github.com/TIBCOSoftware/flogo)|Project Flogo is an open source ecosystem of opinionated  event-driven capabilities to simplify building efficient &amp; modern serverless functions, microservices &amp; edge apps.|1887|256|153|2016-07-10T02:57:43Z|2020-11-30T17:38:34Z|
[periph](https://github.com/google/periph)|Go·Hardware·Lean|1697|174|42|2016-10-13T16:53:51Z|2021-08-30T20:45:54Z|
[mainflux](https://github.com/mainflux/mainflux)|Industrial IoT Messaging and Device Management Platform|1626|476|104|2015-07-06T20:31:50Z|2021-12-10T12:12:34Z|
[gatt](https://github.com/paypal/gatt)|Gatt is a Go package for building Bluetooth Low Energy peripherals|1001|273|52|2014-04-23T13:45:27Z|2020-07-15T05:47:19Z|
[heedy](https://github.com/heedy/heedy)|An aggregator for personal metrics, and an extensible analysis engine|305|30|9|2015-01-16T19:44:21Z|2021-12-06T03:02:34Z|
[devices](https://github.com/goiot/devices)|Suite of libraries for IoT devices (written in Go), experimental for x/exp/io|247|25|9|2016-05-30T08:07:02Z|2016-07-10T00:46:08Z|
[sensorbee](https://github.com/sensorbee/sensorbee)|Lightweight stream processing engine for IoT|208|37|39|2016-02-19T07:49:56Z|2019-11-04T22:46:34Z|
[huego](https://github.com/amimof/huego)|An extensive Philips Hue client library for Go with an emphasis on simplicity|189|32|7|2017-05-16T05:31:45Z|2021-12-01T18:36:37Z|
[iot](https://github.com/vaelen/iot)|A Go client for Google IoT Core|54|8|0|2018-03-08T06:51:51Z|2019-11-08T18:32:28Z|
[eywa](https://github.com/xcodersun/eywa)|Make IoT a lot more fun with data. |50|14|9|2016-02-20T17:02:16Z|2017-04-12T07:41:51Z|


## Job Scheduler
*Libraries for scheduling jobs.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[gocron](https://github.com/go-co-op/gocron)|Easy and fluent Go cron scheduling. This is a fork from https://github.com/jasonlvhit/gocron|1415|112|14|2020-03-20T15:33:05Z|2021-12-11T21:00:08Z|
[jobrunner](https://github.com/bamzi/jobrunner)|Framework for performing work asynchronously, outside of the request flow|895|83|10|2015-10-21T04:17:01Z|2020-11-14T21:03:29Z|
[gron](https://github.com/roylee0704/gron)|gron, Cron Jobs in Go.|876|53|8|2016-06-04T08:02:22Z|2021-01-14T08:44:12Z|
[go-quartz](https://github.com/reugn/go-quartz)|Minimalist and zero-dependency scheduling library for Go|561|31|1|2019-04-14T18:57:51Z|2021-12-03T13:24:13Z|
[jobs](https://github.com/albrow/jobs)|A persistent and flexible background jobs library for go.|484|42|17|2015-02-09T22:13:29Z|2018-06-16T21:00:16Z|
[scheduler](https://github.com/carlescere/scheduler)|Job scheduling made easy.|371|52|6|2015-02-03T17:10:23Z|2020-12-27T08:18:20Z|
[go-cron](https://github.com/rk/go-cron)|A simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.|208|17|0|2011-04-15T14:50:49Z|2020-02-10T17:52:36Z|
[gronx](https://github.com/adhocore/gronx)|Lightweight, fast and dependency-free Cron expression parser (due checker), task scheduler and/or daemon for Golang (tested on v1.13 and above) and standalone usage|180|12|2|2021-04-21T06:14:03Z|2021-10-17T14:47:44Z|
[clockwerk](https://github.com/onatm/clockwerk)|Job Scheduling Library|113|11|0|2017-04-09T23:10:48Z|2019-11-08T07:51:19Z|
[leprechaun](https://github.com/kilgaloon/leprechaun)|You had one job, or more then one, which can be done in steps|84|12|13|2018-04-08T13:44:04Z|2021-11-15T12:40:00Z|
[tasks](https://github.com/madflojo/tasks)|Package tasks is an easy to use in-process scheduler for recurring tasks in Go|65|4|1|2019-12-24T18:26:18Z|2021-03-27T16:07:18Z|
[sched](https://github.com/romshark/sched)|A job scheduler for Go with the ability to fast-forward time.|20|0|0|2021-06-19T22:57:48Z|2021-07-09T14:15:46Z|
[cronticker](https://github.com/krayzpipes/cronticker)|Golang ticker that works with Cron scheduling.|1|2|0|2020-11-28T20:59:38Z|2021-01-02T01:57:05Z|


## JSON
*Libraries for working with JSON.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[gjson](https://github.com/tidwall/gjson)|Get JSON values quickly - JSON parser for Go|9411|631|34|2016-08-11T03:08:47Z|2021-12-10T16:23:12Z|
[json-to-go](https://github.com/mholt/json-to-go)|Translates JSON into a Go type in your browser instantly (original)|3383|406|14|2014-01-21T18:11:13Z|2021-11-22T22:44:50Z|
[gojson](https://github.com/ChimeraCoder/gojson)|Automatically generate Go (golang) struct definitions from example JSON|2438|188|41|2012-12-27T19:10:50Z|2021-07-30T03:02:50Z|
[fastjson](https://github.com/valyala/fastjson)|Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection|1396|82|33|2018-05-28T21:41:47Z|2021-07-20T15:54:48Z|
[ojg](https://github.com/ohler55/ojg)|Optimized JSON for Go|438|28|0|2020-04-12T17:17:31Z|2021-12-10T17:12:18Z|
[kazaam](https://github.com/qntfy/kazaam)|Arbitrary transformations of JSON in Golang|215|45|23|2016-07-19T14:19:03Z|2021-07-05T18:29:50Z|
[gojq](https://github.com/elgs/gojq)|JSON query in Golang|177|21|1|2015-12-30T09:02:13Z|2020-11-20T03:35:26Z|
[jsondiff](https://github.com/wI2L/jsondiff)|Compute the diff between two JSON documents as a series of RFC6902 (JSON Patch) operations|139|13|0|2020-11-28T19:05:16Z|2021-08-30T23:34:13Z|
[jettison](https://github.com/wI2L/jettison)|Highly configurable, fast JSON encoder for Go|113|8|0|2019-08-30T13:28:03Z|2021-11-02T10:13:35Z|
[jsongo](https://github.com/ricardolonga/jsongo)|Fluent API to make it easier to create Json objects.|98|15|2|2015-08-07T23:23:17Z|2021-10-04T03:26:13Z|
[gjo](https://github.com/skanehira/gjo)|Small utility to create JSON objects|97|12|1|2019-02-23T01:54:21Z|2021-04-18T16:48:02Z|
[json2go](https://github.com/m-zajac/json2go)|Create go type representation from json|93|13|1|2017-06-10T23:55:07Z|2021-05-03T18:35:18Z|
[ajson](https://github.com/spyzhov/ajson)|Abstract JSON for Golang with JSONPath support |88|10|10|2019-03-07T20:47:38Z|2021-11-21T23:22:49Z|
[jaydiff](https://github.com/yazgazan/jaydiff)|A JSON diff utility|83|7|2|2017-04-24T16:05:35Z|2021-01-27T19:43:07Z|
[json-to-proto.github.io](https://github.com/json-to-proto/json-to-proto.github.io)|convert JSON to Protocol Buffers online in your browser instantly|64|7|1|2020-04-18T20:42:45Z|2020-11-23T17:24:38Z|
[jsonf](https://github.com/miolini/jsonf)|Console JSON formatter with query feature|62|10|0|2015-05-25T04:53:32Z|2020-12-13T21:45:56Z|
[ujson](https://github.com/olvrng/ujson)|µjson - A fast and minimal JSON parser and transformer that works on unstructured JSON|52|6|0|2019-02-27T12:58:07Z|2021-08-06T04:09:15Z|
[go-respond](https://github.com/nicklaw5/go-respond)|A Go package for handling common HTTP JSON responses.|43|8|1|2017-03-12T21:00:54Z|2021-09-24T20:08:26Z|
[mp](https://github.com/sanbornm/mp)|Simple Email Parser|43|5|1|2014-06-15T21:14:39Z|2016-05-11T19:40:58Z|
[jsoncolor](https://github.com/neilotoole/jsoncolor)|Colorized JSON output for Go|25|4|2|2021-09-13T01:44:14Z|2021-10-05T15:20:09Z|
[vjson](https://github.com/miladibra10/vjson)|vjson is a golang package that helps to validate JSON objects|23|2|3|2021-04-29T16:47:50Z|2021-11-15T05:55:42Z|
[ask](https://github.com/simonnilsson/ask)|A Go package that provides a simple way of accessing nested properties in maps and slices.|13|0|0|2020-09-13T13:53:31Z|2021-02-19T18:47:59Z|
[go-jsonerror](https://github.com/ddymko/go-jsonerror)|Small package which wraps error responses to follow jsonapi.org|10|1|0|2018-10-18T15:03:45Z|2019-10-09T11:56:05Z|
[jsonhal](https://github.com/RichardKnop/jsonhal)|A simple Go package to make custom structs marshal into HAL compatible JSON responses.|9|5|1|2016-01-15T11:38:40Z|2020-03-24T12:17:52Z|
[dynjson](https://github.com/cocoonspace/dynjson)|Client-customizable JSON formats for dynamic APIs|9|4|0|2020-05-06T07:10:02Z|2021-10-11T15:25:37Z|
[jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors)|Go bindings based on the JSON API errors reference|9|3|0|2016-07-08T10:08:58Z|2016-11-17T16:02:12Z|
[mapslice-json](https://github.com/ake-persson/mapslice-json)|Go MapSlice for ordered marshal/ unmarshal of maps in JSON|8|3|0|2020-02-19T11:01:48Z|2021-07-20T08:19:13Z|
[epoch](https://github.com/vtopc/epoch)|Contains primitives for marshaling/unmarshaling Unix timestamp/epoch to/from built-in time.Time type in JSON|7|2|0|2019-12-15T12:54:37Z|2021-03-28T14:59:09Z|
[ej](https://github.com/lucassscaravelli/ej)|Write and read JSON from different sources in one line|6|1|0|2020-01-04T17:39:35Z|2020-04-07T00:36:07Z|
[jsonic](https://github.com/sinhashubham95/jsonic)|All you need with JSON|5|1|0|2021-01-09T06:21:59Z|2021-01-15T08:00:58Z|
[jzon](https://github.com/zerosnake0/jzon)|A golang json library inspired by jsoniter|5|1|0|2019-11-12T10:42:41Z|2021-03-22T11:24:48Z|
[omg.jsonparser](https://github.com/dedalqq/omg.jsonparser)|The simple JSON parser with validation by condition|2|1|0|2021-07-08T23:59:21Z|2021-10-12T12:34:19Z|


## Logging
*Libraries for generating and working with log files.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[logrus](https://github.com/sirupsen/logrus)|Structured, pluggable logging for Go.|19417|2056|74|2013-10-16T19:08:55Z|2021-11-24T12:21:27Z|
[zap](https://github.com/uber-go/zap)|Blazing fast, structured, leveled logging in Go.|14308|1058|89|2016-02-18T19:52:56Z|2021-12-08T21:17:56Z|
[zerolog](https://github.com/rs/zerolog)|Zero Allocation JSON Logger|5614|336|88|2017-05-12T05:24:39Z|2021-12-04T16:11:24Z|
[go-spew](https://github.com/davecgh/go-spew)|Implements a deep pretty printer for Go data structures to aid in debugging|4776|309|55|2013-01-09T05:18:22Z|2020-12-02T06:27:18Z|
[glog](https://github.com/golang/glog)|Leveled execution logs for Go|3062|825|1|2013-07-16T04:33:04Z|2021-08-20T20:52:44Z|
[lumberjack](https://github.com/natefinch/lumberjack)|lumberjack is a log rolling package for Go|2952|379|56|2014-06-14T11:55:47Z|2021-08-23T20:11:36Z|
[tail](https://github.com/hpcloud/tail)|Go package for reading from continously updated files (tail -f)|2232|451|71|2013-02-05T00:28:03Z|2021-07-17T01:27:24Z|
[seelog](https://github.com/cihub/seelog)|Seelog is a native Go logging library that provides flexible asynchronous dispatching, filtering, and formatting.|1567|247|39|2011-11-17T09:43:15Z|2019-03-04T07:03:16Z|
[log](https://github.com/apex/log)|Structured logging package for Go.|1189|105|38|2015-12-21T20:27:48Z|2021-11-11T18:28:36Z|
[log15](https://github.com/inconshreveable/log15)|Structured, composable logging for Go|1033|144|44|2014-05-20T00:11:52Z|2021-10-31T02:28:23Z|
[log](https://github.com/phuslu/log)|Structured Logging Made Easy|401|30|4|2019-07-07T09:40:38Z|2021-11-24T17:28:15Z|
[onelog](https://github.com/francoispqt/onelog)|Dead simple, super fast, zero allocation and modular logger for Golang|400|14|1|2018-05-06T14:32:10Z|2019-03-06T04:37:07Z|
[logxi](https://github.com/mgutz/logxi)|A 12-factor app logger built for performance and happy development|348|41|23|2015-03-01T22:13:45Z|2020-04-14T15:56:24Z|
[logutils](https://github.com/hashicorp/logutils)|Utilities for slightly better logging in Go (Golang).|300|32|2|2013-10-09T07:31:15Z|2021-11-08T05:38:47Z|
[log](https://github.com/go-playground/log)|:green_book: Simple, configurable and scalable Structured Logging for Go.|276|22|1|2016-02-07T16:17:48Z|2019-11-11T18:44:02Z|
[go-logger](https://github.com/apsdehal/go-logger)|Simple logger for Go programs. Allows custom formats for messages.|271|51|2|2014-09-26T04:57:06Z|2019-05-15T21:27:11Z|
[httpretty](https://github.com/henvic/httpretty)|Package httpretty prints the HTTP requests you make with Go pretty on your terminal.|244|7|1|2020-01-24T18:17:16Z|2020-12-16T21:42:18Z|
[sqldb-logger](https://github.com/simukti/sqldb-logger)|A logger for Go SQL database driver without modifying existing *sql.DB stdlib usage.|199|6|4|2019-11-02T17:28:03Z|2021-05-25T05:56:30Z|
[rollingwriter](https://github.com/arthurkiller/rollingwriter)|Rolling writer is an IO util for auto rolling write in go.|182|29|8|2017-02-12T12:05:26Z|2021-08-25T08:22:01Z|
[logger](https://github.com/azer/logger)|Minimalistic logging library for Go.|148|16|0|2014-09-30T06:45:09Z|2021-11-22T15:36:32Z|
[logur](https://github.com/logur/logur)|Logur is an opinionated collection of logging best practices|144|10|8|2018-12-09T16:43:11Z|2020-10-04T16:49:57Z|
[xlog](https://github.com/rs/xlog)|xlog is a logger for net/context aware HTTP applications|135|13|3|2015-10-22T09:26:45Z|2021-02-17T06:17:46Z|
[glg](https://github.com/kpango/glg)|Simple and blazing fast lockfree logging library for golang|131|12|0|2017-06-21T13:26:16Z|2021-07-24T05:29:05Z|
[ozzo-log](https://github.com/go-ozzo/ozzo-log)|A Go (golang) package providing high-performance asynchronous logging, message filtering by severity and category, and multiple message targets.|114|33|9|2015-10-22T22:29:02Z|2021-01-07T10:03:10Z|
[logvoyage](https://github.com/firstrow/logvoyage)|LogVoyage - logging SaaS written in GoLang|89|12|9|2015-03-29T11:05:09Z|2017-05-24T19:48:17Z|
[go-cronowriter](https://github.com/utahta/go-cronowriter)|Time based rotating file writer|46|8|3|2017-02-04T09:02:55Z|2021-03-16T17:25:35Z|
[log](https://github.com/alexcesaro/log)|Logging packages for Go|45|4|1|2014-04-19T14:31:56Z|2015-09-15T22:13:22Z|
**[ARCHIVED]**  [gologger](https://github.com/sadlil/gologger)|The Simplest and worst logging library ever written|39|10|2|2015-09-02T08:52:26Z|2018-01-31T03:17:58Z|
[go-log](https://github.com/ian-kent/go-log)|A logger, for Go|38|19|3|2014-05-02T00:34:09Z|2018-03-31T02:06:55Z|
[logex](https://github.com/chzyer/logex)|An golang log lib, supports tracking and level, wrap by standard log lib|38|10|2|2014-10-10T06:38:39Z|2017-03-29T06:49:08Z|
[gone](https://github.com/One-com/gone)|Golang packages for writing small daemons and servers.|36|6|0|2016-09-05T09:39:11Z|2021-05-24T14:23:37Z|
[noodlog](https://github.com/gyozatech/noodlog)|🍜 Parametrized JSON logging library in Golang which lets you obfuscate sensitive data and marshal any kind of content.|34|8|7|2021-04-09T08:38:54Z|2021-10-06T16:10:24Z|
[go-log](https://github.com/siddontang/go-log)|a golang log lib supports level and multi handlers|28|15|1|2014-05-18T03:41:55Z|2019-02-21T02:24:31Z|
[journald](https://github.com/ssgreg/journald)|Go implementation of systemd Journal&#39;s native API for logging|28|2|0|2017-08-23T07:06:09Z|2021-03-05T18:33:46Z|
[logrusly](https://github.com/sebest/logrusly)|Loggly Hooks for GO Logrus logger|27|18|3|2014-09-11T23:27:11Z|2021-07-27T21:32:29Z|
[distillog](https://github.com/amoghe/distillog)|Logging, distilled|25|7|0|2015-10-12T16:32:21Z|2018-07-26T23:35:13Z|
[mlog](https://github.com/jbrodriguez/mlog)|A simple logging module for go, with a rotating file feature and console logging.|24|19|1|2014-10-20T15:06:26Z|2018-08-05T17:35:46Z|
[log](https://github.com/teris-io/log)|Structured log interface|24|3|0|2017-10-28T19:57:55Z|2017-12-04T18:53:45Z|
[gomol](https://github.com/aphistic/gomol)|Gomol is a library for structured, multiple-output logging for Go with extensible logging outputs|17|1|3|2015-08-30T15:51:46Z|2019-03-14T03:15:36Z|
[glo](https://github.com/lajosbencz/glo)|Logging library for Golang|14|1|0|2019-01-19T22:10:42Z|2019-01-23T11:35:10Z|
[zkits-logger](https://github.com/edoger/zkits-logger)|A powerful zero-dependency json logger.|14|1|3|2020-03-31T14:23:40Z|2021-12-03T04:15:32Z|
[logrusiowriter](https://github.com/cabify/logrusiowriter)|io.Writer implementation using logrus logger|13|1|0|2019-08-09T08:58:25Z|2020-07-15T09:10:12Z|
[go-log](https://github.com/subchen/go-log)|Simple and configurable Logging in Go, with level, formatters and writers|11|6|0|2017-05-07T08:09:24Z|2018-05-19T08:03:37Z|
[logmatic](https://github.com/mborders/logmatic)|Colorized logger for Golang with dynamic log level configuration|11|4|1|2018-11-07T01:52:45Z|2021-01-11T03:10:50Z|
[logdump](https://github.com/ewwwwwqm/logdump)|Package for multi-level logging|9|3|0|2017-01-13T15:34:31Z|2018-04-02T00:28:16Z|
[logo](https://github.com/mbndr/logo)|Golang logger to different configurable writers.|9|2|0|2017-02-07T18:02:55Z|2020-12-27T10:33:21Z|
[log](https://github.com/aerogo/log)|:memo: Logging with multiple output targets.|9|1|0|2017-06-10T09:54:08Z|2019-10-26T04:19:45Z|
[go-log](https://github.com/pieterclaerhout/go-log)|A logging library with strack traces, object dumping and optional timestamps|8|4|0|2019-10-01T08:55:38Z|2020-07-08T07:39:26Z|
[xlog](https://github.com/xfxdev/xlog)|plugin architecture and flexible log system for golang|6|5|0|2016-05-05T16:47:45Z|2019-01-15T10:17:30Z|
[kemba](https://github.com/clok/kemba)|A tiny debug logging tool. Ideal for CLI tools and command applications. Inspired by https://github.com/visionmedia/debug|5|2|2|2020-07-13T03:10:54Z|2021-10-20T06:46:33Z|
[yell](https://github.com/jfcg/yell)|Yet another minimalist logging library|0|0|0|2021-02-07T16:07:27Z|2021-08-07T13:23:51Z|


## Machine Learning
*Libraries for Machine Learning.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[golearn](https://github.com/sjwhitworth/golearn)|Machine Learning for Go|8104|1127|68|2013-12-26T13:06:14Z|2021-10-14T19:37:59Z|
[gorse](https://github.com/zhenghaoz/gorse)|An open source recommender system service written in Go|5019|414|26|2018-08-14T11:01:09Z|2021-12-11T07:15:14Z|
[gorgonia](https://github.com/gorgonia/gorgonia)|Gorgonia is a library that helps facilitate machine learning in Go.|4270|363|82|2016-09-14T23:19:43Z|2021-11-15T23:38:40Z|
[m2cgen](https://github.com/BayesWitnesses/m2cgen)|Transform ML models into a native code (Java, C, Python, Go, JavaScript, Visual Basic, C#, R, PowerShell, PHP, Dart, Haskell, Ruby, F#, Rust) with zero dependencies|1953|168|37|2019-01-13T02:32:55Z|2021-11-25T12:36:04Z|
[tfgo](https://github.com/galeone/tfgo)|Tensorflow &#43; Go, the gopher way|1842|135|10|2017-05-23T13:27:39Z|2021-09-14T07:21:22Z|
[gosseract](https://github.com/otiai10/gosseract)|Go package for OCR (Optical Character Recognition), by using Tesseract C&#43;&#43; library|1620|199|14|2013-10-11T07:27:53Z|2021-12-02T10:16:37Z|
[goml](https://github.com/cdipaolo/goml)|On-line Machine Learning in Go (and so much more)|1271|117|7|2015-06-27T05:52:01Z|2021-10-30T12:24:02Z|
[eaopt](https://github.com/MaxHalford/eaopt)|:four_leaf_clover: Evolutionary optimization library for Go (genetic algorithm, partical swarm optimization, differential evolution)|759|86|7|2016-01-31T00:04:52Z|2021-04-05T09:12:42Z|
[bayesian](https://github.com/jbrukh/bayesian)|Naive Bayesian Classification for Golang.|717|121|8|2011-11-23T04:17:00Z|2020-07-24T17:41:07Z|
[CloudForest](https://github.com/ryanbressler/CloudForest)|Ensembles of decision trees in go/golang.|696|90|35|2012-10-22T17:38:16Z|2020-12-08T20:22:57Z|
[gobrain](https://github.com/goml/gobrain)|Neural Networks written in go|495|57|1|2014-04-29T13:32:36Z|2020-12-12T12:34:25Z|
[ocrserver](https://github.com/otiai10/ocrserver)|A simple OCR API server, seriously easy to be deployed by Docker, on Heroku as well|443|98|1|2015-11-15T07:57:42Z|2021-08-05T08:20:24Z|
[onnx-go](https://github.com/owulveryck/onnx-go)|onnx-go gives the ability to import a pre-trained neural network within Go without being linked to a framework or library.|357|38|20|2018-08-28T07:39:20Z|2021-10-26T15:36:06Z|
[go-deep](https://github.com/patrikeh/go-deep)|Artificial Neural Network|334|35|7|2017-12-09T15:10:06Z|2021-12-07T13:11:42Z|
[regommend](https://github.com/muesli/regommend)|Recommendation engine for Go|295|27|0|2014-02-05T17:00:49Z|2019-08-07T04:55:12Z|
[goptuna](https://github.com/c-bata/goptuna)|A hyperparameter optimization framework, inspired by Optuna.|194|13|11|2019-07-24T12:03:05Z|2021-11-18T00:34:10Z|
[go-galib](https://github.com/thoj/go-galib)|Genetic Algorithms library written in Go / golang|187|40|0|2009-11-30T10:46:58Z|2015-12-28T16:27:45Z|
[goRecommend](https://github.com/timkaye11/goRecommend)|Collaborative Filtering (CF) Algorithms in Go! |177|20|0|2014-07-16T05:32:23Z|2014-07-29T04:49:57Z|
[shield](https://github.com/eaigner/shield)|Bayesian text classifier with flexible tokenizers and storage backends for Go|144|30|5|2013-04-10T19:38:16Z|2020-03-04T03:41:47Z|
[goga](https://github.com/tomcraven/goga)|Golang Genetic Algorithm|119|12|0|2015-10-20T12:50:51Z|2017-01-16T15:29:16Z|
[go-fann](https://github.com/vksnk/go-fann)|Go bindings for FANN, library for artificial neural networks|104|20|2|2011-03-10T21:08:27Z|2015-02-03T21:53:31Z|
[goscore](https://github.com/asafschers/goscore)|Go Scoring API for PMML|71|21|3|2017-08-19T11:08:39Z|2019-08-23T11:21:08Z|
[gonet](https://github.com/dathoangnd/gonet)|Neural Network for Go.|70|6|0|2020-01-11T18:27:28Z|2020-04-05T16:08:18Z|
[libsvm](https://github.com/datastream/libsvm)|libsvm go version|68|13|1|2012-07-31T07:57:47Z|2016-05-09T03:47:11Z|
[go-featureprocessing](https://github.com/nikolaydubina/go-featureprocessing)|🔥 Fast, simple sklearn-like feature processing for Go|66|7|3|2020-12-18T13:09:18Z|2021-06-17T16:30:19Z|
[neural-go](https://github.com/schuyler/neural-go)|A multilayer perceptron network implemented in Go, with training via backpropagation.|61|15|1|2011-10-17T09:31:33Z|2020-08-31T10:58:21Z|
[go-pr](https://github.com/daviddengcn/go-pr)|Pattern recognition package in Go lang.|59|13|0|2013-06-07T02:36:20Z|2013-06-08T10:17:05Z|
**[ARCHIVED]**  [neat](https://github.com/jinyeom/neat)|NEAT (NeuroEvolution of Augmenting Topologies) implemented in Go|59|12|4|2016-11-17T04:23:14Z|2018-07-04T20:45:55Z|
[fonet](https://github.com/Fontinalis/fonet)|fonet is a deep neural network package for Go.|51|13|2|2017-10-03T15:57:15Z|2021-06-01T10:04:04Z|
[golinear](https://github.com/danieldk/golinear)|liblinear bindings for Go|42|11|0|2013-04-05T15:37:01Z|2018-08-29T10:30:44Z|
[Varis](https://github.com/Xamber/Varis)|Golang Neural Network |38|7|0|2017-10-10T08:43:27Z|2018-08-02T13:47:14Z|
[godist](https://github.com/e-dard/godist)|Probability distributions and associated methods in Go|31|6|0|2014-09-05T09:48:51Z|2015-05-11T10:38:48Z|
[go-cluster](https://github.com/e-XpertSolutions/go-cluster)|k-modes and k-prototypes clustering algorithms implementation in Go|29|7|0|2017-10-04T12:24:52Z|2018-08-06T07:35:27Z|
[gomind](https://github.com/surenderthakran/gomind)|A simplistic Neural Network Library in Go|19|3|6|2017-10-19T03:48:51Z|2018-07-31T12:57:31Z|
[evoli](https://github.com/khezen/evoli)|Genetic Algorithm and Particle Swarm Optimization|18|9|21|2015-06-12T06:58:30Z|2021-10-27T10:26:23Z|
[randomForest](https://github.com/malaschitz/randomForest)|Random Forest implementation in golang|18|3|0|2018-10-25T07:05:29Z|2021-10-16T20:42:15Z|
[probab](https://github.com/ThePaw/probab)|Automatically exported from code.google.com/p/probab|16|5|3|2015-09-14T12:07:52Z|2015-09-14T12:08:34Z|
[ddt](https://github.com/sgrodriguez/ddt)|Golang Dynamic Decision Tree|16|1|0|2020-05-20T13:51:42Z|2021-02-22T12:47:34Z|


## Messaging
*Libraries that implement messaging systems.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[sarama](https://github.com/Shopify/sarama)|Sarama is a Go library for Apache Kafka 0.8, and up.|7938|1377|213|2013-07-05T18:52:38Z|2021-12-11T10:38:38Z|
[gorush](https://github.com/appleboy/gorush)|A push notification server written in Go (Golang).|5905|671|40|2016-03-22T07:15:20Z|2021-12-09T16:17:26Z|
[machinery](https://github.com/RichardKnop/machinery)|Machinery is an asynchronous task queue/job queue based on distributed message passing.|5806|724|177|2015-04-05T19:46:34Z|2021-12-09T16:22:25Z|
[centrifugo](https://github.com/centrifugal/centrifugo)|Scalable real-time messaging server in a language-agnostic way. Set up once and forever.|5625|467|4|2015-03-31T20:26:49Z|2021-12-08T08:45:11Z|
[go-socket.io](https://github.com/googollee/go-socket.io)|socket.io library for golang, a realtime application framework.|4409|712|95|2013-07-13T13:04:38Z|2021-12-07T21:48:59Z|
[nats.go](https://github.com/nats-io/nats.go)|Golang client for NATS, the cloud native messaging system.|3679|481|51|2012-08-15T12:54:59Z|2021-11-22T17:04:19Z|
[benthos](https://github.com/Jeffail/benthos)|Fancy stream processing made operationally mundane|3670|348|173|2016-03-22T01:18:48Z|2021-12-11T16:09:02Z|
[confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go)|Confluent&#39;s Apache Kafka Golang client|3035|464|191|2016-07-12T22:23:34Z|2021-12-08T11:49:51Z|
[mercure](https://github.com/dunglas/mercure)|Server-sent live updates: protocol and reference implementation|2604|193|15|2018-07-14T13:47:14Z|2021-12-08T16:28:42Z|
[apns2](https://github.com/sideshow/apns2)|⚡ HTTP/2 Apple Push Notification Service (APNs) push provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps, using the APNs HTTP/2 protocol.|2569|291|24|2016-01-05T00:56:53Z|2021-09-23T03:18:36Z|
[melody](https://github.com/olahol/melody)|:notes: Minimalist websocket framework for Go|2294|280|26|2015-05-13T20:38:32Z|2021-05-20T11:57:30Z|
[go-nsq](https://github.com/nsqio/go-nsq)|The official Go package for NSQ|2034|386|22|2013-08-29T01:18:32Z|2021-11-28T18:07:40Z|
[gopush-cluster](https://github.com/Terry-Mao/gopush-cluster)|Golang push server cluster|2015|556|4|2013-12-27T08:56:10Z|2017-06-07T12:18:31Z|
[asynq](https://github.com/hibiken/asynq)|Asynq: simple, reliable, and efficient distributed task queue in Go|1927|160|20|2019-11-15T05:17:55Z|2021-12-11T18:29:46Z|
[uniqush-push](https://github.com/uniqush/uniqush-push)|Uniqush is a free and open source software system which provides a unified push service for server side notification to apps on mobile devices.|1294|201|72|2011-08-29T08:42:37Z|2020-04-09T17:28:55Z|
[Beaver](https://github.com/Clivern/Beaver)|💨 A real time messaging system to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps.|1284|70|4|2018-10-20T21:10:43Z|2021-11-24T19:56:05Z|
[EventBus](https://github.com/asaskevich/EventBus)|[Go] Lightweight eventbus with async compatibility for Go|1037|124|17|2014-12-19T16:38:39Z|2021-06-22T14:07:38Z|
[zmq4](https://github.com/pebbe/zmq4)|A Go interface to ZeroMQ version 4|944|151|41|2013-10-18T11:48:51Z|2021-09-24T10:29:48Z|
[gollum](https://github.com/trivago/gollum)|An n:m message multiplexer written in Go|906|74|20|2015-06-20T21:51:20Z|2021-07-01T10:05:31Z|
[chanify](https://github.com/chanify/chanify)|Chanify is a safe and simple notification tools. This repository is command line tools for Chanify.|688|60|5|2021-02-25T17:20:04Z|2021-10-11T10:19:27Z|
[dbus](https://github.com/godbus/dbus)|Native Go bindings for D-Bus|646|166|44|2014-03-27T19:07:41Z|2021-12-06T09:08:53Z|
[golongpoll](https://github.com/jcuga/golongpoll)|golang long polling library.  Makes web pub-sub easy via HTTP long-poll servers and clients :smiley: :coffee: :computer:|584|48|1|2015-11-02T00:32:56Z|2021-04-29T12:16:50Z|
[mangos](https://github.com/nanomsg/mangos)|mangos is a pure Golang implementation of nanomsg&#39;s &#34;Scalablilty Protocols&#34;|462|62|26|2018-10-12T17:35:46Z|2021-12-07T22:50:13Z|
[emitter](https://github.com/olebedev/emitter)|Emits events in Go way, with wildcard, predicates, cancellation possibilities and many other good wins|413|32|4|2015-11-10T20:56:36Z|2020-02-05T13:10:15Z|
[glue](https://github.com/desertbit/glue)|Glue - Robust Go and Javascript Socket Library (Alternative to Socket.io)|390|32|6|2015-06-07T10:21:15Z|2020-05-20T06:46:44Z|
[pubsub](https://github.com/cskr/pubsub)|A simple pubsub package for go.|364|57|1|2012-04-01T06:31:43Z|2020-07-15T03:43:26Z|
[bus](https://github.com/mustafaturan/bus)|🔊Minimalist message bus implementation for internal communication with zero-allocation magic on Emit|235|16|0|2019-04-27T06:41:53Z|2021-05-11T03:36:00Z|
[message-bus](https://github.com/vardius/message-bus)|Go simple async message bus|202|34|2|2017-10-04T09:18:34Z|2021-01-14T22:04:03Z|
[rabtap](https://github.com/jandelgado/rabtap)|RabbitMQ wire tap and swiss army knife|196|14|3|2017-11-11T11:32:39Z|2021-11-28T18:37:37Z|
[amqp091-go](https://github.com/rabbitmq/amqp091-go)|An AMQP 0-9-1 Go client maintained by the RabbitMQ team. Originally by @streadway: `streadway/amqp`|158|19|8|2021-06-09T11:03:48Z|2021-12-10T09:34:24Z|
[guble](https://github.com/smancke/guble)|websocket based messaging server written in golang|149|21|5|2015-11-15T20:32:42Z|2017-10-31T19:15:41Z|
[hub](https://github.com/leandro-lugaresi/hub)|:incoming_envelope: A fast Message/Event Hub using publish/subscribe pattern with support for topics like* rabbitMQ exchanges for Go applications|111|11|2|2018-04-13T23:47:13Z|2020-10-26T14:23:55Z|
[oplog](https://github.com/dailymotion/oplog)|A generic oplog/replication system for microservices|110|12|2|2014-11-06T09:11:15Z|2015-11-07T00:51:48Z|
[rabbus](https://github.com/rafaeljesus/rabbus)|A tiny wrapper over amqp exchanges and queues 🚌 ✨|92|23|6|2017-05-07T08:51:11Z|2019-07-23T10:48:01Z|
[drone-line](https://github.com/appleboy/drone-line)|Sending line notifications using a binary, docker or Drone CI.|76|17|0|2016-09-13T05:21:44Z|2021-06-18T00:53:29Z|
[go-mq](https://github.com/cheshir/go-mq)|Declare AMQP entities like queues, producers, and consumers in a declarative way. Can be used to work with RabbitMQ.|72|12|3|2017-06-19T16:16:30Z|2021-11-30T12:40:58Z|
[nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus)|A tiny wrapper around NSQ topic and channel :rocket:|70|13|2|2017-01-15T22:05:13Z|2018-02-15T22:30:14Z|
[redisqueue](https://github.com/robinjoseph08/redisqueue)|redisqueue provides a producer and consumer of a queue that uses Redis streams|69|22|4|2019-07-07T04:36:54Z|2021-07-19T21:16:33Z|
[RapidMQ](https://github.com/sybrexsys/RapidMQ)|RapidMQ is a pure, extremely productive, lightweight and reliable library for managing of the local messages queue|63|10|1|2016-10-04T21:07:48Z|2017-12-07T08:34:10Z|
[commander](https://github.com/jeroenrinzema/commander)|Build event-driven and event streaming applications with ease|57|4|2|2018-04-20T12:30:51Z|2021-04-28T21:55:28Z|
[go-notify](https://github.com/TheCreeper/go-notify)|Package notify provides an implementation of the Gnome DBus Notifications Specification.|56|11|1|2015-03-01T19:21:44Z|2020-12-11T18:09:42Z|
[go-res](https://github.com/jirenius/go-res)|RES Service protocol library for Go|52|7|6|2018-07-15T09:10:11Z|2021-09-17T15:16:12Z|
[event](https://github.com/agoalofalife/event)|The implementation of the pattern observer|43|9|0|2017-07-02T12:19:56Z|2018-02-19T12:11:32Z|
[hare](https://github.com/leozz37/hare)|🐇  Easy to use socket lib for Golang and CLI tool|32|5|0|2020-12-01T22:30:27Z|2021-08-18T23:40:25Z|
[ami](https://github.com/kak-tus/ami)|Go client to reliable queues based on Redis Cluster Streams|21|7|0|2018-10-27T10:38:16Z|2020-04-02T22:56:51Z|
[gosd](https://github.com/alexsniffin/gosd)|A library for scheduling when to dispatch a message to a channel|18|3|0|2020-05-17T23:19:51Z|2020-11-16T03:32:07Z|
[go-vitotrol](https://github.com/maxatome/go-vitotrol)|golang client library to Viessmann Vitotrol web service|16|4|1|2016-11-03T19:59:43Z|2021-02-19T21:40:40Z|
[rmqconn](https://github.com/sbabiv/rmqconn)|RabbitMQ Reconnection client|15|1|0|2019-01-14T16:05:44Z|2020-01-27T09:57:25Z|
[jazz](https://github.com/socifi/jazz)|Abstraction layer for simple rabbitMQ connection, messaging and administration|13|2|1|2018-10-22T12:28:15Z|2019-03-21T11:10:11Z|
[gaurun-client](https://github.com/osamingo/gaurun-client)|Gaurun Client written in Go|9|3|0|2017-06-29T02:50:51Z|2021-08-03T07:04:33Z|


## Microsoft Office

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[unioffice](https://github.com/unidoc/unioffice)|Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents|3103|349|27|2017-08-29T01:25:48Z|2021-11-10T13:02:12Z|


### Microsoft Excel
*Libraries for working with Microsoft Excel.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[excelize](https://github.com/qax-os/excelize)|Golang library for reading and writing Microsoft Excel™ (XLSX) files.|10229|1070|78|2016-08-29T12:32:12Z|2021-12-11T16:07:11Z|
[xlsx](https://github.com/tealeg/xlsx)|Go (golang) library for reading and writing XLSX files. |5184|780|51|2011-06-28T15:20:28Z|2021-11-22T12:16:09Z|
[xlsx](https://github.com/plandem/xlsx)|Fast and reliable way to work with Microsoft Excel™ [xlsx] files in Golang|145|20|11|2017-08-26T23:11:38Z|2020-11-04T15:00:26Z|
[go-excel](https://github.com/szyhf/go-excel)|A simple and light excel file reader to read a standard excel as a table faster   一个轻量级的Excel数据读取库，用一种更`关系数据库`的方式解析Excel。|130|28|2|2017-09-03T11:51:58Z|2021-12-06T09:50:22Z|
[goxlsxwriter](https://github.com/fterrag/goxlsxwriter)|Golang bindings for libxlsxwriter for writing XLSX files|17|5|1|2017-03-13T04:15:17Z|2018-07-31T21:24:17Z|


### Dependency Injection
*Libraries for working with dependency injection.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[fx](https://github.com/uber-go/fx)|A dependency injection based application framework for Go.|2363|161|24|2016-10-27T00:25:00Z|2021-12-10T03:35:18Z|
[dig](https://github.com/uber-go/dig)|A reflection based dependency injection toolkit for Go.|2246|144|18|2017-03-21T23:55:50Z|2021-12-11T00:24:17Z|
[container](https://github.com/golobby/container)|A lightweight yet powerful IoC dependency injection container for Go projects|279|17|3|2019-09-23T16:12:50Z|2021-12-03T13:20:59Z|
[di](https://github.com/goioc/di)|Simple and yet powerful Dependency Injection for Go|122|6|0|2020-06-11T12:28:06Z|2021-08-01T09:49:45Z|
[dingo](https://github.com/i-love-flamingo/dingo)|Go Dependency Injection Framework|116|8|9|2018-10-29T08:55:18Z|2021-04-30T09:02:16Z|
[di](https://github.com/goava/di)|🛠 A full-featured dependency injection container for go programming language.|108|9|1|2020-02-03T19:06:39Z|2021-11-30T00:02:18Z|
[alice](https://github.com/magic003/alice)|An additive dependency injection container for Golang.|43|4|0|2017-04-08T16:25:21Z|2017-04-26T06:08:23Z|
[wire](https://github.com/Fs02/wire)|Strict Runtime Dependency Injection for Golang|34|8|1|2018-07-05T10:42:24Z|2021-08-22T07:00:18Z|
[linker](https://github.com/logrange/linker)|Dependency Injection and Inversion of Control package|32|6|0|2018-12-04T23:56:34Z|2020-06-25T19:18:10Z|
[gocontainer](https://github.com/vardius/gocontainer)|Simple Dependency Injection Container|14|2|0|2019-06-06T08:18:07Z|2020-03-23T09:12:06Z|
[nject](https://github.com/muir/nject)|Golang type-safe dependency injection|5|1|1|2021-09-15T03:48:32Z|2021-11-18T08:29:37Z|
[kinit](https://github.com/go-kata/kinit)|GO Dependency Injection|5|0|0|2021-01-24T13:41:41Z|2021-06-12T14:27:19Z|
[di](https://github.com/HnH/di)|DI container library that is focused on clean API and flexibility.|3|2|0|2021-10-13T07:09:09Z|2021-12-01T09:10:26Z|


### Project Layout
*Unofficial set of patterns for structuring projects.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[project-layout](https://github.com/golang-standards/project-layout)|Standard Go Project Layout|28148|3056|70|2017-09-09T16:33:26Z|2021-12-06T03:46:22Z|
[service](https://github.com/ardanlabs/service)|Starter code for writing web services in Go using Kubernetes.|2073|362|0|2017-11-20T14:51:17Z|2021-12-08T14:55:09Z|
[modern-go-application](https://github.com/sagikazarmark/modern-go-application)|Modern Go Application example|1106|106|16|2018-09-14T12:19:02Z|2021-11-24T18:43:45Z|
[cookiecutter-golang](https://github.com/lacion/cookiecutter-golang)|A Go project template|483|121|4|2016-12-18T18:22:26Z|2021-02-06T19:16:10Z|
[seed](https://github.com/golang-templates/seed)|Go application GitHub repository template.|230|20|0|2020-04-30T21:31:36Z|2021-09-20T06:33:11Z|
[go-todo-backend](https://github.com/Fs02/go-todo-backend)|Go Todo Backend example using modular project layout for product microservice.|107|14|0|2020-06-25T14:28:50Z|2021-09-25T06:49:10Z|
[scaffold](https://github.com/catchplay/scaffold)|Generate scaffold project layout for Go.|101|23|2|2018-12-11T10:36:03Z|2019-01-10T04:00:20Z|
[go-starter](https://github.com/allaboutapps/go-starter)|An opinionated production-ready SQL-/Swagger-first RESTful JSON API written in Go, highly integrated with VSCode DevContainers by allaboutapps.|99|16|8|2020-05-08T14:22:49Z|2021-12-10T04:03:36Z|
[go-sample](https://github.com/zitryss/go-sample)|Go Project Sample Layout|95|22|0|2019-01-24T23:41:46Z|2019-01-24T23:54:54Z|
[gobase](https://github.com/wajox/gobase)|This is a simple skeleton for golang applications|16|3|0|2020-12-15T16:54:20Z|2021-09-20T22:40:52Z|
[go-project-layout](https://github.com/wangyoucao577/go-project-layout)|My understanding of how to structure a golang project. |11|1|0|2019-10-06T12:59:24Z|2021-05-16T01:32:02Z|
[inizio](https://github.com/insidieux/inizio)|Golang project standard layout generator|9|0|1|2021-03-02T20:59:22Z|2021-05-07T17:46:30Z|


### Strings
*Libraries for working with strings.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[xstrings](https://github.com/huandu/xstrings)|Implements string functions widely used in other languages but absent in Go.|951|64|0|2015-01-06T07:25:26Z|2020-12-26T16:26:15Z|
[sttr](https://github.com/abhimanyu003/sttr)|cross-platform, cli app to perform various operations on string|174|12|1|2021-09-18T14:00:40Z|2021-12-05T13:36:11Z|
[strutil](https://github.com/ozgio/strutil)|String utilities for Go|143|16|0|2018-08-16T06:56:15Z|2021-10-26T00:13:52Z|
[stringy](https://github.com/gobeam/stringy)|Convert string to camel case, snake case, kebab case / slugify, custom delimiter, pad string, tease string and many other functionalities with help of by Stringy package.|94|9|0|2020-04-03T03:34:10Z|2021-11-09T05:49:51Z|
[bexp](https://github.com/mkungla/bexp)|Go implementation of Brace Expansion mechanism to generate arbitrary strings.|3|0|0|2020-12-15T17:11:43Z|2021-09-30T02:14:00Z|
[go-formatter](https://gitlab.com/tymonx/go-formatter)|Implements replacement fields surrounded by curly braces {} format strings.|-|-|-|-|-|


### Uncategorized
*These libraries were placed here because none of the other categories seemed to fit.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[gopsutil](https://github.com/shirou/gopsutil)|psutil for golang|7084|1179|123|2014-04-18T07:35:28Z|2021-12-11T11:07:59Z|
[archiver](https://github.com/mholt/archiver)|Easily create &amp; extract archives, and compress &amp; decompress files of various formats|3368|318|70|2016-04-08T22:46:55Z|2021-11-22T18:55:45Z|
[gofakeit](https://github.com/brianvoe/gofakeit)|Random fake data generator written in go|2188|131|4|2015-04-24T04:45:59Z|2021-12-01T04:34:31Z|
[gatus](https://github.com/TwiN/gatus)|⛑ Gatus - Automated service health dashboard|2027|118|45|2019-09-04T02:35:40Z|2021-12-10T21:15:06Z|
[gosms](https://github.com/haxpax/gosms)|:mailbox_closed: Your own local SMS gateway in Go|1358|140|4|2015-01-23T19:25:55Z|2021-02-05T19:15:02Z|
[go-resiliency](https://github.com/eapache/go-resiliency)|Resiliency patterns for golang|1288|108|4|2014-11-29T04:11:32Z|2021-09-17T10:55:35Z|
[base64Captcha](https://github.com/mojocn/base64Captcha)|captcha of base64 image string|1278|201|7|2017-12-12T12:17:07Z|2021-12-07T07:35:42Z|
[go-commons-pool](https://github.com/jolestar/go-commons-pool)|a generic object pool for golang|996|136|3|2015-12-28T14:26:23Z|2021-04-30T02:11:16Z|
[llvm](https://github.com/llir/llvm)|Library for interacting with LLVM IR in pure Go.|826|56|20|2014-09-19T11:18:44Z|2021-11-24T05:12:10Z|
[shortid](https://github.com/teris-io/shortid)|Super short, fully unique, non-sequential and URL friendly Ids|722|57|0|2016-01-04T01:17:10Z|2020-11-17T13:42:43Z|
[health](https://github.com/alexliesenfeld/health)|A simple and flexible health check library for Go.|467|18|0|2021-07-02T11:27:34Z|2021-12-10T20:11:14Z|
[health](https://github.com/dimiro1/health)|An easy to use, extensible health check library for Go applications.|425|43|4|2016-03-08T23:04:43Z|2019-10-21T10:50:08Z|
[stateless](https://github.com/qmuntal/stateless)|Go library for creating state machines|386|23|6|2019-09-11T08:19:18Z|2021-09-24T07:51:52Z|
[go-conv](https://github.com/cstockton/go-conv)|Fast conversions across various Go types with a simple API.|373|17|0|2016-10-11T07:41:41Z|2021-08-23T21:52:24Z|
[banner](https://github.com/dimiro1/banner)|An easy way to add useful startup banners into your Go applications|368|21|1|2016-03-25T21:28:44Z|2021-01-04T09:25:38Z|
[gountries](https://github.com/pariz/gountries)|Gountries provides: Countries (ISO-3166-1), Country Subdivisions(ISO-3166-2), Currencies (ISO 4217), Geo Coordinates(ISO-6709) as well as translations, country borders and other stuff exposed as struct data.|330|58|14|2016-01-13T08:04:18Z|2021-11-29T14:09:08Z|
[shoutrrr](https://github.com/containrrr/shoutrrr)|Notification library for gophers and their furry friends.|270|37|13|2019-04-11T06:49:34Z|2021-11-27T13:44:28Z|
[ffmt](https://github.com/go-ffmt/ffmt)|Golang beautify data display for Humans|244|19|2|2015-02-14T15:19:45Z|2021-11-19T15:22:56Z|
[lk](https://github.com/hyperboloide/lk)|Simple licensing library for golang.|227|35|1|2016-07-14T16:06:07Z|2020-05-04T06:08:01Z|
[antch](https://github.com/antchfx/antch)|Antch, a fast, powerful and extensible web crawling &amp; scraping framework for Go|222|41|4|2017-09-28T05:44:17Z|2020-05-31T15:12:21Z|
[healthcheck](https://github.com/etherlabsio/healthcheck)|An simple, easily extensible and concurrent health-check library for Go services|208|29|1|2017-08-18T12:48:40Z|2021-06-17T16:33:44Z|
[battery](https://github.com/distatus/battery)|cross-platform, normalized battery information library|195|28|9|2016-03-12T23:03:40Z|2021-10-14T20:55:32Z|
[bitio](https://github.com/icza/bitio)|Optimized bit-level Reader and Writer for Go.|171|23|1|2016-05-31T10:02:30Z|2021-12-07T07:02:48Z|
[go-unarr](https://github.com/gen2brain/go-unarr)|Go bindings for unarr (decompression library for RAR, TAR, ZIP and 7z archives)|158|25|5|2015-11-01T09:38:37Z|2021-10-14T00:29:46Z|
[stats](https://github.com/go-playground/stats)|:chart_with_upwards_trend: Monitors Go MemStats &#43; System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...|157|20|1|2015-09-14T20:20:20Z|2016-09-07T12:51:16Z|
[turtle](https://github.com/hackebrot/turtle)|Emojis for Go 😄🐢🚀|128|11|1|2017-09-08T22:25:32Z|2021-10-04T08:23:47Z|
[gommit](https://github.com/antham/gommit)|Enforce git message commit consistency|95|3|1|2016-08-30T11:10:11Z|2021-11-24T07:46:44Z|
[gotoprom](https://github.com/cabify/gotoprom)|Type-safe Prometheus metrics builder library for golang|92|2|0|2018-10-10T16:07:33Z|2020-01-29T09:07:33Z|
[indigo](https://github.com/osamingo/indigo)|A distributed unique ID generator of using Sonyflake and encoded by Base58|91|11|0|2016-08-31T14:17:45Z|2021-02-01T17:53:46Z|
[captcha](https://github.com/steambap/captcha)|:sunglasses:Package captcha provides an easy to use, unopinionated API for captcha generation|86|16|0|2017-09-12T06:52:15Z|2021-09-28T08:44:34Z|
[morse](https://github.com/alwindoss/morse)|Morse Code Library in Go|70|12|3|2018-08-15T05:31:31Z|2021-08-28T21:07:18Z|
[persian](https://github.com/mavihq/persian)|Some utilities for Persian language in Go (Golang)|61|9|1|2017-10-16T16:16:56Z|2021-06-17T05:22:01Z|
[pdfgen](https://github.com/hyperboloide/pdfgen)|HTTP service to generate PDF from Json requests|54|10|0|2015-11-30T19:27:26Z|2018-02-19T15:49:42Z|
[xkg](https://github.com/go-xkg/xkg)|User level X Keyboard Grabber|52|6|1|2015-01-05T01:04:43Z|2015-01-08T04:01:03Z|
[faker](https://github.com/pioz/faker)|Random fake data and struct generator for Go.|51|4|0|2020-07-22T20:09:46Z|2021-10-05T06:59:29Z|
[browscap_go](https://github.com/digitalcrab/browscap_go)|GoLang Library for Browser Capabilities Project|39|24|11|2014-09-18T04:47:42Z|2021-09-15T05:39:42Z|
[datacounter](https://github.com/miolini/datacounter)|Golang counters for readers/writers|37|6|2|2015-10-14T19:15:50Z|2020-02-06T09:12:12Z|
[autoflags](https://github.com/artyom/autoflags)|Populate go command line app flags from config struct|36|3|0|2014-05-15T19:00:29Z|2021-04-29T21:03:09Z|
[sandid](https://github.com/aofei/sandid)|Every grain of sand on Earth has its own ID.|33|6|0|2018-06-12T01:24:14Z|2021-11-11T11:48:29Z|
[url-shortener](https://github.com/pantrif/url-shortener)|A golang URL Shortener|32|6|0|2018-06-04T05:57:45Z|2018-06-09T14:39:44Z|
[gosh](https://github.com/osamingo/gosh)|Provide Go Statistics Handler, Struct, Measure Method|28|2|0|2018-05-25T08:55:55Z|2021-01-08T10:30:51Z|
[xdg](https://github.com/rkoesters/xdg)|FreeDesktop.org (xdg) Specs implemented in Go|27|8|1|2013-12-15T09:51:51Z|2021-07-18T10:03:54Z|
[metrics](https://github.com/pascaldekloe/metrics)|atomic measures &#43; Prometheus exposition library|22|5|2|2019-01-29T09:39:18Z|2021-03-08T20:02:13Z|
[gtree](https://github.com/ddddddO/gtree)|Generate tree🌳 from Markdown or Programmatically. Provide CLI and Package for Go, output is JSON or YAML or TOML or  tree command.|21|1|2|2021-05-30T01:51:22Z|2021-12-11T12:24:21Z|
[shellwords](https://github.com/Wing924/shellwords)|A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell.|16|3|0|2017-09-28T09:08:28Z|2017-10-03T02:04:28Z|
[anagent](https://github.com/mudler/anagent)|Minimalistic, pluggable Golang evloop/timer handler with dependency-injection|13|4|0|2017-12-29T17:16:25Z|2018-08-12T17:51:33Z|
[avgRating](https://github.com/kirillDanshin/avgRating)|Calculate average score and rating based on Wilson Score Equation|11|3|0|2017-08-05T19:04:30Z|2017-08-05T19:37:44Z|
[hostutils](https://github.com/Wing924/hostutils)|A golang library for packing and unpacking hosts list|10|5|0|2017-09-26T03:47:32Z|2019-01-06T09:46:53Z|
[numa](https://github.com/lrita/numa)|NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code.|6|3|0|2018-12-10T09:59:13Z|2020-12-16T02:23:07Z|
[generators](https://github.com/azr/generators)|#golang generator|4|2|0|2016-02-29T14:29:02Z|2016-12-30T13:41:30Z|
[go-commandbus](https://github.com/lana/go-commandbus)|Simple command bus for GO|3|3|0|2019-10-03T20:08:22Z|2020-04-21T09:20:43Z|
[varint](https://github.com/chmike/varint)|variable length integer encoding using prefix code|2|0|0|2021-11-30T11:29:34Z|2021-12-01T10:24:46Z|
[basexx](https://github.com/bobg/basexx)|Convert digit strings between arbitrary bases.|1|0|0|2019-06-08T17:46:13Z|2021-10-02T14:57:12Z|
[go-openapi](https://github.com/go-openapi)|Collection of packages to parse and utilize open-api schemas.|-|-|-|-|-|


## Natural Language Processing
*Libraries for working with human languages.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[prose](https://github.com/jdkato/prose)|:book: A Golang library for text processing, including tokenization, part-of-speech tagging, and named-entity extraction.|2838|137|17|2017-02-17T17:08:22Z|2021-09-21T20:53:23Z|
[go-i18n](https://github.com/nicksnyder/go-i18n)|Translate your Go program into multiple languages.|1830|186|15|2012-01-14T21:44:37Z|2021-06-23T23:51:28Z|
[gojieba](https://github.com/yanyiwu/gojieba)|&#34;结巴&#34;中文分词的Golang版本|1701|240|48|2015-09-12T01:30:44Z|2021-06-05T14:05:43Z|
[gse](https://github.com/go-ego/gse)|Go efficient multilingual NLP and text segmentation; support english, chinese, japanese and other. Go 高性能多语言 NLP 和分词|1692|148|4|2017-06-23T15:42:35Z|2021-12-09T21:48:13Z|
[when](https://github.com/olebedev/when)|A natural language date/time parser with pluggable rules|1135|66|14|2016-12-27T13:11:46Z|2021-12-11T00:46:37Z|
[spago](https://github.com/nlpodyssey/spago)|Self-contained Machine Learning and Natural Language Processing library in Go|1051|51|12|2020-01-05T20:39:29Z|2021-12-07T13:56:57Z|
[go-pinyin](https://github.com/mozillazg/go-pinyin)|汉字转拼音|1044|161|13|2014-11-09T14:04:33Z|2021-12-11T14:54:43Z|
[kagome](https://github.com/ikawaha/kagome)|Self-contained Japanese Morphological Analyzer written in pure Go|594|41|5|2014-06-26T04:38:13Z|2021-11-23T14:10:05Z|
[whatlanggo](https://github.com/abadojack/whatlanggo)|Natural language detection library for Go|512|51|11|2017-02-20T17:32:01Z|2021-01-15T09:31:00Z|
[nlp](https://github.com/shixzie/nlp)|[UNMANTEINED] Extract values from strings and fill your structs with nlp.|376|31|3|2017-01-25T07:19:03Z|2017-09-18T14:32:30Z|
[nlp](https://github.com/james-bowman/nlp)|Selected Machine Learning algorithms for natural language processing and semantic analysis in Golang|349|41|4|2017-03-15T08:28:05Z|2021-05-11T12:03:06Z|
[sentences](https://github.com/neurosnap/sentences)|A multilingual command line sentence tokenizer in Golang|314|30|2|2015-08-07T01:08:20Z|2021-06-18T16:19:34Z|
[getlang](https://github.com/rylans/getlang)|Natural language detection package in pure Go|126|18|4|2018-03-01T21:27:30Z|2020-12-27T07:47:21Z|
[go-unidecode](https://github.com/mozillazg/go-unidecode)|ASCII transliterations of Unicode text.|89|14|4|2016-07-08T13:15:10Z|2021-04-29T19:33:56Z|
**[ARCHIVED]**  [go-nlp](https://github.com/nuance/go-nlp)|Utilities for working with discrete probability distributions and other tools useful for doing NLP work|89|10|0|2011-05-02T06:43:36Z|2011-11-15T17:49:45Z|
[RAKE.Go](https://github.com/afjoseph/RAKE.Go)|A Go port of the Rapid Automatic Keyword Extraction algorithm (RAKE)|84|17|4|2016-12-17T13:36:25Z|2020-02-27T08:40:40Z|
[gounidecode](https://github.com/fiam/gounidecode)|Unicode transliterator for #golang|73|20|2|2012-05-01T11:59:34Z|2015-09-23T21:17:29Z|
[segment](https://github.com/blevesearch/segment)|A Go library for performing Unicode Text Segmentation as described in Unicode Standard Annex #29|69|11|5|2014-10-16T19:24:26Z|2021-01-13T19:12:27Z|
[textcat](https://github.com/pebbe/textcat)|A Go package for n-gram based text categorization, with support for utf-8 and raw text|65|9|1|2012-09-21T15:04:45Z|2021-02-20T13:40:48Z|
[go-stem](https://github.com/agonopol/go-stem)|Word Stemming in Go|64|15|1|2011-09-23T19:07:23Z|2018-06-16T22:48:56Z|
[MMSEGO](https://github.com/awsong/MMSEGO)|Chinese word splitting algorithm MMSEG in GO|59|14|0|2012-04-18T04:06:21Z|2012-04-18T04:18:51Z|
[stemmer](https://github.com/dchest/stemmer)|Stemmer packages for Go programming language. Includes English, German and Dutch stemmers.|50|5|0|2011-03-21T02:08:12Z|2016-12-07T10:24:03Z|
[address](https://github.com/bojanz/address)|Address handling for Go.|42|1|0|2020-10-07T18:15:27Z|2021-11-08T11:03:07Z|
[porter2](https://github.com/zentures/porter2)|High Performance Porter2 Stemmer|41|6|1|2015-01-21T07:30:32Z|2020-10-07T17:10:59Z|
[go2vec](https://github.com/danieldk/go2vec)|Read and use word2vec vectors in Go|41|5|0|2015-01-27T12:02:04Z|2018-08-30T05:34:08Z|
[petrovich](https://github.com/striker2000/petrovich)|Golang port of Petrovich - an inflector for Russian anthroponyms.|35|4|0|2016-12-26T22:50:38Z|2021-02-22T18:27:56Z|
[go-localize](https://github.com/m1/go-localize)|i18n (Internationalization and localization) engine written in Go, used for translating locale strings. |29|10|1|2019-12-23T12:02:51Z|2021-10-29T18:23:38Z|
[snowball](https://github.com/goodsign/snowball)|Cgo binding for Snowball C library|28|4|0|2012-12-11T12:42:19Z|2017-06-27T08:13:41Z|
[mystem](https://github.com/dveselov/mystem)|CGo bindings to Yandex.Mystem|27|7|0|2016-08-30T14:55:39Z|2016-10-05T05:53:17Z|
[paicehusk](https://github.com/rookii/paicehusk)|Golang implementation of the Paice/Husk Stemming Algorithm|27|6|2|2012-09-29T16:06:58Z|2013-12-16T12:45:11Z|
[iuliia-go](https://github.com/mehanizm/iuliia-go)|Transliterate Cyrillic → Latin in every possible way|26|3|0|2020-04-27T09:29:40Z|2021-06-15T16:27:22Z|
[golibstemmer](https://github.com/rjohnsondev/golibstemmer)|Go bindings for the snowball libstemmer library including porter 2|19|5|0|2012-08-06T19:31:05Z|2014-06-17T16:04:56Z|
[icu](https://github.com/goodsign/icu)|Cgo binding for icu4c library|19|5|2|2012-12-11T13:09:41Z|2017-03-29T16:17:26Z|
[govader](https://github.com/jonreiter/govader)|vader sentiment analysis in go|18|3|0|2020-01-19T10:06:15Z|2021-02-24T07:24:41Z|
[transliterator](https://github.com/alexsergivan/transliterator)|Golang text Transliterator (i.e München -&gt; Muenchen)|15|7|1|2020-04-17T14:19:55Z|2020-05-08T16:48:36Z|
[gotokenizer](https://github.com/xujiajun/gotokenizer)|A tokenizer based on the dictionary and Bigram language models for Go. (Now only support chinese segmentation)|13|5|0|2018-10-11T03:22:36Z|2019-04-10T09:39:09Z|
[shamoji](https://github.com/osamingo/shamoji)|The shamoji (杓文字) is a word filtering package|12|1|0|2017-07-23T06:38:42Z|2021-01-14T18:13:56Z|
[libtextcat](https://github.com/goodsign/libtextcat)|Cgo binding for libtextcat C library|10|7|0|2012-12-10T21:21:47Z|2012-12-27T17:23:35Z|
[detectlanguage-go](https://github.com/detectlanguage/detectlanguage-go)|Detect Language API Go Client|10|1|0|2019-12-14T23:30:44Z|2020-10-11T14:32:38Z|
[porter](https://github.com/a2800276/porter)|porter stemmer|8|1|0|2013-09-17T11:10:16Z|2013-10-03T11:10:18Z|
[t](https://github.com/youthlin/t)|t: translation util for go, using GNU gettext|7|1|0|2021-06-04T07:22:41Z|2021-10-29T02:26:36Z|
[gosentiwordnet](https://github.com/dinopuguh/gosentiwordnet)|💬 Sentiment analyzer library using SentiWordnet in Go|7|1|0|2020-04-21T09:09:28Z|2021-03-11T05:01:50Z|


## Networking
*Libraries for working with various layers of the network.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[fasthttp](https://github.com/valyala/fasthttp)|Fast HTTP package for Go. Tuned for high performance. Zero memory allocations in hot paths. Up to 10x faster than net/http|16559|1387|51|2015-10-18T22:19:57Z|2021-12-11T18:23:51Z|
[kcptun](https://github.com/xtaci/kcptun)|A Stable &amp; Secure Tunnel based on KCP with N:M multiplexing and FEC. Available for ARM, MIPS, 386 and AMD64。KCPプロトコルに基づく安全なトンネル。KCP 프로토콜을 기반으로 하는 보안 터널입니다。|12705|2458|78|2016-02-26T09:54:46Z|2021-11-30T15:40:09Z|
[webrtc](https://github.com/pion/webrtc)|Pure Go implementation of the WebRTC API|8363|1062|62|2018-05-18T23:10:05Z|2021-12-10T16:53:40Z|
[quic-go](https://github.com/lucas-clemente/quic-go)|A QUIC implementation in pure go|6049|810|96|2016-04-06T20:16:27Z|2021-10-31T16:35:52Z|
[dns](https://github.com/miekg/dns)|DNS library in Go|5910|930|16|2010-08-03T21:56:23Z|2021-12-03T17:56:41Z|
[gnet](https://github.com/panjf2000/gnet)|🚀 gnet is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go./ gnet 是一个高性能、轻量级、非阻塞的事件驱动 Go 网络框架。|5639|650|32|2019-02-24T03:48:45Z|2021-12-10T02:05:04Z|
[gopacket](https://github.com/google/gopacket)|Provides packet processing capabilities for Go|4535|879|207|2015-03-16T20:46:00Z|2021-12-10T10:56:32Z|
[httplab](https://github.com/qustavo/httplab)|The interactive web server|3750|122|12|2017-02-08T17:13:19Z|2019-06-05T15:10:46Z|
[kcp-go](https://github.com/xtaci/kcp-go)| A Crypto-Secure, Production-Grade Reliable-UDP Library for golang with FEC |3172|582|31|2015-06-16T06:15:55Z|2021-10-22T12:23:18Z|
[gobgp](https://github.com/osrg/gobgp)|BGP implemented in the Go Programming Language|2357|524|105|2014-09-14T01:51:58Z|2021-12-11T10:30:56Z|
[ssh](https://github.com/gliderlabs/ssh)|Easy SSH servers in Golang|2233|280|40|2016-10-03T21:53:44Z|2021-11-25T22:16:27Z|
[fortio](https://github.com/fortio/fortio)|Fortio load testing library, command line tool, advanced echo server and web UI in go (golang). Allows to specify a set query-per-second load and record latency histograms and other useful stats.|2192|183|77|2017-10-10T01:01:39Z|2021-12-06T21:51:35Z|
[paho.mqtt.golang](https://github.com/eclipse/paho.mqtt.golang)||1806|426|17|2016-02-03T19:03:35Z|2021-12-08T20:54:15Z|
[water](https://github.com/songgao/water)|A simple TUN/TAP library written in native Go.|1338|203|20|2013-03-25T20:06:52Z|2021-04-16T14:47:27Z|
[gev](https://github.com/Allenxuxu/gev)|🚀Gev is a lightweight, fast non-blocking TCP network library / websocket server based on Reactor mode. Support custom protocols to quickly and easily build high-performance servers. |1302|158|6|2019-09-01T12:16:18Z|2021-12-05T07:58:34Z|
[go-getter](https://github.com/hashicorp/go-getter)|Package for downloading things from a string URL using a variety of protocols.|1263|160|102|2015-10-12T23:17:07Z|2021-12-03T15:46:00Z|
[nff-go](https://github.com/intel-go/nff-go)|NFF-Go -Network Function Framework for GO (former YANFF)|1134|134|63|2017-03-29T17:07:29Z|2021-09-07T16:07:05Z|
[sftp](https://github.com/pkg/sftp)|SFTP support for the go.crypto/ssh package|1072|310|27|2013-11-05T04:36:00Z|2021-12-05T13:34:52Z|
[grab](https://github.com/cavaliergopher/grab)|A download manager package for Go|947|112|25|2016-01-05T12:46:35Z|2021-09-12T12:59:36Z|
[ftp](https://github.com/jlaffaye/ftp)|FTP client package for Go|867|290|9|2011-05-06T18:31:51Z|2021-11-26T14:51:31Z|
[mdns](https://github.com/hashicorp/mdns)|Simple mDNS client/server library in Golang|831|181|33|2014-01-29T19:39:18Z|2021-11-05T13:26:16Z|
[gosnmp](https://github.com/gosnmp/gosnmp)|An SNMP library written in Go|781|257|27|2012-08-27T05:59:24Z|2021-11-17T11:28:02Z|
[vssh](https://github.com/yahoo/vssh)|Go Library to Execute Commands Over SSH at Scale|758|57|0|2020-06-09T16:19:22Z|2020-11-22T02:34:52Z|
[cidranger](https://github.com/yl2chen/cidranger)|Fast IP to CIDR lookup in Golang|651|76|5|2017-08-21T05:50:14Z|2021-11-17T04:43:33Z|
[lhttp](https://github.com/fanux/lhttp)|go websocket, a better way to buid your IM server|637|132|6|2015-12-29T01:13:36Z|2018-04-08T08:06:09Z|
[gmqtt](https://github.com/DrmagicE/gmqtt)|Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.x and V5 in golang|523|105|7|2018-09-16T11:46:17Z|2021-11-25T11:14:34Z|
[peerdiscovery](https://github.com/schollz/peerdiscovery)|Pure-Go library for cross-platform local peer discovery using UDP multicast :woman: :repeat: :woman:|514|37|6|2018-04-22T23:59:37Z|2021-10-02T17:54:31Z|
[gotcp](https://github.com/gansidui/gotcp)|A Go package for quickly building tcp servers|486|157|0|2014-04-13T14:54:01Z|2017-04-18T07:26:13Z|
**[ARCHIVED]**  [stun](https://github.com/gortc/stun)|Fast RFC 5389 STUN implementation in go|475|50|4|2016-04-24T17:46:38Z|2021-05-17T05:47:09Z|
[go-stun](https://github.com/ccding/go-stun)|A go implementation of the STUN client (RFC 3489 and RFC 5389)|461|88|3|2013-08-17T22:16:33Z|2021-04-13T15:42:25Z|
[gopcap](https://github.com/akrennmair/gopcap)|A simple wrapper around libpcap for the Go programming language|435|137|12|2009-11-19T10:13:48Z|2021-05-17T13:59:53Z|
[nbio](https://github.com/lesismal/nbio)|Pure Go 1000k&#43; connections solution, support tls/http1.x/websocket and basically compatible with net/http, with high-performance and low memory cost, non-blocking, event-driven, easy-to-use.|415|43|5|2020-01-25T11:46:54Z|2021-12-06T12:46:23Z|
[raw](https://github.com/mdlayher/raw)|Package raw enables reading and writing data at the device driver level for a network interface.  MIT Licensed.|411|72|15|2015-07-06T16:11:47Z|2021-11-26T14:27:49Z|
[gaio](https://github.com/xtaci/gaio)|High performance async-io(proactor) networking for Golang。golangのための高性能非同期io(proactor)ネットワーキング|404|46|14|2019-12-20T05:19:00Z|2021-09-17T12:24:53Z|
[tcp_server](https://github.com/firstrow/tcp_server)|golang tcp server|400|138|4|2014-10-13T20:38:42Z|2021-11-10T09:30:31Z|
[winrm](https://github.com/masterzen/winrm)|Command-line tool and library for Windows remote command execution in Go|340|94|26|2013-12-30T18:29:15Z|2021-11-06T15:32:15Z|
[ftpserverlib](https://github.com/fclairamb/ftpserverlib)|golang ftp server library|272|64|4|2016-09-25T12:05:29Z|2021-11-25T11:29:19Z|
[arp](https://github.com/mdlayher/arp)|Package arp implements the ARP protocol, as described in RFC 826. MIT Licensed.|265|43|4|2015-07-06T18:50:34Z|2019-12-13T14:26:04Z|
[buffstreams](https://github.com/StabbyCutyou/buffstreams)|A library to simplify writing applications using TCP sockets to stream protobuff messages|247|33|7|2015-06-29T19:07:31Z|2020-08-14T20:02:54Z|
[ethernet](https://github.com/mdlayher/ethernet)|Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. MIT Licensed.|233|32|0|2015-07-03T00:15:18Z|2019-06-06T14:27:57Z|
[easytcp](https://github.com/DarthPestilane/easytcp)|:sparkles: :rocket: EasyTCP is a light-weight TCP framework written in Go (Golang), built with message router. EasyTCP helps you build a TCP server easily fast and less painful.|209|8|0|2021-04-26T10:11:59Z|2021-12-06T08:05:47Z|
[gnxi](https://github.com/google/gnxi)|gNXI Tools - gRPC Network Management/Operations Interface Tools|197|95|16|2017-09-26T08:19:41Z|2021-12-09T15:22:47Z|
[jazigo](https://github.com/udhos/jazigo)|Jazigo is a tool written in Go for retrieving configuration for multiple devices, similar to rancid, fetchconfig, oxidized, Sweet.|171|16|3|2016-06-07T19:53:53Z|2019-09-17T18:31:17Z|
[utp](https://github.com/anacrolix/utp)|Use anacrolix/go-libutp instead|161|34|4|2015-03-20T04:39:22Z|2021-01-29T09:58:07Z|
[canopus](https://github.com/zubairhamed/canopus)|CoAP Client/Server implementing RFC 7252 for the Go Language|147|41|43|2015-02-24T04:12:20Z|2018-03-25T17:28:53Z|
[sslb](https://github.com/eduardonunesp/sslb)|Golang Super Simple Load Balance|133|26|10|2015-10-18T21:31:09Z|2019-09-24T22:03:37Z|
[xtcp](https://github.com/xfxdev/xtcp)|A TCP Server Framework with graceful shutdown, custom protocol.|130|28|0|2016-03-31T16:50:14Z|2020-02-29T18:57:41Z|
[dnsmonster](https://github.com/mosajjal/dnsmonster)|Passive DNS Capture/Monitoring Framework|108|19|1|2020-02-09T01:10:39Z|2021-12-02T09:50:22Z|
[dhcp6](https://github.com/mdlayher/dhcp6)|Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. MIT Licensed.|71|18|2|2015-05-22T04:13:30Z|2019-03-11T16:24:02Z|
[ether](https://github.com/songgao/ether)|A Go package for sending and receiving ethernet frames. Currently supporting Linux, Freebsd, and OS X.|71|6|0|2014-05-21T03:46:30Z|2016-04-05T03:04:14Z|
[iplib](https://github.com/c-robinson/iplib)|A library  for working with IP addresses and networks in Go|62|8|0|2019-05-06T06:23:41Z|2021-11-02T05:39:49Z|
[packet](https://github.com/aerogo/packet)|:package: Send network packets over a TCP or UDP connection.|59|14|1|2017-10-29T05:46:44Z|2019-11-20T22:35:38Z|
[linkio](https://github.com/ian-kent/linkio)|Simulate network link speed|49|6|0|2014-12-24T10:50:03Z|2017-08-07T20:57:56Z|
[go-powerdns](https://github.com/joeig/go-powerdns)|Go PowerDNS 4.x API Client|45|14|0|2018-06-21T21:37:33Z|2021-11-12T18:57:10Z|
[portproxy](https://github.com/aybabtme/portproxy)|TCP proxy, highjacks HTTP to allow CORS|45|12|0|2014-12-13T02:57:36Z|2014-12-13T03:05:07Z|
[panoptes-stream](https://github.com/yahoo/panoptes-stream)|A cloud native distributed streaming network telemetry.|29|6|1|2020-10-09T04:26:26Z|2021-03-04T03:28:51Z|
[graval](https://github.com/koofr/graval)|An experimental go FTP server framework|27|7|0|2014-04-22T19:17:18Z|2020-10-02T13:42:14Z|
[publicip](https://github.com/polera/publicip)|Go pkg for returning your public facing IP address.|24|7|0|2016-12-28T19:31:07Z|2016-12-29T04:30:29Z|
[golibwireshark](https://github.com/sunwxg/golibwireshark)||21|6|0|2015-11-16T06:48:41Z|2017-10-24T11:56:01Z|
[goshark](https://github.com/sunwxg/goshark)||14|4|0|2015-11-01T07:23:09Z|2017-10-24T11:36:13Z|
[gohooks](https://github.com/averageflow/gohooks)|GoHooks make it easy to send and consume secured web-hooks from a Go application|14|2|0|2020-10-30T17:20:36Z|2021-07-16T09:57:04Z|
[llb](https://github.com/kirillDanshin/llb)||11|2|0|2016-02-21T06:30:17Z|2016-04-04T04:13:06Z|
[tspool](https://github.com/two/tspool)|tcp server pool|10|2|0|2018-10-27T01:05:03Z|2018-10-29T01:55:10Z|
[httpproxy](https://github.com/wzshiming/httpproxy)|HTTP proxy handler and dialer|10|2|1|2018-07-18T09:42:34Z|2021-11-13T08:25:28Z|


### HTTP Clients
*Libraries for making HTTP requests.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[resty](https://github.com/go-resty/resty)|Simple HTTP and REST client library for Go|5331|435|43|2015-08-28T17:48:47Z|2021-12-10T02:07:03Z|
[heimdall](https://github.com/gojek/heimdall)|An enhanced HTTP client for Go|2130|179|39|2018-01-19T09:32:26Z|2021-10-11T23:05:21Z|
[grequests](https://github.com/levigross/grequests)|A Go &#34;clone&#34; of the great and famous Requests library|1844|113|29|2015-06-11T16:41:48Z|2020-12-03T02:31:16Z|
[sling](https://github.com/dghubble/sling)|A Go HTTP client library for creating and sending API requests|1392|106|1|2015-04-02T08:42:52Z|2021-10-28T20:00:12Z|
[gentleman](https://github.com/h2non/gentleman)|Plugin-driven, extensible HTTP client toolkit for Go|926|51|19|2016-02-21T23:00:24Z|2021-02-18T19:34:43Z|
[pester](https://github.com/sethgrid/pester)|Go (golang) http calls with retries and backoff |566|64|4|2015-05-20T13:50:49Z|2020-06-17T17:44:02Z|
[request](https://github.com/monaco-io/request)|go request, go http client|152|20|0|2020-03-25T06:24:18Z|2021-10-03T02:28:50Z|
[rq](https://github.com/ddo/rq)|A nicer interface for golang stdlib HTTP client|38|3|1|2017-12-26T10:48:27Z|2019-08-28T17:45:31Z|
[go-http-client](https://github.com/bozd4g/go-http-client)|An enhanced http client for Golang|32|9|0|2019-12-14T11:22:19Z|2021-05-02T18:35:32Z|
[requests](https://github.com/carlmjohnson/requests)|HTTP requests for Gophers|24|0|0|2021-05-20T19:20:29Z|2021-10-21T13:04:32Z|
[httpretry](https://github.com/ybbus/httpretry)|Enriches the standard go http client with retry functionality.|14|2|0|2020-02-05T10:17:42Z|2020-02-14T08:20:21Z|
[go-req](https://github.com/wenerme/go-req)|Declarative golang HTTP client|12|1|1|2021-07-11T10:42:40Z|2021-09-07T16:14:09Z|
**[ARCHIVED]**  [httpc](https://github.com/valord577/httpc)|A customizable and simple HTTP client library. Only depend on the stdlib HTTP client.|2|0|0|2021-08-11T12:26:27Z|2021-11-22T04:21:25Z|


## OpenGL
*Libraries for using OpenGL in Go.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[glfw](https://github.com/go-gl/glfw)|Go bindings for GLFW 3|1190|146|13|2013-05-19T06:38:45Z|2021-12-11T16:28:36Z|
[gl](https://github.com/go-gl/gl)|Go bindings for OpenGL (generated via glow)|851|64|10|2015-02-22T03:29:45Z|2021-12-10T17:28:15Z|
[mathgl](https://github.com/go-gl/mathgl)|A pure Go 3D math library.|401|52|9|2013-02-13T14:18:55Z|2020-09-18T17:18:08Z|
[gl](https://github.com/goxjs/gl)|Go cross-platform OpenGL bindings.|154|17|9|2015-05-18T08:10:15Z|2021-01-04T18:53:21Z|
[glfw](https://github.com/goxjs/glfw)|Go cross-platform glfw library for creating an OpenGL context and receiving events.|71|16|10|2014-12-27T22:40:24Z|2020-03-30T04:26:18Z|
[go-glmatrix](https://github.com/technohippy/go-glmatrix)|go-glmatrix is a golang version of glMatrix, which is &#34;designed to perform vector and matrix operations stupidly fast&#34;.|2|3|0|2020-07-02T13:40:40Z|2021-02-05T02:33:06Z|


## ORM
*Libraries that implement Object-Relational Mapping or datamapping techniques.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[beego](https://github.com/beego/beego)|beego is an open-source, high-performance web framework for the Go programming language.|27360|5363|26|2012-02-29T02:32:08Z|2021-12-09T14:30:13Z|
[gorm](https://github.com/go-gorm/gorm)|The fantastic ORM library for Golang, aims to be developer friendly|26132|2961|27|2013-10-25T08:31:38Z|2021-12-10T09:51:15Z|
[ent](https://github.com/ent/ent)|An entity framework for Go|9393|509|200|2019-06-12T22:53:55Z|2021-12-11T12:25:02Z|
**[ARCHIVED]**  [xorm](https://github.com/go-xorm/xorm)|Simple and Powerful ORM for Go, support mysql,postgres,tidb,sqlite3,mssql,oracle, Moved to https://gitea.com/xorm/xorm|6465|799|308|2013-05-09T02:35:04Z|2020-04-03T01:12:12Z|
[pg](https://github.com/go-pg/pg)|Golang ORM with focus on PostgreSQL features and performance|4909|365|104|2013-04-24T12:31:41Z|2021-12-01T22:44:18Z|
[sqlboiler](https://github.com/volatiletech/sqlboiler)|Generate a Go ORM tailored to your database schema.|4488|412|111|2016-02-21T06:18:25Z|2021-12-04T11:06:08Z|
[gorp](https://github.com/go-gorp/gorp)|Go Relational Persistence - an ORM-ish library for Go|3533|375|136|2012-01-04T19:50:09Z|2021-03-04T16:05:59Z|
[db](https://github.com/upper/db)|Data access layer for PostgreSQL, CockroachDB, MySQL, SQLite and MongoDB with ORM-like features.|2822|199|127|2013-10-23T02:04:36Z|2021-12-11T13:27:00Z|
[gormt](https://github.com/xxjwxc/gormt)|database to golang struct|1576|258|35|2019-05-05T13:10:26Z|2021-12-08T11:05:42Z|
[reform](https://github.com/go-reform/reform)|A better ORM for Go, based on non-empty interfaces and code generation.|1211|55|74|2016-02-25T09:41:09Z|2021-12-06T13:13:14Z|
[pop](https://github.com/gobuffalo/pop)|A Tasty Treat For All Your Database Needs|1145|224|122|2018-02-07T21:13:46Z|2021-12-10T17:25:51Z|
[prisma-client-go](https://github.com/prisma/prisma-client-go)|Prisma Client Go is an auto-generated and fully type-safe database client|1066|48|88|2019-09-24T12:17:03Z|2021-12-10T13:24:59Z|
[go-sqlbuilder](https://github.com/huandu/go-sqlbuilder)|A flexible and powerful SQL string builder library plus a zero-config ORM.|673|63|1|2017-12-27T16:37:48Z|2021-10-21T10:45:08Z|
[go-queryset](https://github.com/jirfag/go-queryset)|100% type-safe ORM for Go (Golang) with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support. GORM under the hood.|643|63|18|2017-09-03T17:29:30Z|2021-07-18T08:49:34Z|
[rel](https://github.com/go-rel/rel)|:gem: Modern ORM for Golang - Testable, Extendable and Crafted Into a Clean and Elegant API|454|44|17|2019-10-06T07:08:01Z|2021-11-16T16:56:25Z|
[zoom](https://github.com/albrow/zoom)|A blazing-fast datastore and querying engine for Go built on Redis.|283|25|2|2013-07-17T00:32:34Z|2020-05-06T18:52:16Z|
[grimoire](https://github.com/Fs02/grimoire)|Database access layer for golang|156|16|0|2018-03-05T16:52:20Z|2021-10-25T23:52:11Z|
[gosql](https://github.com/rushteam/gosql)|golang orm and sql builder|149|17|2|2020-04-27T09:16:29Z|2021-06-21T07:03:35Z|
[go-store](https://github.com/gosuri/go-store)|A simple and fast Redis backed key-value store library for Go|105|8|1|2015-03-22T12:07:29Z|2017-02-23T15:11:42Z|
[go-firestorm](https://github.com/jschoedt/go-firestorm)|Simple Go ORM for Google/Firebase Cloud Firestore|26|7|0|2018-12-04T14:53:53Z|2021-11-30T13:29:42Z|
[cacheme-go](https://github.com/Yiling-J/cacheme-go)|🚀 Schema based, typed Redis caching/memoize framework for Go|17|0|0|2021-10-03T08:44:28Z|2021-12-06T03:45:10Z|
[lore](https://github.com/abrahambotros/lore)|Light Object-Relational Environment (LORE) provides a simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go|10|2|0|2017-04-29T03:57:15Z|2017-10-21T18:26:41Z|
[marlow](https://github.com/marlow/marlow)|persistence layer code generation for golang|9|2|0|2020-08-11T13:34:00Z|2020-08-18T14:06:35Z|


## Package Management
*Official tooling for dependency and package management*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)|Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.|-|-|-|-|-|


## Performance

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[jaeger](https://github.com/jaegertracing/jaeger)|CNCF Jaeger, a Distributed Tracing Platform|14772|1771|365|2016-04-15T18:49:02Z|2021-12-10T20:30:50Z|
[pixie](https://github.com/pixie-io/pixie)|Instant Kubernetes-Native Application Observability|2619|141|66|2020-02-27T00:22:45Z|2021-12-10T21:13:32Z|
[statsviz](https://github.com/arl/statsviz)|:rocket: Instant live visualization of your Go application runtime statistics (GC, MemStats, etc.) in the browser|1742|59|6|2020-08-14T00:00:41Z|2021-11-16T17:26:16Z|
[profile](https://github.com/pkg/profile)|Simple profiling for Go|1630|113|9|2014-10-22T01:35:18Z|2021-11-02T20:31:47Z|
[tracer](https://github.com/kamilsk/tracer)|🧶 Dead simple, lightweight tracing.|60|3|11|2019-06-22T13:23:27Z|2021-02-27T09:49:34Z|


## Query Language

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[graphql](https://github.com/graphql-go/graphql)|An implementation of GraphQL for Go / Golang|8168|721|176|2015-07-19T12:25:43Z|2021-11-10T02:41:09Z|
[gqlgen](https://github.com/99designs/gqlgen)|go generate based graphql server library|6860|762|227|2018-02-11T04:54:11Z|2021-12-09T18:29:31Z|
[graphql-go](https://github.com/graph-gophers/graphql-go)|GraphQL server with a focus on ease of use|4000|440|86|2016-10-18T13:57:24Z|2021-12-11T20:09:11Z|
[gojsonq](https://github.com/thedevsaddam/gojsonq)|A simple Go package to Query over JSON/YAML/XML/CSV Data |1791|109|17|2018-05-19T16:15:18Z|2021-08-16T08:34:28Z|
[dasel](https://github.com/TomWright/dasel)|Select, put and delete data from JSON, TOML, YAML, XML and CSV files with a single tool. Supports conversion between formats and can be used as a Go package.|1717|40|16|2020-09-22T10:33:56Z|2021-12-08T20:54:41Z|
[jsonql](https://github.com/elgs/jsonql)|JSON query expression library in Golang.|248|37|5|2015-12-29T11:24:04Z|2020-11-20T03:19:00Z|
[rql](https://github.com/a8m/rql)|Resource Query Language for REST|241|30|13|2018-06-05T18:37:29Z|2021-12-04T21:35:55Z|
[jsonslice](https://github.com/bhmj/jsonslice)|json slicer|57|6|3|2018-05-02T00:33:15Z|2021-11-24T23:16:21Z|
[graphql](https://github.com/tmc/graphql)|graphql parser &#43; utilities|54|7|3|2015-04-18T21:05:52Z|2017-06-02T05:21:03Z|
[api-fu](https://github.com/ccbrown/api-fu)|A collection of Go packages for creating robust GraphQL APIs|37|3|2|2019-07-30T05:18:43Z|2021-09-29T21:53:36Z|
[straf](https://github.com/SonicRoshan/straf)|Convert Golang Struct To GraphQL Object On The Fly|32|5|0|2019-08-16T13:31:39Z|2020-05-16T13:22:22Z|
[rest-query-parser](https://github.com/timsolov/rest-query-parser)|Query Parser for REST|28|7|2|2020-02-10T17:58:42Z|2021-11-23T13:57:49Z|
[jsonpath](https://github.com/AsaiYusuke/jsonpath)|A query library for retrieving part of JSON based on JSONPath syntax.|8|1|0|2020-11-29T05:37:26Z|2021-10-30T22:19:48Z|
[gws](https://github.com/Zaba505/gws)|A WebSocket client and server for GraphQL|4|2|2|2020-06-08T19:51:36Z|2020-09-04T06:02:11Z|


## Resource Embedding

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[statik](https://github.com/rakyll/statik)|Embed files into a Go executable|3340|209|33|2014-02-04T14:54:51Z|2020-11-08T13:41:59Z|
[packr](https://github.com/gobuffalo/packr)|The simple and easy way to embed static files into Go binaries.|3328|190|67|2017-03-15T22:24:53Z|2021-12-04T19:53:01Z|
[go.rice](https://github.com/GeertJohan/go.rice)|go.rice is a Go package that makes working with resources such as html,js,css,images,templates, etc very easy.|2265|142|39|2013-10-23T21:29:34Z|2021-10-19T21:45:05Z|
[vfsgen](https://github.com/shurcooL/vfsgen)|Takes an input http.FileSystem (likely at go generate time) and generates Go code that statically implements it.|935|80|32|2015-05-18T13:03:02Z|2020-08-24T05:29:21Z|
[esc](https://github.com/mjibson/esc)|A simple file embedder for Go|605|66|11|2014-01-26T05:08:04Z|2019-11-14T16:22:26Z|
[fileb0x](https://github.com/UnnoTed/fileb0x)|a better customizable tool to embed files in go; also update embedded files remotely without restarting the server|601|52|9|2016-01-23T20:19:33Z|2021-02-14T13:05:35Z|
[go-resources](https://github.com/omeid/go-resources)|Unfancy resources embedding for Go with out of box http.FileSystem support.|173|19|3|2015-02-21T15:40:17Z|2021-05-30T03:53:52Z|
[statics](https://github.com/go-playground/statics)|:file_folder: Embeds static resources into go files for single binary compilation &#43; works with http.FileSystem &#43; symlinks|64|5|0|2015-10-07T11:49:52Z|2016-10-05T01:27:05Z|
[templify](https://github.com/wlbr/templify)|A tool to be used with &#39;go generate&#39; to embed external template files into Go code.|26|5|1|2016-05-22T16:42:47Z|2021-08-16T20:22:50Z|
[rebed](https://github.com/soypat/rebed)|Recreates directory and files from embedded filesystem using Go 1.16 embed.FS type.|19|2|0|2021-02-17T18:19:49Z|2021-03-30T01:47:16Z|
[debme](https://github.com/leaanthony/debme)|embed.FS wrapper providing additional functionality|11|3|0|2021-04-16T00:25:13Z|2021-06-06T02:03:03Z|
[mule](https://github.com/wlbr/mule)|mule is a tool to be used with &#39;go generate&#39; to embed external resources files into Go code.|9|2|1|2020-01-17T10:56:00Z|2021-08-16T20:23:29Z|


## Science and Data Analysis
*Libraries for scientific computing and data analyzing.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[gonum](https://github.com/gonum/gonum)|Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more|5363|429|221|2017-03-25T14:54:38Z|2021-11-24T10:55:01Z|
[stats](https://github.com/montanaflynn/stats)|A well tested and comprehensive Golang statistics library package with no dependencies.|2192|145|16|2014-12-16T03:25:19Z|2021-11-06T19:43:30Z|
[plot](https://github.com/gonum/plot)|A repository for plotting and visualizing data|2052|184|89|2013-07-23T07:01:13Z|2021-09-27T18:02:17Z|
[gosl](https://github.com/cpmech/gosl)|Linear algebra, eigenvalues, FFT, Bessel, elliptic, orthogonal polys, geometry, NURBS, numerical quadrature, 3D transfinite interpolation, random numbers, Mersenne twister, probability distributions, optimisation, differential equations.|1628|144|0|2015-02-09T23:00:38Z|2021-11-11T22:17:12Z|
[streamtools](https://github.com/nytlabs/streamtools)|tools for working with streams of data|1315|110|47|2013-07-05T18:58:45Z|2015-07-17T13:38:10Z|
[go-dsp](https://github.com/mjibson/go-dsp)|Digital Signal Processing for Go|759|77|6|2011-11-02T06:28:41Z|2018-05-08T04:29:40Z|
[chart](https://github.com/vdobler/chart)|Provide basic charts in go|703|104|6|2011-06-27T12:19:42Z|2021-06-03T05:17:13Z|
[dataframe-go](https://github.com/rocketlaunchr/dataframe-go)|DataFrames for Go: For statistics, machine-learning, and data manipulation/exploration|678|62|6|2018-10-01T12:19:31Z|2021-10-25T05:27:10Z|
[goraph](https://github.com/gyuho/goraph)|Package goraph implements graph data structure and algorithms.|650|74|6|2014-02-27T03:15:55Z|2017-10-01T06:05:15Z|
[graph](https://github.com/yourbasic/graph)|Graph algorithms and data structures|498|53|3|2017-04-27T18:43:54Z|2021-09-23T06:27:31Z|
[orb](https://github.com/paulmach/orb)|Types and utilities for working with 2d geometry in Golang|447|61|9|2016-03-28T01:19:01Z|2021-11-12T04:05:47Z|
[ewma](https://github.com/VividCortex/ewma)|Exponentially Weighted Moving Average algorithms for Go.|342|25|4|2013-07-05T21:33:25Z|2021-08-14T11:56:33Z|
[calendarheatmap](https://github.com/nikolaydubina/calendarheatmap)|📅 Calendar heatmap inspired by GitHub contribution activity |331|14|12|2020-07-01T18:30:48Z|2021-10-01T11:28:16Z|
[gohistogram](https://github.com/VividCortex/gohistogram)|Streaming approximate histograms in Go|158|29|1|2013-07-02T12:53:22Z|2020-12-15T17:33:31Z|
[TextRank](https://github.com/DavidBelicza/TextRank)|:wink: :cyclone: :strawberry: TextRank implementation in Golang with extendable features (summarization, phrase extraction) and multithreading (goroutine).|145|19|5|2018-01-09T19:36:17Z|2021-07-08T17:29:28Z|
[sparse](https://github.com/james-bowman/sparse)|Sparse matrix formats for linear algebra supporting scientific and machine learning applications|122|20|4|2017-05-16T13:54:36Z|2021-07-29T09:01:28Z|
[go-estimate](https://github.com/milosgajdos/go-estimate)|State estimation and filtering algorithms in Go|87|7|2|2018-11-04T22:32:52Z|2021-08-21T16:16:55Z|
[pagerank](https://github.com/alixaxel/pagerank)|Weighted PageRank implementation in Go|72|19|3|2015-08-06T01:33:34Z|2021-06-19T22:18:08Z|
[jsonl-graph](https://github.com/nikolaydubina/jsonl-graph)|🏝 JSONL Graph Tools|53|3|4|2021-06-26T06:37:03Z|2021-11-20T15:57:06Z|
[geom](https://github.com/skelterjohn/geom)|2d geometry for golang|48|17|1|2011-06-07T17:49:11Z|2018-01-03T14:24:18Z|
[evaler](https://github.com/soniah/evaler)|Implements a simple floating point arithmetic expression evaluator in Go (golang).|45|13|5|2012-09-04T23:37:58Z|2018-07-27T12:02:52Z|
[goent](https://github.com/kzahedi/goent)|GO Implementation of Entropy Measures|25|2|0|2017-08-08T05:37:12Z|2019-04-03T09:41:55Z|
[triangolatte](https://github.com/tchayen/triangolatte)|2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs.|23|2|4|2018-07-18T21:17:09Z|2021-08-04T11:33:07Z|
[decimal](https://github.com/db47h/decimal)|An arbitrary-precision decimal floating-point arithmetic package for Go|23|1|0|2020-05-27T15:23:59Z|2020-07-06T12:23:53Z|
[piecewiselinear](https://github.com/sgreben/piecewiselinear)|tiny linear interpolation library for go (factored out from https://github.com/sgreben/yeetgif)|20|2|0|2018-10-21T13:19:44Z|2020-12-01T19:30:38Z|
[GoStats](https://github.com/OGFris/GoStats)|GoStats is a go library for math statistics mostly used in ML domains, it covers most of the statistical measures functions.|17|1|0|2018-07-22T20:55:16Z|2019-01-14T16:50:38Z|
[godesim](https://github.com/soypat/godesim)|ODE system solver made simple. For IVPs (initial value problems).|17|0|1|2020-12-16T01:02:26Z|2021-11-27T16:09:34Z|
[PiHex](https://github.com/claygod/PiHex)|PiHex Library, written in Go, generates a hexadecimal number sequence in the number Pi in the range from 0 to 10,000,000.|14|3|0|2016-07-22T11:21:37Z|2020-09-16T19:31:47Z|
[ode](https://github.com/ChristopherRabotin/ode)|An ordinary differential equation solving library in golang.|13|1|1|2016-11-11T22:40:21Z|2017-03-18T01:10:01Z|
[rootfinding](https://github.com/khezen/rootfinding)|root-finding library|6|1|0|2018-10-30T22:31:48Z|2020-03-22T09:14:10Z|
[assocentity](https://github.com/ndabAP/assocentity)|Package assocentity returns the average distance from words to a given entity|5|2|6|2018-12-21T07:17:09Z|2020-10-27T12:49:40Z|
[go-gt](https://github.com/ThePaw/go-gt)|Automatically exported from code.google.com/p/go-gt|5|1|2|2015-09-14T12:05:37Z|2015-09-14T12:08:59Z|
[bradleyterry](https://github.com/seanhagen/bradleyterry)|Package to do Bradley-Terry Model pairwise compairsons|4|1|0|2019-04-30T00:28:13Z|2019-05-02T18:10:35Z|


## Security
*Libraries that are used to help make your application more secure.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[age](https://github.com/FiloSottile/age)|A simple, modern and secure encryption tool (and Go library) with small explicit keys, no config options, and UNIX-style composability.|9306|290|19|2019-05-18T20:44:54Z|2021-12-07T16:14:50Z|
[lego](https://github.com/go-acme/lego)|Let&#39;s Encrypt client and ACME library written in Go|4967|655|127|2015-06-08T00:36:41Z|2021-12-09T17:36:05Z|
[certmagic](https://github.com/caddyserver/certmagic)|Automatic HTTPS for any Go program: fully-managed TLS certificate issuance and renewal|3854|199|2|2018-12-10T03:12:30Z|2021-11-23T16:48:41Z|
[cameradar](https://github.com/Ullaakut/cameradar)|Cameradar hacks its way into RTSP videosurveillance cameras|2770|390|21|2016-05-20T11:35:41Z|2021-11-08T16:58:28Z|
[crypto](https://github.com/golang/crypto)|[mirror] Go supplementary cryptography libraries|2303|1273|50|2014-12-04T04:02:55Z|2021-12-09T19:37:09Z|
[memguard](https://github.com/awnumar/memguard)|Secure software enclave for storage of sensitive information in memory.|2015|93|4|2017-04-22T07:40:40Z|2021-03-16T05:26:30Z|
[acmetool](https://github.com/hlandau/acmetool)|:lock: acmetool, an automatic certificate acquisition tool for ACME (Let&#39;s Encrypt)|1882|123|69|2015-11-15T01:56:02Z|2021-04-01T13:13:57Z|
[secure](https://github.com/unrolled/secure)|HTTP middleware for Go that facilitates some quick security wins.|1854|122|3|2014-05-20T19:46:28Z|2021-11-11T13:13:35Z|
[themis](https://github.com/cossacklabs/themis)|Easy to use cryptographic framework for data protection: secure messaging with forward secrecy and secure data storage. Has unified APIs across 14 platforms.|1380|119|10|2015-05-06T13:25:25Z|2021-12-03T14:13:29Z|
[acra](https://github.com/cossacklabs/acra)|Database security suite. Database proxy with field-level encryption, search through encrypted data, SQL injections prevention, intrusion detection, honeypots. Supports client-side and proxy-side (&#34;transparent&#34;) encryption. SQL, NoSQL.|880|96|1|2016-11-14T16:23:25Z|2021-12-07T18:12:24Z|
[nacl](https://github.com/kevinburke/nacl)|Pure Go implementation of the NaCL set of API&#39;s|517|30|3|2017-07-20T19:07:19Z|2021-04-05T17:38:05Z|
[ssh-vault](https://github.com/ssh-vault/ssh-vault)|🌰  encrypt/decrypt using ssh keys|329|23|8|2016-09-29T14:46:30Z|2021-07-12T08:00:17Z|
[firewalld-rest](https://github.com/prashantgupta24/firewalld-rest)|A rest application to update firewalld rules on a linux server |311|13|2|2020-06-12T20:16:33Z|2020-09-04T18:10:18Z|
[go-password-validator](https://github.com/wagslane/go-password-validator)|Validate the Strength of a Password in Go|310|24|1|2020-10-14T15:52:14Z|2021-04-13T14:19:46Z|
[badactor](https://github.com/jaredfolkins/badactor)|BadActor.org An in-memory application driven jailer written in Go|305|16|1|2014-12-12T20:05:20Z|2020-05-28T22:21:02Z|
[optimus-go](https://github.com/pjebs/optimus-go)|ID hashing and Obfuscation using Knuth&#39;s Algorithm|290|20|0|2015-05-05T10:12:38Z|2020-05-04T00:14:25Z|
[passlib](https://github.com/hlandau/passlib)|:key: Idiotproof golang password validation library inspired by Python&#39;s passlib|257|26|1|2014-12-21T17:45:52Z|2021-03-23T06:03:00Z|
[go-yara](https://github.com/hillu/go-yara)|Go bindings for YARA|238|84|6|2015-01-25T01:01:11Z|2021-11-05T09:22:05Z|
[simple-scrypt](https://github.com/elithrar/simple-scrypt)|A convenience library for generating, comparing and inspecting password hashes using the scrypt KDF in Go 🔑|173|25|4|2015-04-14T06:52:21Z|2021-04-12T20:33:15Z|
[argon2pw](https://github.com/raja/argon2pw)|Argon2 password hashing package for go with constant time hash comparison|88|8|0|2018-03-13T13:56:36Z|2021-09-10T18:37:55Z|
[dongle](https://github.com/golang-module/dongle)|A simple, semantic and developer-friendly golang package for encoding&amp;decoding and encryption&amp;decryption|66|3|1|2021-08-11T07:11:54Z|2021-08-30T03:09:57Z|
[goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword)|A probably paranoid Golang utility library for securely hashing and encrypting passwords based on the Dropbox method. This implementation uses Blake2b, Scrypt and XSalsa20-Poly1305 (via NaCl SecretBox) to create secure password hashes that are also encrypted using a master passphrase.|44|7|0|2017-10-19T19:34:45Z|2020-12-11T04:22:56Z|
[go-generate-password](https://github.com/m1/go-generate-password)|Password generator written in Go|29|4|0|2019-11-14T17:57:19Z|2021-08-14T10:24:52Z|
[certificates](https://github.com/mvmaasakkers/certificates)|An opinionated helper for generating tls certificates|21|4|0|2019-03-04T07:20:36Z|2020-12-09T19:49:59Z|
[secureio](https://github.com/xaionaro-go/secureio)|An easy-to-use XChaCha20-encryption wrapper for io.ReadWriteCloser (even lossy UDP) using ECDH key exchange algorithm, ED25519 signatures and Blake3&#43;Poly1305 checksums/message-authentication for Go (golang). Also a multiplexer.|21|2|1|2018-12-25T14:20:59Z|2020-06-28T16:32:59Z|
[go-htpasswd](https://github.com/tg123/go-htpasswd)|Apache htpasswd Parser for Go.|18|9|0|2015-06-18T06:50:27Z|2021-10-20T22:22:00Z|
[argon2-hashing](https://github.com/andskur/argon2-hashing)|A light package for generating and comparing password hashing with argon2 in Go|14|2|0|2019-01-02T20:41:02Z|2020-04-05T22:12:45Z|
[goArgonPass](https://github.com/dwin/goArgonPass)|goArgonPass is a Argon2 Password utility package for Go using the crypto library package Argon2 designed to be compatible with Passlib for Python and Argon2 PHP. Argon2 was the winner of the most recent Password Hashing Competition. This is designed for use anywhere password hashing and verification might be needed and is intended to replace implementations using bcrypt or Scrypt.|13|6|1|2018-05-30T01:32:10Z|2020-12-11T04:07:56Z|
[sslmgr](https://github.com/adrianosela/sslmgr)|A layer of abstraction the around acme/autocert certificate manager (Golang)|11|3|0|2019-04-02T17:35:38Z|2019-07-27T18:49:03Z|
[Interpol](https://bitbucket.org/vahidi/interpol)|Rule-based data generator for fuzzing and penetration testing.|-|-|-|-|-|


## Serialization
*Libraries and tools for binary serialization.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[go](https://github.com/json-iterator/go)|A high-performance 100% compatible drop-in replacement of &#34;encoding/json&#34;|10204|826|173|2016-11-30T00:30:24Z|2021-11-20T16:00:40Z|
[protobuf](https://github.com/golang/protobuf)|Go support for Google&#39;s protocol buffers|8105|1488|63|2014-11-23T23:07:23Z|2021-09-16T07:58:48Z|
[mapstructure](https://github.com/mitchellh/mapstructure)|Go library for decoding generic map values into native Go structures and vice versa.|5160|528|37|2013-05-20T05:24:34Z|2021-12-01T18:06:27Z|
[protobuf](https://github.com/gogo/protobuf)|[Looking for new ownership] Protocol Buffers for Go with Gadgets|4986|669|216|2014-12-03T11:27:10Z|2021-11-02T08:13:40Z|
[go](https://github.com/ugorji/go)|idiomatic codec and rpc lib for msgpack, cbor, json, etc. msgpack.org[Go]|1607|268|2|2013-05-30T02:13:13Z|2021-11-06T18:45:48Z|
[colfer](https://github.com/pascaldekloe/colfer)|binary serialization format|640|49|12|2015-09-05T16:42:41Z|2021-09-06T12:24:35Z|
[csvutil](https://github.com/jszwec/csvutil)|csvutil provides fast and idiomatic mapping between CSV and Go (golang) values.|630|41|1|2017-10-30T04:09:48Z|2021-11-25T20:35:16Z|
[cbor](https://github.com/fxamacker/cbor)|CBOR codec (in Go) with CBOR tags, Go struct tags (toarray/keyasint/omitempty), float64/32/16, big.Int, and fuzz tested billions of execs for reliable RFC 7049 &amp; RFC 8949. |353|32|15|2019-05-15T21:22:15Z|2021-12-06T03:40:05Z|
[go-capnproto](https://github.com/glycerine/go-capnproto)|Cap&#39;n Proto library and parser for go. This is go-capnproto-1.0, and does not have rpc. See https://github.com/zombiezen/go-capnproto2 for 2.0 which has rpc and capabilities.|279|19|1|2013-11-07T20:28:12Z|2020-01-29T18:25:38Z|
[php_session_decoder](https://github.com/yvasiyarov/php_session_decoder)|PHP session encoder/decoder written in Go|151|42|3|2012-12-23T14:04:25Z|2018-11-02T07:23:13Z|
[structomap](https://github.com/danhper/structomap)|Easily and dynamically generate maps from Go static structures|130|10|0|2015-05-13T08:54:11Z|2019-05-24T14:07:40Z|
[bambam](https://github.com/glycerine/bambam)|auto-generate capnproto schema from your golang source files. Depends on go-capnproto-1.0 at https://github.com/glycerine/go-capnproto|62|11|3|2014-09-17T14:39:12Z|2016-10-07T18:28:00Z|
**[ARCHIVED]**  [asn1](https://github.com/Logicalis/asn1)|Asn.1 BER and DER encoding library for golang.|47|23|6|2016-02-29T13:00:25Z|2019-03-12T17:35:41Z|
[binstruct](https://github.com/ghostiam/binstruct)|Golang binary decoder for mapping data into the structure|36|8|0|2018-10-23T15:42:22Z|2021-09-18T08:48:53Z|
[bel](https://github.com/csweichel/bel)|Generate TypeScript interfaces from Go structs/interfaces - useful for JSON RPC|16|5|2|2019-02-20T20:48:24Z|2020-08-05T08:59:23Z|
[fwencoder](https://github.com/o1egl/fwencoder)|Fixed width file parser (encoder/decoder) in GO (golang)|16|4|0|2017-12-25T12:55:29Z|2020-02-13T14:05:52Z|
[elastic](https://github.com/epiclabs-io/elastic)|Converts go types no matter what|15|2|1|2020-02-25T19:55:00Z|2021-05-21T12:32:58Z|
[pletter](https://github.com/vimeda/pletter)|A standard way to wrap a proto message|14|2|3|2019-07-09T14:02:08Z|2021-09-29T11:56:45Z|
[fixedwidth](https://github.com/huydang284/fixedwidth)|A Go package for encode/decode fixed-width data|5|1|0|2019-08-11T03:42:24Z|2019-12-20T03:18:01Z|
[unitpacking](https://github.com/recolude/unitpacking)|A library for storing unit vectors in a representation that lends itself to saving space on disk.|3|0|0|2021-01-17T22:31:41Z|2021-04-17T17:32:33Z|
[go-lctree](https://github.com/sbourlon/go-lctree)|go-lctree provides a CLI and Go primitives to serialize and deserialize LeetCode binary trees (e.g. &#34;[5,4,7,3,null,2,null,-1,null,9]&#34;).|2|1|0|2020-05-04T05:39:46Z|2020-06-03T21:19:42Z|


## Server Applications

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[etcd](https://github.com/etcd-io/etcd)|Distributed reliable key-value store for the most critical data of a distributed system|38156|8185|192|2013-07-06T21:57:21Z|2021-12-10T08:21:35Z|
[caddy](https://github.com/caddyserver/caddy)|Fast, multi-platform web server with automatic HTTPS|35784|2865|98|2015-01-13T19:45:03Z|2021-12-11T14:26:23Z|
[minio](https://github.com/minio/minio)|High Performance, Kubernetes Native Object Storage|30621|3475|28|2015-01-14T19:23:58Z|2021-12-11T19:59:47Z|
[consul](https://github.com/hashicorp/consul)|Consul is a distributed, highly available, and data center aware solution to connect and configure applications across dynamic, distributed infrastructure.|23674|3927|1038|2013-11-04T22:15:27Z|2021-12-11T06:46:00Z|
[roadrunner](https://github.com/spiral/roadrunner)|🤯 High-performance PHP application server, load-balancer and process manager written in Golang|6111|334|8|2017-12-26T16:13:10Z|2021-12-09T11:03:54Z|
[easegress](https://github.com/megaease/easegress)|A Cloud Native traffic orchestration system|3997|332|41|2021-05-28T03:02:42Z|2021-12-10T08:19:55Z|
[sftpgo](https://github.com/drakkan/sftpgo)|Fully featured and highly configurable SFTP server with optional HTTP, FTP/S and WebDAV support - S3, Google Cloud Storage, Azure Blob|3498|280|8|2019-07-20T10:18:31Z|2021-12-11T10:09:53Z|
[devd](https://github.com/cortesi/devd)| A local webserver for developers|3183|141|22|2015-09-27T22:43:00Z|2021-08-19T16:52:00Z|
[algernon](https://github.com/xyproto/algernon)|:tophat: Small self-contained pure-Go web server with Lua, Markdown, HTTP/2, QUIC, Redis and PostgreSQL support|1879|105|9|2015-03-10T11:25:30Z|2021-12-06T11:05:42Z|
[flagr](https://github.com/checkr/flagr)|Flagr is a feature flagging, A/B testing and dynamic configuration microservice|1775|148|71|2017-10-03T19:07:32Z|2021-10-05T22:13:06Z|
[fider](https://github.com/getfider/fider)|Open platform to collect and prioritize feedback|1768|511|34|2017-01-17T22:55:19Z|2021-12-09T07:02:24Z|
[flipt](https://github.com/markphelps/flipt)|An open-source, on-prem feature flag solution|1619|80|11|2016-11-05T00:09:07Z|2021-12-08T17:39:47Z|
[trickster](https://github.com/trickstercache/trickster)|Open Source HTTP Reverse Proxy Cache and Time Series Dashboard Accelerator|1588|152|25|2018-03-29T20:31:44Z|2021-11-13T02:08:25Z|
[discovery](https://github.com/bilibili/discovery)|A registry for resilient mid-tier load balancing and failover.|1554|358|23|2018-04-20T12:57:50Z|2021-11-16T10:34:44Z|
[jackal](https://github.com/ortuman/jackal)|💬 Instant messaging server for the Extensible Messaging and Presence Protocol (XMPP).|1164|102|17|2017-11-13T18:17:48Z|2021-12-10T14:37:20Z|
[go-feature-flag](https://github.com/thomaspoignant/go-feature-flag)|A simple and complete feature flag solution, without any complex backend system to install, all you need is a file as your backend. 🎛️|379|15|8|2020-12-11T13:19:17Z|2021-12-01T08:23:35Z|
[euterpe](https://github.com/ironsmile/euterpe)|Self-hosted music streaming server 🎶 with RESTful API and Web interface. Think of it as your very own Spotify!|354|15|13|2014-01-01T12:51:54Z|2021-10-26T20:19:11Z|
[dudeldu](https://github.com/krotik/dudeldu)|A simple SHOUTcast server.|130|13|0|2016-09-07T19:11:04Z|2019-09-22T09:17:43Z|
[lets-proxy2](https://github.com/rekby/lets-proxy2)|Reverse proxy with automatically obtains TLS certificates from Let&#39;s Encrypt|59|9|31|2019-04-12T05:39:43Z|2021-12-03T07:07:48Z|
[psql-streamer](https://github.com/blind-oracle/psql-streamer)|Stream database events from PostgreSQL to Kafka|33|6|2|2019-04-28T21:55:31Z|2020-03-10T09:59:38Z|
[go-proxy-cache](https://github.com/fabiocicerchia/go-proxy-cache)|Simple Reverse Proxy with Caching, written in Go, using Redis.|32|7|23|2020-11-12T15:10:40Z|2021-12-10T15:04:53Z|
[cortex-tenant](https://github.com/blind-oracle/cortex-tenant)|Prometheus remote write proxy that adds Cortex tenant ID based on metric labels|28|12|1|2020-10-06T16:52:25Z|2021-11-16T08:12:00Z|
[nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus)|Turn Nginx logs into Prometheus metrics|27|4|0|2018-10-23T09:10:27Z|2020-09-16T09:07:15Z|
[simple-jwt-provider](https://github.com/leberKleber/simple-jwt-provider)||19|3|3|2019-12-18T12:48:14Z|2021-10-06T21:28:50Z|
[protoxy](https://github.com/camgraff/protoxy)|A proxy server than converts JSON request bodies to protocol buffers|18|2|0|2020-09-03T23:24:34Z|2020-11-08T21:25:43Z|
[moxy](https://github.com/sinhashubham95/moxy)|Mocker &#43; Proxy Application|2|0|0|2021-07-17T05:21:41Z|2021-08-12T17:30:51Z|
[riemann-relay](https://github.com/blind-oracle/riemann-relay)|Service for relaying Riemann events to Riemann/Carbon destinations|0|1|0|2019-04-23T14:17:12Z|2019-10-29T15:00:14Z|
[nsq](https://nsq.io/)|A realtime distributed messaging platform.|-|-|-|-|-|


## Stream Processing
*Libraries and tools for stream processing and reactive programming.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[go-streams](https://github.com/reugn/go-streams)|A lightweight stream processing library for Go|802|63|2|2019-04-30T17:28:15Z|2021-11-27T15:43:06Z|
[machine](https://github.com/whitaker-io/machine)|Machine is a workflow/pipeline library for processing data|100|7|5|2020-10-13T04:24:19Z|2021-12-10T19:23:46Z|
[stream](https://github.com/youthlin/stream)|Go Stream, like Java 8 Stream.|49|5|0|2020-11-12T03:52:50Z|2020-12-08T03:14:39Z|


## Template Engines
*Libraries and tools for templating and lexing.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
**[ARCHIVED]**  [gofpdf](https://github.com/jung-kurt/gofpdf)|A PDF document generator with high level support for text, drawing and images|3907|649|56|2015-03-13T11:57:30Z|2021-11-13T13:53:41Z|
[sprig](https://github.com/Masterminds/sprig)|Useful template functions for Go templates.|2695|301|85|2013-11-22T01:20:40Z|2021-11-18T11:40:13Z|
[quicktemplate](https://github.com/valyala/quicktemplate)|Fast, powerful, yet easy to use template engine for Go. Optimized for speed, zero memory allocations in hot paths. Up to 20x faster than html/template|2283|130|29|2016-03-06T21:42:01Z|2021-09-15T06:31:36Z|
[pongo2](https://github.com/flosch/pongo2)|Django-syntax like template-engine for Go|2105|202|62|2013-08-23T01:00:08Z|2021-12-10T13:53:35Z|
[hero](https://github.com/shiyanhui/hero)|A handy, fast and powerful go template engine.|1484|102|27|2017-01-15T13:31:50Z|2020-01-09T01:41:22Z|
[mustache](https://github.com/hoisie/mustache)|The mustache template language in Go|1027|195|32|2009-12-30T21:05:05Z|2021-12-07T17:55:51Z|
[amber](https://github.com/eknkc/amber)|Amber is an elegant templating engine for Go Programming Language, inspired from HAML and Jade|885|59|23|2012-10-31T20:27:24Z|2020-10-13T09:28:15Z|
[jet](https://github.com/CloudyKit/jet)|Jet  template engine|855|84|18|2016-03-31T16:53:36Z|2021-10-27T04:41:38Z|
[ace](https://github.com/yosssi/ace)|HTML template engine for Go|805|47|29|2014-07-13T13:39:19Z|2018-06-17T06:37:00Z|
[gorazor](https://github.com/sipin/gorazor)|Razor view engine for go|790|88|2|2014-05-01T05:30:31Z|2020-11-24T14:24:29Z|
[fasttemplate](https://github.com/valyala/fasttemplate)|Simple and fast template engine for Go|552|69|9|2015-08-19T12:44:22Z|2021-01-11T18:21:27Z|
[ego](https://github.com/benbjohnson/ego)|An ERB-style templating language for Go.|510|37|11|2014-02-23T18:14:41Z|2021-11-22T14:54:10Z|
[maroto](https://github.com/johnfercher/maroto)|A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple.|488|80|20|2019-05-20T23:27:47Z|2021-10-25T19:53:39Z|
[raymond](https://github.com/aymerick/raymond)|Handlebars for golang|452|65|19|2015-04-22T13:07:59Z|2021-11-05T10:39:38Z|
[goview](https://github.com/foolin/goview)|Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application.|255|28|10|2019-04-14T11:22:41Z|2020-12-30T02:16:57Z|
[soy](https://github.com/robfig/soy)|Go implementation for Soy templates (Google Closure templates)|157|39|6|2013-12-15T01:14:48Z|2021-09-28T19:46:32Z|
[liquid](https://github.com/osteele/liquid)|A Liquid template engine in Go|142|35|19|2017-06-26T14:39:52Z|2021-10-14T13:45:53Z|
[velvet](https://github.com/gobuffalo/velvet)|A sweet velvety templating package|73|8|2|2016-12-29T16:46:57Z|2017-03-20T14:41:20Z|
[kasia.go](https://github.com/ziutek/kasia.go)|Templating system for HTML and other text documents - go implementation|72|8|2|2010-12-07T10:46:05Z|2015-08-31T17:06:10Z|
[extemplate](https://github.com/dannyvankooten/extemplate)|Wrapper package for Go&#39;s template/html to allow for easy file-based template inheritance.|41|11|1|2018-08-10T20:34:19Z|2021-06-15T11:58:56Z|
[gospin](https://github.com/m1/gospin)|Article spinning and spintax/spinning syntax engine written in Go, useful for A/B, testing pieces of text/articles and creating more natural conversations|28|7|3|2019-02-22T17:04:51Z|2021-05-12T09:29:11Z|
[damsel](https://github.com/dskinner/damsel)|Package damsel provides html outlining via css-selectors and common template functionality.|24|5|1|2012-05-02T23:06:48Z|2016-04-07T02:54:55Z|
[tbd](https://github.com/lucasepe/tbd)|&#34;to be defined&#34; - a really simple way to create text templates with placeholders|16|0|0|2021-05-21T13:11:33Z|2021-08-29T07:51:06Z|


## Testing
*Libraries for testing codebases and generating test data.*
	

### Testing Frameworks



	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[testify](https://github.com/stretchr/testify)|A toolkit with common assertions and mocks that plays nicely with the standard library|14971|1167|302|2012-10-16T16:43:17Z|2021-12-11T15:35:07Z|
[goconvey](https://github.com/smartystreets/goconvey)|Go testing in the browser. Integrates with `go test`. Write behavioral tests in Go.|6852|490|140|2013-08-21T04:52:28Z|2021-11-01T23:20:26Z|
[go-cmp](https://github.com/google/go-cmp)|Package for comparing Go values in tests|2606|159|14|2017-07-07T19:28:22Z|2021-12-07T21:01:45Z|
[httpexpect](https://github.com/gavv/httpexpect)|End-to-end HTTP and REST API testing for Go.|1818|142|10|2016-04-29T17:05:20Z|2021-08-20T11:24:48Z|
[godog](https://github.com/cucumber/godog)|Cucumber for golang|1513|164|47|2015-06-10T13:16:31Z|2021-12-03T21:37:27Z|
[is](https://github.com/matryer/is)|Professional lightweight testing mini-framework for Go.|1278|49|5|2016-12-06T13:24:01Z|2021-09-14T08:39:23Z|
[goblin](https://github.com/franela/goblin)|Minimal and Beautiful Go testing framework|817|74|19|2013-09-19T02:34:24Z|2021-10-03T14:34:22Z|
[go-vcr](https://github.com/dnaeon/go-vcr)|Record and replay your HTTP interactions for fast, deterministic and accurate tests|796|53|4|2015-12-14T12:52:17Z|2021-10-11T14:28:09Z|
[testfixtures](https://github.com/go-testfixtures/testfixtures)|Ruby on Rails like test fixtures for Go. Write tests against a real database|723|54|18|2016-04-05T11:33:28Z|2021-12-05T09:29:57Z|
[baloo](https://github.com/h2non/baloo)|Expressive end-to-end HTTP API testing made easy in Go|717|29|8|2016-05-29T21:40:58Z|2019-01-09T23:32:04Z|
[gnomock](https://github.com/orlangure/gnomock)|Test your code without writing mocks with ephemeral Docker containers 📦 Setup popular services with just a couple lines of code ⏱️ No bash, no yaml, only code 💻|601|30|13|2020-01-31T14:50:52Z|2021-12-09T13:36:17Z|
[go-mutesting](https://github.com/zimmski/go-mutesting)|Mutation testing for Go source code|492|38|36|2014-12-26T22:23:44Z|2021-12-10T15:54:07Z|
[goc](https://github.com/qiniu/goc)|A Comprehensive Coverage Testing System for The Go Programming Language|458|66|23|2020-05-07T03:46:25Z|2021-11-29T06:55:36Z|
[gofight](https://github.com/appleboy/gofight)|Testing API Handler written in Golang.|397|40|6|2016-03-29T00:13:21Z|2021-06-27T15:34:44Z|
[embedded-postgres](https://github.com/fergusstrange/embedded-postgres)|Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test|310|26|3|2019-11-16T23:49:40Z|2021-12-10T09:13:24Z|
[testza](https://github.com/MarvinJWendt/testza)|Full-featured test framework for Go! Assertions, mocking, input testing, output capturing, and much more! 🍕|309|12|8|2021-07-05T16:21:38Z|2021-11-19T14:25:06Z|
[frisby](https://github.com/hofstadter-io/frisby)|API testing framework inspired by frisby-js|271|28|13|2015-09-15T14:35:58Z|2020-03-03T23:49:00Z|
[gotest.tools](https://github.com/gotestyourself/gotest.tools)|A collection of packages to augment the go testing package and support common patterns.|260|35|22|2017-08-08T21:28:54Z|2021-11-16T17:30:11Z|
[go-testdeep](https://github.com/maxatome/go-testdeep)|Extremely flexible golang deep comparison, extends the go testing package, tests HTTP APIs and provides tests suite|235|10|2|2018-05-26T15:03:28Z|2021-11-08T11:51:43Z|
[go-carpet](https://github.com/msoap/go-carpet)|go-carpet - show test coverage in terminal for Go source files|217|8|2|2016-02-28T12:02:51Z|2021-11-13T22:39:46Z|
[endly](https://github.com/viant/endly)|End to end functional test and automation framework|200|24|0|2017-08-28T20:24:43Z|2021-12-10T21:33:35Z|
[charlatan](https://github.com/percolate/charlatan)|Go Interface Mocking Tool|194|9|2|2017-10-06T21:55:14Z|2019-09-05T21:25:40Z|
[cupaloy](https://github.com/bradleyjkemp/cupaloy)|Simple Go snapshot testing|193|27|10|2017-08-07T18:30:05Z|2021-11-12T11:53:41Z|
[commander](https://github.com/commander-cli/commander)|Test your command line interfaces on windows, linux and osx and nodes viá ssh and docker|191|13|32|2019-02-22T16:35:16Z|2021-12-03T15:15:04Z|
[dbcleaner](https://github.com/khaiql/dbcleaner)|Clean database for testing, inspired by database_cleaner for Ruby|133|12|0|2017-01-17T18:18:40Z|2021-11-10T01:57:55Z|
[gospec](https://github.com/luontola/gospec)|Testing framework for Go. Allows writing self-documenting tests/specifications, and executes them concurrently and safely isolated. [UNMAINTAINED]|115|17|3|2009-11-24T13:59:26Z|2014-07-31T18:59:25Z|
[wstest](https://github.com/posener/wstest)|go websocket client for unit testing of a websocket handler|87|13|1|2017-03-31T21:06:18Z|2020-12-30T21:32:28Z|
[gocrest](https://github.com/corbym/gocrest)|GoCrest - Hamcrest-like matchers for Go|81|5|2|2017-12-23T23:27:00Z|2020-12-21T16:05:30Z|
[testcase](https://github.com/adamluzsi/testcase)|testcase is an opinionated testing framework based on BDD principles.|80|5|0|2019-04-22T21:20:51Z|2021-11-16T01:11:22Z|
[jsonassert](https://github.com/kinbiko/jsonassert)|A Go test assertion library for verifying that two representations of JSON are semantically equal|71|14|3|2018-10-26T20:31:01Z|2021-12-11T12:23:28Z|
[go-hit](https://github.com/Eun/go-hit)|http integration test framework|56|3|7|2019-06-04T16:28:23Z|2021-12-06T04:06:25Z|
[restit](https://github.com/go-restit/restit)|A Go library help testing your RESTful API application|55|6|4|2014-06-25T10:25:46Z|2019-10-18T03:18:17Z|
[gospecify](https://github.com/stesla/gospecify)|A BDD library for Go|52|7|1|2009-11-20T06:34:29Z|2011-10-18T02:38:16Z|
[covergates](https://github.com/covergates/covergates)|The portal gates to coverage reports|46|9|9|2020-05-29T04:02:01Z|2021-01-06T05:19:11Z|
[gomatch](https://github.com/jfilipczyk/gomatch)|Library created for testing JSON against patterns.|41|4|0|2019-01-27T20:19:06Z|2021-01-15T13:14:48Z|
[dsunit](https://github.com/viant/dsunit)|Datastore Testibility|39|7|0|2016-06-13T20:20:52Z|2020-02-04T18:38:22Z|
[assert](https://github.com/go-playground/assert)|:exclamation:Basic Assertion Library used along side native go testing, with building blocks for custom assertions|35|12|2|2015-07-20T17:53:45Z|2020-11-04T12:21:01Z|
[hamcrest](https://github.com/rdrdr/hamcrest)|Hamcrest matchers for the Go programming language|27|5|2|2010-12-22T04:49:44Z|2021-01-07T21:29:48Z|
[schema](https://github.com/jgroeneveld/schema)|Quick and easy expression matching for JSON schemas used in requests and responses|15|1|0|2015-08-13T13:36:54Z|2019-10-13T10:57:48Z|
[flute](https://github.com/suzuki-shunsuke/flute)|Golang HTTP client testing framework|15|1|2|2019-07-06T04:32:03Z|2021-12-09T21:58:16Z|
[gogiven](https://github.com/corbym/gogiven)|gogiven - BDD testing framework for go that generates readable output directly from source code|11|3|4|2017-12-31T22:33:37Z|2021-07-28T06:23:41Z|
[testsql](https://github.com/zhulongcheng/testsql)|Generate test data from SQL files before testing and clear it after finished.|11|2|3|2018-09-22T12:03:50Z|2019-09-26T07:23:40Z|
[gosuite](https://github.com/pavlo/gosuite)|Test suites support for standard Go1.7 &#34;testing&#34; by leveraging Subtests feature|10|4|1|2016-10-15T19:28:14Z|2016-10-18T16:53:21Z|
[biff](https://github.com/fulldump/biff)|Bifurcation Framework for testing and use cases|9|2|0|2018-03-28T18:35:53Z|2021-07-18T09:38:46Z|
[badio](https://github.com/cavaliergopher/badio)|Extensions to Go&#39;s testing/iotest package|9|2|0|2016-02-11T10:29:25Z|2016-02-13T15:00:58Z|
[stop-and-go](https://github.com/elgohr/stop-and-go)|Testing helper for concurrency|6|4|0|2020-11-06T09:04:58Z|2021-08-28T20:30:46Z|
[test](https://github.com/tvastar/test)|test utilities for golang|6|1|0|2019-03-23T21:47:36Z|2019-09-23T01:09:27Z|
[trial](https://github.com/jgroeneveld/trial)|A simple assertion library for go|5|1|0|2015-06-18T09:01:30Z|2019-10-13T10:54:15Z|
[tt](https://github.com/vcaesar/tt)|Simple and colorful test tools|4|1|0|2018-04-03T11:47:21Z|2021-09-24T15:54:19Z|
[go-testpredicate](https://github.com/maargenton/go-testpredicate)|Unit-testing predicates for Go.|3|0|0|2018-11-23T21:39:11Z|2021-11-20T03:04:15Z|
[fixenv](https://github.com/rekby/fixenv)||2|0|0|2021-08-27T22:33:04Z|2021-09-10T02:57:46Z|
[omg.testingtools](https://github.com/dedalqq/omg.testingtools)|This tool can be useful for writing a tests. If you want change private field in struct from imported libraries than it can help you.|1|0|0|2021-10-13T13:49:30Z|2021-10-14T23:05:20Z|
[apitest](https://apitest.dev)|Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.|-|-|-|-|-|
[gomega](https://onsi.github.io/gomega/)|Rspec like matcher/assertion library.|-|-|-|-|-|
[gocheck](https://labix.org/gocheck)|More advanced testing framework alternative to gotest.|-|-|-|-|-|
[ginkgo](https://onsi.github.io/ginkgo/)|BDD Testing Framework for Go.|-|-|-|-|-|


### Mock



	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[mock](https://github.com/golang/mock)|GoMock is a mocking framework for the Go programming language.|6512|474|43|2015-06-12T17:15:11Z|2021-12-10T11:21:55Z|
[go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)|Sql mock driver for golang to test database interactions|3982|312|56|2014-02-07T07:59:29Z|2021-12-08T03:31:00Z|
[mockery](https://github.com/vektra/mockery)|A mock code autogenerator for Golang|3114|254|70|2014-09-02T16:49:01Z|2021-11-17T13:39:37Z|
[hoverfly](https://github.com/SpectoLabs/hoverfly)|Lightweight service virtualization/API simulation tool for developers and testers|1813|175|33|2015-11-30T16:36:31Z|2021-11-23T20:56:37Z|
[gock](https://github.com/h2non/gock)|HTTP traffic mocking and testing made easy in Go ༼ʘ̚ل͜ʘ̚༽|1532|80|29|2016-03-02T16:20:26Z|2021-08-03T15:24:55Z|
[httpmock](https://github.com/jarcoal/httpmock)|HTTP mocking for Golang|1259|88|3|2014-02-24T16:47:59Z|2021-11-05T07:35:28Z|
[counterfeiter](https://github.com/maxbrunsfeld/counterfeiter)|A tool for generating self-contained, type-safe test doubles in go|599|72|16|2014-05-21T00:12:54Z|2021-12-06T14:57:57Z|
[minimock](https://github.com/gojuno/minimock)|Powerful mock generation tool for Go programming language|435|28|12|2016-08-03T16:01:35Z|2021-09-22T20:55:37Z|
[go-txdb](https://github.com/DATA-DOG/go-txdb)|Immutable transaction isolated sql driver for golang|398|34|5|2015-07-08T07:34:53Z|2021-10-27T07:09:30Z|
[govcr](https://github.com/seborama/govcr)|HTTP mock for Golang: record and replay HTTP/HTTPS interactions for offline testing|99|14|4|2016-07-10T17:47:41Z|2019-09-24T07:17:55Z|
[timex](https://github.com/cabify/timex)|A test-friendly replacement for golang&#39;s time package|58|3|0|2020-01-02T18:06:48Z|2020-08-03T08:54:37Z|
[go-localstack](https://github.com/elgohr/go-localstack)|Go Wrapper for using localstack|25|5|0|2020-03-18T07:13:02Z|2021-12-10T19:33:41Z|
[mockhttp](https://github.com/tv42/mockhttp)|Mock object for Go http.ResponseWriter|22|6|0|2011-06-11T16:03:01Z|2014-10-29T22:14:22Z|
[mockit](https://github.com/pasdam/mockit)|Library that make mocking of Go functions/methods easy|8|3|1|2020-01-01T08:46:09Z|2021-09-22T18:58:20Z|
[genmock](https://gitlab.com/so_literate/genmock)|Go mocking system with code generator for building calls of the interface methods.|-|-|-|-|-|


### Fuzzing and delta-debugging/reducing/shrinking.



	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[go-fuzz](https://github.com/dvyukov/go-fuzz)|Randomized testing for Go|4215|249|55|2015-04-15T13:07:50Z|2021-12-02T20:44:49Z|
[gofuzz](https://github.com/google/gofuzz)|Fuzz testing for go.|1197|112|12|2014-07-31T16:21:29Z|2021-09-04T11:39:44Z|
[tavor](https://github.com/zimmski/tavor)|A generic fuzzing and delta-debugging framework|231|10|53|2014-05-18T14:59:14Z|2018-10-31T19:43:32Z|


### Selenium and browser control tools.



	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[chromedp](https://github.com/chromedp/chromedp)|A faster, simpler way to drive browsers supporting the Chrome DevTools Protocol.|7036|586|52|2017-01-24T14:54:30Z|2021-12-09T08:41:40Z|
[selenoid](https://github.com/aerokube/selenoid)|Selenium Hub successor running browsers within containers. Scalable, immutable, self hosted Selenium-Grid on any platform with single binary.|2053|268|177|2016-08-22T09:11:16Z|2021-11-21T05:48:29Z|
[rod](https://github.com/go-rod/rod)|A Devtools driver for web automation and scraping|1973|132|75|2020-01-21T20:09:45Z|2021-11-26T06:36:50Z|
[cdp](https://github.com/mafredri/cdp)|Package cdp provides type-safe bindings for the Chrome DevTools Protocol (CDP), written in the Go programming language.|586|41|12|2017-03-12T10:25:41Z|2021-07-17T10:41:56Z|
[playwright-go](https://github.com/mxschmitt/playwright-go)|Playwright for Go a browser automation library to control Chromium, Firefox and WebKit with a single API.|577|58|15|2020-08-16T12:46:14Z|2021-11-30T09:18:13Z|
[ggr](https://github.com/aerokube/ggr)|A lightweight load balancer used to create big Selenium clusters|276|58|16|2016-06-16T15:33:24Z|2021-11-02T22:09:37Z|


### Fail injection



	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[failpoint](https://github.com/pingcap/failpoint)|An implementation of failpoints for Golang.|649|58|11|2019-04-02T07:48:18Z|2021-09-21T02:55:16Z|


## Text Processing
*Libraries for parsing and manipulating texts.*
	

### Specific Formats



	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[colly](https://github.com/gocolly/colly)|Elegant Scraper and Crawler Framework for Golang|15498|1329|139|2017-09-29T14:08:49Z|2021-10-17T18:23:56Z|
[goquery](https://github.com/PuerkitoBio/goquery)|A little like that j-thing, only in Go.|10884|818|4|2012-08-29T02:14:59Z|2021-10-26T00:31:08Z|
[blackfriday](https://github.com/russross/blackfriday)|Blackfriday: a markdown processor for Go|4839|578|201|2011-05-27T22:28:58Z|2021-10-08T02:26:56Z|
[sh](https://github.com/mvdan/sh)|A shell parser, formatter, and interpreter with bash support; includes shfmt|4329|225|66|2016-01-16T08:39:09Z|2021-12-11T14:08:58Z|
[toml](https://github.com/BurntSushi/toml)|TOML parser for Golang with reflection.|3700|476|20|2013-02-26T05:05:48Z|2021-11-25T12:10:41Z|
[go-humanize](https://github.com/dustin/go-humanize)|Go Humans! (formatters for units to human friendly sizes)|2942|191|32|2012-01-13T03:48:55Z|2021-10-19T10:19:36Z|
[bluemonday](https://github.com/microcosm-cc/bluemonday)|bluemonday: a fast golang HTML sanitizer (inspired by the OWASP Java HTML Sanitizer) to scrub user generated content of XSS|2131|133|18|2013-11-20T22:15:49Z|2021-10-18T13:39:21Z|
[gofeed](https://github.com/mmcdole/gofeed)|Parse RSS, Atom and JSON feeds in Go|1753|162|40|2016-01-23T02:44:34Z|2021-12-02T12:45:28Z|
**[ARCHIVED]**  [inject](https://github.com/facebookarchive/inject)|Package inject provides a reflect based injector.|1357|120|9|2013-10-21T01:51:46Z|2019-01-14T04:05:17Z|
[go-toml](https://github.com/pelletier/go-toml)|Go library for the TOML file format|1135|164|34|2013-02-24T17:45:51Z|2021-12-06T01:02:20Z|
[commonregex](https://github.com/mingrammer/commonregex)|🍫 A collection of common regular expressions for Go|787|59|2|2017-03-23T14:33:18Z|2019-11-12T07:22:40Z|
[slug](https://github.com/gosimple/slug)|URL-friendly slugify with multiple languages support.|756|76|6|2014-03-31T06:24:51Z|2021-12-03T10:20:02Z|
[dataflowkit](https://github.com/slotix/dataflowkit)|Extract structured data from web sites. Web sites scraping.  |503|67|0|2017-02-09T15:08:15Z|2020-06-12T20:57:30Z|
[mxj](https://github.com/clbanning/mxj)|Decode / encode XML to/from map[string]interface{} (or JSON); extract values with dot-notation paths and wildcards.  Replaces x2j and j2x packages.|491|89|0|2014-02-03T13:39:16Z|2021-03-09T16:59:04Z|
[gographviz](https://github.com/awalterschulze/gographviz)|Parses the Graphviz DOT language in golang|456|69|9|2015-03-14T18:27:00Z|2021-08-24T16:28:52Z|
[htmlquery](https://github.com/antchfx/htmlquery)|htmlquery is golang XPath package for HTML query.|426|52|8|2017-12-05T01:08:41Z|2021-11-25T07:43:33Z|
[gommon](https://github.com/labstack/gommon)|Common packages for Go|420|94|16|2015-03-12T22:35:57Z|2021-11-29T15:58:08Z|
[omniparser](https://github.com/jf-tech/omniparser)|omniparser: a native Golang ETL streaming parser and transform library for CSV, JSON, XML, EDI, text, etc.|400|20|0|2020-08-16T22:22:21Z|2021-11-18T19:43:55Z|
[go-runewidth](https://github.com/mattn/go-runewidth)|wcwidth for golang|395|70|9|2013-06-21T04:56:50Z|2021-11-26T05:17:01Z|
[gotext](https://github.com/leonelquinteros/gotext)|Go (Golang) GNU gettext utilities package |310|37|6|2016-06-19T20:14:43Z|2021-09-26T14:00:56Z|
[html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown)|⚙️ Convert HTML to Markdown. Even works with entire websites and can be extended through rules.|291|40|7|2018-05-15T13:26:26Z|2021-09-07T09:34:23Z|
[goq](https://github.com/andrewstuart/goq)|A declarative struct-tag-based HTML unmarshaling or scraping package for Go built on top of the goquery library|203|17|2|2017-02-20T02:54:40Z|2021-09-02T04:20:26Z|
[go-nmea](https://github.com/adrianmo/go-nmea)|A NMEA parser library in pure Go|160|58|3|2015-07-22T08:55:54Z|2021-12-10T13:49:11Z|
[github_flavored_markdown](https://github.com/shurcooL/github_flavored_markdown)|GitHub Flavored Markdown renderer with fenced code block highlighting, clickable header anchor links.|135|34|10|2015-05-16T04:09:07Z|2021-02-28T21:35:44Z|
[gospider](https://github.com/zhshch2002/gospider)|⚡ Light weight Golang spider framework   轻量的 Golang 爬虫框架|135|8|0|2020-06-17T06:01:39Z|2021-03-16T07:18:08Z|
**[ARCHIVED]**  [sdp](https://github.com/gortc/sdp)|RFC 4566 SDP implementation in go|110|31|5|2016-05-13T14:35:11Z|2020-05-03T07:27:16Z|
[podcast](https://github.com/eduncan911/podcast)|iTunes and RSS 2.0 Podcast Generator in Golang|99|27|5|2017-02-02T12:45:04Z|2020-11-04T21:44:28Z|
[go-zero-width](https://github.com/trubitsyn/go-zero-width)|Zero-width character detection and removal for Go|96|8|0|2018-06-18T13:55:09Z|2020-08-06T14:29:12Z|
[editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go)|EditorConfig Core written in Go|89|30|4|2016-07-05T03:50:41Z|2021-12-06T10:21:56Z|
[align](https://github.com/Guitarbum722/align)|A general purpose application and library for aligning text.|71|7|0|2017-04-29T23:22:22Z|2021-09-12T16:21:36Z|
[go-slugify](https://github.com/mozillazg/go-slugify)|Pretty Slug.|69|5|1|2016-07-16T11:55:15Z|2020-05-13T18:54:09Z|
[go-vcard](https://github.com/emersion/go-vcard)|A Go library to parse and format vCard|63|19|4|2017-03-21T08:30:36Z|2021-05-21T07:54:10Z|
[genex](https://github.com/alixaxel/genex)|Genex package for Go|62|6|0|2015-03-09T19:24:16Z|2020-01-05T18:10:35Z|
[goregen](https://github.com/zach-klippenstein/goregen)|randexp for Go.|60|7|4|2014-12-27T00:19:39Z|2021-07-12T07:39:42Z|
[did](https://github.com/ockam-network/did)|A golang package to work with Decentralized Identifiers (DIDs) |56|14|4|2018-11-02T17:49:14Z|2021-01-03T17:25:37Z|
[go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth)|Encoding and decoding for fixed-width formatted data|54|21|9|2017-11-15T21:05:44Z|2021-06-18T08:48:31Z|
[guesslanguage](https://github.com/endeveit/guesslanguage)|Guess the natural language of a text in Go|52|2|1|2014-12-16T10:58:47Z|2017-11-08T02:01:01Z|
[allot](https://github.com/sbstjn/allot)|Parse placeholder and wildcard text commands|49|7|1|2016-10-16T15:49:08Z|2021-04-16T13:30:34Z|
[pagser](https://github.com/foolin/pagser)|Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler|44|3|2|2020-04-19T09:22:00Z|2020-05-09T14:25:27Z|
[bafi](https://github.com/mmalcek/bafi)|Universal JSON, BSON, YAML, CSV, XML converter with templates|43|2|0|2021-07-13T10:48:40Z|2021-12-10T08:23:23Z|
[gonameparts](https://github.com/polera/gonameparts)|Takes a full name and splits it into individual name parts|32|3|2|2015-05-17T05:20:17Z|2019-08-09T10:09:36Z|
[slugify](https://github.com/avelino/slugify)|A Go slugify application that handles string|28|3|0|2015-04-13T01:54:30Z|2018-05-01T14:59:21Z|
[normalize](https://github.com/avito-tech/normalize)||22|1|0|2021-03-22T09:25:14Z|2021-04-01T08:47:45Z|
[codetree](https://github.com/aerogo/codetree)|:evergreen_tree: Parses indented code and returns a tree structure.|17|4|0|2016-11-26T02:50:38Z|2019-10-26T04:19:45Z|
[enca](https://github.com/endeveit/enca)|Minimal cgo bindings for libenca|10|4|0|2014-12-17T04:55:16Z|2016-03-15T07:18:17Z|
[syndfeed](https://github.com/zhengchun/syndfeed)|A syndication feed parser for Atom 1.0 and RSS 2.0 in Go|7|3|0|2017-04-07T09:30:55Z|2018-03-13T02:31:36Z|
[bbConvert](https://github.com/CalebQ42/bbConvert)|Converter from BBCode to HTML|5|2|0|2016-04-15T14:35:38Z|2016-09-14T13:04:30Z|
[doi](https://github.com/hscells/doi)|Parse and check doi objects in go.|5|1|0|2017-08-02T05:58:01Z|2017-08-21T05:50:49Z|
[ltsv](https://github.com/Wing924/ltsv)|High performance LTSV (Labeled Tab Separeted Value) reader for Go.|5|0|0|2019-05-12T06:11:04Z|2019-06-23T05:47:44Z|
[encoding](https://github.com/ake-persson/encoding)|Go package provides a generic interface to encoders and decoders|4|2|1|2018-04-06T20:48:00Z|2019-11-12T13:29:42Z|
[go-wildcard](https://github.com/IGLOU-EU/go-wildcard)||4|0|0|2021-03-28T16:31:41Z|2021-04-09T10:53:55Z|
[go-output-format](https://github.com/drewstinnett/go-output-format)|Output go objects in standard formats, such as YAML, JSON, etc|3|0|0|2021-04-08T20:48:17Z|2021-10-18T23:14:38Z|


### Utility



	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[xurls](https://github.com/mvdan/xurls)|Extract urls from text|845|99|2|2015-01-12T01:28:46Z|2021-09-30T19:33:01Z|
[gotabulate](https://github.com/bndr/gotabulate)|Gotabulate - Easily pretty-print your tabular data with Go|272|28|6|2014-08-21T07:44:28Z|2021-02-09T14:02:15Z|
[radix](https://github.com/yourbasic/radix)|A fast string sorting algorithm (MSD radix sort)|171|10|0|2017-06-09T14:38:58Z|2018-03-08T12:29:25Z|
[regroup](https://github.com/oriser/regroup)|Match regex group into go struct using struct tags and automatic parsing|97|7|1|2020-09-08T19:04:42Z|2021-07-30T15:53:28Z|
[gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself)|A sanitization-based swear filter for Go.|40|5|2|2018-09-09T00:07:26Z|2021-06-23T18:34:01Z|
[parth](https://github.com/codemodus/parth)|Path parsing for segment unmarshaling and slicing.|39|5|0|2015-04-06T22:53:59Z|2019-02-01T00:16:42Z|
[xj2go](https://github.com/wk30/xj2go)|Convert xml and json to go struct|21|7|0|2017-09-19T13:20:57Z|2021-10-12T17:03:04Z|
[tagify](https://github.com/zoomio/tagify)|Tagify produces a set of tags from a given source. Source can be either an HTML page, a Markdown document or a plain text. Supports English, Russian, Chinese, Hindi, Spanish, Arabic, Japanese, German, Hebrew, French and Korean languages.|17|1|0|2018-03-20T10:30:11Z|2021-11-11T05:12:55Z|
[kace](https://github.com/codemodus/kace)|Common case conversions covering common initialisms.|16|2|1|2015-06-04T20:36:49Z|2018-08-26T21:35:11Z|
[TySug](https://github.com/Dynom/TySug)|A project around helping to prevent typing typos. TySug (Typo Suggestions) suggests alternative words with respect to keyboard layouts|9|2|0|2018-06-05T19:46:29Z|2020-08-03T09:26:45Z|
[parseargs-go](https://github.com/txgruppi/parseargs-go)|A string argument parser that understands quotes and backslashes|8|4|1|2016-02-24T00:53:38Z|2017-01-24T21:54:06Z|
[textwrap](https://github.com/isbm/textwrap)|Port of Python&#39;s &#34;textwrap&#34; module to Go|1|1|1|2019-07-26T17:57:55Z|2019-08-03T19:01:29Z|


## Third-party APIs
*Libraries for accessing third party APIs.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[go-github](https://github.com/google/go-github)|Go library for accessing the GitHub API|7980|1651|32|2013-05-24T16:42:58Z|2021-12-08T17:22:43Z|
[aws-sdk-go](https://github.com/aws/aws-sdk-go)|AWS SDK for the Go programming language.|7256|1772|112|2014-12-05T05:29:41Z|2021-12-10T01:01:31Z|
[slack](https://github.com/slack-go/slack)|Slack API in Go - community-maintained fork created by the original author, @nlopes|3735|918|73|2015-01-24T14:19:00Z|2021-11-26T21:50:22Z|
[google-api-go-client](https://github.com/googleapis/google-api-go-client)|Auto-generated Google APIs for Go.|2859|894|27|2014-11-24T21:45:36Z|2021-12-11T13:12:15Z|
[google-cloud-go](https://github.com/googleapis/google-cloud-go)|Google Cloud Client Libraries for Go.|2713|979|225|2014-05-09T11:11:58Z|2021-12-10T02:05:54Z|
[discordgo](https://github.com/bwmarrin/discordgo)| (Golang) Go bindings for Discord|2562|506|129|2015-11-01T20:51:01Z|2021-12-11T16:08:08Z|
[stripe-go](https://github.com/stripe/stripe-go)|Go library for the Stripe API.    |1470|390|11|2014-06-05T23:38:14Z|2021-12-09T20:42:54Z|
[minio-go](https://github.com/minio/minio-go)|MinIO Client SDK for Go|1467|449|4|2015-05-02T02:36:46Z|2021-12-10T18:40:41Z|
[go-twitter](https://github.com/dghubble/go-twitter)|Go Twitter REST and Streaming API v1.1|1357|266|32|2015-04-11T23:26:07Z|2021-11-15T16:04:50Z|
[anaconda](https://github.com/ChimeraCoder/anaconda)|A Go client library for the Twitter 1.1 API|1104|254|72|2013-03-04T22:46:07Z|2021-07-28T08:02:55Z|
[go-jira](https://github.com/andygrunwald/go-jira)|Go client library for Atlassian Jira|1031|345|83|2015-08-20T15:02:46Z|2021-12-06T15:57:05Z|
[facebook](https://github.com/huandu/facebook)|A Facebook Graph API SDK For Go.|1019|394|0|2012-07-28T19:05:56Z|2021-04-30T03:47:52Z|
[githubv4](https://github.com/shurcooL/githubv4)|Package githubv4 is a client library for accessing GitHub GraphQL API v4 (https://docs.github.com/en/graphql).|826|69|35|2017-05-27T05:05:31Z|2021-11-17T02:00:13Z|
[webhooks](https://github.com/go-playground/webhooks)|:fishing_pole_and_fish: Webhook receiver for GitHub, Bitbucket, GitLab, Gogs|680|169|25|2015-10-25T17:38:13Z|2021-11-16T17:03:02Z|
[paypal](https://github.com/plutov/paypal)|Golang client for PayPal REST API|458|207|5|2015-10-14T04:57:49Z|2021-11-24T15:00:11Z|
[geo-golang](https://github.com/codingsince1985/geo-golang)|Go library to access geocoding and reverse geocoding APIs|419|49|6|2014-12-04T08:18:31Z|2021-08-14T08:19:58Z|
[ethrpc](https://github.com/onrik/ethrpc)|Golang client for ethereum json rpc api|218|91|10|2017-01-24T09:47:00Z|2020-08-24T04:49:02Z|
[go-marathon](https://github.com/gambol99/go-marathon)|A GO API library for working with Marathon|196|132|27|2015-02-11T13:25:26Z|2020-10-01T16:32:07Z|
[trello](https://github.com/adlio/trello)|Trello API wrapper for Go|188|68|12|2016-09-24T04:36:10Z|2021-06-18T23:24:22Z|
[twitter-scraper](https://github.com/n0madic/twitter-scraper)|Scrape the Twitter Frontend API without authentication with Golang.|154|39|1|2018-11-29T15:31:50Z|2021-12-07T08:18:05Z|
[medium-sdk-go](https://github.com/Medium/medium-sdk-go)|A Golang SDK for Medium&#39;s OAuth2 API|130|21|6|2015-09-26T23:45:46Z|2018-10-26T20:37:15Z|
[gostorm](https://github.com/jsgilmore/gostorm)|GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.|128|21|5|2013-07-22T12:43:41Z|2017-10-09T12:00:28Z|
[go-trending](https://github.com/andygrunwald/go-trending)|Go library for accessing trending repositories and developers at Github.|118|17|3|2015-07-04T08:06:48Z|2021-10-26T04:15:47Z|
[hipchat](https://github.com/daneharrigan/hipchat)|A golang package to communicate with HipChat over XMPP|111|37|3|2013-04-28T02:16:21Z|2017-06-12T14:49:06Z|
[wit-go](https://github.com/wit-ai/wit-go)|Go client for wit.ai HTTP API|110|27|0|2018-08-20T07:18:40Z|2021-08-31T20:18:19Z|
[pushover](https://github.com/gregdel/pushover)|Go wrapper for the Pushover API|109|10|1|2015-02-19T15:30:05Z|2021-10-21T12:21:35Z|
[hipchat](https://github.com/andybons/hipchat)|This project implements a Go client library for the Hipchat API.|105|22|0|2012-10-20T18:34:06Z|2016-03-24T19:12:10Z|
[cachet](https://github.com/andygrunwald/cachet)|Go(lang) client library for Cachet (open source status page system).|89|13|1|2015-10-31T12:30:07Z|2021-06-22T17:03:41Z|
[simples3](https://github.com/rhnvrm/simples3)|Simple no frills AWS S3 Golang Library using REST with V4 Signing (without AWS Go SDK)|70|13|2|2018-12-06T10:24:21Z|2021-11-24T17:36:36Z|
[gosip](https://github.com/koltyakov/gosip)|⚡️ SharePoint authentication, HTTP client &amp; fluent API wrapper for Go (Golang)|69|23|10|2019-01-26T08:48:48Z|2021-11-18T07:11:42Z|
[igdb](https://github.com/Henry-Sarabia/igdb)|Go client for the Internet Game Database API|68|13|3|2017-08-24T08:31:53Z|2021-03-15T21:23:29Z|
[lark](https://github.com/go-lark/lark)|An easy-to-use unofficial SDK for Feishu and Lark Open Platform (Messaging API only)|68|3|3|2021-04-20T12:09:03Z|2021-12-10T09:21:06Z|
[go-circleci](https://github.com/jszwedko/go-circleci)|Go library for interacting with CircleCI|61|49|5|2015-08-14T21:19:36Z|2019-11-21T00:02:51Z|
[gogtrends](https://github.com/groovili/gogtrends)|Unofficial Google Trends API for Go|59|18|0|2018-12-27T13:50:34Z|2021-09-07T06:44:09Z|
[go-unsplash](https://github.com/hbagdi/go-unsplash)|Go Client for the Unsplash API |59|11|8|2017-01-19T07:04:04Z|2021-03-30T15:17:24Z|
[clarifai-go](https://github.com/Clarifai/clarifai-go)|DEPRECATED: please use https://github.com/Clarifai/clarifai-go-grpc|55|13|8|2015-09-28T23:33:59Z|2017-08-28T17:25:50Z|
[megos](https://github.com/andygrunwald/megos)|Go(lang) client library for accessing information of an Apache Mesos cluster.|54|11|0|2015-10-02T14:29:20Z|2021-06-22T17:06:10Z|
[go-amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api)|Go Client Library for Amazon Product Advertising API|51|15|3|2016-11-15T15:37:32Z|2018-04-05T22:06:29Z|
[gads](https://github.com/emiddleton/gads)|Google Adwords API for Go|49|58|8|2014-01-20T02:22:15Z|2019-09-23T04:43:04Z|
[ynab.go](https://github.com/brunomvsouza/ynab.go)|Go client for the YNAB API. Unofficial. It covers 100% of the resources made available by the YNAB API.|48|14|6|2018-07-13T11:10:54Z|2021-09-15T04:45:36Z|
[uptimerobot](https://github.com/bitfield/uptimerobot)|Client library for UptimeRobot v2 API|45|11|12|2018-05-29T10:27:19Z|2020-12-28T14:49:04Z|
[golang-tmdb](https://github.com/cyruzin/golang-tmdb)|This is a Golang wrapper for working with TMDb API. It aims to support version 3.|45|10|1|2019-01-11T22:59:33Z|2021-12-11T12:49:53Z|
[go-xkcd](https://github.com/nishanths/go-xkcd)|xkcd.com API client in Go|45|5|1|2016-02-26T05:14:31Z|2021-10-27T13:26:22Z|
[gomusicbrainz](https://github.com/michiwend/gomusicbrainz)|a Go (Golang) MusicBrainz WS2 client library - work in progress|44|17|5|2014-09-10T16:42:33Z|2021-02-09T23:41:40Z|
[mixpanel](https://github.com/dukex/mixpanel)|Golang Mixpanel Client|42|24|3|2014-05-20T03:50:34Z|2021-08-11T14:13:14Z|
[fcm](https://github.com/maddevsio/fcm)|Firebase Cloud Messaging for application servers implemented using the Go programming language.|41|14|2|2017-01-06T08:30:57Z|2020-03-06T05:17:46Z|
[go-spotify](https://github.com/rapito/go-spotify)|Go library for the Spotify Web API|39|6|0|2014-10-30T02:52:04Z|2020-12-03T22:51:03Z|
[golyrics](https://github.com/mamal72/golyrics)|A simple Go package to fetch lyrics from Wikia|36|2|0|2016-11-18T04:40:37Z|2018-06-30T08:33:13Z|
[go-postman-collection](https://github.com/rbretecher/go-postman-collection)|Go module to work with Postman Collections|34|11|1|2019-11-16T12:13:32Z|2021-11-24T23:16:26Z|
**[ARCHIVED]**  [translate](https://github.com/nuveo/translate)|Go online translation package|31|6|0|2015-07-13T15:42:13Z|2016-02-28T15:13:19Z|
[gcm](https://github.com/TheOrioli/gcm)|Google Cloud Messaging for application servers implemented using the Go programming language.|30|4|0|2015-11-09T16:16:25Z|2015-12-04T14:37:11Z|
**[ARCHIVED]**  [gami](https://github.com/bit4bit/gami)|GO - Asterisk AMI Interface|29|25|1|2014-05-14T16:11:37Z|2018-06-26T10:42:14Z|
[airtable](https://github.com/mehanizm/airtable)|Simple golang airtable API wrapper|28|11|0|2020-04-12T10:05:07Z|2021-10-02T11:27:01Z|
[patreon-go](https://github.com/mxpv/patreon-go)|Patreon Go API client|27|14|1|2017-08-06T21:15:14Z|2019-09-17T02:27:28Z|
[go-myanimelist](https://github.com/nstratos/go-myanimelist)|Go library for accessing the MyAnimeList API: https://myanimelist.net/apiconfig/references/api/v2|25|1|0|2015-05-03T10:07:05Z|2021-11-27T00:34:32Z|
[go-steam](https://github.com/sostronk/go-steam)|Go library for querying Source servers|24|7|2|2014-11-23T16:34:56Z|2021-09-07T16:30:55Z|
[lastpass-go](https://github.com/ansd/lastpass-go)|Golang client for LastPass|23|4|1|2019-07-11T14:26:39Z|2021-11-08T17:16:04Z|
[google-play-scraper](https://github.com/n0madic/google-play-scraper)|Golang scraper to get data from Google Play Store|22|9|0|2019-09-20T14:03:01Z|2021-10-26T10:12:10Z|
[go-shopify](https://github.com/rapito/go-shopify)|Simple Shopify API for the Go Programming Language|22|7|2|2014-10-28T02:53:25Z|2020-12-03T22:50:32Z|
[go-imgur](https://github.com/koffeinsource/go-imgur)|Go library to use the imgur.com API|21|5|1|2016-03-30T22:05:35Z|2021-04-30T12:05:19Z|
[go-twitch](https://github.com/knspriggs/go-twitch)|A golang client for the Twitch v3 API - public APIs only (for now)|21|3|3|2016-06-28T20:54:34Z|2017-08-23T16:28:21Z|
[codeship-go](https://github.com/codeship/codeship-go)|Go library for accessing the Codeship API v2|18|8|1|2017-09-08T16:49:59Z|2020-11-03T16:20:17Z|
[textbelt](https://github.com/farmergreg/textbelt)|golang library for textbelt.com|17|1|0|2015-09-01T22:46:42Z|2015-09-04T14:12:39Z|
[brewerydb](https://github.com/naegelejd/brewerydb)|Go library for http://www.brewerydb.com/ API|17|1|5|2015-04-15T02:59:41Z|2015-06-18T19:34:13Z|
[coinpaprika-api-go-client](https://github.com/coinpaprika/coinpaprika-api-go-client)|Go client library for interacting with Coinpaprika&#39;s API|13|5|1|2018-09-25T07:34:50Z|2020-09-16T05:09:30Z|
[go-aws-news](https://github.com/circa10a/go-aws-news)|Go app &#43; library to fetch what&#39;s new from AWS|13|4|0|2020-01-08T00:59:39Z|2021-11-24T15:26:46Z|
[go-hacknews](https://github.com/PaulRosset/go-hacknews)|📟  Tiny utility Go client for HackerNews API.|13|1|0|2017-08-10T20:44:02Z|2017-08-15T07:51:32Z|
[go-google-analytics](https://github.com/chonthu/go-google-analytics)|Simple Reporting for Google Analytics|12|3|0|2015-06-01T13:50:00Z|2015-06-09T11:38:07Z|
[device-check-go](https://github.com/rinchsan/device-check-go)|:iphone: iOS DeviceCheck SDK for Go - query and modify the per-device bits|11|5|0|2019-04-11T13:09:11Z|2021-10-09T05:28:32Z|
[gopaapi5](https://github.com/utekaravinash/gopaapi5)|Go Client Library for Amazon&#39;s Product Advertising API 5.0|11|5|0|2020-02-15T06:21:31Z|2020-04-03T18:38:34Z|
[go-here](https://github.com/abdullahselek/go-here)|Go client library around the HERE location based APIs.|10|5|0|2019-07-07T12:14:34Z|2020-06-23T13:20:37Z|
[smitego](https://github.com/sergiotapia/smitego)|SmiteGo is an API wrapper for the Smite game from HiRez. It is written in Go!|10|1|0|2013-12-11T02:38:19Z|2014-07-18T15:51:45Z|
[go-openproject](https://github.com/manuelbcd/go-openproject)|Go client library for OpenProject|9|4|5|2021-02-13T23:23:13Z|2021-04-09T08:39:38Z|
[gomalshare](https://github.com/MonaxGT/gomalshare)|Go library MalShare API|9|3|0|2019-03-01T09:33:41Z|2019-04-29T08:00:01Z|
[go-sophos](https://github.com/esurdam/go-sophos)|Sophos UTM 9 REST API Client in Golang|8|4|0|2018-09-05T04:37:25Z|2020-10-04T01:07:21Z|
**[ARCHIVED]**  [rrdaclient](https://github.com/Omie/rrdaclient)|Go bindings for RRDA https://github.com/fcambus/rrda|8|0|0|2014-09-15T21:06:16Z|2014-09-19T16:36:10Z|
[go-zooz](https://github.com/gojuno/go-zooz)|Zooz API client for Go|7|6|0|2017-07-04T09:28:23Z|2021-11-22T10:54:08Z|
[go-google-email-audit-api](https://github.com/ngs/go-google-email-audit-api)|Go Client Library for G Suite Email Audit API|7|5|0|2016-10-24T02:34:29Z|2016-10-26T12:55:17Z|
[gumblr](https://github.com/mattcunningham/gumblr)|A Go Wrapper for the Tumblr v2 API|6|6|0|2015-07-09T23:13:51Z|2016-10-30T23:45:20Z|
[go-sptrans](https://github.com/sergioaugrod/go-sptrans)|Go client library for the SPTrans Olho Vivo API. :bus:|6|1|0|2017-09-11T01:21:28Z|2020-09-16T22:40:59Z|
[bqwriter](https://github.com/OTA-Insight/bqwriter)|Stream data into Google BigQuery concurrently using InsertAll() or BQ Storage.|6|1|0|2021-10-12T13:58:18Z|2021-12-06T10:33:09Z|
[go-chronos](https://github.com/axelspringer/go-chronos)|:dancers: Go Chronos 3.x REST API Client|4|3|0|2017-10-23T12:19:01Z|2018-01-23T14:00:43Z|
[libgoffi](https://github.com/clevabit/libgoffi)|libgoffi - libffi adapter library for Go|4|1|0|2019-08-03T17:05:34Z|2020-08-23T13:02:21Z|
[go-swagger-ui](https://github.com/esurdam/go-swagger-ui)|Golang package which provides http Handlers to serve the swagger ui|4|0|0|2021-05-25T01:26:09Z|2021-06-04T20:38:49Z|
[kanka](https://github.com/Henry-Sarabia/kanka)|Go client for the Kanka API|3|4|2|2019-12-26T00:07:57Z|2020-08-06T19:59:39Z|
[rawg-sdk-go](https://github.com/dimuska139/rawg-sdk-go)|This is RAWG SDK GO. This library contains methods for interacting with RAWG API.|3|1|0|2020-10-16T15:31:37Z|2021-07-06T20:46:00Z|
[go-restcountries](https://github.com/chriscross0/go-restcountries)|Go wrapper for the REST Countries API.|3|0|0|2021-08-01T17:49:51Z|2021-10-27T15:38:43Z|
[appstore-sdk-go](https://github.com/Kachit/appstore-sdk-go)|Golang SDK for AppStore Connect API (Unofficial)|2|0|0|2020-06-11T10:05:56Z|2021-02-21T19:19:33Z|
[tripadvisor-golang](https://github.com/mrbenosborne/tripadvisor-golang)|A TripAdvisor API wrapper for Golang.|1|1|0|2019-04-15T18:12:11Z|2019-10-23T15:20:38Z|
[vl-go](https://github.com/verifid/vl-go)|Go client library around the VerifID identity verification layer API.|1|1|0|2019-02-09T12:46:53Z|2021-05-30T19:02:02Z|
[playlyfe-go-sdk](https://github.com/playlyfe/playlyfe-go-sdk)|This is the official Playlyfe Golang Sdk|1|1|0|2015-05-25T09:34:47Z|2016-03-06T10:21:06Z|
[go-telegraph](https://gitlab.com/toby3d/telegraph)|Telegraph publishing platform API client.|-|-|-|-|-|
[go-yapla](https://git.iglou.eu/Production/go-yapla)|Go client library for the Yapla v2.0 API.|-|-|-|-|-|


## Utilities
*General utilities and tools to make your life easier.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[fzf](https://github.com/junegunn/fzf)|:cherry_blossom: A command-line fuzzy finder|40812|1775|258|2013-10-23T16:04:23Z|2021-12-09T02:05:12Z|
[hub](https://github.com/github/hub)|A command-line tool that makes git easier to use with GitHub.|21413|2261|289|2009-12-05T22:15:25Z|2021-10-10T18:44:04Z|
[ctop](https://github.com/bcicen/ctop)|Top-like interface for container metrics|12173|472|64|2016-12-27T02:25:57Z|2021-12-08T01:45:58Z|
[sqlx](https://github.com/jmoiron/sqlx)|general purpose extensions to golang&#39;s database/sql|11175|868|290|2013-01-28T19:40:00Z|2021-12-07T12:43:38Z|
[wuzz](https://github.com/asciimoo/wuzz)|Interactive cli tool for HTTP inspection|9839|400|37|2017-01-30T21:22:00Z|2021-08-29T20:46:21Z|
[goreleaser](https://github.com/goreleaser/goreleaser)|Deliver Go binaries as fast and easily as possible|9272|617|36|2016-12-21T17:13:39Z|2021-12-10T13:50:29Z|
[usql](https://github.com/xo/usql)|Universal command-line interface for SQL databases|6862|251|53|2017-03-02T13:03:21Z|2021-12-02T03:40:19Z|
[peco](https://github.com/peco/peco)|Simplistic interactive filtering tool|6644|219|41|2014-06-06T06:06:32Z|2021-07-30T03:30:09Z|
[godropbox](https://github.com/dropbox/godropbox)|Common libraries for writing Go services/applications.|3999|425|5|2014-06-22T23:09:29Z|2020-07-07T19:02:22Z|
[hystrix-go](https://github.com/afex/hystrix-go)|Netflix&#39;s Hystrix latency and fault tolerance library, for Go |3473|403|54|2013-12-15T08:51:23Z|2021-09-13T14:48:27Z|
[go-funk](https://github.com/thoas/go-funk)|A modern Go utility library which provides helpers (map, find, contains, filter, ...)|3074|185|6|2016-12-30T13:55:15Z|2021-12-05T20:45:01Z|
[panicparse](https://github.com/maruel/panicparse)|Crash your app in style (Golang)|2935|88|3|2015-02-02T02:14:41Z|2021-12-06T15:17:15Z|
[goreporter](https://github.com/qax-os/goreporter)|A Golang tool that does static analysis, unit testing, code review and generate code quality report.|2934|266|29|2017-03-27T08:46:38Z|2018-10-27T22:30:57Z|
[minify](https://github.com/tdewolff/minify)|Go minifiers for web formats|2819|175|8|2014-05-21T09:03:48Z|2021-12-07T21:22:21Z|
[mc](https://github.com/minio/mc)|MinIO Client is a replacement for ls, cp, mkdir, diff and rsync commands for filesystems and object storage.|1960|363|26|2015-01-16T02:56:51Z|2021-12-11T17:04:32Z|
[storm](https://github.com/asdine/storm)|Simple and powerful toolkit for BoltDB|1810|128|65|2016-01-10T12:55:59Z|2021-05-14T06:46:07Z|
[mergo](https://github.com/imdario/mergo)|Mergo: merging Go structs and maps since 2013.|1805|207|27|2013-03-11T22:51:11Z|2021-11-09T19:01:16Z|
[mole](https://github.com/davrodpin/mole)|CLI application to create ssh tunnels focused on resiliency and user experience.|1518|83|16|2018-10-04T02:38:00Z|2021-10-06T07:10:54Z|
[filetype](https://github.com/h2non/filetype)|Fast, dependency-free Go package to infer binary file types based on the magic numbers header signature|1450|132|26|2015-09-24T09:15:51Z|2021-11-21T11:26:41Z|
[boilr](https://github.com/tmrts/boilr)|:zap: boilerplate template manager that generates files or directories from template repositories|1412|109|44|2015-12-19T16:57:26Z|2021-07-30T23:38:48Z|
[jump](https://github.com/gsamokovarov/jump)|Jump helps you navigate faster by learning your habits. ✌️|1278|49|1|2015-08-16T22:07:17Z|2021-12-06T09:25:37Z|
[cli](https://github.com/create-go-app/cli)|✨ Create a new production-ready project with backend, frontend and deploy automation by running one CLI command!|1186|118|0|2019-12-30T22:08:38Z|2021-11-24T14:42:27Z|
[circuitbreaker](https://github.com/rubyist/circuitbreaker)|Circuit Breakers in Go|969|110|18|2014-07-17T22:41:33Z|2019-10-21T12:27:19Z|
[gtm](https://github.com/git-time-metric/gtm)|Simple, seamless, lightweight time tracking for Git|883|49|48|2016-06-19T21:17:04Z|2020-09-27T19:02:27Z|
[immortal](https://github.com/immortal/immortal)|⭕  A *nix cross-platform (OS agnostic) supervisor|717|51|1|2016-06-30T17:02:27Z|2020-06-23T11:52:32Z|
[hostctl](https://github.com/guumaster/hostctl)|Your dev tool to manage /etc/hosts like a pro!|704|33|8|2020-03-14T11:29:02Z|2021-12-08T10:07:37Z|
[mimetype](https://github.com/gabriel-vasile/mimetype)|A fast Golang library for media type and file extension detection, based on magic numbers|620|92|45|2018-07-02T07:15:29Z|2021-12-10T19:37:12Z|
[circuit](https://github.com/cep21/circuit)|An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern.|617|36|7|2017-12-23T22:17:43Z|2021-09-01T21:26:30Z|
[htcat](https://github.com/htcat/htcat)|Parallel and Pipelined HTTP GET Utility|536|30|5|2013-08-05T11:17:01Z|2019-02-26T22:54:07Z|
[ergo](https://github.com/cristianoliveira/ergo)|The management of multiple apps running over different ports made easy|488|49|17|2017-08-19T18:41:56Z|2021-11-19T15:16:11Z|
[godaemon](https://github.com/VividCortex/godaemon)|Daemonize Go applications deviously.|481|55|8|2013-08-01T17:16:30Z|2021-06-29T04:55:28Z|
[koazee](https://github.com/wesovilabs/koazee)|A StreamLike, Immutable, Lazy Loading and smart Golang Library to deal with slices.|475|27|16|2018-11-09T09:49:19Z|2020-11-18T17:04:42Z|
[go-dry](https://github.com/ungerik/go-dry)|DRY (don&#39;t repeat yourself) package for Go|472|33|0|2014-02-28T13:49:31Z|2021-11-26T07:58:43Z|
[scany](https://github.com/georgysavva/scany)|Library for scanning data from a database into Go structs and more|452|33|19|2020-07-02T11:02:58Z|2021-12-11T20:54:50Z|
[gopencils](https://github.com/bndr/gopencils)|Easily consume REST APIs with Go (golang)|438|42|7|2014-06-23T11:41:24Z|2019-02-18T01:03:37Z|
[delve](https://github.com/derekparker/delve)|Delve is a debugger for the Go programming language.|422|89|0|2020-02-18T18:03:33Z|2021-11-24T21:10:23Z|
[request](https://github.com/mozillazg/request)|A developer-friendly HTTP request library for Gopher.|401|39|6|2014-12-21T04:30:42Z|2019-12-05T09:11:26Z|
[gubrak](https://github.com/novalagung/gubrak)|⚙️ Golang functional utility library with syntactic sugar. It&#39;s like lodash, but for Go|382|33|0|2018-03-09T11:28:05Z|2020-05-26T11:07:56Z|
[deepcopier](https://github.com/ulule/deepcopier)|simple struct copying for golang|370|51|7|2015-07-24T18:01:01Z|2020-04-30T08:31:45Z|
[clockwork](https://github.com/jonboulle/clockwork)|a fake clock for golang|363|44|5|2014-09-09T18:24:00Z|2021-11-29T22:30:20Z|
[go-rate](https://github.com/beefsack/go-rate)|A timed rate limiter for Go|344|28|1|2014-08-25T04:42:34Z|2020-08-27T23:24:07Z|
[retry](https://github.com/kamilsk/retry)|♻️ The most advanced interruptible mechanism to perform actions repetitively until successful.|310|14|9|2016-11-02T20:20:43Z|2021-02-23T07:20:20Z|
**[ARCHIVED]**  [gohper](https://github.com/cosiner/gohper)|[UNMATAINED] common libs here.|255|47|0|2015-03-23T22:46:12Z|2017-08-12T06:53:29Z|
[serve](https://github.com/syntaqx/serve)|🍽️ a static http server anywhere you need one.|253|16|4|2019-01-10T23:31:52Z|2021-07-27T15:44:11Z|
[scan](https://github.com/blockloop/scan)|Scan database/sql rows directly to structs, slices, and primitive types|237|14|0|2017-11-27T23:22:18Z|2021-10-23T00:51:00Z|
[util](https://github.com/shomali11/util)|A collection of useful utility functions|227|36|0|2017-05-24T00:21:29Z|2020-03-29T02:14:23Z|
[go-trigger](https://github.com/sadlil/go-trigger)|A Global event triggerer for golang. Defines functions as event with id string. Trigger the event anywhere from your project.|226|38|1|2015-10-19T09:42:17Z|2017-03-28T16:18:42Z|
[gotenv](https://github.com/subosito/gotenv)|Load environment variables from `.env` or `io.Reader` in Go.|206|24|5|2013-08-27T12:56:47Z|2021-09-20T08:18:21Z|
[grofer](https://github.com/pesos/grofer)|A system and resource monitoring tool written in Golang!|184|42|11|2020-08-01T16:26:03Z|2021-12-01T19:08:13Z|
[go-bind-plugin](https://github.com/wendigo/go-bind-plugin)|go-bind-plugin generates API for exported plugin symbols (-buildmode=plugin) - go1.8&#43; only (http://golang.org/pkg/plugin)|175|10|0|2016-11-08T14:40:26Z|2019-08-29T11:59:32Z|
[death](https://github.com/vrecan/death)|Managing go application shutdown with signals.|172|18|0|2015-03-09T03:50:40Z|2021-04-27T22:12:42Z|
[mani](https://github.com/alajmo/mani)|CLI tool to help you manage multiple repositories|170|5|2|2019-10-22T20:05:11Z|2021-11-10T14:19:46Z|
[toolbox](https://github.com/viant/toolbox)|Toolbox - go utility library|168|20|2|2016-06-13T19:33:35Z|2021-08-09T16:08:17Z|
[go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator)|go-sitemap-generator is the easiest way to generate Sitemaps in Go|164|50|22|2015-10-12T16:23:13Z|2021-09-01T14:21:44Z|
[rerun](https://github.com/ivpusic/rerun)|Configurable recompiling and rerunning go apps when source changes|161|10|2|2014-12-10T00:29:54Z|2018-03-22T19:46:51Z|
[moldova](https://github.com/StabbyCutyou/moldova)|A lightweight templating system for generating random data|160|6|0|2016-01-30T05:25:39Z|2017-09-04T15:06:03Z|
[apm](https://github.com/topfreegames/apm)|APM is a process manager for Golang applications.|155|75|9|2015-11-18T16:56:48Z|2016-11-24T20:58:45Z|
[robustly](https://github.com/VividCortex/robustly)|Run functions resiliently in Go, catching and restarting panics|151|7|1|2013-07-08T13:27:10Z|2021-04-26T21:51:12Z|
[chyle](https://github.com/antham/chyle)|Changelog generator : use a git repository and various data sources and publish the result on external services|138|11|0|2016-11-17T21:14:44Z|2021-12-01T07:53:12Z|
[countries](https://github.com/biter777/countries)|Countries - ISO 3166 (ISO3166-1, ISO3166, Digit, Alpha-2 and Alpha-3) countries codes and names (on eng and rus), ISO 4217 currency designators, ITU-T E.164 IDD calling phone codes, countries capitals, UN M.49 regions codes, ccTLD countries domains, IOC/NOC and FIFA letters codes, VERY FAST, NO maps[], NO slices[], NO init() funcs, NO external links/files/data, NO interface{}, NO specific dependencies, Databases/JSON/GOB/XML/CSV compatible, Emoji countries flags and currencies support, full support ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and ccTLD standarts.|135|28|1|2019-04-22T14:47:11Z|2021-11-03T19:40:15Z|
[changie](https://github.com/miniscruff/changie)|Automated changelog tool for preparing releases with lots of customization options|134|10|1|2020-12-05T19:38:33Z|2021-12-07T17:21:02Z|
[onecache](https://github.com/adelowo/onecache)|One caching API, Multiple backends|124|7|0|2017-04-14T21:49:15Z|2020-05-25T15:44:21Z|
[go-bsdiff](https://github.com/gabstv/go-bsdiff)|Pure Go bsdiff and bspatch libraries and CLI tools.|122|15|0|2019-02-23T23:33:50Z|2019-03-21T12:35:11Z|
[clipboard](https://github.com/golang-design/clipboard)|📋 cross-platform clipboard package that supports accessing text and image in Go (macOS/Linux/Windows/Android/iOS) |121|15|1|2020-11-19T11:42:08Z|2021-09-22T16:01:33Z|
[lrserver](https://github.com/jaschaephraim/lrserver)|LiveReload server for Go [golang]|120|12|0|2014-07-15T05:36:53Z|2017-11-29T20:31:22Z|
[go-pattern-match](https://github.com/alexpantyukhin/go-pattern-match)|Pattern matchings for Go.|102|7|0|2018-12-11T20:11:17Z|2020-06-28T20:14:38Z|
[cryptgo](https://github.com/Gituser143/cryptgo)|A terminal application to watch crypto prices!|98|15|6|2021-05-20T06:36:28Z|2021-10-17T13:19:34Z|
[sorty](https://github.com/jfcg/sorty)|Fast Concurrent / Parallel Sorting in Go|94|2|0|2019-02-18T21:05:45Z|2021-11-10T22:03:40Z|
[nostromo](https://github.com/pokanop/nostromo)|CLI for building powerful aliases|91|5|10|2019-07-13T04:51:46Z|2021-11-24T08:25:01Z|
[goseaweedfs](https://github.com/linxGnu/goseaweedfs)|A complete Golang client for SeaweedFS|90|27|1|2017-07-20T04:35:39Z|2021-11-01T11:33:51Z|
[mssqlx](https://github.com/linxGnu/mssqlx)|Database client library, proxy for any master slave, master master structures. Lightweight, performant and auto balancing in mind.|89|11|0|2016-12-26T04:05:09Z|2021-09-17T07:45:51Z|
[xferspdy](https://github.com/monmohan/xferspdy)|Xferspdy provides binary diff and patch library in golang. [Mentioned in Awesome Go, https://github.com/avelino/awesome-go]|89|11|3|2015-05-22T13:23:34Z|2021-04-04T09:44:40Z|
[go-health](https://github.com/Talento90/go-health)|:heart: Health check your applications and dependencies|85|5|1|2018-02-13T18:40:54Z|2018-06-13T18:40:20Z|
[rospo](https://github.com/ferama/rospo)|🐸 Simple and reliable ssh tunnels with embedded ssh server|80|4|1|2021-04-02T13:16:14Z|2021-11-04T10:33:20Z|
[mongo-go-pagination](https://github.com/gobeam/mongo-go-pagination)|Golang Mongodb Pagination for official mongodb/mongo-go-driver package which supports both normal queries and Aggregation pipelines with all information like Total records, Page, Per Page, Previous, Next, Total Page and query results.|78|28|0|2020-02-04T08:23:33Z|2021-11-09T05:46:26Z|
[pm](https://github.com/VividCortex/pm)|Processlist manager with TCP listener|77|7|2|2013-11-17T19:17:01Z|2020-12-15T17:40:41Z|
[cmd](https://github.com/commander-cli/cmd)|A simple package to execute shell commands on linux, windows and osx|75|9|3|2019-09-27T13:22:06Z|2020-07-08T09:11:50Z|
[repeat](https://github.com/ssgreg/repeat)|Go implementation of different backoff strategies useful for retrying operations and heartbeating.|75|5|0|2017-11-22T07:06:47Z|2020-07-24T08:18:11Z|
[netbug](https://github.com/e-dard/netbug)|Package netbug provides a handler for registering profilers on your own ServeMux.|69|5|0|2015-03-05T19:27:29Z|2015-10-29T17:28:38Z|
[mimemagic](https://github.com/zRedShift/mimemagic)|Powerful and versatile MIME sniffing package using pre-compiled glob patterns, magic number signatures, XML document namespaces, and tree magic for mounted volumes, generated from the XDG shared-mime-info database.|68|8|2|2018-10-11T16:12:54Z|2021-11-24T13:46:28Z|
[unis](https://github.com/esemplastic/unis)|UNIS: A Common Architecture for String Utilities within the Go Programming Language.|67|4|2|2017-05-06T05:01:03Z|2017-05-09T16:17:33Z|
[handy](https://github.com/miguelpragier/handy)|GO Golang Utilities and helpers like validators and string formatters|66|7|0|2018-06-13T13:10:07Z|2020-09-30T01:22:20Z|
[multitick](https://github.com/VividCortex/multitick)|A multiplexor for aligned time.Time tickers in Go|65|2|1|2013-12-10T16:47:26Z|2021-04-26T21:18:13Z|
[goreadability](https://github.com/philipjkim/goreadability)|Webpage summary extractor using Facebook Open Graph and arc90&#39;s readability|62|8|2|2016-04-20T01:40:14Z|2019-04-22T09:46:39Z|
[pgo](https://github.com/arthurkushman/pgo)|Go library for PHP community with convenient functions|60|13|2|2018-12-26T06:59:47Z|2021-09-26T13:12:45Z|
[minquery](https://github.com/icza/minquery)|MongoDB / mgo query that supports efficient pagination (cursors to continue listing documents where we left off).|59|21|4|2016-11-16T12:23:07Z|2021-07-26T20:21:21Z|
[limiters](https://github.com/mennanov/limiters)|Golang rate limiters for distributed applications|58|11|1|2019-08-28T18:09:54Z|2021-08-24T05:39:27Z|
[go-astitodo](https://github.com/asticode/go-astitodo)|Parse TODOs in your GO code|58|9|2|2016-10-17T19:51:36Z|2020-08-17T22:56:15Z|
[goval](https://github.com/maja42/goval)|Expression evaluation in golang|58|6|0|2018-06-17T15:43:44Z|2021-02-02T17:11:01Z|
[golog](https://github.com/mlimaloureiro/golog)|Easy and simple CLI time tracker for your tasks|55|12|15|2016-01-09T15:43:47Z|2019-01-22T17:34:26Z|
[go-lock](https://github.com/viney-shih/go-lock)|go-lock is a lock library implementing read-write mutex and read-write trylock without starvation|52|5|0|2020-04-30T11:40:21Z|2021-07-26T14:06:14Z|
[copy-pasta](https://github.com/jutkko/copy-pasta)|Universal copy paste service, works across different machines!|49|10|10|2017-01-28T15:35:24Z|2020-06-20T13:33:28Z|
[filter](https://github.com/gookit/filter)|⏳ Provide filtering, sanitizing, and conversion of Golang data. 提供对Golang数据的过滤，净化，转换。|49|6|1|2018-09-26T09:11:13Z|2021-10-19T09:58:06Z|
[retry](https://github.com/thedevsaddam/retry)|Simple and easy retry mechanism package for Go|49|5|0|2018-02-25T19:08:03Z|2021-01-18T17:17:54Z|
[beyond](https://github.com/wesovilabs/beyond)|The Go library that will drive you to AOP world!|45|11|8|2019-10-18T05:41:45Z|2021-09-27T23:38:40Z|
[golarm](https://github.com/msempere/golarm)|Fire alarms with system events|45|9|0|2015-08-14T16:51:53Z|2015-08-24T13:33:34Z|
[slice](https://github.com/psampaz/slice)|Type-safe functions for common Go slice operations|45|5|1|2019-11-26T05:20:39Z|2020-04-09T15:24:07Z|
[goback](https://github.com/carlescere/goback)|Golang simple exponential backoff package.|44|8|6|2015-03-13T16:09:18Z|2021-03-09T23:40:57Z|
[dbt](https://github.com/nikogura/dbt)|Dynamic Binary Toolkit- A framework for running self-updating signed binaries from a central, trusted repository.|43|5|6|2017-11-30T22:53:17Z|2021-03-03T20:39:42Z|
[retry-go](https://github.com/rafaeljesus/retry-go)|Retrying made simple and easy for golang :repeat: |43|4|2|2017-06-09T16:07:37Z|2018-10-25T12:14:03Z|
**[ARCHIVED]**  [intrinsic](https://github.com/mengzhuo/intrinsic)|Provide Golang native SIMD intrinsics on x86/amd64 platform|41|1|1|2017-06-13T09:26:34Z|2017-06-23T01:17:03Z|
[gpath](https://github.com/tenntenn/gpath)|gpath is a Go package to access a field by a path using reflect pacakge|40|4|0|2017-05-24T06:24:18Z|2017-06-04T08:31:39Z|
[go-httpheader](https://github.com/mozillazg/go-httpheader)|A Go library for encoding structs into Header fields.|36|9|0|2017-06-24T11:28:06Z|2021-12-05T02:16:14Z|
**[ARCHIVED]**  [myhttp](https://github.com/inancgumus/myhttp)|Simplest HTTP GET requester for Go with timeout support|35|13|1|2017-09-13T15:48:47Z|2018-05-06T18:25:10Z|
[evaluator](https://github.com/nullne/evaluator)||33|9|0|2017-04-27T18:31:46Z|2021-07-25T13:59:51Z|
[backscanner](https://github.com/icza/backscanner)|A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward.|33|8|0|2017-10-18T07:59:07Z|2021-10-12T15:39:54Z|
[equalizer](https://github.com/reugn/equalizer)|A rate limiters package for Go|33|1|0|2019-06-14T09:25:13Z|2021-02-16T13:50:24Z|
[rclient](https://github.com/zpatrick/rclient)|Minimalistic REST client for Go applications|32|3|2|2017-02-28T01:07:25Z|2019-11-28T00:03:52Z|
[gostrutils](https://github.com/ik5/gostrutils)|Collections of string utils I have created over the years|31|6|1|2018-09-19T11:06:11Z|2021-09-11T08:18:12Z|
[slicer](https://github.com/leaanthony/slicer)|Utility class for handling slices|29|3|0|2019-01-10T09:55:25Z|2021-08-08T01:34:54Z|
[set](https://github.com/nofeaturesonlybugs/set)|Package set is a small wrapper around the official reflect package that facilitates loose type conversion and assignment into native Go types.|28|2|0|2020-12-16T22:12:18Z|2021-10-18T19:29:31Z|
[tome](https://github.com/cyruzin/tome)|Package tome was designed to paginate simple RESTful APIs.|27|3|1|2019-04-12T16:49:45Z|2020-04-09T13:51:46Z|
[throttle](https://github.com/yudppp/throttle)|lodash throttle like Go library|27|1|0|2019-10-25T14:30:38Z|2021-08-24T15:15:43Z|
[ugo](https://github.com/alxrm/ugo)|Simple and expressive toolbox written in Go|26|5|0|2016-02-17T19:41:57Z|2016-06-30T19:18:16Z|
[generate](https://github.com/go-playground/generate)|:runner:runs go generate recursively on a specified path or environment variable and can filter by regex|24|5|0|2015-11-15T01:52:04Z|2017-01-10T00:20:55Z|
[shutdown](https://github.com/ztrue/shutdown)|Golang app shutdown hooks.|24|5|0|2018-11-17T17:56:03Z|2019-02-08T22:37:34Z|
[sshman](https://github.com/shoobyban/sshman)|SSH Manager - manage authorized_keys file on remote servers|24|0|0|2021-08-27T13:04:47Z|2021-09-12T21:09:30Z|
[goplaceholder](https://github.com/michiwend/goplaceholder)|a small golang lib to generate placeholder images|22|7|1|2014-10-12T00:50:46Z|2016-01-17T18:24:14Z|
[rerate](https://github.com/abo/rerate)|redis-based rate counter and rate limiter|21|5|1|2016-05-24T08:59:00Z|2017-03-28T02:22:25Z|
[ghokin](https://github.com/antham/ghokin)|Parallelized formatter with no external dependencies for gherkin (cucumber, behat...)|20|1|3|2018-08-03T11:36:35Z|2021-09-22T06:45:19Z|
[structs](https://github.com/PumpkinSeed/structs)|Golang struct operations.|18|3|0|2017-08-26T09:59:00Z|2017-10-23T13:03:17Z|
[ctxutil](https://github.com/posener/ctxutil)|utils for Go context|17|3|1|2018-07-30T11:28:57Z|2020-03-01T00:49:08Z|
[mimesniffer](https://github.com/aofei/mimesniffer)|A MIME type sniffer for Go.|17|1|0|2018-12-20T03:40:20Z|2021-11-11T11:49:18Z|
[filler](https://github.com/yaronsumel/filler)|fill struct data easily with fill tags|16|4|0|2017-04-05T08:14:04Z|2017-04-10T08:03:38Z|
[okrun](https://github.com/xta/okrun)|ok, run your gofile|15|3|0|2014-10-01T06:18:56Z|2014-10-06T01:15:31Z|
[dlog](https://github.com/kirillDanshin/dlog)|Simple build-time controlled debug log with ability to log where the logger was called|15|2|0|2016-07-04T19:59:09Z|2017-07-28T00:08:08Z|
[rest-go](https://github.com/edermanoel94/rest-go)|A package that provide many helpful methods for working with rest api.|15|2|1|2019-07-29T18:56:08Z|2020-08-16T04:57:53Z|
[go-convert](https://github.com/Eun/go-convert)|Convert a value into another type|15|2|3|2019-06-07T16:56:38Z|2021-05-12T04:06:18Z|
[jsend](https://github.com/clevergo/jsend)|:100: JSend&#39;s implementation writen in Go(golang)|14|5|0|2020-01-14T04:41:36Z|2021-06-29T03:46:18Z|
[copy](https://github.com/gotidy/copy)|Package for fast copying structs of different types|14|2|4|2020-10-09T06:59:08Z|2020-12-28T08:02:43Z|
[command](https://github.com/txgruppi/command)|Command pattern for Go with thread safe serial and parallel dispatcher|13|4|0|2015-08-24T09:43:50Z|2016-04-20T17:06:57Z|
[go-types](https://github.com/mikekonan/go-types)|Library providing opanapi3 and Go types for store/validation and transfer of ISO-4217, ISO-3166, and other types.|12|6|0|2021-04-21T11:34:25Z|2021-11-18T14:15:12Z|
[ptr](https://github.com/gotidy/ptr)|Contains functions for simplified creation of pointers from constants of basic types|11|3|0|2019-12-25T15:29:48Z|2020-11-03T12:03:47Z|
[go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails)|Problem json implementation (https://tools.ietf.org/html/rfc7807) package for go|11|1|0|2019-05-16T05:42:14Z|2020-02-17T11:12:12Z|
[silk](https://github.com/chrispassas/silk)|Read Silk Flow Files|10|2|0|2018-12-18T04:23:35Z|2020-11-13T14:12:01Z|
[retry](https://github.com/shafreeck/retry)|A pretty simple library to ensure your work to be done|10|2|1|2018-07-18T09:48:33Z|2020-02-11T03:47:03Z|
[go-countries](https://github.com/mikekonan/go-countries)||9|4|0|2020-10-27T12:56:40Z|2020-12-17T15:41:16Z|
[statiks](https://github.com/janiltonmaciel/statiks)|Fast, zero-configuration, static HTTP filer server.|9|1|0|2018-06-26T23:42:33Z|2020-10-06T20:27:09Z|
[retry](https://github.com/percolate/retry)|Percolate&#39;s Go retry package|7|2|0|2018-06-15T19:23:36Z|2019-09-05T21:13:28Z|
[sliceconv](https://github.com/Henry-Sarabia/sliceconv)|Slice conversion between primitive types|7|1|0|2019-02-15T06:50:34Z|2020-02-03T04:41:41Z|
[go-clip](https://github.com/prashantgupta24/go-clip)|A minimalistic clipboard manager for Mac.|7|0|2|2020-11-18T22:19:01Z|2021-02-05T17:37:54Z|
[nfdump](https://github.com/chrispassas/nfdump)|NFDump File Reader|6|1|0|2020-04-08T01:01:22Z|2021-08-11T17:23:13Z|
[blank](https://github.com/Henry-Sarabia/blank)|Detect blank strings or remove whitespace from strings|6|1|0|2019-02-13T00:07:27Z|2019-07-31T23:16:14Z|
[go-safe](https://github.com/kenkyu392/go-safe)|This Go package provides a sandbox for the safe execution of panic-inducing programs|5|1|0|2019-10-29T15:20:37Z|2021-11-30T08:24:38Z|
[lets-go](https://github.com/aplescia/lets-go)|Go module that provides common utilities for Cloud Native development|3|1|0|2020-02-19T16:32:41Z|2021-04-24T17:30:07Z|
[tik](https://github.com/andy2046/tik)|hierarchical timing wheel|3|1|0|2020-07-04T09:13:49Z|2020-10-17T03:23:45Z|
[go-actuator](https://github.com/sinhashubham95/go-actuator)|Golang production-ready features|3|0|0|2021-07-17T05:47:50Z|2021-08-26T09:59:54Z|
[goctx](https://github.com/zerosnake0/goctx)|Get your context value faster|2|2|0|2020-11-14T14:16:09Z|2020-11-24T14:42:11Z|
[bleep](https://github.com/sinhashubham95/bleep)|OS Signal Handlers in Go|2|1|0|2021-01-02T05:22:08Z|2021-01-06T03:41:42Z|
[olaf](https://github.com/btnguyen2k/olaf)|Twitter Snowflake implemented in Go|2|1|0|2019-01-03T13:31:10Z|2019-04-10T08:59:20Z|


## UUID
*Libraries for working with UUIDs.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[uuid](https://github.com/google/uuid)|Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.|3219|282|19|2016-02-12T22:17:59Z|2021-12-06T11:29:24Z|
[ulid](https://github.com/oklog/ulid)|Universally Unique Lexicographically Sortable Identifier (ULID) in Go|2461|95|1|2016-12-06T15:26:52Z|2021-10-20T22:07:29Z|
[xid](https://github.com/rs/xid)|xid is a globally unique id generator thought for the web|2335|145|10|2015-11-10T20:32:24Z|2021-05-03T15:45:47Z|
[uuid](https://github.com/gofrs/uuid)|A UUID package originally forked from github.com/satori/go.uuid|1035|72|7|2018-07-13T02:13:28Z|2021-12-10T01:04:18Z|
[wuid](https://github.com/edwingeng/wuid)|An extremely fast UUID alternative written in golang|442|40|0|2018-01-27T01:16:28Z|2021-08-17T16:55:10Z|
[sno](https://github.com/muyo/sno)|Compact, sortable and fast unique IDs with embedded metadata.|54|3|0|2019-05-26T22:05:26Z|2021-11-12T01:59:41Z|
[nanoid](https://github.com/aidarkhanov/nanoid)|A tiny and fast Go unique string generator|39|5|0|2019-07-02T12:15:56Z|2021-09-15T22:25:23Z|
[Goid](https://github.com/JakeHL/Goid)|A UUIDv4 generation package written in go|29|2|1|2017-05-19T10:40:45Z|2019-02-18T15:50:01Z|
[uuid](https://github.com/agext/uuid)|Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier.|13|4|0|2016-02-03T03:02:51Z|2020-03-12T22:02:03Z|
[gouid](https://github.com/twharmon/gouid)|Fast, dependable universally unique ids|10|1|0|2020-10-08T19:54:41Z|2021-05-24T14:06:53Z|
[goflake](https://github.com/Hart87/goflake)|A highly scalable and serverless unique ID generator for use in distributed systems. Written in GoLang. Inspired by Twitters Snowflake.|8|0|0|2021-05-03T14:44:19Z|2021-05-17T13:58:55Z|
[uniq](https://gitlab.com/skilstak/code/go/uniq)|No hassle safe, fast unique identifiers with commands.|-|-|-|-|-|


## Validation
*Libraries for validation.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[validator](https://github.com/go-playground/validator)|:100:Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving|9218|826|172|2015-02-12T16:32:22Z|2021-11-18T19:44:49Z|
[govalidator](https://github.com/asaskevich/govalidator)|[Go] Package of validators and sanitizers for strings, numerics, slices and structs|5153|506|143|2014-06-20T10:45:23Z|2021-10-18T13:22:29Z|
[ozzo-validation](https://github.com/go-ozzo/ozzo-validation)|An idiomatic Go (golang) validation package. Supports configurable and extensible validation rules (validators) using normal language constructs instead of error-prone struct tags.|2463|155|28|2016-06-22T03:47:43Z|2021-12-02T08:01:32Z|
[govalidator](https://github.com/thedevsaddam/govalidator)|Validate Golang request data with simple rules. Highly inspired by Laravel&#39;s request validation.|1030|87|30|2017-09-13T16:42:20Z|2021-01-18T17:16:29Z|
[validate](https://github.com/gookit/validate)|⚔ Go package for data validation and filtering. support Map, Struct, Form data. Go通用的数据验证与过滤库，使用简单，内置大部分常用验证、过滤器，支持自定义验证器、自定义消息、字段翻译。|497|71|13|2018-07-16T08:23:49Z|2021-11-13T13:03:40Z|
[checkdigit](https://github.com/osamingo/checkdigit)|Provide check digit algorithms and calculators written in Go|83|4|0|2019-04-05T09:46:36Z|2021-01-01T07:36:39Z|
[terraform-validator](https://github.com/thazelart/terraform-validator)|A norms and conventions validator for Terraform|73|5|6|2019-05-29T11:37:15Z|2020-09-20T13:52:37Z|
[validate](https://github.com/gobuffalo/validate)|This package provides a framework for writing validations for Go applications.|65|20|4|2018-02-10T18:25:55Z|2021-11-11T18:51:11Z|
[jio](https://github.com/faceair/jio)|jio is a json schema validator similar to joi|62|10|0|2018-10-28T11:02:45Z|2020-05-08T16:22:47Z|
[gody](https://github.com/guiferpa/gody)|:balloon: A lightweight struct validator for Go|52|4|1|2018-11-01T21:08:16Z|2021-02-02T15:18:35Z|
[govalid](https://github.com/twharmon/govalid)|Struct validation using tags|24|5|1|2019-02-17T23:25:43Z|2021-10-14T17:46:17Z|


## Version Control
*Libraries for version control.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[go-git](https://github.com/go-git/go-git)|A highly extensible Git implementation in pure Go.|2849|329|243|2019-12-19T10:27:02Z|2021-12-11T20:32:32Z|
[git2go](https://github.com/libgit2/git2go)|Git to Go; bindings for libgit2. Like McDonald&#39;s but tastier.|1682|293|43|2013-03-05T19:50:43Z|2021-11-28T20:20:59Z|
[glab](https://github.com/profclems/glab)|An open-source GitLab command line tool bringing GitLab&#39;s cool features to your command line|1605|134|89|2020-07-24T20:36:56Z|2021-12-02T16:49:14Z|
[hercules](https://github.com/src-d/hercules)|Gaining advanced insights from Git repository history.|1502|129|40|2016-12-12T17:30:29Z|2021-11-08T12:45:48Z|
[gh](https://github.com/rjeczalik/gh)|Scriptable server and net/http middleware for GitHub Webhooks.|76|13|2|2015-03-08T21:04:05Z|2018-10-28T15:27:35Z|
[go-vcs](https://github.com/sourcegraph/go-vcs)|manipulate and inspect VCS repositories in Go|74|22|23|2013-06-02T02:36:18Z|2021-03-31T12:37:46Z|
[Githooks](https://github.com/gabyx/Githooks)|🦎 Githooks: per-repo and shared Git hooks with version control and auto update. |29|0|3|2019-06-28T06:28:55Z|2021-12-11T21:07:49Z|
[hgo](https://github.com/beyang/hgo)|Hgo is a collection of Go packages providing read-access to local Mercurial repositories.|13|4|0|2014-06-18T03:54:40Z|2015-08-25T03:56:31Z|


## Video
*Libraries for manipulating video.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[goav](https://github.com/giorgisio/goav)|Golang bindings for FFmpeg|1757|320|47|2015-05-21T05:31:14Z|2021-06-11T10:20:05Z|
[m3u8](https://github.com/grafov/m3u8)|Parser and generator of M3U8-playlists for Apple HLS. Library for Go language. :cinema:|878|234|46|2013-02-05T22:26:30Z|2021-12-01T12:25:08Z|
[gmf](https://github.com/3d0c/gmf)|Go Media Framework|724|145|42|2013-04-03T09:07:47Z|2021-10-01T08:04:59Z|
[go-astits](https://github.com/asticode/go-astits)|Demux and mux MPEG Transport Streams (.ts) natively in GO|381|36|5|2017-07-04T13:06:15Z|2021-10-01T11:22:09Z|
[go-astisub](https://github.com/asticode/go-astisub)|Manipulate subtitles in GO (.srt, .ssa/.ass, .stl, .ttml, .vtt (webvtt), teletext, etc.)|361|74|8|2016-12-16T14:47:59Z|2021-12-04T13:01:08Z|
[libvlc-go](https://github.com/adrg/libvlc-go)|Go bindings for libVLC and high-level media player interface|261|34|3|2015-01-06T14:01:50Z|2021-09-24T12:00:06Z|
[gst](https://github.com/ziutek/gst)|Go bindings for GStreamer (retired: currently I don&#39;t use/develop this package)|162|46|9|2011-07-26T00:44:40Z|2021-01-07T12:04:16Z|
[gortsplib](https://github.com/aler9/gortsplib)|RTSP 1.0 client and server library for the Go programming language|158|50|6|2020-01-20T09:08:24Z|2021-12-10T22:02:02Z|
[go-m3u8](https://github.com/quangngotan95/go-m3u8)|Parse and generate m3u8 playlists for Apple HTTP Live Streaming (HLS) in Golang (ported from gem https://github.com/sethdeckard/m3u8)|83|13|1|2018-11-06T02:42:27Z|2020-05-14T04:36:59Z|
[v4l](https://github.com/korandiz/v4l)|Facade to the Video4Linux video capture interface. |63|12|0|2016-10-25T10:50:25Z|2021-10-06T17:03:25Z|
[libgosubs](https://github.com/wargarblgarbl/libgosubs)|golang library to read and write various subtitle formats|16|4|0|2017-05-03T21:05:25Z|2020-05-13T06:18:07Z|
[go-mpd](https://github.com/unki2aut/go-mpd)|Go library for parsing and generating MPEG-DASH Media Presentation Description (MPD) files|10|5|0|2018-11-02T19:09:07Z|2020-08-18T09:32:36Z|


## Web Frameworks
*Full stack web frameworks.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[gin](https://github.com/gin-gonic/gin)|Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.|53757|6107|460|2014-06-16T23:57:25Z|2021-12-11T08:24:10Z|
[beego](https://github.com/beego/beego)|beego is an open-source, high-performance web framework for the Go programming language.|27360|5363|26|2012-02-29T02:32:08Z|2021-12-09T14:30:13Z|
[echo](https://github.com/labstack/echo)|High performance, minimalist Go web framework|21242|1886|80|2015-03-01T17:43:01Z|2021-12-09T19:57:26Z|
[fiber](https://github.com/gofiber/fiber)|⚡️ Express inspired web framework written in Go|17198|872|41|2020-01-16T03:59:20Z|2021-12-11T14:40:47Z|
[revel](https://github.com/revel/revel)|A high productivity, full-stack web framework for the Go language.|12460|1403|108|2011-12-09T04:10:26Z|2021-11-03T09:35:38Z|
[goa](https://github.com/goadesign/goa)|Design-based APIs and microservices in Go|4476|465|11|2014-12-05T07:17:53Z|2021-12-07T20:25:19Z|
[gizmo](https://github.com/nytimes/gizmo)|A Microservice Toolkit from The New York Times|3563|231|26|2015-12-15T18:09:36Z|2021-08-03T10:55:58Z|
[go-json-rest](https://github.com/ant0ine/go-json-rest)|A quick and easy way to setup a RESTful JSON API|3490|389|49|2013-02-19T03:15:45Z|2021-01-23T18:47:50Z|
[macaron](https://github.com/go-macaron/macaron)|Package macaron is a high productive and modular web framework in Go.|3232|286|6|2014-07-10T03:13:30Z|2021-11-28T06:52:28Z|
[utron](https://github.com/gernest/utron)|A lightweight MVC framework for Go(Golang)|2204|158|9|2015-09-16T07:55:54Z|2018-10-28T20:04:59Z|
[go-tigertonic](https://github.com/rcrowley/go-tigertonic)|A Go framework for building JSON web services inspired by Dropwizard|997|75|28|2013-02-09T21:16:13Z|2018-07-24T09:26:32Z|
[goyave](https://github.com/go-goyave/goyave)|🍐 Elegant Golang REST API Framework|917|42|5|2019-10-21T09:44:34Z|2021-11-23T13:18:41Z|
**[ARCHIVED]**  [tango](https://github.com/lunny/tango)|This is only a mirror and Moved to https://gitea.com/lunny/tango|834|106|9|2014-12-17T03:07:09Z|2019-05-17T03:31:14Z|
[aah](https://github.com/go-aah/aah)|A secure, flexible, rapid Go web framework|659|37|16|2016-06-27T04:47:45Z|2020-09-02T02:31:21Z|
[gearbox](https://github.com/gogearbox/gearbox)|Gearbox :gear: is a web framework written in Go with a focus on high performance|580|46|2|2020-04-25T01:28:37Z|2021-12-05T08:35:45Z|
[gongular](https://github.com/mustafaakin/gongular)|A different approach to Go web frameworks|446|16|8|2016-06-22T11:52:42Z|2020-07-05T14:40:50Z|
[aero](https://github.com/aerogo/aero)|:bullettrain_side: High-performance web server for Go.|418|25|4|2016-11-09T13:02:13Z|2021-11-20T11:42:50Z|
[neo](https://github.com/ivpusic/neo)|Go Web Framework|410|41|6|2015-02-04T19:16:06Z|2017-12-30T07:35:36Z|
[air](https://github.com/aofei/air)|An ideally refined web framework for Go.|407|40|4|2016-07-20T12:09:48Z|2021-04-18T10:29:01Z|
[mango](https://github.com/paulbellamy/mango)|Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.|359|38|9|2011-05-25T07:26:46Z|2017-10-17T08:18:44Z|
[confetti](https://github.com/confetti-framework/confetti)|Confetti is a web application framework with an expressive, elegant syntax. This repository contains configuration files and is intended as a template for your codebase. Download these configuration files and include them in your git repository.|323|16|0|2019-11-01T23:14:21Z|2021-11-17T22:06:10Z|
[gondola](https://github.com/rainycape/gondola)|The web framework for writing faster sites, faster|308|23|8|2014-07-25T21:28:55Z|2019-02-19T00:41:28Z|
[golf](https://github.com/dinever/golf)|:golf: The Golf web framework|250|29|6|2015-11-18T15:10:14Z|2021-08-27T22:20:34Z|
[flamingo](https://github.com/i-love-flamingo/flamingo)|Flamingo Framework and Core Library. Flamingo is a go based framework for pluggable web projects. It is used to build scalable and maintainable (web)applications.|244|33|19|2019-04-02T12:24:02Z|2021-11-30T16:34:06Z|
[flamingo-commerce](https://github.com/i-love-flamingo/flamingo-commerce)|Flexible E-Commerce Framework on top of Flamingo. Used to build E-Commerce &#34;Portals&#34; and connect it with the help of individual Adapters to other services. |240|41|18|2019-04-02T15:11:57Z|2021-12-01T15:25:00Z|
[webgo](https://github.com/bnkamalesh/webgo)|A microframework to build web apps; with handler chaining, middleware support, and most of all; standard library compliant HTTP handlers(i.e. http.HandlerFunc).|201|20|4|2015-12-16T07:35:02Z|2021-10-12T12:39:44Z|
[ginrpc](https://github.com/xxjwxc/ginrpc)|gin auto binding,grpc, and annotated route,gin 注解路由, grpc,自动参数绑定工具|195|24|7|2019-06-22T12:03:53Z|2021-11-12T07:20:54Z|
[uadmin](https://github.com/uadmin/uadmin)|The web framework for Golang|162|37|15|2018-10-05T09:00:17Z|2021-11-03T05:22:35Z|
[hiboot](https://github.com/hidevopsio/hiboot)|hiboot is a high performance web and cli application framework with dependency injection support|162|28|6|2018-03-16T11:21:46Z|2021-06-29T05:34:35Z|
[go-rest](https://github.com/ungerik/go-rest)|A small and evil REST framework for Go|125|14|2|2012-07-13T10:02:15Z|2017-01-20T13:26:12Z|
[appy](https://github.com/appist/appy)|An opinionated productive web framework that helps scaling business easier.|112|12|1|2019-05-27T04:48:59Z|2021-11-25T09:01:09Z|
[patron](https://github.com/beatlabs/patron)|Microservice framework following best cloud practices with a focus on productivity.|84|52|24|2019-01-30T13:49:54Z|2021-12-10T20:40:39Z|
[microservice](https://github.com/claygod/microservice)|This library provides a simple microservice framework based on clean architecture principles with a working example implemented.|74|13|0|2016-12-15T09:07:04Z|2021-12-05T14:55:39Z|
[vox](https://github.com/aisk/vox)|Simple and lightweight Go web framework inspired by koa|74|5|7|2014-12-24T11:22:08Z|2021-05-31T16:20:33Z|
[golax](https://github.com/fulldump/golax)|Golax, a go implementation for the Lax framework.|72|7|6|2016-01-30T19:11:39Z|2018-06-03T17:25:13Z|
[rux](https://github.com/gookit/rux)|⚡ Rux is an simple and fast web framework. support param route binding, middleware, compatible http.Handler interface. 简单且快速的 Go api/web 框架，支持参数路由绑定，中间件，兼容 http.Handler 接口|69|12|1|2018-08-05T06:13:57Z|2021-10-19T09:57:43Z|
[rk-boot](https://github.com/rookie-ninja/rk-boot)|Bootstrapper for golang application. See https://rkdev.info/docs/ for details.|63|9|3|2020-07-31T02:36:56Z|2021-12-10T19:46:42Z|
[yarf](https://github.com/yarf-framework/yarf)|Yet Another REST Framework|63|7|2|2015-09-02T13:56:47Z|2019-03-07T20:28:46Z|
[fireball](https://github.com/zpatrick/fireball)|Go web framework with a natural feel|55|5|1|2016-07-20T05:04:54Z|2018-10-03T21:26:08Z|
[goa](https://github.com/goa-go/goa)|Goa is a  web framework based on middleware, like koa.js.|45|2|0|2019-07-26T07:12:23Z|2019-12-06T10:29:45Z|
[gotuna](https://github.com/gotuna/gotuna)|GoTuna a lightweight web framework for Go with mux router, middlewares, user sessions, templates, embedded views, and static file server.|36|4|1|2021-04-08T14:08:08Z|2021-07-23T09:10:58Z|
[api](https://github.com/resoursea/api)|A REST framework for quickly writing resource based services in Golang.|31|3|0|2015-01-24T18:45:30Z|2015-02-01T22:58:21Z|
[rex](https://github.com/goanywhere/rex)|Pleasures for Web in Golang|31|2|0|2014-10-16T02:26:18Z|2017-12-22T03:25:41Z|
[goweb](https://github.com/twharmon/goweb)|Lightweight web framework based on net/http.|25|2|1|2019-05-07T21:04:43Z|2021-02-11T14:31:28Z|
[banjo](https://github.com/n4Zz2/banjo)|BANjO is a simple web framework written in Go (golang)|18|6|4|2017-12-09T13:35:31Z|2018-01-31T16:42:14Z|
[REST Layer](https://rest-layer.io)|Framework to build REST/GraphQL API on top of databases with mostly configuration over code.|-|-|-|-|-|
[Buffalo](https://gobuffalo.io)|Bringing the productivity of Rails to Go!|-|-|-|-|-|


#### Actual middlewares

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[tollbooth](https://github.com/didip/tollbooth)|Simple middleware to rate-limit HTTP requests.|2070|188|6|2015-05-17T15:20:03Z|2021-12-03T15:48:35Z|
[cors](https://github.com/rs/cors)|Go net/http configurable handler to handle CORS requests|1975|178|19|2014-10-25T03:49:45Z|2021-11-18T14:45:08Z|
[limiter](https://github.com/ulule/limiter)|Dead simple rate limit middleware for Go.|1455|112|6|2015-10-02T08:12:38Z|2021-11-26T12:04:37Z|
[go-server-timing](https://github.com/mitchellh/go-server-timing)|Go (golang) library for creating and consuming HTTP Server-Timing headers|821|29|5|2018-02-12T03:56:02Z|2020-11-08T05:50:53Z|
[go-fault](https://github.com/github/go-fault)|Fault injection library in Go using standard http middleware|419|20|0|2020-05-14T16:13:17Z|2021-09-15T16:13:09Z|
[ln-paywall](https://github.com/philippgille/ln-paywall)|Go middleware for monetizing your API on a per-request basis with Bitcoin and Lightning ⚡️|116|8|17|2018-06-29T21:51:00Z|2019-02-24T19:40:57Z|
[xff](https://github.com/sebest/xff)|A Golang Middleware to handle X-Forwarded-For Header|85|20|5|2014-12-22T10:29:05Z|2021-01-08T15:26:50Z|
[formjson](https://github.com/rs/formjson)|Go net/http handler to transparently manage posted JSON|35|2|0|2015-03-19T23:52:28Z|2015-12-17T09:35:29Z|
[client-timing](https://github.com/posener/client-timing)|An HTTP client for go-server-timing middleware. Enables automatic timing propagation through HTTP calls between servers.|19|5|1|2018-02-23T01:52:45Z|2020-03-13T18:47:59Z|
[rk-grpc](https://github.com/rookie-ninja/rk-grpc)|grpc related entry. See https://rkdev.info/docs/ for details.|18|3|2|2020-07-25T20:33:46Z|2021-12-09T13:39:33Z|
[rk-gin](https://github.com/rookie-ninja/rk-gin)|Bootstrapper and interceptor for gin framework. See https://rkdev.info/docs/ for details.|13|4|1|2020-10-12T16:48:48Z|2021-12-09T13:38:36Z|
[mid](https://github.com/bobg/mid)|Middleware for HTTP services in Go|2|0|0|2020-07-13T14:53:59Z|2021-08-19T00:34:16Z|


#### Libraries for creating HTTP middlewares

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[negroni](https://github.com/urfave/negroni)|Idiomatic HTTP Middleware for Golang|7112|570|11|2014-05-18T22:09:10Z|2021-11-29T15:52:40Z|
[alice](https://github.com/justinas/alice)|Painless middleware chaining for Go|2432|135|6|2014-05-25T07:27:41Z|2021-11-04T17:40:07Z|
[render](https://github.com/unrolled/render)|Go package for easily rendering JSON, XML, binary data, and HTML templates responses.|1562|126|1|2014-06-10T16:20:35Z|2021-11-11T13:22:41Z|
[stats](https://github.com/thoas/stats)|A Go middleware that stores various information about your web application (response time, status code count, etc.)|581|49|8|2015-03-05T18:02:50Z|2019-04-07T19:46:42Z|
[interpose](https://github.com/carbocation/interpose)|Minimalist net/http middleware for golang|295|15|1|2014-07-20T00:19:52Z|2016-12-06T21:52:53Z|
[renderer](https://github.com/thedevsaddam/renderer)|Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go|229|25|0|2017-11-07T18:53:49Z|2021-01-18T17:17:13Z|
[muxchain](https://github.com/stephens2424/muxchain)|Lightweight Middleware for net/http|209|14|1|2014-05-03T17:14:17Z|2019-03-19T21:44:51Z|
[rye](https://github.com/InVisionApp/rye)|A tiny http middleware for Golang with added handlers for common needs.|95|13|0|2016-10-06T19:51:59Z|2018-10-04T15:00:04Z|
[gores](https://github.com/alioygur/gores)|Go package that handles HTML, JSON, XML and etc. responses|95|3|0|2015-12-25T12:41:01Z|2021-01-01T12:48:26Z|
[mediary](https://github.com/HereMobilityDevelopers/mediary)|Add interceptors to GO http.Client|78|6|0|2020-03-23T18:54:56Z|2020-06-24T14:38:59Z|
[chain](https://github.com/codemodus/chain)|Composable chains of nested http.Handler instances.|63|4|0|2015-05-14T19:52:58Z|2018-08-25T20:35:40Z|
**[ARCHIVED]**  [wrap](https://github.com/go-on/wrap)|Go http.Hander based middleware stack with context sharing|59|5|0|2014-02-16T07:12:36Z|2018-08-15T19:29:53Z|
[catena](https://github.com/codemodus/catena)|gRPC interceptor catenation.|7|1|0|2015-07-30T19:07:01Z|2018-08-25T22:06:48Z|


### Routers

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[mux](https://github.com/gorilla/mux)|A powerful HTTP router and URL matcher for building Go web servers with 🦍|15597|1449|31|2012-10-02T21:32:24Z|2021-12-08T21:25:43Z|
[httprouter](https://github.com/julienschmidt/httprouter)|A high performance HTTP request router that scales well|13477|1291|65|2013-12-05T15:10:55Z|2021-11-10T03:59:23Z|
[chi](https://github.com/go-chi/chi)|lightweight, idiomatic and composable router for building Go HTTP services|10541|722|38|2015-10-15T20:46:29Z|2021-12-03T11:23:28Z|
[web](https://github.com/gocraft/web)|Go Router &#43; Middleware. Your Contexts.|1450|123|24|2013-11-16T20:48:20Z|2020-10-01T09:54:18Z|
[bone](https://github.com/go-zoo/bone)|Lightning Fast HTTP Multiplexer|1275|86|3|2014-11-19T02:16:36Z|2019-05-06T14:37:24Z|
[goji](https://github.com/goji/goji)|Goji is a minimalistic and flexible HTTP request multiplexer for Go (golang)|892|64|5|2015-11-16T00:52:41Z|2019-08-01T15:36:58Z|
[fasthttprouter](https://github.com/buaazp/fasthttprouter)|A high performance fasthttp request router that scales well|871|91|19|2015-12-13T09:32:30Z|2019-04-25T14:24:36Z|
[httptreemux](https://github.com/dimfeld/httptreemux)|High-speed, flexible tree-based HTTP router for Go.|515|48|4|2014-05-14T20:10:20Z|2021-11-07T07:42:14Z|
[gorouter](https://github.com/xujiajun/gorouter)|xujiajun/gorouter is a simple and fast HTTP router for Go. It is easy to build RESTful APIs and your web framework.|508|83|0|2018-01-29T09:28:28Z|2019-09-27T07:07:43Z|
[ozzo-routing](https://github.com/go-ozzo/ozzo-routing)|An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs.|420|49|10|2015-10-27T01:03:14Z|2021-09-11T21:43:46Z|
[lars](https://github.com/go-playground/lars)|:rotating_light: Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.|384|24|1|2015-12-24T17:28:45Z|2019-05-15T21:58:32Z|
[siesta](https://github.com/VividCortex/siesta)|Composable framework for writing HTTP handlers in Go.|348|15|0|2014-09-23T13:55:56Z|2021-04-26T21:52:25Z|
[vestigo](https://github.com/husobee/vestigo)|Echo Inspired Stand Alone URL Router|265|30|14|2015-09-22T03:08:03Z|2020-10-08T16:23:52Z|
[router](https://github.com/gowww/router)|⚡️ A lightning fast HTTP router|159|13|0|2017-05-25T10:29:27Z|2020-05-04T16:39:26Z|
[alien](https://github.com/gernest/alien)|A lightweight and  fast http router from outer space|120|11|3|2016-01-30T23:23:10Z|2019-03-23T07:13:30Z|
[pure](https://github.com/go-playground/pure)|:non-potable_water: Is a lightweight  HTTP router that sticks to the std &#34;net/http&#34; implementation|119|10|0|2016-09-23T19:57:58Z|2020-11-19T05:20:04Z|
[gorouter](https://github.com/vardius/gorouter)|Go Server/API micro framework, HTTP request router, multiplexer, mux|103|14|5|2016-07-14T13:13:34Z|2021-11-14T03:53:56Z|
[violetear](https://github.com/nbari/violetear)|Go HTTP router|102|9|1|2015-06-19T16:49:41Z|2021-05-25T14:39:05Z|
[Bxog](https://github.com/claygod/Bxog)|Bxog is a simple and fast HTTP router for Go (HTTP request multiplexer).|101|7|0|2016-05-19T12:20:08Z|2020-06-12T14:56:00Z|
[xmux](https://github.com/rs/xmux)|xmux is a httprouter fork on top of xhandler (net/context aware)|92|10|2|2015-12-14T19:01:05Z|2017-06-09T18:54:18Z|
[bellt](https://github.com/GuilhermeCaruso/bellt)|:bell: A simple Go router|51|5|0|2019-02-21T13:13:52Z|2020-06-18T03:03:14Z|
[ngamux](https://github.com/ngamux/ngamux)|Simple HTTP router for Go|45|16|1|2021-08-22T08:31:40Z|2021-10-26T04:06:54Z|
[fastrouter](https://github.com/razonyang/fastrouter)|FastRouter is a fast, flexible HTTP router written in Go.|19|4|0|2017-11-01T08:52:52Z|2017-11-03T15:05:25Z|
[route](https://github.com/goroute/route)|Go Route - Simple yet powerful HTTP request multiplexer|6|2|1|2019-07-06T18:47:38Z|2019-12-23T20:20:48Z|


## WebAssembly

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[tinygo](https://github.com/tinygo-org/tinygo)|Go compiler for small places. Microcontrollers, WebAssembly (WASM/WASI), and command-line tools. Based on LLVM.|9032|497|315|2018-06-07T16:39:19Z|2021-12-11T20:32:25Z|
[dom](https://github.com/dennwc/dom)|DOM library for Go and WASM|437|52|11|2018-06-30T18:37:35Z|2019-09-26T14:33:41Z|
[go-canvas](https://github.com/markfarnan/go-canvas)|Library to use HTML5 Canvas  from Go-WASM, with all drawing within go code|148|11|5|2019-05-05T14:05:55Z|2020-12-09T22:42:50Z|
[wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest)|Run WASM tests inside your browser|109|13|1|2018-07-14T18:42:24Z|2021-10-31T08:29:53Z|
[webapi](https://github.com/gowebapi/webapi)|Go Lang Web Assembly bindings for DOM, HTML etc|91|10|2|2019-02-08T05:58:35Z|2021-01-30T09:07:16Z|
[vert](https://github.com/norunners/vert)|WebAssembly interop between Go and JS values.|57|9|2|2018-03-25T17:26:47Z|2021-03-20T05:09:53Z|


## Windows

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[go-ole](https://github.com/go-ole/go-ole)|win32 ole implementation for golang|813|151|61|2011-01-21T12:45:20Z|2021-11-24T05:58:10Z|
[d3d9](https://github.com/gonutz/d3d9)|Direct3D9 wrapper for Go.|128|11|1|2015-12-12T21:24:38Z|2021-12-10T17:39:50Z|
[gosddl](https://github.com/MonaxGT/gosddl)|GoSDDL converter|8|2|0|2018-12-04T08:36:11Z|2019-04-30T10:04:14Z|


## XML
*Libraries and tools for manipulating XML.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[zek](https://github.com/miku/zek)|Generate a Go struct from XML.|515|44|8|2017-11-23T19:03:11Z|2021-09-23T09:10:24Z|
[xpath](https://github.com/antchfx/xpath)|XPath package for Golang, supports HTML, XML, JSON document query.|435|63|10|2016-10-09T05:51:24Z|2021-09-25T03:05:03Z|
**[ARCHIVED]**  [xquery](https://github.com/antchfx/xquery)|Extract data or evaluate value from HTML/XML documents using XPath|153|26|0|2016-10-09T05:54:10Z|2018-05-15T05:19:11Z|
[xml2map](https://github.com/sbabiv/xml2map)|XML to MAP converter written Golang|37|9|1|2018-08-06T17:51:46Z|2021-12-07T20:49:48Z|
[xmlwriter](https://github.com/shabbyrobe/xmlwriter)|xmlwriter is a pure-Go library providing procedural XML generation based on libxml2&#39;s xmlwriter module|20|4|1|2017-04-11T04:43:26Z|2021-03-24T11:07:52Z|
[XML-Comp](https://github.com/XML-Comp/XML-Comp)|Compare ANY markup documents.|15|10|8|2016-10-25T22:09:12Z|2018-07-19T12:21:08Z|


## Zero Trust
*Libraries and tools to implement Zero Trust architectures.*
	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[cosign](https://github.com/sigstore/cosign)|Container Signing|1343|143|154|2021-02-04T12:49:39Z|2021-12-11T17:52:45Z|
[spire](https://github.com/spiffe/spire)|The SPIFFE Runtime Environment|1011|273|216|2017-08-11T18:46:51Z|2021-12-10T21:16:08Z|
[in-toto-golang](https://github.com/in-toto/in-toto-golang)|A Go implementation of in-toto. in-toto is a framework to protect software supply chain integrity.|43|35|23|2018-10-15T15:18:06Z|2021-12-11T12:43:59Z|
[spiffe-vault](https://github.com/philips-labs/spiffe-vault)|Integrates Spiffe and Vault to have secretless authentication|8|0|1|2021-08-26T10:53:00Z|2021-11-30T11:20:22Z|


## Code Analysis

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[tools](https://github.com/golang/tools)|[mirror] Go Tools|5734|1899|57|2014-11-25T21:07:26Z|2021-12-11T05:39:52Z|
[go-tools](https://github.com/dominikh/go-tools)|Staticcheck - The advanced Go linter|4296|279|441|2017-01-24T21:11:01Z|2021-12-01T16:45:04Z|
[go-tools](https://github.com/dominikh/go-tools)|Staticcheck - The advanced Go linter|4296|279|441|2017-01-24T21:11:01Z|2021-12-01T16:45:04Z|
[go-tools](https://github.com/dominikh/go-tools)|Staticcheck - The advanced Go linter|4296|279|441|2017-01-24T21:11:01Z|2021-12-01T16:45:04Z|
**[ARCHIVED]**  [lint](https://github.com/golang/lint)|[mirror] This is a linter for Go source code. (deprecated)|3925|527|0|2013-06-02T22:45:37Z|2021-05-08T22:21:20Z|
[errcheck](https://github.com/kisielk/errcheck)|errcheck checks that you checked errors.|1760|128|9|2013-02-24T22:32:02Z|2021-06-25T16:39:53Z|
[go-critic](https://github.com/go-critic/go-critic)|The most opinionated Go source code linter for code audit.|1188|84|108|2018-05-05T09:17:26Z|2021-12-10T05:40:43Z|
[gcvis](https://github.com/davecheney/gcvis)|Visualise Go program GC trace data in real time|1036|69|10|2014-07-10T12:34:07Z|2019-03-13T01:20:26Z|
[php-parser](https://github.com/z7zmey/php-parser)|PHP parser written in Go|840|63|18|2017-11-07T06:20:46Z|2021-04-28T03:22:19Z|
[goplantuml](https://github.com/jfeliu007/goplantuml)|PlantUML Class Diagram Generator for golang projects|591|62|23|2019-05-26T15:43:12Z|2021-12-03T09:12:15Z|
[goast-viewer](https://github.com/yuroyoro/goast-viewer)|Golang AST visualizer|562|51|1|2014-06-30T11:09:01Z|2019-05-31T02:48:19Z|
[go-mod-outdated](https://github.com/psampaz/go-mod-outdated)|Find outdated dependencies of your Go projects. go-mod-outdated provides a table view of the go list -u -m -json all command which lists all dependencies of a Go project and their available minor and patch updates. It also provides a way to filter indirect dependencies and dependencies without updates.|551|21|4|2019-04-19T07:12:13Z|2021-10-11T17:54:10Z|
[go-cleanarch](https://github.com/roblaszczak/go-cleanarch)|Clean architecture validator for go, like a The Dependency Rule and interaction between packages in your Go projects.|538|39|3|2017-04-12T21:59:16Z|2021-11-08T16:18:42Z|
[goreturns](https://github.com/sqs/goreturns)|A gofmt/goimports-like tool for Go programmers that fills in Go return statements with zero values to match the func return types|505|56|28|2014-10-07T15:48:08Z|2020-10-17T19:35:15Z|
[todocheck](https://github.com/preslavmihaylov/todocheck)|A static code analyser for annotated TODO comments|363|29|11|2020-07-18T16:19:00Z|2021-10-11T07:47:06Z|
[unconvert](https://github.com/mdempsky/unconvert)|Remove unnecessary type conversions from Go source|313|19|6|2016-02-19T21:59:07Z|2020-05-18T20:43:04Z|
[dupl](https://github.com/mibk/dupl)|a tool for code clone detection|266|17|2|2015-05-20T15:45:15Z|2020-12-19T20:18:10Z|
[tickgit](https://github.com/augmentable-dev/tickgit)|Manage your repository&#39;s TODOs, tickets and checklists as config in your codebase.|266|15|9|2019-10-12T00:49:10Z|2020-06-30T17:31:38Z|
[golines](https://github.com/segmentio/golines)|A golang formatter that fixes long lines|264|16|8|2019-10-01T00:34:25Z|2021-11-18T19:03:07Z|
[gostatus](https://github.com/shurcooL/gostatus)|A command line tool that shows the status of Go repositories.|245|11|1|2013-11-27T04:06:35Z|2019-02-03T17:04:19Z|
[apicompat](https://github.com/bradleyfalzon/apicompat)|apicompat checks recent changes to a Go project for backwards incompatible changes|175|5|7|2016-07-10T13:39:02Z|2017-02-05T09:57:05Z|
[checkstyle](https://github.com/qiniu/checkstyle)|checkstyle for go|118|17|5|2014-01-01T01:09:27Z|2021-03-10T02:55:53Z|
[lint](https://github.com/surullabs/lint)|Run linters from Go code - |65|10|1|2016-07-09T09:52:39Z|2018-10-28T00:00:40Z|
[validate](https://github.com/mccoyst/validate)|A Go package to automatically validate fields with tags|60|14|1|2013-11-22T01:28:40Z|2016-03-28T22:03:18Z|
**[ARCHIVED]**  [go-outdated](https://github.com/firstrow/go-outdated)|Find outdated golang packages|44|2|0|2015-06-29T06:10:39Z|2019-01-15T09:49:38Z|
**[ARCHIVED]**  [blanket](https://github.com/verygoodsoftwarenotvirus/blanket)|MOVED TO GITLAB|14|0|1|2017-09-04T13:09:28Z|2018-07-22T18:28:33Z|
[golang-ifood-sdk](https://github.com/arxdsilva/golang-ifood-sdk)|Golang Ifood API SDK |7|2|0|2021-03-13T15:15:45Z|2021-11-03T05:50:34Z|
[GoCover.io](https://gocover.io/)|GoCover.io offers the code coverage of any golang package as a service.|-|-|-|-|-|
[Golint online](http://go-lint.appspot.com/)|Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.|-|-|-|-|-|


## Editor Plugins

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[vim-go](https://github.com/fatih/vim-go)|Go development plugin for Vim|14052|1398|29|2014-03-24T13:12:26Z|2021-11-25T19:23:04Z|
[gocode](https://github.com/nsf/gocode)|An autocompletion daemon for the Go programming language|4943|681|67|2010-07-05T00:13:16Z|2021-10-27T23:29:18Z|
[GoSublime](https://github.com/DisposaBoy/GoSublime)|A Golang plugin collection for SublimeText 3, providing code completion and other IDE-like features.|3420|318|86|2011-08-27T22:24:39Z|2020-07-21T18:51:34Z|
[vscode-go](https://github.com/golang/vscode-go)|Go extension for VS Code|2243|400|214|2020-03-06T17:52:04Z|2021-12-09T17:38:52Z|
**[ARCHIVED]**  [go-plus](https://github.com/joefitzgerald/go-plus)|An Enhanced Go Experience For The Atom Editor|1519|140|92|2014-03-13T19:19:18Z|2021-05-04T12:16:23Z|
[go-mode.el](https://github.com/dominikh/go-mode.el)|Emacs mode for the Go programming language|1210|195|42|2013-01-30T23:47:03Z|2021-11-13T07:05:46Z|
[coc-go](https://github.com/josa42/coc-go)|Go language server extension using gopls for coc.nvim.|400|20|6|2019-04-25T09:08:04Z|2021-12-03T19:06:26Z|
[Watch](https://github.com/eaburns/Watch)|Watches for changes in a directory tree and reruns a command in an acme win or just on the terminal.|188|45|7|2013-08-08T17:10:22Z|2018-03-25T14:15:49Z|
[goimports-reviser](https://github.com/incu6us/goimports-reviser)|Right imports sorting &amp; code formatting tool (goimports alternative)|92|21|12|2020-04-08T14:49:07Z|2021-12-02T11:59:13Z|
[vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go)|Vim compiler plugin for Go (golang)|88|17|0|2012-11-25T18:15:52Z|2016-06-28T22:00:12Z|
**[ARCHIVED]**  [go-language-server](https://github.com/theia-ide/go-language-server)|A Go language server.|31|10|3|2017-11-21T13:10:33Z|2019-03-25T14:30:07Z|
[gounit-vim](https://github.com/hexdigest/gounit-vim)|Vim plugin for https://github.com/hexdigest/gounit|22|1|0|2018-02-21T18:27:17Z|2018-10-29T11:14:49Z|
**[ARCHIVED]**  [theia-go-extension](https://github.com/theia-ide/theia-go-extension)|Theia Go Extension|16|6|4|2017-11-30T15:15:39Z|2019-03-14T08:06:45Z|
[vscode-go-prof](https://github.com/MaxM65dia/vscode-go-prof)|Go language profiling|5|0|3|2019-04-18T06:40:25Z|2019-06-04T07:46:34Z|
[vscode-go-doc](https://github.com/msyrus/vscode-go-doc)|An Microsoft Visual Code extension for Golang to print symbol definition to output|4|0|3|2018-03-15T08:53:19Z|2021-06-08T03:55:27Z|
[Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go)|Go plugin for JetBrains IDEs.|-|-|-|-|-|


## Go Generate Tools

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[gotests](https://github.com/cweill/gotests)|Automatically generate Go test boilerplate from your source code.|3589|265|44|2016-01-19T05:06:02Z|2021-10-09T05:15:21Z|
[genny](https://github.com/cheekybits/genny)|Elegant generics for Go|1610|127|29|2014-10-27T22:03:45Z|2021-08-24T18:48:42Z|
**[ARCHIVED]**  [re2dfa](https://github.com/opennota/re2dfa)|Transform regular expressions into finite state machines and output Go source code. This repository has migrated to https://gitlab.com/opennota/re2dfa|189|16|4|2015-06-20T10:56:24Z|2018-09-11T05:52:06Z|
[toml-to-go](https://github.com/xuri/toml-to-go)|Translates TOML into a Go type in your browser instantly|138|33|0|2016-08-03T06:26:02Z|2020-01-09T16:08:01Z|
[xgen](https://github.com/xuri/xgen)|XSD (XML Schema Definition) parser and Go/C/Java/Rust/TypeScript code generator|123|25|11|2019-06-22T13:56:05Z|2021-11-30T16:15:55Z|
[gonerics](https://github.com/bouk/gonerics)|Generics for go|113|9|0|2014-09-29T00:47:23Z|2014-09-29T15:04:55Z|
[hasgo](https://github.com/DylanMeeus/hasgo)|Haskell-flavoured functions for Go :smiley:|108|6|16|2019-05-16T22:14:08Z|2021-04-29T20:23:38Z|
[gocontracts](https://github.com/Parquery/gocontracts)|A tool for design-by-contract in Go|78|5|1|2018-08-13T17:33:48Z|2019-01-26T07:32:40Z|
[gounit](https://github.com/hexdigest/gounit)|Unit tests generator for Go programming language|60|9|0|2018-02-05T00:08:30Z|2018-08-17T09:38:42Z|
[generic](https://github.com/usk81/generic)|flexible data type for Go|41|7|2|2016-06-15T14:00:36Z|2021-01-13T20:33:15Z|
[godal](https://github.com/mafulong/godal)|godal provides the ability to generate specific golang code. The godal is to enable developers to write fast code in an expressive way.|11|0|0|2021-03-16T03:09:34Z|2021-10-23T04:38:11Z|


## Go Tools

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[go-swagger](https://github.com/go-swagger/go-swagger)|Swagger 2.0 implementation for go|7042|1052|503|2014-11-16T20:13:15Z|2021-12-09T23:37:47Z|
[OctoLinker](https://github.com/OctoLinker/OctoLinker)|OctoLinker — Links together, what belongs together|4851|307|57|2013-12-27T18:01:52Z|2021-12-10T22:43:50Z|
[go-callvis](https://github.com/ofabry/go-callvis)|Visualize call graph of a Go program using Graphviz|3677|277|36|2016-09-03T11:31:46Z|2021-11-25T07:59:41Z|
[depth](https://github.com/KyleBanks/depth)|Visualize Go Dependency Trees|656|47|8|2017-03-04T15:42:23Z|2021-12-10T15:51:45Z|
[richgo](https://github.com/kyoh86/richgo)|Enrich `go test` outputs with text decorations.|632|14|3|2017-01-04T17:05:57Z|2021-12-07T17:38:58Z|
[rts](https://github.com/galeone/rts)|RTS: request to struct. Generates Go structs from JSON server responses.|226|11|0|2016-04-04T13:17:19Z|2021-09-26T08:39:38Z|
[godbg](https://github.com/tylerwince/godbg)|Go implementation of the Rust `dbg` macro|180|10|2|2019-01-23T23:51:43Z|2019-04-20T00:52:22Z|
[typex](https://github.com/dtgorski/typex)|[TOOL, CLI] - Filter and examine Go type structures, interfaces and their transitive dependencies and relationships. Export structural types as TypeScript value object or bare type representations.|136|10|1|2020-03-24T21:02:44Z|2021-01-18T00:58:25Z|
[colorgo](https://github.com/songgao/colorgo)|Colorize (highlight) `go build` command output|108|15|3|2013-02-14T18:06:10Z|2020-07-18T23:02:45Z|
[gothanks](https://github.com/psampaz/gothanks)|GoThanks automatically stars Go&#39;s official repository and your go.mod github dependencies, providing a simple way  to say thanks to the maintainers of the modules you use and the contributors of Go itself.|107|7|1|2019-11-10T07:48:02Z|2021-03-01T23:15:34Z|
[roumon](https://github.com/becheran/roumon)|Universal goroutine monitor using pprof and termui |66|3|0|2021-03-02T18:02:41Z|2021-03-29T21:27:13Z|
[igo](https://github.com/rocketlaunchr/igo)|Improved Go Syntax (transpiler)|51|3|0|2018-11-17T05:34:03Z|2020-04-06T07:25:36Z|
[go-james](https://github.com/pieterclaerhout/go-james)|James is your butler and helps you to create, build, debug, test and run your Go projects|47|3|10|2019-10-14T16:00:14Z|2021-10-26T11:45:43Z|
[go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete)|bash completion for go and wgo|38|7|0|2015-05-22T03:07:05Z|2017-11-17T14:00:35Z|
**[ARCHIVED]**  [generator-go-lang](https://github.com/axelspringer/generator-go-lang)|:guardsman: A teeny tiny and somewhat opinionated generator for your next golang project|23|4|0|2017-09-13T11:33:06Z|2020-04-06T07:02:29Z|
[gomodrun](https://github.com/dustinblackman/gomodrun)|The forgotten go tool that executes and caches binaries included in go.mod files.|18|3|0|2020-01-26T15:33:18Z|2021-09-18T18:40:24Z|
[docs](https://github.com/go-oas/docs)|Automatically generate RESTful API documentation for GO projects - aligned with Open API Specification standard|9|1|8|2021-01-28T18:51:47Z|2021-03-06T11:31:16Z|
[modver](https://github.com/bobg/modver)||0|0|0|2021-07-17T15:05:52Z|2021-08-19T14:28:20Z|


### DevOps Tools

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[kubernetes](https://github.com/kubernetes/kubernetes)|Production-Grade Container Scheduling and Management|83531|30552|2459|2014-06-06T22:56:04Z|2021-12-11T15:48:03Z|
[moby](https://github.com/moby/moby)|Moby Project - a collaborative project for the container ecosystem to assemble container-based systems|61752|17849|4187|2013-01-18T18:10:57Z|2021-12-11T02:05:14Z|
[gogs](https://github.com/gogs/gogs)|Gogs is a painless self-hosted Git service|38119|4446|805|2014-02-12T01:57:08Z|2021-11-28T09:03:18Z|
[traefik](https://github.com/traefik/traefik)|The Cloud Native Application Proxy|35958|3973|608|2015-09-13T19:04:02Z|2021-12-10T16:56:26Z|
[gitea](https://github.com/go-gitea/gitea)|Git with a cup of tea, painless self-hosted git service|27251|3421|1865|2016-11-01T02:13:26Z|2021-12-11T17:26:15Z|
[vegeta](https://github.com/tsenart/vegeta)|HTTP load testing tool and library. It&#39;s over 9000!|18750|1167|86|2013-08-13T11:45:21Z|2021-09-23T13:14:57Z|
[packer](https://github.com/hashicorp/packer)|Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.|13341|3177|290|2013-03-23T05:43:03Z|2021-12-11T20:39:20Z|
[hey](https://github.com/rakyll/hey)|HTTP load generator, ApacheBench (ab) replacement, formerly known as rakyll/boom|12484|877|147|2016-09-02T10:24:09Z|2021-11-29T11:43:40Z|
[webhook](https://github.com/adnanh/webhook)|webhook is a lightweight incoming webhook server to run shell commands|7178|603|59|2015-01-12T20:59:11Z|2021-10-11T10:47:23Z|
[gvm](https://github.com/moovweb/gvm)|Go Version Manager|6828|382|142|2011-12-03T02:34:04Z|2021-11-22T22:04:12Z|
[gaia](https://github.com/gaia-pipeline/gaia)|Build powerful pipelines in any programming language.|4535|220|29|2017-12-28T11:01:31Z|2021-09-17T06:48:38Z|
[gox](https://github.com/mitchellh/gox)|A dead simple, no frills Go cross compile tool|4152|323|67|2013-11-17T03:11:35Z|2021-03-11T18:25:16Z|
[bosun](https://github.com/bosun-monitor/bosun)|Time Series Alerting Framework|3220|505|10|2013-11-15T00:12:27Z|2021-10-18T07:44:06Z|
[ddosify](https://github.com/ddosify/ddosify)|High-performance load testing tool, written in Golang.|3121|103|4|2021-08-04T07:43:53Z|2021-11-22T07:40:14Z|
[bombardier](https://github.com/codesenberg/bombardier)|Fast cross-platform HTTP benchmarking tool written in Go|2930|185|11|2016-05-29T15:16:30Z|2021-07-11T21:55:07Z|
[pomerium](https://github.com/pomerium/pomerium)|Pomerium is an identity-aware access proxy.|2844|220|74|2019-01-01T08:04:37Z|2021-12-10T22:27:15Z|
[script](https://github.com/bitfield/script)|Making it easy to write shell-like scripts in Go|1942|163|28|2019-04-20T14:37:03Z|2021-12-09T20:57:01Z|
[fac](https://github.com/mkchoi212/fac)|Easy-to-use CUI for fixing git conflicts|1738|47|9|2017-12-29T19:11:45Z|2019-10-09T10:24:03Z|
[kala](https://github.com/ajvb/kala)|Modern Job Scheduler|1731|167|20|2015-03-19T04:24:19Z|2021-11-19T20:48:24Z|
[goxc](https://github.com/laher/goxc)|a build tool for Go, with a focus on cross-compiling, packaging and deployment|1672|78|12|2013-02-11T08:49:53Z|2019-09-30T08:22:07Z|
[statusok](https://github.com/sanathp/statusok)|Monitor your Website and APIs from your Computer. Get Notified through Slack, E-mail when your server is down or response time is more than expected. |1520|196|41|2015-08-26T17:39:48Z|2021-08-11T16:30:28Z|
[s3gof3r](https://github.com/rlmcpherson/s3gof3r)|Fast, concurrent, streaming access to Amazon S3, including gof3r, a CLI. http://godoc.org/github.com/rlmcpherson/s3gof3r|1109|188|55|2013-08-02T13:11:39Z|2021-08-28T17:43:13Z|
[go-selfupdate](https://github.com/sanbornm/go-selfupdate)|Enable your Go applications to self update|890|125|13|2013-11-13T06:17:43Z|2021-09-23T14:42:59Z|
[s5cmd](https://github.com/peak/s5cmd)|Parallel S3 and local filesystem execution tool.|864|85|47|2016-11-16T10:31:15Z|2021-12-09T15:04:44Z|
[skm](https://github.com/TimothyYe/skm)|A simple and powerful SSH keys manager|716|43|0|2017-10-11T06:52:55Z|2021-09-26T04:09:24Z|
[scaleway-cli](https://github.com/scaleway/scaleway-cli)|Command Line Interface for Scaleway|709|118|100|2015-03-20T09:45:50Z|2021-12-10T20:43:11Z|
[ghorg](https://github.com/gabrie30/ghorg)|Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, Bitbucket, and more 🥚|632|77|7|2018-03-29T02:53:05Z|2021-12-11T19:36:13Z|
[utask](https://github.com/ovh/utask)|µTask is an automation engine that models and executes business processes declared in yaml. ✏️📋|567|44|41|2019-11-05T12:59:55Z|2021-11-30T12:06:09Z|
[kool](https://github.com/kool-dev/kool)|From local development to the cloud: development workflow made easy.|566|42|9|2020-07-06T22:25:04Z|2021-12-11T11:24:06Z|
[cassowary](https://github.com/rogerwelin/cassowary)|:rocket: Modern cross-platform HTTP load-testing tool written in Go|545|22|9|2019-08-25T21:28:34Z|2021-11-25T06:18:26Z|
[aurora](https://github.com/xuri/aurora)|Cross-platform beanstalkd queue server admin console.|541|77|7|2016-10-09T03:17:51Z|2021-08-19T16:05:21Z|
[govvv](https://github.com/ahmetb/govvv)|&#34;go build&#34; wrapper to add version info to Golang applications|508|38|1|2016-08-02T22:30:23Z|2020-02-03T18:05:00Z|
[gonative](https://github.com/inconshreveable/gonative)|Build Go Toolchains /w native libs for cross-compilation|325|33|7|2014-05-01T01:43:15Z|2016-07-21T19:34:23Z|
[trubka](https://github.com/xitonix/trubka)|A CLI tool for Kafka|307|17|3|2019-07-05T02:02:25Z|2020-12-09T22:31:40Z|
[pewpew](https://github.com/bengadbois/pewpew)|Flexible HTTP command line stress tester for websites and web services|304|27|1|2016-10-12T22:59:40Z|2021-07-08T21:43:11Z|
[mora](https://github.com/emicklei/mora)|MongoDB generic REST server in Go|298|54|9|2013-07-12T09:07:01Z|2021-04-11T12:45:54Z|
[jenkins-cli](https://github.com/jenkins-zh/jenkins-cli)|Jenkins CLI allows you to manage your Jenkins in an easy way. Jenkins 命令行客户端|295|72|70|2019-06-21T10:19:34Z|2021-12-08T14:14:16Z|
[lstags](https://github.com/ivanilves/lstags)|Explore Docker registries and manipulate Docker images!|282|25|6|2017-08-15T05:25:17Z|2021-07-21T05:39:05Z|
[dogo](https://github.com/liudng/dogo)|Monitoring changes in the source file and automatically compile and run (restart).|238|42|5|2014-11-19T10:16:35Z|2019-03-15T05:14:19Z|
[balerter](https://github.com/balerter/balerter)|Script Based Alerting Manager|238|14|2|2019-12-30T09:25:01Z|2021-12-08T08:19:21Z|
[manssh](https://github.com/xwjdsh/manssh)|Manage your ssh alias configs easily.|233|24|3|2017-10-08T06:52:42Z|2018-06-21T18:02:43Z|
[godbg](https://github.com/sirnewton01/godbg)|Web-based gdb front-end application|224|25|5|2013-08-09T01:02:00Z|2018-07-09T13:50:41Z|
[easyssh-proxy](https://github.com/appleboy/easyssh-proxy)|easyssh-proxy provides a simple implementation of some SSH protocol features in Go|211|47|12|2017-03-03T02:58:14Z|2021-12-09T13:47:47Z|
[blast](https://github.com/dave/blast)|Blast is a simple tool for API load testing and batch jobs|202|9|1|2017-10-21T17:13:09Z|2018-03-01T09:57:41Z|
[gobrew](https://github.com/cryptojuice/gobrew)|Shell script to download and set GO environmental paths to allow multiple versions.|185|16|4|2013-11-13T00:32:18Z|2020-05-21T03:38:51Z|
[abbreviate](https://github.com/dnnrly/abbreviate)|Supporting your devops by shortening your strings using common abbreviations and clever guesswork|174|13|4|2018-11-23T23:05:15Z|2021-09-29T22:07:49Z|
[terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi)|OpenAPI Terraform Provider that configures itself at runtime with the resources exposed by the service provider (defined in a swagger file)|173|32|15|2017-10-17T03:47:09Z|2021-11-13T20:18:06Z|
[ostent](https://github.com/ostrost/ostent)|Ostent is a server tool to collect, display and report system metrics.|171|12|0|2014-03-31T04:52:10Z|2018-04-03T20:54:21Z|
[kcli](https://github.com/cswank/kcli)|A kafka command line browser|162|13|1|2017-03-25T20:41:22Z|2020-01-04T00:26:19Z|
[grapes](https://github.com/yaronsumel/grapes)|easy way to distribute commands over ssh.|152|8|1|2016-09-01T11:28:47Z|2020-12-21T15:58:45Z|
[winrm-cli](https://github.com/masterzen/winrm-cli)|Command-line tool to remotely execute commands on Windows machines through WinRM|130|17|1|2016-05-23T09:03:15Z|2021-11-06T15:33:17Z|
[dockerfile-generator](https://github.com/ozankasikci/dockerfile-generator)|dfg - Generates dockerfiles based on various input channels. |119|11|0|2019-08-14T20:03:37Z|2020-01-14T02:56:23Z|
[drone-scp](https://github.com/appleboy/drone-scp)|Copy files and artifacts via SSH using a binary, docker or Drone CI.|92|22|20|2016-10-16T13:35:56Z|2021-10-23T10:43:33Z|
[go-furnace](https://github.com/go-furnace/go-furnace)|Go Hosting Solution for AWS, Google Could and Digital Ocean|85|25|12|2016-10-09T11:17:20Z|2021-10-28T07:50:11Z|
[wide](https://github.com/88250/wide)|🌈 一款基于 Web 的 Go 语言 IDE，随时随地玩 golang。|72|29|5|2019-12-01T11:30:46Z|2021-08-11T18:17:43Z|
[s3-proxy](https://github.com/oxyno-zeta/s3-proxy)|S3 Reverse Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth)|69|13|6|2019-09-22T14:17:39Z|2021-12-10T09:31:06Z|
[dropship](https://github.com/ChrisMcKenzie/dropship)|Super simple deployment tool|56|11|10|2015-09-03T23:09:19Z|2018-07-25T21:03:58Z|
[drone-jenkins](https://github.com/appleboy/drone-jenkins)|Drone plugin for trigger Jenkins jobs.|31|13|4|2016-10-15T00:53:03Z|2021-10-05T06:51:03Z|
[rodent](https://github.com/alouche/rodent)|Manage Go Versions/Projects/Dependencies|31|2|6|2014-06-01T21:08:42Z|2017-04-22T07:47:52Z|
[awsenv](https://github.com/soniah/awsenv)|AWS environment config loader|27|6|0|2015-08-05T07:21:24Z|2018-07-17T14:05:46Z|
[lwc](https://github.com/timdp/lwc)|A live-updating version of the UNIX wc command.|25|3|0|2018-04-22T09:23:44Z|2020-05-03T16:25:01Z|
[docker-go-mingw](https://github.com/x1unix/docker-go-mingw)|Docker image for building Go binaries with MinGW toolchain|21|3|0|2020-09-16T14:02:35Z|2021-12-10T08:22:32Z|
[httpref](https://github.com/dnnrly/httpref)|Command line, offline, access to HTTP status code, common header, and port references|17|8|3|2020-01-10T22:00:47Z|2021-10-25T22:34:05Z|
[depcharge](https://github.com/centerorbit/depcharge)|DepCharge is a tool designed to help orchestrate the execution of commands across many directories at once.|17|4|2|2018-07-25T04:02:09Z|2020-02-05T20:47:07Z|
[sg](https://github.com/ChristopherRabotin/sg)|Stress gauge allows one to gauge response times of an HTTP service under stress.|5|0|2|2015-08-19T15:06:32Z|2016-10-28T23:18:00Z|
[aptly-fork](https://github.com/smira/aptly-fork)|aptly - Debian repository management tool (fork of aptly-dev/aptly)|3|2|0|2019-07-04T16:45:46Z|2019-09-27T12:21:26Z|
[gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator)|Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.|-|-|-|-|-|


### Other Software

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[syncthing](https://github.com/syncthing/syncthing)|Open Source Continuous File Synchronization|41811|3272|304|2013-11-26T09:48:21Z|2021-12-10T13:20:40Z|
[croc](https://github.com/schollz/croc)|Easily and securely send things from one computer to another :crocodile: :package:|17681|743|74|2017-10-17T15:20:18Z|2021-11-18T15:04:05Z|
[lime](https://github.com/limetext/lime)|Open source API-compatible alternative to the text editor Sublime Text|15283|1125|22|2012-10-03T18:10:02Z|2021-01-02T13:10:47Z|
[restic](https://github.com/restic/restic)|Fast, secure, efficient backup program|15045|1027|457|2014-04-27T14:07:58Z|2021-12-08T09:17:11Z|
[goreplay](https://github.com/buger/goreplay)|GoReplay is an open-source tool for capturing and replaying live HTTP traffic into a test environment in order to continuously test your system with real data. It can be used to increase confidence in code deployments, configuration changes and infrastructure changes.|14964|1510|229|2013-05-30T09:29:07Z|2021-12-01T23:43:41Z|
[seaweedfs](https://github.com/chrislusf/seaweedfs)|SeaweedFS is a fast distributed storage system for blobs, objects, files, and data lake, for billions of files! Blob store has O(1) disk seek, cloud tiering. Filer supports Cloud Drive, cross-DC active-active replication, Kubernetes, POSIX FUSE mount, S3 API, S3 Gateway, Hadoop, WebDAV, encryption, Erasure Coding.|13328|1644|55|2014-07-14T16:41:37Z|2021-12-10T21:24:42Z|
[comcast](https://github.com/tylertreat/comcast)|Simulating shitty network connections so you can build better systems.|7701|331|23|2014-11-12T03:15:58Z|2021-06-02T17:45:03Z|
[confd](https://github.com/kelseyhightower/confd)|Manage local application configuration files using templates and data from etcd or consul|7670|1327|156|2013-10-01T04:06:09Z|2021-10-05T14:12:59Z|
[toxiproxy](https://github.com/Shopify/toxiproxy)|:alarm_clock: :fire: A TCP proxy to simulate network and system conditions for chaos and resiliency testing|7503|347|60|2014-09-04T13:56:38Z|2021-11-07T15:14:06Z|
[liteide](https://github.com/visualfc/liteide)|LiteIDE is a simple, open source, cross-platform Go IDE. |6649|912|359|2012-11-19T01:54:25Z|2021-12-05T02:54:48Z|
[drive](https://github.com/odeke-em/drive)|Google Drive client for the commandline|6242|423|270|2014-11-03T08:18:11Z|2021-02-08T10:45:18Z|
[nes](https://github.com/fogleman/nes)|NES emulator written in Go.|4955|461|7|2015-03-02T22:16:13Z|2021-06-05T21:50:16Z|
[duplicacy](https://github.com/gilbertchen/duplicacy)|A new generation cloud backup tool |3981|288|295|2016-02-23T01:28:10Z|2021-10-25T03:55:26Z|
[tsuru](https://github.com/tsuru/tsuru)|Open source and extensible Platform as a Service (PaaS).|3752|496|169|2012-03-05T21:41:08Z|2021-12-10T20:42:17Z|
[scc](https://github.com/boyter/scc)|Sloc, Cloc and Code: scc is a very fast accurate code counter with complexity calculations and COCOMO estimates written in pure Go|2918|133|40|2018-03-01T06:44:25Z|2021-11-11T05:25:51Z|
[mylg](https://github.com/mehrdadrad/mylg)|Network Diagnostic Tool|2536|221|13|2016-06-21T19:39:58Z|2020-02-26T22:39:02Z|
[goboy](https://github.com/Humpheh/goboy)|Multi-platform Nintendo Game Boy Color emulator written in Go|2401|103|7|2017-08-20T14:59:05Z|2020-08-09T11:00:27Z|
[sup](https://github.com/pressly/sup)|Super simple deployment tool - think of it like &#39;make&#39; for a network of servers|2303|168|54|2015-02-23T23:04:21Z|2021-06-07T09:42:58Z|
[lgo](https://github.com/yunabe/lgo)|Interactive Go programming with Jupyter|2222|108|26|2017-10-05T15:29:10Z|2020-11-20T07:01:33Z|
[circuit](https://github.com/gocircuit/circuit)|Circuit: Dynamic cloud orchestration http://gocircuit.org|1921|159|12|2014-04-10T20:46:06Z|2020-05-03T14:20:23Z|
[snap](https://github.com/intelsdi-x/snap)|The open telemetry framework|1796|302|150|2014-08-13T21:04:51Z|2018-12-20T01:29:47Z|
[borg](https://github.com/ok-borg/borg)|Search and save shell snippets without leaving your terminal|1528|60|13|2016-09-10T20:20:42Z|2018-02-07T19:40:06Z|
[community](https://github.com/documize/community)|Modern Confluence alternative designed for internal &amp; external docs, built with Golang &#43; EmberJS|1443|146|44|2016-04-29T23:35:07Z|2021-11-19T13:34:21Z|
[blocky](https://github.com/0xERR0R/blocky)|Fast and lightweight DNS proxy as ad-blocker for local network with many features|1147|78|23|2019-11-06T09:03:31Z|2021-12-11T03:13:20Z|
[Go-Package-Store](https://github.com/shurcooL/Go-Package-Store)|An app that displays updates for the Go packages in your GOPATH.|886|29|10|2014-01-24T06:02:09Z|2020-03-07T22:35:33Z|
[shell2http](https://github.com/msoap/shell2http)|Executing shell commands via HTTP server|868|92|3|2015-03-11T19:39:09Z|2021-10-30T17:52:00Z|
[vflow](https://github.com/EdgeCast/vflow)| Enterprise Network Flow Collector (IPFIX, sFlow, Netflow) |855|185|39|2017-02-24T21:28:21Z|2021-12-09T05:04:07Z|
[peg](https://github.com/pointlander/peg)|Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.|825|104|31|2010-04-25T21:20:46Z|2021-08-22T22:12:48Z|
[leaps](https://github.com/Jeffail/leaps)|A pair programming service using operational transforms|709|53|13|2014-06-19T20:33:05Z|2021-02-22T08:51:54Z|
[gfile](https://github.com/Antonito/gfile)|Direct file transfer over WebRTC|637|39|5|2019-03-08T06:02:16Z|2021-02-23T09:43:17Z|
[guora](https://github.com/meloalright/guora)|🖖🏻 A self-hosted Quora like web application written in Go  基于 Golang 类似知乎的私有部署问答应用 包含问答、评论、点赞、管理后台等功能|567|84|7|2020-08-13T16:05:08Z|2020-11-28T03:25:36Z|
[gebug](https://github.com/moshebe/gebug)|Debug Dockerized Go applications better|557|19|3|2020-07-20T13:43:42Z|2021-12-11T00:22:05Z|
[gocc](https://github.com/goccmack/gocc)|Parser / Scanner Generator|505|43|33|2015-06-05T13:08:21Z|2021-12-04T15:02:25Z|
[mockingjay-server](https://github.com/quii/mockingjay-server)|Fake server, Consumer Driven Contracts and help with testing performance from one configuration file with zero system dependencies and no coding whatsoever|500|60|9|2015-04-04T19:18:02Z|2021-01-15T09:44:20Z|
[go-playground](https://github.com/x1unix/go-playground)|Better Go Playground powered by React and Monaco editor|485|26|6|2020-01-16T19:03:35Z|2021-09-21T23:01:17Z|
[go-peerflix](https://github.com/Sioro-Neoku/go-peerflix)|Go Peerflix|440|74|11|2015-10-08T19:44:47Z|2021-08-04T03:42:32Z|
[ipe](https://github.com/dimiro1/ipe)|An open source Pusher server implementation compatible with Pusher client libraries written in GO|341|66|1|2015-01-13T11:49:19Z|2021-03-28T13:07:21Z|
[IDE](https://github.com/thestrukture/IDE)|Web based, Go IDE. |322|20|0|2017-09-09T19:49:57Z|2021-07-30T08:21:00Z|
[tcpprobe](https://github.com/mehrdadrad/tcpprobe)|Modern TCP tool and service for network performance observability.|315|17|0|2020-10-26T00:27:20Z|2021-02-21T22:15:21Z|
[wellington](https://github.com/wellington/wellington)|Spriting that sass has been missing|301|16|26|2014-12-08T18:08:59Z|2020-10-30T00:02:54Z|
[woke](https://github.com/get-woke/woke)|Detect non-inclusive language in your source code.|293|50|13|2020-08-31T17:21:07Z|2021-12-07T22:12:02Z|
[cherry](https://github.com/rafael-santiago/cherry)|A tiny webchat server in Go.|265|41|0|2015-10-24T20:56:23Z|2017-06-24T10:34:24Z|
[joincap](https://github.com/assafmo/joincap)|Merge multiple pcap files together, gracefully.|171|14|2|2018-05-31T16:57:22Z|2021-03-15T16:44:16Z|
[tcpdog](https://github.com/mehrdadrad/tcpdog)|eBPF based TCP observability.|168|16|0|2020-12-30T00:10:39Z|2021-07-21T14:36:31Z|
[orbit](https://github.com/gulien/orbit)|:satellite: A cross-platform task runner for executing commands and generating files from templates|160|9|2|2017-05-13T11:25:00Z|2021-01-18T11:35:49Z|
[vaku](https://github.com/lingrino/vaku)|Vaku extends the Vault API &amp; CLI|126|15|0|2018-04-24T04:52:10Z|2021-11-23T16:56:04Z|
[dp](https://github.com/scryinfo/dp)|Scry Data Protocol|83|38|44|2018-12-12T03:14:22Z|2021-11-23T18:15:30Z|
[boxed](https://github.com/tejo/boxed)|dropbox based blog engine, written in go.|75|9|0|2015-04-18T20:48:46Z|2018-08-09T20:27:08Z|
[crawley](https://github.com/s0rg/crawley)|The unix-way web crawler|44|1|0|2021-10-27T18:48:51Z|2021-12-08T18:07:10Z|
[naclpipe](https://github.com/unix4fun/naclpipe)|NaCL pipe|22|2|0|2015-05-05T23:16:39Z|2018-11-18T14:43:21Z|
[term-quiz](https://github.com/crazcalm/term-quiz)|Terminal Quiz Application Written in Go|19|5|0|2017-12-26T07:36:40Z|2018-10-24T22:46:25Z|
[snitch](https://github.com/lucasgomide/snitch)|Keep updated about all deploys on Tsuru|15|1|5|2017-04-06T21:02:05Z|2018-07-23T18:16:30Z|
[GoDocTooltip](https://github.com/diankong/GoDocTooltip)|A Chrome extension for golang users.When you&#39;re at golang&#39;s official doc site, it will show function&#39;s description as tooltip on function list|12|2|0|2016-01-21T12:06:55Z|2021-01-30T02:10:33Z|
[hoofli](https://github.com/dnnrly/hoofli)|Generate PlantUML diagrams from Chrome or Firefox network inspections|1|0|1|2021-04-23T20:36:56Z|2021-09-29T22:23:16Z|
[Juju](https://jujucharms.com/)|Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.|-|-|-|-|-|
[GoLand](https://jetbrains.com/go)|Full featured cross-platform Go IDE.|-|-|-|-|-|
[hugo](https://gohugo.io/)|Fast and Modern Static Website Engine.|-|-|-|-|-|
[Docker](https://www.docker.com/)|Open platform for distributed applications for developers and sysadmins.|-|-|-|-|-|


## Benchmarks

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark)|:zap: Go web framework benchmark|1596|179|6|2016-04-06T06:27:33Z|2021-08-07T23:51:55Z|
[go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark)|Go HTTP request router and web framework benchmark|1536|216|22|2013-12-16T21:28:47Z|2021-07-26T23:43:58Z|
[go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)|Benchmarks of Go serialization methods|1231|121|8|2013-01-18T16:03:58Z|2021-09-30T21:27:42Z|
[skynet](https://github.com/atemerev/skynet)|Skynet 1M threads microbenchmark|1000|133|31|2016-02-14T13:59:19Z|2021-05-23T18:11:12Z|
[speedtest-resize](https://github.com/fawick/speedtest-resize)|Compare various Image resize algorithms for the Go language|214|17|2|2013-09-16T12:40:05Z|2020-10-28T16:26:39Z|
[go-benchmarks](https://github.com/tylertreat/go-benchmarks)|A few miscellaneous Go microbenchmarks.|141|26|2|2016-02-25T01:00:38Z|2016-02-25T05:42:36Z|
[gospeed](https://github.com/feyeleanor/gospeed)|Go micro-benchmarks for calculating the speed of language constructs|108|6|0|2011-05-23T21:16:11Z|2021-03-02T15:18:25Z|
[autobench](https://github.com/davecheney/autobench)|Go benchmark harness. |92|28|2|2013-03-28T11:17:01Z|2014-07-28T04:52:21Z|
[golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark)|A benchmarking shootout of various db/SQL utilities for Go|60|14|2|2014-09-24T20:47:26Z|2018-03-22T01:42:17Z|
[gocostmodel](https://github.com/mna/gocostmodel)|Benchmarks of common basic operations for the Go language.|57|5|0|2014-12-19T02:54:45Z|2021-05-19T15:19:44Z|
[kvbench](https://github.com/jimrobinson/kvbench)|Key/Value database benchmark|24|2|0|2014-04-15T09:59:27Z|2019-09-28T10:24:57Z|
[go-benchmark-app](https://github.com/mrLSD/go-benchmark-app)|Application for HTTP benchmarking via different rules and configs|22|4|0|2017-01-24T12:24:08Z|2017-03-17T11:40:10Z|
[go-ml-benchmarks](https://github.com/nikolaydubina/go-ml-benchmarks)|⏱ Benchmarks of machine learning inference for Go|19|1|1|2021-02-09T10:20:46Z|2021-02-19T02:59:16Z|
[go-json-benchmark](https://github.com/zerosnake0/go-json-benchmark)|Benchmark of Golang JSON Libraries|6|1|0|2019-11-10T08:00:15Z|2020-10-08T08:21:03Z|


## Conferences

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[Capital Go](http://www.capitalgolang.com)|Washington, D.C., USA.|-|-|-|-|-|
[dotGo](https://www.dotgo.eu)|Paris, France.|-|-|-|-|-|
[GoCon](https://gocon.connpass.com/)|Tokyo, Japan.|-|-|-|-|-|
[GoDays](https://www.godays.io/)|Berlin, Germany.|-|-|-|-|-|
[GoLab](https://golab.io/)|Florence, Italy.|-|-|-|-|-|
[GopherChina](https://gopherchina.org)|Shanghai, China.|-|-|-|-|-|
[GopherCon](https://www.gophercon.com/)|Denver, USA.|-|-|-|-|-|
[GopherCon Australia](https://gophercon.com.au/)|Sydney, Australia.|-|-|-|-|-|
[GopherCon Brazil](https://gopherconbr.org)|Florianópolis, Brazil.|-|-|-|-|-|
[GopherCon Europe](https://gophercon.is/)|Berlin, Germany.|-|-|-|-|-|
[GopherCon India](https://www.gophercon.in/)|Pune, India.|-|-|-|-|-|
[GopherCon Israel](https://www.gophercon.org.il/)|Tel Aviv, Israel.|-|-|-|-|-|
[GopherCon Russia](https://www.gophercon-russia.ru)|Moscow, Russia.|-|-|-|-|-|
[GopherCon Singapore](https://gophercon.sg)|Mapletree Business City, Singapore.|-|-|-|-|-|
[GopherCon UK](https://www.gophercon.co.uk/)|London, UK.|-|-|-|-|-|
[GopherCon Vietnam](https://gophercon.vn/)|Ho Chi Minh City, Vietnam.|-|-|-|-|-|
[GoWayFest](https://goway.io/)|Minsk, Belarus.|-|-|-|-|-|
[GoWest Conference](https://www.gowestconf.com/)|Lehi, USA.|-|-|-|-|-|


## E-Books

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[GoBooks](https://github.com/dariubs/GoBooks)|List of Golang books|10458|1533|2|2015-05-05T10:45:36Z|2021-12-09T01:04:47Z|
[The-Golang-Standard-Library-by-Example](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)|Golang标准库。对于程序员而言，标准库与语言本身同样重要，它好比一个百宝箱，能为各种常见的任务提供完美的解决方案。以示例驱动的方式讲解Golang的标准库。|8258|1892|29|2013-04-14T02:21:23Z|2021-10-19T03:41:33Z|
[web-dev-golang-anti-textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook)|Learn how to write webapps without a framework in Go.|2908|275|9|2016-01-01T07:49:17Z|2021-10-19T11:14:43Z|
[gosuccinctly](https://github.com/thedevsir/gosuccinctly)| This is the companion repo for Go Succinctly by Amir Irani.|21|1|0|2018-09-02T05:36:10Z|2018-10-03T07:03:46Z|
[100 Go Mistakes: How to Avoid Them](https://www.manning.com/books/100-go-mistakes-how-to-avoid-them)||-|-|-|-|-|
[Building Web Apps With Go](https://codegangsta.gitbooks.io/building-web-apps-with-go/content/)||-|-|-|-|-|
[Continuous Delivery in Go](https://www.manning.com/books/continuous-delivery-in-go)| This practical guide to continuous delivery shows you how to rapidly establish an automated pipeline that will improve your testing, code quality, and final product.|-|-|-|-|-|
[For the Love of Go](https://bitfieldconsulting.com/books/love)|An introductory book for Go beginners.|-|-|-|-|-|
[Go 101](https://go101.org)|A book focusing on Go syntax/semantics and all kinds of details.|-|-|-|-|-|
[Go Bootcamp](http://golangbootcamp.com)||-|-|-|-|-|
[Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/)||-|-|-|-|-|
[Build an Orchestrator in Go](https://www.manning.com/books/build-an-orchestrator-in-go)||-|-|-|-|-|
[How To Code in Go eBook](https://www.digitalocean.com/community/books/how-to-code-in-go-ebook)|A 600 page introduction to Go aimed at first time developers.|-|-|-|-|-|
[Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)||-|-|-|-|-|
[Network Programming With Go](https://jan.newmarch.name/go/)||-|-|-|-|-|
[Practical Go Lessons](https://www.practical-go-lessons.com/)||-|-|-|-|-|
[Spaceship Go A Journey to the Standard Library](https://blasrodri.github.io/spaceship-go-gh-pages/)||-|-|-|-|-|
[The Go Programming Language](https://www.gopl.io/)||-|-|-|-|-|
[An Introduction to Programming in Go](http://www.golang-book.com/)||-|-|-|-|-|
[A Go Developer’s Notebook](https://leanpub.com/GoNotebook/read)||-|-|-|-|-|
[Writing A Compiler In Go](https://compilerbook.com)||-|-|-|-|-|
[Writing An Interpreter In Go](https://interpreterbook.com)|Book that introduces dozens of techniques for writing idiomatic, expressive, and efficient Go code that avoids common pitfalls.|-|-|-|-|-|


## Gophers

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[gophers](https://github.com/ashleymcnamara/gophers)|Gopher Artwork by Ashley McNamara|2537|125|13|2017-02-15T14:29:00Z|2019-04-12T18:38:12Z|
[gophers](https://github.com/egonelbre/gophers)|Free gophers|2535|131|5|2015-06-03T06:34:42Z|2020-06-18T06:10:29Z|
[free-gophers-pack](https://github.com/MariaLetta/free-gophers-pack)|✨ This pack of 100&#43; gopher pictures and elements will help you to build own design of almost anything related to Go Programming Language: presentations, posts in blogs or social media, courses, videos and many, many more.|2336|135|1|2019-04-02T22:11:29Z|2020-06-30T10:59:42Z|
[gophericons](https://github.com/shalakhin/gophericons)|34 gopher images for Go developers community|590|25|2|2015-08-22T14:41:34Z|2018-03-23T23:10:38Z|
[gopherize.me](https://github.com/matryer/gopherize.me)|Gopherize.me app|538|45|17|2017-01-25T12:51:35Z|2021-08-23T21:46:57Z|
[gopher-stickers](https://github.com/tenntenn/gopher-stickers)|gopher stickers|517|36|7|2014-11-09T16:41:03Z|2019-12-03T14:50:44Z|
[gopher-vector](https://github.com/golang-samples/gopher-vector)|Vector data of gopher|401|54|0|2013-03-31T05:45:36Z|2016-07-27T02:12:04Z|
[go-gopher](https://github.com/sillecelik/go-gopher)|The Go Gopher Amigurumi Pattern|98|13|1|2018-03-28T22:54:06Z|2021-08-14T08:22:14Z|
[gopher-logos](https://github.com/GolangUA/gopher-logos)|adorable gopher logos|96|8|1|2017-07-27T14:27:20Z|2021-06-24T19:17:44Z|
[gophers](https://github.com/rogeralsing/gophers)|random gopher graphics|54|3|2|2017-01-28T23:58:35Z|2020-08-06T15:16:29Z|
[Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector)|Go gopher Vector Data [.ai, .svg]|41|4|0|2014-09-03T17:29:51Z|2018-03-04T07:19:54Z|
[gophers](https://github.com/scraly/gophers)|Gopher artwork (Golang mascot)|12|5|0|2021-06-23T16:36:58Z|2021-10-28T12:12:38Z|


## Meetups

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[Basel Go Meetup](https://www.meetup.com/Basel-Go-Meetup/)||-|-|-|-|-|
[Belgrade Golang Meetup](https://www.meetup.com/golang-serbia/)||-|-|-|-|-|
[Belfast Gophers](https://www.meetup.com/Belfast-Gophers/)||-|-|-|-|-|
[Berlin Golang](https://www.meetup.com/golang-users-berlin/)||-|-|-|-|-|
[Brisbane Gophers](https://www.meetup.com/Brisbane-Golang-Meetup/)||-|-|-|-|-|
[Canberra Gophers](https://www.meetup.com/Canberra-Gophers/)||-|-|-|-|-|
[Go Language NYC](https://www.meetup.com/golanguagenewyork/)||-|-|-|-|-|
[Go London User Group](https://www.meetup.com/Go-London-User-Group/)||-|-|-|-|-|
[Go Remote Meetup](https://www.meetup.com/Go-Remote-Meetup/)||-|-|-|-|-|
[Go Toronto](https://www.meetup.com/go-toronto/)||-|-|-|-|-|
[Go User Group Atlanta](https://www.meetup.com/Go-Users-Group-Atlanta/)||-|-|-|-|-|
[GoBandung](https://www.meetup.com/GoBandung/)||-|-|-|-|-|
[GoBridge, San Francisco, CA](https://www.meetup.com/gobridge/)||-|-|-|-|-|
[GoCracow - Krakow, Poland](https://www.meetup.com/GoCracow/)||-|-|-|-|-|
[GoJakarta](https://www.meetup.com/GoJakarta/)||-|-|-|-|-|
[Golang Amsterdam](https://www.meetup.com/golang-amsterdam/)||-|-|-|-|-|
[Golang Argentina](https://www.meetup.com/Golang-Argentina/)||-|-|-|-|-|
[Golang Athens](https://www.meetup.com/Athens-Gophers/)||-|-|-|-|-|
[Golang Baltimore, MD](https://www.meetup.com/BaltimoreGolang/)||-|-|-|-|-|
[Golang Bangalore](https://www.meetup.com/Golang-Bangalore/)||-|-|-|-|-|
[Golang Belo Horizonte - Brazil](https://www.meetup.com/go-belo-horizonte/)||-|-|-|-|-|
[Golang Boston](https://www.meetup.com/bostongo/)||-|-|-|-|-|
[Golang Bulgaria](https://www.meetup.com/Golang-Bulgaria/)||-|-|-|-|-|
[Golang Cardiff, UK](https://www.meetup.com/Cardiff-Go-Meetup/)||-|-|-|-|-|
[Golang Copenhagen](https://www.meetup.com/Go-Cph/)||-|-|-|-|-|
[Golang Curitiba - Brazil](https://www.meetup.com/GolangCWB/)||-|-|-|-|-|
[Golang DC, Arlington, VA](https://www.meetup.com/Golang-DC/)||-|-|-|-|-|
[Golang Dorset, UK](https://www.meetup.com/golang-dorset/)||-|-|-|-|-|
[Golang Estonia](https://www.meetup.com/Golang-Estonia/)||-|-|-|-|-|
[Golang Gurgaon, India](https://www.meetup.com/Gurgaon-Go-Meetup/)||-|-|-|-|-|
[Golang Hamburg - Germany](https://www.meetup.com/Go-User-Group-Hamburg/)||-|-|-|-|-|
[Golang Israel](https://www.meetup.com/Go-Israel/)||-|-|-|-|-|
[Golang Joinville - Brazil](https://www.meetup.com/Joinville-Go-Meetup/)||-|-|-|-|-|
[Golang Kathmandu](https://www.meetup.com/Golang-Kathmandu/)||-|-|-|-|-|
[Golang Korea](https://www.meetup.com/GDG-Golang-Korea/)||-|-|-|-|-|
[Golang Lima - Peru](https://www.meetup.com/Golang-Peru/)||-|-|-|-|-|
[Golang Lyon](https://www.meetup.com/Golang-Lyon/)||-|-|-|-|-|
[Golang Marseille](https://www.meetup.com/fr-FR/Golang-Marseille/)||-|-|-|-|-|
[Golang Melbourne](https://www.meetup.com/golang-mel/)||-|-|-|-|-|
[Golang Mountain View](https://www.meetup.com/Golang-Mountain-View/)||-|-|-|-|-|
[Golang New York](https://www.meetup.com/nycgolang/)||-|-|-|-|-|
[Golang North East](https://www.meetup.com/en-AU/Golang-North-East/)||-|-|-|-|-|
[Golang Paris](https://www.meetup.com/Golang-Paris/)||-|-|-|-|-|
[Golang Poland](https://www.meetup.com/Golang-Poland/)||-|-|-|-|-|
[Golang Pune](https://www.meetup.com/Golang-Pune/)||-|-|-|-|-|
[Golang Singapore](https://www.meetup.com/golangsg/)||-|-|-|-|-|
[Golang Stockholm](https://www.meetup.com/Go-Stockholm/)||-|-|-|-|-|
[Golang Sydney, AU](https://www.meetup.com/golang-syd/)||-|-|-|-|-|
[Golang São Paulo - Brazil](https://www.meetup.com/golangbr/)||-|-|-|-|-|
[Golang Taipei](https://www.meetup.com/golang-taipei-meetup/)||-|-|-|-|-|
[Golang Thessaloniki](https://www.meetup.com/thessaloniki-golang-meetup/)||-|-|-|-|-|
[Golang Turkey](https://kommunity.com/goturkiye)||-|-|-|-|-|
[Golang Vancouver, BC](https://www.meetup.com/golangvan/)||-|-|-|-|-|
[Golang Vienna, Austria](https://www.meetup.com/viennago/)||-|-|-|-|-|
[Golang Казань](https://www.meetup.com/GolangKazan/)||-|-|-|-|-|
[Golang Москва](https://www.meetup.com/Golang-Moscow/)||-|-|-|-|-|
[Golang Питер](https://www.meetup.com/Golang-Peter/)||-|-|-|-|-|
[GoSF - San Francisco, CA](https://www.meetup.com/golangsf)||-|-|-|-|-|
[Istanbul Golang](https://www.meetup.com/Istanbul-Golang/)||-|-|-|-|-|
[Seattle Go Programmers](https://www.meetup.com/golang/)||-|-|-|-|-|
[Ukrainian Golang User Groups](https://www.meetup.com/uagolang/)||-|-|-|-|-|
[Utah Go User Group](https://www.meetup.com/utahgophers/)||-|-|-|-|-|
[Women Who Go - San Francisco, CA](https://www.meetup.com/Women-Who-Go/)||-|-|-|-|-|


## Style Guides

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[cockroach](https://github.com/cockroachdb/cockroach)|CockroachDB - the open source, cloud-native distributed SQL database.|22659|2950|4449|2014-02-06T00:18:47Z|2021-12-11T20:38:45Z|
[fabric](https://github.com/hyperledger/fabric)|Hyperledger Fabric is an enterprise-grade permissioned distributed ledger framework for developing solutions and applications. Its modular and versatile design satisfies a broad range of industry use cases. It offers a unique approach to consensus that enables performance at scale while preserving privacy.|12883|7569|52|2016-08-25T16:05:27Z|2021-12-11T16:53:14Z|
[guide](https://github.com/uber-go/guide)|The Uber Go Style Guide.|10100|1087|9|2018-11-10T18:14:59Z|2021-12-03T18:11:29Z|
[magnetico](https://github.com/boramalper/magnetico)|Autonomous (self-hosted) BitTorrent DHT search engine suite.|2621|318|75|2017-03-05T11:10:57Z|2021-01-03T10:58:00Z|
[go-styleguide](https://github.com/bahlo/go-styleguide)|🏆 Opinionated Styleguide for the Go language|1184|114|0|2017-07-29T10:03:30Z|2021-09-29T07:55:00Z|
[GitLab](https://docs.gitlab.com/ee/development/go_guide/)||-|-|-|-|-|
[Sourcegraph](https://about.sourcegraph.com/handbook/engineering/go_style_guide)||-|-|-|-|-|
[Thanos](https://thanos.io/tip/contributing/coding-style-guide.md/)||-|-|-|-|-|


### Twitter

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[@golang](https://twitter.com/golang)||-|-|-|-|-|
[@golang_news](https://twitter.com/golang_news)||-|-|-|-|-|
[@golangch](https://twitter.com/golangch)||-|-|-|-|-|
[@golangflow](https://twitter.com/golangflow)||-|-|-|-|-|
[@golangweekly](https://twitter.com/golangweekly)||-|-|-|-|-|


### Reddit

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[r/golang](https://www.reddit.com/r/golang/)||-|-|-|-|-|


## Websites

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[go](https://github.com/golang/go)|The Go programming language|93143|13807|7607|2014-08-19T04:33:40Z|2021-12-11T20:41:34Z|
[awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness)|A curated list of awesome awesomeness|28301|3449|16|2014-07-08T05:44:19Z|2021-12-08T03:17:29Z|
[awesome-remote-job](https://github.com/lukasz-madon/awesome-remote-job)|A curated list of awesome remote jobs and resources. Inspired by https://github.com/vinta/awesome-python|21427|2130|15|2015-01-02T00:31:34Z|2021-12-02T10:29:35Z|
[awesome-golang-workshops](https://github.com/amit-davidson/awesome-golang-workshops)|A curated list of awesome golang workshops.|452|20|0|2021-06-27T01:06:03Z|2021-07-13T14:14:28Z|
**[ARCHIVED]**  [golang-graphics](https://github.com/mholt/golang-graphics)|Community-contributed Go graphics files|139|10|1|2014-03-24T23:10:53Z|2015-08-24T21:30:06Z|
[gocryforhelp](https://github.com/ninedraft/gocryforhelp)|List of opensource projects looking for help|40|2|0|2016-05-09T14:30:41Z|2017-09-23T14:04:04Z|
[awesome-go-extra](https://github.com/xwjdsh/awesome-go-extra)|Parse awesome-go README file and generate a new README file with repo info.|19|2|0|2021-06-01T17:55:30Z|2021-12-10T21:10:43Z|
[Awesome Go @LibHunt](https://go.libhunt.com)|Your go-to Go Toolbox.|-|-|-|-|-|
[Go Community on Hashnode](https://hashnode.com/n/go)|Community of Gophers on Hashnode.|-|-|-|-|-|
[Go Forum](https://forum.golangbridge.org)|Forum to discuss Go.|-|-|-|-|-|
[Go In 5 Minutes](https://www.goin5minutes.com/)|5 minute screencasts focused on getting one thing done.|-|-|-|-|-|
[Go Code Club](https://www.youtube.com/watch?v=nvoIPQYdx9g&amp;list=PLEcwzBXTPUE_YQR7R0BRtHBYJ0LN3Y0i3)|A group of Gophers read and discuss a different Go project every week.|-|-|-|-|-|
[Go Proverbs](https://go-proverbs.github.io/)|Go Proverbs by Rob Pike.|-|-|-|-|-|
[Go Report Card](https://goreportcard.com)|A report card for your Go package.|-|-|-|-|-|
[go.dev](https://go.dev/)|A hub for Go developers.|-|-|-|-|-|
[Go Blog](https://blog.golang.org)|The official Go blog.|-|-|-|-|-|
[godoc.org](https://godoc.org/)|Documentation for open source Go packages.|-|-|-|-|-|
[Golang Developer Jobs](https://golangjob.xyz)|Developer Jobs exclusivly for Golang related Roles.|-|-|-|-|-|
[Golang Flow](https://golangflow.io)|Post Updates, News, Packages and more.|-|-|-|-|-|
[Golang News](https://golangnews.com)|Links and news about Go programming.|-|-|-|-|-|
[Golang Resources](https://golangresources.com)|A curation of the best articles, exercises, talks and videos to learn Go.|-|-|-|-|-|
[CodinGame](https://www.codingame.com/)|Learn Go by solving interactive tasks using small games as practical examples.|-|-|-|-|-|
[golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts)|Go mailing list.|-|-|-|-|-|
[Google Plus Community](https://plus.google.com/communities/114112804251407510571)|The Google&#43; community for #golang enthusiasts.|-|-|-|-|-|
[Gopher Community Chat](https://invite.slack.golangbridge.org)|Join Our New Slack Community For Gophers (Understand how it came).|-|-|-|-|-|
[Gophercises](https://gophercises.com/)|Free coding exercises for budding gophers.|-|-|-|-|-|
[gowalker.org](https://gowalker.org)|Go Project API documentation.|-|-|-|-|-|
[json2go](https://m-zajac.github.io/json2go)|Advanced JSON to Go struct conversion - online tool.|-|-|-|-|-|
[justforfunc](https://www.youtube.com/c/justforfunc)|Youtube channel dedicated to Go programming language tips and tricks, hosted by  Francesc Campoy @francesc.|-|-|-|-|-|
[Learn Go Programming](https://blog.learngoprogramming.com)|Learn Go concepts with illustrations.|-|-|-|-|-|
[Lille Gophers](https://lille-gophers.loscrackitos.codes/)|Golang talks community in Lille, France (@LilleGophers).|-|-|-|-|-|
[Made with Golang](https://madewithgolang.com/?ref=awesome-go)||-|-|-|-|-|
[r/Golang](https://www.reddit.com/r/golang)|News about Go.|-|-|-|-|-|
[studygolang](https://studygolang.com)|The community of studygolang in China.|-|-|-|-|-|
[Trending Go repositories on GitHub today](https://github.com/trending?l=go)|Good place to find new Go libraries.|-|-|-|-|-|
[TutorialEdge - Golang](https://tutorialedge.net/course/golang/)||-|-|-|-|-|


### Tutorials

	
|Name|Desc|Star|Fork|Issue|Created|Pushed|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
[build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang)|A golang ebook intro how to build a web with golang|39235|10364|105|2012-08-02T11:49:35Z|2021-11-28T00:14:58Z|
[go-patterns](https://github.com/tmrts/go-patterns)|Curated list of Go design patterns, recipes and idioms|17785|1642|60|2015-12-14T22:05:06Z|2021-08-12T14:21:12Z|
[learn-go-with-tests](https://github.com/quii/learn-go-with-tests)|Learn Go with test-driven development|15978|2059|30|2018-03-02T11:41:14Z|2021-11-29T22:52:50Z|
[learngo](https://github.com/inancgumus/learngo)|1000&#43; Hand-Crafted Go Examples, Exercises, and Quizzes|11825|1525|7|2018-10-15T11:12:00Z|2021-12-10T09:54:28Z|
[golang-cheat-sheet](https://github.com/a8m/golang-cheat-sheet)|An overview of Go syntax and features.|6033|830|15|2014-02-13T11:24:58Z|2021-12-11T10:13:28Z|
[golang-for-nodejs-developers](https://github.com/miguelmota/golang-for-nodejs-developers)|Examples of Golang compared to Node.js for learning|2687|185|0|2019-01-03T05:30:44Z|2021-09-11T10:20:38Z|
**[ARCHIVED]**  [working-with-go](https://github.com/mkaz/working-with-go)|A set of example golang code to start learning Go|1157|179|0|2014-05-04T21:29:05Z|2020-02-03T19:45:18Z|
[ethereum-development-with-go-book](https://github.com/miguelmota/ethereum-development-with-go-book)|📖 A little guide book on Ethereum Development with Go (golang)|1061|247|7|2018-05-16T09:22:56Z|2021-11-07T02:57:49Z|
[go-clean-template](https://github.com/evrone/go-clean-template)|Clean Architecture template for Golang services|459|58|0|2021-01-18T09:29:43Z|2021-12-10T11:00:02Z|
[goapp](https://github.com/bnkamalesh/goapp)|An opinionated guideline to structure &amp; develop a Go web application/service|310|23|0|2020-07-04T11:47:44Z|2021-11-04T13:29:00Z|
[design-patterns](https://github.com/shubhamzanwar/design-patterns)|common creational, behavioural and structural patterns implemented in go 🤩|62|5|0|2020-09-24T05:48:15Z|2020-11-07T17:58:20Z|
[go-patterns](https://github.com/haveyoudebuggedit/go-patterns)||0|1|0|2021-06-25T14:06:07Z|2021-06-25T14:08:21Z|
[Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)|Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.|-|-|-|-|-|
[Go By Example](https://gobyexample.com/)|Hands-on introduction to Go using annotated example programs.|-|-|-|-|-|
[Go database/sql tutorial](http://go-database-sql.org/)|Introduction to database/sql.|-|-|-|-|-|
[Go Playground for iOS](https://codeplayground.app)|Interactively edit &amp; play Go snippets on your mobile device.|-|-|-|-|-|
[Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)||-|-|-|-|-|
[Games With Go](https://gameswithgo.org/)|A video series teaching programming and game development.|-|-|-|-|-|
[50 Shades of Go](https://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)|Traps, Gotchas, and Common Mistakes for New Golang Devs.|-|-|-|-|-|
[Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30)|How to cancel MySQL queries.|-|-|-|-|-|
[Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9)|How to cache slow database queries.|-|-|-|-|-|
[Golang Tutorial Guide](https://www.freecodecamp.org/news/golang-tutorial-list-free-courses-learn-go-programming-language/)|A List of Free Courses to Learn the Go Programming Language|-|-|-|-|-|
[Golangbot](https://golangbot.com/learn-golang-series/)|Tutorials to get started with programming in Go.|-|-|-|-|-|
[GolangCode](https://golangcode.com/)|Collection of code snippets and tutorials to help tackle every day issues.|-|-|-|-|-|
[GopherSnippets](https://gophersnippets.com/)|Code snippets with tests and testable examples for the Go programming language.|-|-|-|-|-|
[Gosamples](https://gosamples.dev/)|Collection of code snippets that let you solve everyday code problems.|-|-|-|-|-|
[Hackr.io](https://hackr.io/tutorials/learn-golang)|Learn Go from the best online golang tutorials submitted &amp; voted by the golang programming community.|-|-|-|-|-|
[How to Benchmark: dbq vs sqlx vs GORM](https://medium.com/@rocketlaunchr.cloud/how-to-benchmark-dbq-vs-sqlx-vs-gorm-e814caacecb5)|Learn how to benchmark in Go. As a case-study, we will benchmark dbq, sqlx and GORM.|-|-|-|-|-|
[How To Deploy a Go Web Application with Docker](https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker)|Learn how to use Docker for Go development and how to build production Docker images.|-|-|-|-|-|
[How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go)|Get started with Godog — a Behavior-driven development framework for building and testing Go applications.|-|-|-|-|-|
[Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)|We’ll write an API with the help of the powerful Gorilla Mux.|-|-|-|-|-|
[Go Language Tutorial](https://www.javatpoint.com/go-tutorial)| Learn Go language Tutorial|-|-|-|-|-|
[Learning Golang - From zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/)|Getting started with golang for beginner.|-|-|-|-|-|
[package main](https://www.youtube.com/packagemain)|YouTube channel about Programming in Go.|-|-|-|-|-|
[Programming with Google Go](https://www.coursera.org/specializations/google-golang)|Coursera Specialization to learn about Go from scratch.|-|-|-|-|-|
[Saving a Third of Our Memory by Re-ordering Go Struct Fields](https://qvault.io/golang/struct-field-ordering-memory/)|How inefficient field ordering in Go structs|-|-|-|-|-|
[A Tour of Go](https://tour.golang.org/)|Interactive tour of Go.|-|-|-|-|-|
[A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo)|Building a Golang site for e-commerce (demo included).|-|-|-|-|-|
[Your basic Go](https://yourbasic.org/golang)|Huge collection of tutorials and how to’s.|-|-|-|-|-|
[Learning Go by examples](https://dev.to/aurelievache/learning-go-by-examples-introduction-448n)| Serie of article in order to learn Golang language by concrete applications as example|-|-|-|-|-|
[Go Tutorial](https://www.tutorialspoint.com/go/index.htm)| Learn Go programming|-|-|-|-|-|
[The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)||-|-|-|-|-|

