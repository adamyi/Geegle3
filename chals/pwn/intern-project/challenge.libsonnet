{
  services: [
    {
      name: "intern-project",
      category: "pwn",
      clustertype: "team",
    },
  ],
  flags: [
    {
      Flag: "GEEGLE{DM98DAMW98AMD98WAMD98M}",
      Points: 250,
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
        Please find the code here: <a href="https://intern-project.corp.geegle.org">https://intern-project.corp.geegle.org</a>.

        Good luck, have fun.

        CodeReview Team
      |||,
      "DependsOnPoints": 1100,
      "Delay": 900000
    },
  ],
}
