{
  // TODO(lachlan): you need the onboarding service here
  flags: [
    {
      Flag: "GEEGLE{ISND3N98DN19D1NDND339D}",
      Points: 100,
    },
  ],
  emails: [
    {
      Sender: "hr@geegle.org",
      Title: "Geegle Onboarding Survey",
      Body: |||
        Thanks for joining us at Geegle!
        
        It's important to us that you enjoy your time here with us, so we'd love to to complete a survey for us.
        In order to make the experience as seamless as possible for you, your work machine should have TellGeegle installed on it already. In case it doesn't or you're one of our BYOD workers, you can download the executable from https://sffe.corp.geegle.org/s/d3230df6465754f85ee9e9c300f26c50/tellgeegle.
        
        Remember, your experience matters to us!
      |||,
      DependsOnPoints: 1,
      Delay: 10000,
    },
  ],
}
