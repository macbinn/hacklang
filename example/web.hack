web.route({
  url: `^/$`,

  get: i => {
    print(i.query.name)
    i.resp(200, `text/html`, `hello world`)
  },

  post: i => {
    print(i.jsonBody)
    i.json(true)
  }
})

web.run(`:8083`)