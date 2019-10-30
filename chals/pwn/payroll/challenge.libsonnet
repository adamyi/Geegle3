{
services: [
   {
     name: "payroll",
     category: "pwn",
     clustertype: "team",
   },
 ],
 flags: [
    {
      Flag: "GEEGLE{IUDSNAIUDWAND9MADA0D9M}",
      Points: 150,
    },
  ],
  clistaticfiles: [
    {
      filename: "payroll",
      flags: [
        {
          name: "salary",
          value: "underpaid",
        },
      ],
    },
  ],
  emails: [
    {
      "Sender": "payroll@geegle.org",
      "Title": "Attn: Security Team",
      "Body": |||
        Hello Security Team,
        
        We have just received the product handover from the dev team for our new internal payroll system. Would you be able to have a look and check if it works okay?
        Here's the link: https://payroll.corp.geegle.org.
        The last thing we would want is to have some intern submit a ridiculous amount of hours …
        
        Regards,
        Payroll
      |||,
      "DependsOnPoints": 1100,
      "Delay": 0
    },
  ],
}
