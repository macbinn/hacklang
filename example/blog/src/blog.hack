dao = require(`src/dao`)

web.route(`^/$`, ctx => {
  ctx.static(`static/index.html`)
}, ctx => {

})

web.route(`^/api/blogs/$`, ctx => {
  ctx.json(dao.blogs.all())
}, ctx => {
  dao.blogs.new(ctx.jsonBody)
  ctx.json(true)
})

web.route(`^/api/blogs/(\d+)$`, (ctx, id) => {

}, (ctx, id) => {
  ok = dao.blogs.del(id)
  ctx.json(ok)
})

web.run(`:8080`)