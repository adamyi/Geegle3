{
  services: [
    {
      name: "geelang",
      category: "pwn",
      clustertype: "team",
    },
  ],
  flags: [
    {
      Flag: "GEEGLE{03MDD98M3Q09MDAMD093WD}",
      Points: 500,
    },
  ],
  clistaticfiles: [
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
    {
      filename: "/lib/x86_64-linux-gnu/libc.so.6",
    },
  ],
  emails: [
    {
      Sender: "adamyi@geegle.org",
      Title: "This got dumped on my desk",
      Body: |||
        Tanana gave me this stupid language. He says it's the next generation of computer science, but I told him it's just assembly. Anyway, if you want to try it, we're storing the compiler at <a href="https://geelang.corp.geegle.org">https://geelang.corp.geegle.org</a>.
      |||,
      DependsOnPoints: 600,
      Delay: 1000,
    },
  ],
}
