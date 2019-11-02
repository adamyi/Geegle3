# PasteWeb (XSS)

### Payload

Create a paste with the following title:
```
<iframe id="ps"></iframe><script>
```
and the following content:
```
a<form id="fb"><input name="ownerDocument"/><script>alert(document.cookie);</script></form>
```

and access it.

### Idea
1. unfiltered output at title, recreate unsandboxed iframe with same DOM ID (`ps`)
2. `<script>` has highest priority in HTML, use `<script>` to disable DOMPurify
3. Bug in paste.html js -> try catch doesn't halt execution
4. Create element with same DOM Element ID `fb` to mess up with js queryselector
5. jQuery Script Gadget to bypass CSP Strict-Dynamic (https://www.blackhat.com/docs/us-17/thursday/us-17-Lekies-Dont-Trust-The-DOM-Bypassing-XSS-Mitigations-Via-Script-Gadgets.pdf) (https://github.com/jquery/jquery/blob/30e1bfbdcb0ff658f1fa128b72480194e8ecb926/src/manipulation.js#L103)

## author
adamyi
