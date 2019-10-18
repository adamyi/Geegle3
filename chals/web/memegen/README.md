# Memegen

## Vuln
1. PHP Upload - no extension check, only mime check
2. LD_PRELOAD inject RCE

## Write-Up

Upload `exploit/adamyi.jpg.php`

Then POST to `uploads/adamyi.jpg.php` with body param cmd=`content of solve/exploit.php`

`id` is executed with output at `uploads/adamyi.out`

## Alternative Solution

SSRF to PHP-FPM CGI

https://cxsecurity.com/issue/WLB-2013010139

## TODO

put a flag to reflect RCE
