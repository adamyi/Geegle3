{
services: [
   {
     name: "payroll",
     category: "pwn",
     staticfiles: [
       {
         filename: "payroll",
         flags:[
           {
             name: "salary",
             value: "underpaid",
           },
         ],
       },
     ],
   },
 ],
 flags: [
    {
      Flag: "GEEGLE{IUDSNAIUDWAND9MADA0D9M}",
      Points: 150,
    },
  ],
  emails: [
    {
      "Sender": "payroll@geegle.org",
      "Title": "Attn: Security Team",
      "Body": |||
        Hello Security Team,
        
        We have just received the product handover from the dev team for our new internal payroll system. Would you be able to have a look and check if it works okay?
        Here's the link: https://payroll.corp.geegle.org _If you're on the go, you can download a version at https://ssfe.corp.geegle.org/s/31e795a6004cbec960a27195b213b6df/salary=underpaid/payroll).
        The last thing we would want is to have some intern submit a ridiculous amount of hours …
        
        Regards,
        Payroll
      |||,
      "DependsOnPoints": 1100,
      "Delay": 0
    },
  ],
}
