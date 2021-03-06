{
  services: [
    {
      name: "who-is-attacking-me",
      category: "ir",
      clustertype: "team",
    },
  ],
  flags: [
    {
      Flag: "GEEGLE{9UMND39Q39DM39MD0M3D}",
      Points: 100,
    },
  ],
  emails: [
    {
      "Sender": "cyberdefence-noreply@geegle.org",
      "Title": "[Do Not reply] DDoS Detected",
      "Body": |||
        -- Alert --

        We have received a report of an attacker from an unknown origin. Please identify the origin and the nature of the attack.
        Details: <a href="https://who-is-attacking-me.corp.geegle.org">https://who-is-attacking-me.corp.geegle.org</a>

        Please do not reply to this email.
      |||,
      "DependsOnPoints": 600,
      "Delay": 300000
    },
  ],
}
