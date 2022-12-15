const sls = require('serverless-http')
const binaryMimeTypes = require('./binaryMimeTypes')
const axios = require('axios')

const nuxt = require('./nuxt')
module.exports.nuxt = sls(nuxt, {
  binary: binaryMimeTypes
})

module.exports.warmup = async () => {
  
  const domainURL = process.env.DOMAIN_URL || "http://localhost:3000"
  const data = await axios.get(domainURL)
  console.log(`Status: ${data.status}`)
  if (data.status === 200) {
    console.log(`${domainURL} warm up` ) 
  } else {
    console.log(`Error Detail: ${data.statusText}`)
    await lineNotification(data)
  }
}

const lineNotification = async (data) => {
  const qs = require('querystring');
  const LINE_URL = 'https://notify-api.line.me';
  const LINE_API_PATH = '/api/notify';
  const config = {
    baseURL: LINE_URL,
    url: LINE_API_PATH,
    method: 'post',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
      'Authorization': `Bearer ${process.env.LINE_TOKEN}`
    },
    data: qs.stringify({
      message: `Error Detail: ${data.statusText}`,
    })
  };

  const response = await axios.request(config);
  console.log(response)
}