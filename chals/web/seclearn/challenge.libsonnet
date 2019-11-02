{
  services: [
    {
      name: "seclearn",
      category: "web",
      clustertype: "team",
    },
    {
      name: "seclearn-solver",
      category: "web",
      clustertype: "master",
    },
  ],
  flags: [
    {
      Flag: "GEEGLE{MNA9MDW9MAD92D}",
      Points: 500,
    },
  ],
  emails: [
    {
      Sender: "adamyi@geegle.org",
      Title: "SecLearn Testing",
      Body: |||
        Hi there,
        Steve from the security team has just built an internal cyber-security awareness training platform called SecLearn : <a href="https://seclearn.corp.geegle.org">https://seclearn.corp.geegle.org</a>. He asked me to test it, but... to be honest, I think it's kind of lame. It looks just like a flash card app for language learners and I don't see how it's related to security at all. He said he doesn't really know how to code so it lacks functionality (I accidentally pasted my password there and can't even figure out how to delete it from the system…) Anyway, I helped him add some NLP functionality to make it more robust. Hopefully we can roll it out company-wide, cuz I do  understand the importance of internal security training.
        Before that, I'd still like you to take a look and tell me what you think. If you find any bugs in the system, just shoot me an email with the url, and I'll take a look.
        
        Thanks,
        Adam
      |||,
      DependsOnPoints: 1350,
      Delay: 30000,

    },
  ],
}
