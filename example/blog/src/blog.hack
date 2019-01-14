dao = require(`src/dao`)

web.route(`^/$`, ctx => {
  ctx.static(`static/index.html`)
}, ctx => {

})

web.route(`^/api/blogs/$`, ctx => {
  ctx.json(dao.blogs.all())
}, ctx => {
  if ctx.user {
    ctx.jsonBody.authorId = ctx.user.id
    dao.blogs.new(ctx.jsonBody)
    ctx.json(true)
  }
})

web.route(`^/api/blogs/(\d+)$`, (ctx, id) => {

}, (ctx, id) => {
  if ctx.user {
    ok = dao.blogs.del(id)
    ctx.json(ok)
  }
})

require(`src/user`)

web.run(`:8080`)