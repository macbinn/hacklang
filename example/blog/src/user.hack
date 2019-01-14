dao = require(`src/dao`)

secret = `e58c2cee118dffff4f55703b6dbf5f86`

web.route({
  url: `^/api/users/$`,
  post: ctx => {
    dao.users.new({
      name: ctx.jsonBody.name,
      passwordHash: hash.bcrypt.generate(ctx.jsonBody.password)
    })
    ctx.json(true)
  }
})

web.route({
  url: `^/api/users/self$`,
  get: ctx => {
    ctx.user.passwordHash = ``
    ctx.json(ctx.user)
  }
})

web.route({
  url: `^/api/signIn$`,
  post: ctx => {
    user = dao.users.find({
      name: ctx.jsonBody.name
    })
    ok = hash.bcrypt.compare(user.passwordHash, ctx.jsonBody.password)
    if ok {
      user.passwordHash = ``
      ctx.json({
        ticket: hash.ticket.generate(secret, user.name, 1209600),
        user: user
      })
    }
  }
})

web.prepare(ctx => {
  ticket = ctx.header(`x-user-ticket`)
  username = hash.ticket.getData(ticket, secret)
  if username {
    ctx.user = dao.users.find({
      name: username
    })
  }
})