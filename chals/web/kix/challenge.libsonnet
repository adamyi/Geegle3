{
  flags: [
    {
      Flag: "GEEGLE{FVHSIVKNM734O02NFI9HN}",
      Points: 300,
    },
  ],
  emails: [
    {
      Sender: "dogfood@geegle.org",
      Title: "[Announcement] Dogfooding Geegle Docs",
      Body: |||
        At Geegle we encourage everyone to join in on testing our new products before they go to production.\nOur newest product is Docs. We've created it to allow for better formatted pdf documents so you don't have to send your fellow Geeglers ugly markdown files.\nTry it out here: https://docs.corp.geegle.org\n
      |||,
      DependsOnPoints: 600,
      Delay: 10000,
    },
  ],
}
