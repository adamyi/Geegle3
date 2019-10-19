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
          Hello Security Team,\n\nWe have just received the product handover from the dev team for our new internal payroll system. Would you be able to have a look and check if it works okay?\nHere's the link: https://payroll.corp.geegle.org\nThe last thing we would want is to have some intern submit a ridiculous amount of hours … \n\n Regards,\nPayroll
      |||,
      "DependsOnPoints": 1100,
      "Delay": 0
    },
  ],
}
