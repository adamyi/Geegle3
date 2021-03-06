{
  services: [
    {
      name: "tellgeegle",
      category: "re",
      clustertype: "team",
    },
  ],
  flags: [
    {
      Flag: "GEEGLE{ISND3N98DN19D1NDND339D}",
      Points: 300,
    },
  ],
  clistaticfiles: [
    {
      filename: "tellGeegle",
      flags: [
        {
          name: "importance",
          value: "definitelysuperduperhigh",
        }
      ],
    },
  ],
  emails: [
    {
      Sender: "hr@geegle.org",
      Title: "Geegle Onboarding Survey",
      Body: |||
        Thanks for joining us at Geegle!

        It's important to us that you enjoy your time here with us, so we'd love to to complete a survey for us.
        In order to make the experience as seamless as possible for you, your work machine should have TellGeegle installed on it already. In case it doesn't or you're one of our BYOD workers, you can download the executable from the centralised version at <a href="https://tellgeegle.corp.geegle.org">https://tellgeegle.corp.geegle.org</a>.

        Remember, your experience matters to us!
      |||,
      DependsOnPoints: 1350,
      Delay: 10000,
    },
  ],
}
