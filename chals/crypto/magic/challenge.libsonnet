{
  services: [
    {
      name: "magic",
      category: "crypto",
      clustertype: "team",
    },
  ],
  flags: [
    {
      Flag: "GEEGLE{M0DM9ADN33M9DMA9DMMD}",
      Points: 200,
    },
  ],
  emails: [
    {
      "Sender": "alerts@geegle.org",
      "Title": "[Do Not reply] Holiday Season",
      "Body": |||
        -- Alert --

        Welcome to Geegle, we appreciate holidays. Here's a magic trick, can you figure out how it's done?
        Details: <a href="https://magic.corp.geegle.org">https://magic.corp.geegle.org</a>

        Please do not reply to this email.
      |||,
      "DependsOnPoints": 1,
      "Delay": 100000
    },
  ],
}
