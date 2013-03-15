log = -> logger.info arguments...

{argv} = require 'optimist'

{exec} = require 'child_process'
{extend} = require 'underscore'

process.on 'uncaughtException', (err)->
  exec './beep'
  console.log err, err?.stack

Bongo = require 'bongo'
Broker = require 'broker'

KONFIG = require('koding-config-manager').load("main.#{argv.c}")
Object.defineProperty global, 'KONFIG', value: KONFIG
{mq, mongo, email, social} = KONFIG

mongo += '?auto_reconnect'

mqOptions = extend {}, mq
mqOptions.login = social.login if social?.login?

broker = new Broker mqOptions

processMonitor = (require 'processes-monitor').start
  name : "Social Worker #{process.pid}"
  stats_id: "worker.social." + process.pid
  interval : 30000
  librato: KONFIG.librato
  limit_hard  :
    memory   : 300
    callback : (name,msg,details)->
      console.log "[SOCIAL WORKER #{name}] Using excessive memory, exiting."
      process.exit()
  die :
    after: "non-overlapping, random, 3 digits prime-number of minutes"
    middleware : (name,callback) -> koding.disconnect callback
    middlewareTimeout : 15000
  # WE'RE NOT SURE IF THIS SOFT LIMIT WAS A GOOD IDEA OR NOT
  # limit_soft:
  #   memory: 200
  #   callback: (name, msg, details) ->
  #     console.log "[SOCIAL WORKER #{name}] Using too much memory, accepting no more new jobs."
  #     process.send?({pid: process.pid, exiting: yes})
  #     koding.disconnect()
  #     setTimeout ->
  #       process.exit()
  #     , 20000
  # DISABLED TO TEST MEMORY LEAKS
  # die :
  #   after: "non-overlapping, random, 3 digits prime-number of minutes"
  #   middleware : (name,callback) -> koding.disconnect callback
  #   # TEST AMQP WITH THIS CODE. IT THROWS THE CHANNEL ERROR.
  #   # middleware : (name,callback) ->
  #   #   koding.disconnect ->
  #   #     console.log "[SOCIAL WORKER #{name}] is reached end of its life, will die in 10 secs."
  #   #     setTimeout ->
  #   #       callback null
  #   #     ,10*1000
  #   middlewareTimeout : 15000
  # toobusy:
  #   interval: 10000
  #   callback: ->
  #     console.log "[SOCIAL WORKER #{name}] I'm too busy, accepting no more new jobs."
  #     process.send?({pid: process.pid, exiting: yes})
  #     koding.disconnect()
  #     setTimeout ->
  #       process.exit()
  #      , 20000

koding = new Bongo
  root        : __dirname
  mongo       : mongo
  models      : './models'
  resourceName: social.queueName
  mq          : broker
  fetchClient :(sessionToken, context, callback)->
    # console.log {'fetchClient', sessionToken, context, callback}
    [callback, context] = [context, callback] unless callback
    context             ?= group: 'koding'
    callback            ?= ->
    koding.models.JUser.authenticateClient sessionToken, context, (err, account)->
      if err
        koding.emit 'error', err
      else
        callback {sessionToken, context, connection:delegate:account}

koding.on 'authenticateUser', (client, callback)->
  {delegate} = client.connection
  callback delegate

# koding.on 'auth', (exchange, sessionToken)->
#   koding.fetchClient sessionToken, (client)->
#     {delegate} = client.connection

#     # if delegate instanceof koding.models.JAccount
#     #   koding.models.JAccount.emit "AccountAuthenticated", delegate

#     koding.handleResponse exchange, 'changeLoggedInState', [delegate]
koding.connect ->

  # create default roles for groups
  JGroupRole = require './models/group/role'

  JGroupRole.createDefaultRoles (err)->
    if err then console.log err.message
    else console.log "Default group roles created!"

  if KONFIG.misc?.claimGlobalNamesForUsers
    require('./models/account').reserveNames console.log

  if KONFIG.misc?.updateAllSlugs
    require('./traits/slugifiable').updateSlugsByBatch 100, [
      require './models/tag'
      require './models/app'
      require './models/messages/codesnip'
      require './models/messages/discussion'
      require './models/messages/tutorial'
    ]

  if KONFIG.misc?.debugConnectionErrors then
    # console.log 'ffaafafafaf'
    # TEST AMQP WITH THIS CODE. IT THROWS THE CHANNEL ERROR.
    # koding.disconnect ->
    #   console.log "[SOCIAL WORKER #{name}] is reached end of its life, will die in 10 secs."
    #   setTimeout ->
    #     process.exit()
    #   ,10*1000

console.log "Koding Social Worker #{process.pid} has started."
