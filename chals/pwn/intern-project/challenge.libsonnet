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
          Hello and congratulations\n\nYou have been randomly selected as part of our intern improvement scheme to review the code of an intern's summer project.\nPlease find the code here: https://intern-project.corp.geegle.org
      |||,
      "DependsOnPoints": 1100,
      "Delay": 900000
    },
  ],
}
