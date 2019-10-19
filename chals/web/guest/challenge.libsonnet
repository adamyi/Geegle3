{
  services: [
    {
      name: "guest",
      category: "web",
    },
  ],
  flags: [
    {
      Flag: "GEEGLE{UAN398DN398DN93D}",
      Points: 100,
    },
  ],
  emails: [
    {
      Sender: "events@geegle.org",
      Title: "Geegle Summit 2020!",
      Body: |||
        You're invited to the 2nd annual Geegle Summit!
        
        We're bringing together major leaders, industry experts and top men to discuss what it takes to take Geegle.org to the next level.
        Please make sure you upload and scan your QR code here: https://guest.corp.geegle.org/ before you enter the venue. On behalf of everyone at Geegle.org, we look forward to seeing you at our Summit 2020.
        
        Regards,
        Parry Lage
      |||,
      DependsOnPoints: 1,
      Delay: 60000,
    },
  ],
}
