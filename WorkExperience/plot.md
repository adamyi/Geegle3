# Challenge descriptions

And associated copy.

------------------

## Ch 1

### All:

Info:

    Welcome to Geegle!

    We're so happy to have you as one of our newest Work Experience placements. We don't want to give you anything boring to work on, so we're going to throw you into the thick of it and see how you go. After all, the best way to learn how to swim is to jump in the deep end right?
    To start with we've got a couple of basic tasks for you - some general knowledge and basic programming challenges.  Work your way through most of those, and then we'll get you into the real work. Have fun!

    All the best,

    Adam
    HR Manager

trivia + scripting stuff is the first thing they do


## Ch 2

Second email to be sent 10 seconds after the first.
misc-email-intercept is not required to progress, but gains extra points.


### misc-incident-response:

Info:

    Hey newbie,

    So I know you haven't quite finished the intro tasks yet, but we've just got word from our monitoring team that someone is trying to use a Dictionary brute force to login to our staff portal. I've got the traffic logs here and it's all hands on deck, so I'm going to have to pull you away from the intro tasks to help us out. Can you take a look and figure out which IP address the attacker is coming from so we can IP ban them?

    Thanks so much,

    Adam
    Work Experience Manager

You are tasked with finding who is attacking us... Here is a file- email us the ip thats attacking us

### misc-email-intercept:

Info:

    Hi,

    Apologies, I seem to have attached the wrong file to the last email. Please do not read the last file, it was a confidential matter. Instead, please see attached the correct file.

    Thanks,

    Adam
    Work Experience Manager

You intercepted an email between HR and your manager... can you decode what they were saying?

## Ch 3

Both challenges sent together, with the following info.
Solving either challenge grants progression.

Info:

    Hey,

    We forgot to get you setup with access to a few systems unfortunately, which is going to slow work down for a while. Ideally, you should have access to both [Intranet](web-intranet) and [Account Manager](web-intern-account-manager), however neither of these have been setup for you. Unfortunately the HR Rep is away on long service leave so we won't be able to put it through in time - what I've managed to do is organise permission for you to perform a security audit of the [intranet](web-intranet) services - if you can get into it, you can update the records yourself. We do NOT have permission for a similar test on [Account Manager](web-intern-account-manager), but if the HR Rep gets it changed while he's away he'll send you a secret code - send it to me if you get it so I can verify it's sorted thanks.

    I suppose if you want to get paid, you better get onto it!

    Cheers,

    Adam
    Work Experience Manager

### web-intern-account-manager:

Is a SSO Intern Management website to increase their pay, Where if you Login, it will inform you have intern's are not able to access this site.

### web-intranet:

Is a Local intranet browser, the idea is that interns don't have access to the employee VPN, so to access documentation on internal Geegle services, they must access the documents via this page

## Ch 4

### web-filesystem:

Info:

    ATTENTION

    We have logged an unauthorised access to one of either [account manager](web-intern-account-manager) or [intranet](web-intranet). This is an automated alert message informing you to verify that the [filesystem](is still secure).

    Automated Message, please do not reply.

    == This message is from IceEar Monitoring software ==

you found a link to an open file browser on the internal network, are there any cool files here?
      - have etcpasswd in here. The file will have to have the location of a server in it.

### crypto-etcpasswd:

you have leaked two files from a server, can you login to the server?

## Ch 5

### crypto-intercepted:

Info:

    so uhhh we may be in trouble

    they've figured out that we're poking into stuff we're not meant to. i don't think they know who is doing it yet but everyone seems to have switched to using encrypted email. ive intercepted this but it's useless. can you crack the encryption so we can see what they're saying about us?

    ta

    Lachlan
    Work Experience Placement, 2019

So the employees found that one of the interns was intercepting and reading shit they wern't supposed to, they started to communicate using a custom cipher they developed... We managed to intercept both the cipher text and the key.. Can you decode it??

## Ch 6

Info:

    yo its me again

    ive found what looks like a new game geegles developing but its behind some custom crypto stuff. nobody can write decent crypto so odds are its broken somehow. do me a solid and check [Private](web-private-files) to find the source code for [this](pwn-internproject) binary program then break the crypto for me so I can get into the game.

    ta

    Lachlan
    Work Experience Placement, 2019

Challs sent together as described above

### pwn-intern-project:


### web-privatefile:



## Ch 8

Info:

    Hi!

    Thanks for taking part in the Work Experience Program with Geegle! We really hoped you've enoyed working with us!
    Unfortunately all good things must come to an end. We'd love to have you back, so please keep an eye out for our awesome Intern program which runs soon! A few new interns are arriving in the building today, so just be careful as you leave not to disturb the [Guest Kiosks](web-guest-kiosk), as that is how we tell the interns apart from the general public when they sign in!

    Thank you so much for working with us, and we can't wait to see you again soon.

    Thanks,

    Adam
    HR Manager

### web-guest-kiosk:
        
