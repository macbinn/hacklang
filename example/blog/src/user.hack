dao = require(`src/dao`)

secret = `e58c2cee118dffff4f55703b6dbf5f86`

web.route(`^/api/users/$`, ctx => {

}, ctx => {
  user = {}
  user.name = ctx.jsonBody.name
  user.passwordHash = hash.bcrypt.generate(ctx.jsonBody.password)
  dao.users.new(user)
  ctx.json(true)
})

web.route(`^/api/users/self$`, ctx => {
  ctx.user.passwordHash = ``
  ctx.json(ctx.user)
}, ctx => {

})

web.route(`^/api/signIn$`, ctx => {

}, ctx => {
  query = {}
  query.name = ctx.jsonBody.name
  user = dao.users.find(query)
  ok = hash.bcrypt.compare(user.passwordHash, ctx.jsonBody.password)
  if ok {
    resp = {}
    resp.ticket = hash.ticket.generate(secret, user.name, 1209600)
    user.passwordHash = ``
    resp.user = user
    ctx.json(resp)
  }
})

web.prepare(ctx => {
  ticket = ctx.header(`x-user-ticket`)
  username = hash.ticket.getData(ticket, secret)
  if username {
    query = {}
    query.name = username
    ctx.user = dao.users.find(query)
  }
})