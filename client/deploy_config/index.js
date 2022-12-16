const sls = require('serverless-http')
const axios = require('axios')
const binaryMimeTypes = require('./binaryMimeTypes')

const nuxt = require('./nuxt')
module.exports.nuxt = sls(nuxt, {
    binary: binaryMimeTypes,
})

module.exports.warmup = async () => {
    const domainURL = process.env.DOMAIN_URL || 'http://localhost:3000'
    const data = await axios.get(domainURL)
    // eslint-disable-next-line no-console
    console.log(`Status: ${data.status}`)
    if (data.status === 200) {
        // eslint-disable-next-line no-console
        console.log(`${domainURL} warm up`)
    } else {
        // eslint-disable-next-line no-console
        console.log(`Error Detail: ${data.statusText}`)
        await lineNotification(data)
    }
}

const lineNotification = async (data) => {
    const qs = require('querystring')
    const LINE_URL = 'https://notify-api.line.me'
    const LINE_API_PATH = '/api/notify'
    const config = {
        baseURL: LINE_URL,
        url: LINE_API_PATH,
        method: 'post',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            Authorization: `Bearer ${process.env.LINE_TOKEN}`,
        },
        data: qs.stringify({
            message: `Error Detail: ${data.statusText}`,
        }),
    }

    const response = await axios.request(config)
    // eslint-disable-next-line no-console
    console.log(response)
}
