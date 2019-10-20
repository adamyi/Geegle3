{
  services: [
    {
      name: "geelang",
      category: "pwn",
      staticfiles: [
        {
          filename: "geelang-compiler",
          flags:[
            {
              name: "type",
              value: "totallynotassembly",
            },
            {
              name: "assembly",
              value: "stopsayingthat",
            },
          ],
        },
      ],
    },
  ],
  flags: [
    {
      Flag: "GEEGLE{03MDD98M3Q09MDAMD093WD}",
      Points: 500,
    },
  ],
  emails: [
    {
      Sender: "adamyi@geegle.org",
      Title: "This got dumped on my desk",
      Body: |||
        Tanana gave me this stupid language. He says it's the next generation of computer science, but I told him it's just assembly. Anyway, if you want to try it, it's at http://ssfe.corp.geegle.org/s/23dd7dd62a47dde48f6b2998841dfbae/assembly=stopsayingthat/type=totallynotassembly/geelang-compiler
      |||,
      DependsOnPoints: 200,
      Delay: 1000,
    },
  ],
}
