const express = require('express');
// const puppeteer = require('puppeteer');
const {Cluster} = require('puppeteer-cluster');
const jwt = require('jsonwebtoken');

const PORT = process.env.PORT || 8080;
const TASKTIMEOUT = process.env.TASKTIMEOUT || 5000;
const MAXCONCURRENTY = process.env.MAXCONCURRENCY || 2;
const app = express();

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

function waitForNetworkIdle(page, timeout, maxInflightRequests = 0) {
  page.on('request', onRequestStarted);
  page.on('requestfinished', onRequestFinished);
  page.on('requestfailed', onRequestFinished);

  let inflight = 0;
  let fulfill;
  let promise = new Promise(x => fulfill = x);
  let timeoutId = setTimeout(onTimeoutDone, timeout);
  return promise;

  function onTimeoutDone() {
    console.log("network idled");
    page.removeListener('request', onRequestStarted);
    page.removeListener('requestfinished', onRequestFinished);
    page.removeListener('requestfailed', onRequestFinished);
    fulfill();
  }

  function onRequestStarted() {
    ++inflight;
    if (inflight > maxInflightRequests)
      clearTimeout(timeoutId);
  }

  function onRequestFinished() {
    if (inflight === 0)
      return;
    --inflight;
    if (inflight === maxInflightRequests)
      timeoutId = setTimeout(onTimeoutDone, timeout);
  }
}

(async () => {
  const cluster = await Cluster.launch({
    concurrency : Cluster.CONCURRENCY_CONTEXT,
    timeout: TASKTIMEOUT,
    maxConcurrency : 5,
    puppeteerOptions : {
      // headless : false,
      // slowMo : 250,
      executablePath : '/chrome-linux/chrome',
      args : [ '--no-sandbox', '--disable-setuid-sandbox' ],
    }
  });
  await cluster.task(async ({page, data}) => {
    console.log(data);
    page.on('console', msg => console.log('PAGE LOG:', msg.text()));
    await page.setExtraHTTPHeaders({'X-Geegle-SubAcc' : data.subacc});
    await page.goto(data.url);
    await Promise.all([
      page.goto(data.url),
      waitForNetworkIdle(page, 2000, 0),
    ]);
    // await sleep(1500);
    console.log("done");
  });

  // setup server
  app.get('/', async function(req, res) {
    let token = req.headers['x-geegle-jwt'];
    var djwt;
    if (token) {
      jwt.verify(token, "superSecretJWTKEY", (err, decoded) => {
        if (err) {
          return res.json({success : false, message : 'Token is not valid'});
        } else {
          djwt = decoded;
        }
      });
    } else {
      return res.json(
          {success : false, message : 'Auth token is not supplied'});
    }

    if (!req.query.url) {
      return res.end('Please specify url like this: ?url=example.com');
    }
    try {
      cluster.queue(
          {url : req.query.url, subacc : djwt['username'].split('@')[0].split('+')[1]});
    } catch (err) {
      return res.json({success : false, message : err.message});
    }

    return res.json({success : true, message : 'queued'});
  });

  app.listen(PORT,
             function() { console.log('xssbot listening on port ' + PORT); });
})();
