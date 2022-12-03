// nuxt.js
const express = require('express')
const app = express()
const { Nuxt } = require('nuxt')
const path = require('path')

app.use('/_nuxt', express.static(path.join(__dirname, '.nuxt', 'dist')))
const config = require('../nuxt.config.js')
const nuxt = new Nuxt(config)
app.use(async (req, res, next) => {
    await nuxt.ready()
    nuxt.render(req, res, next)
  })

module.exports = app
