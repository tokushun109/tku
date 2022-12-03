// nuxt.js
const express = require('express')
const app = express()
const { Nuxt } = require('nuxt')
const path = require('path')

// additional header
app.use((_req, res, next) => {
    res.removeHeader('x-powered-by')
    res.header('no-cache', 'Set-Cookie')
    res.header('x-xss-protection', '1; mode=block')
    res.header('x-frame-options', 'DENY')
    res.header('x-content-type-options', 'nosniff')
    next()
})

app.use('/_nuxt', express.static(path.join(__dirname, '.nuxt', 'dist')))
const config = require('../nuxt.config.js')
const nuxt = new Nuxt(config)
app.use(async (req, res, next) => {
    await nuxt.ready()
    nuxt.render(req, res, next)
  })

module.exports = app
