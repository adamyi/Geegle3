{
  services: [
    {
      name: "shell",
      category: "pwn",
    },
  ],
  flags: [
    {
      Flag: "GEEGLE{0J398MD93AMD9AID}",
      Points: 100,
    },
  ],
  clistaticfiles: [
    {
      filename: "shell",
      flags: [
        {
          name: "rollyourown",
          value: "always",
        },
      ],
    },
  ],
  emails: [
    {
      "Sender": "dogfood@geegle.org",
      "Title": "[Announcement] Dogfooding Geegle Shell",
      "Body": |||
          At Geegle we encourage everyone to join in on testing our new products before they go to production.
          We've got out new shell that we think is ready to go live, but we want your feedback! We've created it to allow for better interactions with remote services, and easier authentication.
          
          Try it out here: https://shell.corp.geegle.org.
      |||,
      "DependsOnPoints": 1,
      "Delay": 180000
    },
  ],
}
