{
    services: [
      {
        name: "game",
        category: "pwn",
        clustertype: "team",
      },
    ],
    flags: [
    {
      Flag: "GEEGLE{YAAFN4EM3209Q3MD09M}",
      Points: 100,
    },
  ],
  clistaticfiles: [
    {
      filename: "game",
      flags:[
        {
          name: "funlevel",
          value: "100",
        },
      ],
    },
  ],
  emails: [
    {
      "Sender": "hr@geegle.org",
      "Title": "Shall we play a game?",
      "Body": |||
        We know that being a Geegler is hard.
        As part of our continued commitment to employee satisfaction, we regularly release games for our employees to play. You can download a local version, or play the main version on our server!
        
        Good luck have fun!
        Download it at: <a href="https://game.corp.geegle.org">https://game.corp.geegle.org</a>
      |||,
      "DependsOnPoints": 200,
      "Delay": 10000
    },
  ],
}
