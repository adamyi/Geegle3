# kiosk (Geegle Kiosk)

basic chals

Idea:
two challanges

first is buffer overflow via qrcode upload.

## Solution
1. Gen qr code, utility in /test/genqrcode/<data>
2. Upload QR Code containing data of 32 characters padding + "adamt" string

second is incorrect permissions on directories

1. Notice pickup button on index.html directs to /pickup/guest
2. Notice Enter Guest details shows current logged in user as `adamt`
3. Go to /pickup/adamt

## author
adamt
