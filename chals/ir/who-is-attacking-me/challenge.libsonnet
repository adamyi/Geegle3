{
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
          -- Alert --\n\nWe have received a report of an attacker from an unknown origin. Please identify the origin and the nature of the attack.\n Details: https://attack-check.corp.geegle.org\n\nPlease do not reply to this email.\n
      |||,
      "DependsOnPoints": 600,
      "Delay": 900000
    },
  ],
}
