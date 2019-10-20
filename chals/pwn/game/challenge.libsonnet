{
    services: [
      {
        name: "game",
        category: "pwn",
        staticfiles: [
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
      },
    ],
    flags: [
    {
      Flag: "GEEGLE{YAAFN4EM3209Q3MD09M}",
      Points: 100,
    },
  ],
  emails: [
    {
      "Sender": "hr@geegle.org",
      "Title": "Shall we play a game?",
      "Body": |||
        We know that being a Geegler is hard.
        As part of our continued commitment to employee satisfaction, we regularly release games for our employees to play.
        Good luck have fun!
        Play it here: https://ssfe.corp.geegle.org/s/270e32d2aa9faa1785ff58ad3a078966/funlevel=100/game.
      |||,
      "DependsOnPoints": 200,
      "Delay": 10000
    },
  ],
}
