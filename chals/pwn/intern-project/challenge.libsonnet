{
  services: [
    {
      name: "intern-project",
      category: "pwn",
      staticfiles: [
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
    },
  ],
  flags: [
    {
      Flag: "GEEGLE{DM98DAMW98AMD98WAMD98M}",
      Points: 100,
    },
  ],
  emails: [
    {
      "Sender": "codereview@geegle.org",
      "Title": "You have been selected…",
      "Body": |||
        Hello and congratulations
        
        You have been randomly selected as part of our intern improvement scheme to review the code of an intern's summer project.
        Please find the code here: https:// TODO update URL
      |||, 
      "DependsOnPoints": 1100,
      "Delay": 900000
    },
  ],
}
