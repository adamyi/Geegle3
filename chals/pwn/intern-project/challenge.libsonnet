{
  services: [
    {
      name: "intern-project",
      category: "pwn",
    },
  ],
  flags: [
    {
      Flag: "GEEGLE{DM98DAMW98AMD98WAMD98M}",
      Points: 100,
    },
  ],
  clistaticfiles: [
    {
      filename: "intern-project",
      flags:[
        {
          name: "codereview",
          value: "alwaysfun",
        },
        {
          name: "sarcasm",
          value: "maybe",
        },
      ],
    },
  ],
  emails: [
    {
      "Sender": "codereview@geegle.org",
      "Title": "You have been selected…",
      "Body": |||
        Hello and congratulations
        
        You have been randomly selected as part of our intern improvement scheme to review the code of an intern's summer project.
        Please find the code here: https://ssfe.corp.geegle.org/s/6de4afadd2dd3bfd7fea4a51b239423b/codereview=alwaysfun/sarcasm=maybe/intern-project, and if you find anything interesting take a look at on our master version at https://intern.corp.geegle.org.
      |||, 
      "DependsOnPoints": 1100,
      "Delay": 900000
    },
  ],
}
