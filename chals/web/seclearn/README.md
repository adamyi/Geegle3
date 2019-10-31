# SecLearn

Intro:
```
Mail From: adamyi@geegle.org
Mail To: <CTF-Player>@geegle.org
Title: SecLearn Testing

Hi there,
Steve from the security team has just built an internal cyber-security awareness training platform called SecLearn. He asked me to test it, but... to be honest, I think it's kind of lame. It looks just like a flash cards app for language learners and I don't see how it's related to security at all. He said he doesn't really know how to code so it lacks functionalities (I accidentally pasted my password there and can't even find out how to delete it from the system...) Anyway, I helped him add some NLP functionalities to make it robust. Hopefully, we will roll it out company-wide, cuz I do understand the importance of internal security training.

Before that, I'd still like you to take a look and tell me what you think. If you found any bug in the system, just shoot me an email with the url, and I'll take a look.

Thanks,
Adam
```


## Idea
* Send URL to adamyi by email to trigger Headless Chrome
* Abuse Chrome's XSS Auditor 
* Browser Side Channel

## Payload

https://seclearn-solver.corp.geegle.org/b1da4a0e-fbf7-11e9-aad5-362b9e155667/

## author
adamyi
