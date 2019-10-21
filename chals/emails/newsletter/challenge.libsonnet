{
  emails: [
    {
      "Sender": "Geegle-Newsletter@geegle.org",
      "Title": "Geegle News 0x01",
      "Body": |||
        Welcome to Geegle Newsletter!
        
        We work to send out a newsletter every so often to make sure that all our Geeglers are up to date with everything going on at Geegle.
        This is issue 0x01, our introductory edition. Light on content and easy to read, just the way everyone likes newsletters.

        THe big news for this newsletter is our brand new <em>CLI-Relay</em>. This is a brand new way to interact with remote executables, courtesy of Geegle.

        We realise new things are hard to use, so here's some instructions:

        <ol>
        <li>
        Once you access the CLI web interface for a relevant URL, capture your cookie from the browser. 
        </li>
        <li>
        Download the binary from the webpage.
        </li>
        </li>
        Use the binary as decribed on the CLI web page!
        </li>
        </ol>

        Alternatively, you can also interact with the binary through the CLI web interface if that's easier.

        That's all for this edition, enjoy!

        Thats a wrap,
        Geegle Newsletter
      |||,
      "DependsOnPoints": 1,
      "Delay": 600000
    },
    {
      "Sender": "Geegle-Newsletter@geegle.org",
      "Title": "Geegle News 0x02",
      "Body": |||
        Welcome to Geegle Newsletter!
        
        Welcome to the CTF edition of the newsletter! We wanted to share some things we always notice in CTF's so you watch out for them whenever you play!
        
        There's always one of these...
        <img src="https://www.brawl.com/attachments/37697/">
        So you better be ready for it!
        
        
        There's hidden flags:
        <img src="http://i.imgur.com/8NWXTbA.png">
        
        
        And...
        Someone tries to be funny and attack the CTF infrastructure
        <img src="https://media.makeameme.org/created/why-the-fuck-f47ctf.jpg">
        
        
        So yeah, I bet you can't wait to see all of that sometime soon!
        
        Thats a wrap,
        Geegle Newsletter
      |||,
      "DependsOnPoints": 1,
      "Delay": 2400000
    },
    {
    {
      "Sender": "Geegle-Newsletter@geegle.org",
      "Title": "Geegle News 0x03",
      "Body": |||
        Welcome to Geegle Newsletter!
        
        In edition 0x03, we have some fun testing info. We promise it's nothing to do with anything important, just a chance to learn something new.
        
        <h1>Testing on the Toilet</h1>
        <h2>Only Verify State-Changing Method Calls</h2>
        
        <p>
        <b>Method calls on another object fall into one of two categories:</b>
        <ul>
        <li>
        State-changing: methods that have side effects and change the world outside the code under test,
e.g., sendEmail(), saveRecord(), logAccess().
        </li>
        <li>
        Non-state-changing: methods that return information about the world outside the code under test and don't modify anything, e.g., getUser(), findResults(), readFile().
        </li>
        </ul>
        </p>
        
        <p>
        You should usually avoid verifying that non-state-changing methods are called:
        <ul>
        <li>
        It is often redundant: a method call that doesn't change the state of the world is meaningless on its own. The code under test will use the return value of the method call to do other work that you can assert.
        </li>
        <li>
        It makes tests brittle: tests need to be updated whenever method calls change. For example, if a test is expecting mockUserService.isUserActive(USER) to be called, it would fail if the code under test is modified to call user.isActive() instead.
        </li>
        <li>
        It makes tests less readable: the additional assertions in the test make it more difficult to determine which method calls actually affect the state of the world.
        </li>
        <li>
        It gives a false sense of security: just because the code under test called a method does not mean the code under test did the right thing with the method’s return value.
        </li>
        </p>
        
        <p>
        Instead of verifying that they are called, use non-state-changing methods to simulate different conditions in tests, e.g., when(mockUserService.isUserActive(USER)).thenReturn(false). Then write assertions for the return value of the code under test, or verify state-changing method calls.
        </p>
        
        <p>
        Verifying non-state-changing method calls may be useful if there is no other output you can assert. For example, if your code should be caching an RPC result, you can verify that the method that makes the RPC is called only once.
        </p>
        
        <p>
        That’s much simpler! But remember that instead of using a mock to verify that a method was called, it would be even better to use a real or fake object to actually execute the method and check that it works properly. For example, the above test could use a fake database to check that the permission exists in the database rather than just verifying that addPermission() was called.
        <p>
        
        Thats a wrap,
        Geegle Newsletter
      |||,
      "DependsOnPoints": 1,
      "Delay": 4800000
    
    },
    {
      "Sender": "Geegle-Newsletter@geegle.org",
      "Title": "Geegle News 0x04",
      "Body": |||
        Welcome to Geegle Newsletter!
        
        How about a nice game of chess?
        
        ...
        
        Thank you for a very enjoyable game.
        
        Thats a wrap,
        Geegle Newsletter
      |||,
      "DependsOnPoints": 1,
      "Delay": 9600000
    },
  ],
}
