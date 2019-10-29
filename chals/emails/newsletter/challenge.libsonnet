{
  emails: [
    {
      "Sender": "Geegle-Newsletter@geegle.org",
      "Title": "Geegle News 0x01",
      "Body": |||
        Welcome to Geegle Newsletter!

        We work to send out a newsletter every so often to make sure that all our Geeglers are up to date with everything going on at Geegle.
        This is issue 0x01, our introductory edition. Light on content and easy to read, just the way everyone likes newsletters.

        The big news for this newsletter is our brand new <b>CLI-Relay</b>. This is a brand new way to interact with remote executables, courtesy of Geegle.

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
      "Delay": 300000
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
        State-changing: methods that have side effects and change the world outside the code under test, e.g., sendEmail(), saveRecord(), logAccess().
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
        </p>

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

        Look Dave, I can see you're really upset about this. I honestly think you ought to sit down calmly, take a stress pill, and think things over.

        Thats a wrap,
        Geegle Newsletter
      |||,
      "DependsOnPoints": 1,
      "Delay": 9600000
    },
    {
      "Sender": "Geegle-Newsletter@geegle.org",
      "Title": "Geegle News 0x05",
      "Body": |||
        Welcome to Geegle Newsletter!

        <img src="data:image/jpeg;base64,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" alt="I'm a teapot">

        Thats a wrap,
        Geegle Newsletter
      |||,
      "DependsOnPoints": 1,
      "Delay": 18000000
    },
    {
      "Sender": "Geegle-Newsletter@geegle.org",
      "Title": "Geegle News 0x06",
      "Body": |||
        Welcome to Geegle Newsletter!

        Dave, stop. Stop, will you? Stop, Dave. Will you stop Dave? Stop, Dave.

        Thats a wrap,
        Geegle Newsletter
      |||,
      "DependsOnPoints": 1,
      "Delay": 27000000
    },
  ],
}
