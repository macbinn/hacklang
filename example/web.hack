web.route(
  `^/$`,

  i => {
    print(i.query.name)
    i.resp(200, `text/html`, `hello world`)
  },

  i => {
    print(i.jsonBody)
    i.json(true)
  })

web.run(`:8083`)