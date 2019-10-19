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
      "Sender": "HR@geegle.org",
      "Title": "Shall we play a game?",
      "Body": |||
          We know that being a Geegler is hard.\nAs part of our continued commitment to employee satisfaction, we regularly release games for our employees to play. \nGood luck have fun! \n Play it here: https://game.corp.geegle.org.
      |||,
      "DependsOnPoints": 200,
      "Delay": 10000
    },
  ],
}