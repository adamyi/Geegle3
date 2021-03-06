{
  services: [
    {
      name: "employees",
      category: "web",
      clustertype: "team",
    },
  ],
  flags: [
    {
      Flag: "GEEGLE{D9ND9nKJNDNJS032009fjSUS}",
      Points: 100,
    },
  ],
  emails: [
    {
      Sender: "adamt@geegle.org",
      Title: "Employee Database",
      Body: |||
        Hi there,
        
        Seeing as you've been here for a few days now, I'd like to introduce to you out internal Employee Database.
        Running next-gen anti-analysis blockchain algorithms to store data like never before. Feel free to use this to find other employees to communicate with :).
        
        <a href="https://employees.corp.geegle.org">https://employees.corp.geegle.org</a>
        
        Thanks,
        Adam
      |||,
      DependsOnPoints: 200,
      Delay: 0,
    },
  ],
}
