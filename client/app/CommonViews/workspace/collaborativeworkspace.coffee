class CollaborativeWorkspace extends Workspace

  constructor: (options = {}, data) ->

    super options, data

    @sessionData       = []
    @firepadRef        = new Firebase "https://hulogggg.firebaseIO.com/"
    @sessionKey        = options.sessionKey or @createSessionKey()
    @workspaceRef      = @firepadRef.child @sessionKey

    @workspaceRef.once "value", (snapshot) =>
      if @getOptions().sessionKey
        log "user wants to join a session"
        if snapshot.val() is null
          log "session is not active"
          @showNotActiveView()
          return false

        log "session is valid, trying to recover"

      log "everything is something happened", "value", snapshot.val(), snapshot.name()

      keys = snapshot.val()?.keys

      @workspaceRef.onDisconnect().remove()  if @amIHost()

      if keys # if we have keys this means we're about to join an old session
        log "it's an old session, impressed!"
        @sessionData  = keys
        @createPanel()
        @workspaceRef.child("users").push KD.nick()
      else
        log "your awesome new session, saving keys now"
        @createPanel()
        @workspaceRef.set "keys": @sessionData
        @workspaceRef.child("users").push KD.nick()

    @workspaceRef.on "child_added", (snapshot) =>
      log "everything is something happened", "child_added", snapshot.val(), snapshot.name()

    @workspaceRef.on "child_removed", (snapshot) =>
      log "possible disconnection occured"

      @showDisconnectedModal()  unless @disconnectedModal

    @on "NewPanelAdded", (panel) ->
      log "New panel created", panel

    @on "AllPanesAddedToPanel", (panel, panes) ->
      paneSessionKeys = []
      paneSessionKeys.push pane.sessionKey for pane in panes
      @sessionData.push paneSessionKeys

  createPanel: (callback = noop) ->
    panelOptions             = @getOptions().panels[@lastCreatedPanelIndex]
    panelOptions.delegate    = @
    panelOptions.sessionKeys = @sessionData[@lastCreatedPanelIndex]  if @sessionData
    newPanel                 = new CollaborativePanel panelOptions

    log "instantiated a panel with these session keys", panelOptions.sessionKeys

    @container.addSubView newPanel
    @panels.push newPanel
    @activePanel = newPanel

    callback()
    @emit "PanelCreated"

  createSessionKey: ->
    nick = KD.nick()
    u    = KD.utils
    return  "#{nick}:#{u.generatePassword(4)}:#{u.getRandomNumber(100)}"

  ready: -> # have to override for collaborative workspace

  amIHost: ->
    [sessionOwner] = @sessionKey.split ":"
    return sessionOwner == KD.nick()

  showNotActiveView: ->
    notValid = new KDView
      cssClass : "not-valid"
      partial  : "This session is not valid or no longer available."

    notValid.addSubView new KDView
      cssClass : "description"
      partial  : """
        If there is nothing wrong with our servers, this usually means,
        the person who is hosting this session is disconnected or closed the session.
      """

    notValid.addSubView new KDButtonView
      cssClass : "cupid-green"
      title    : "Start New Session"
      callback : @bound "startNewSession"

    @container.addSubView notValid

  startNewSession: ->
    @destroySubViews()
    options = @getOptions()
    delete options.sessionKey
    @addSubView new CollaborativeWorkspace options

  joinSession: (sessionKey) ->
    {parent}           = @
    options            = @getOptions()
    options.sessionKey = sessionKey
    @destroy()
    parent.addSubView new CollaborativeWorkspace options
    log "user joined a new session:", sessionKey

  showDisconnectedModal: ->
    if @amIHost()
      title   = "Disconnected from remote"
      content = "It seems, you have been disconnected from Firebase server. You cannot continue this session."
    else
      title   = "Host disconnected"
      content = "It seems, host is disconnected from Firebase server. You cannot continue this session."

    @disconnectedModal = new KDModalView
      title        : title
      content      : "<p>#{content}</p>"
      cssClass     : "host-disconnected-modal"
      overlay      : yes
      buttons      :
        Start      :
          title    : "Start New Session"
          callback : =>
            @disconnectedModal.destroy()
            delete @disconnectedModal
            @startNewSession()
        Join       :
          title    : "Join Another Session"
          callback : =>
            @disconnectedModal.destroy()
            delete @disconnectedModal
            @showSessionModal (modal) ->
              modal.modalTabs.showPaneByIndex(1)
        Exit       :
          title    : "Exit App"
          cssClass : "modal-cancel"
          callback : =>
            @disconnectedModal.destroy()
            delete @disconnectedModal
            appManager = KD.getSingleton("appManager")
            appManager.quit appManager.frontApp

  showJoinModal: (callback = noop) ->
    modal                       = new KDModalViewWithForms
      title                     : "Join New Session"
      content                   : ""
      overlay                   : yes
      cssClass                  : "workspace-modal"
      width                     : 500
      tabs                      :
        forms                   :
          "Join A Session"      :
            fields              :
              Label             :
                itemClass       : KDCustomHTMLView
                tagName         : "p"
                partial         : "Paste the session ID that you want to join and start collaborating."
              SessionInput      :
                itemClass       : KDHitEnterInputView
                type            : "text"
                callback        : =>
                  sessionKey = modal.modalTabs.forms["Join A Session"].inputs.SessionInput.getValue()
                  @joinSession sessionKey
                  modal.destroy()
            buttons             :
              Join              :
                title           : "Join Session"
                cssClass        : "modal-clean-green"
                callback        : =>
                  sessionKey = modal.modalTabs.forms["Join A Session"].inputs.SessionInput.getValue()
                  @joinSession sessionKey
                  modal.destroy()
              Close             :
                title           : "Close"
                cssClass        : "modal-cancel"
                callback        : -> modal.destroy()

    callback modal

  showShareModal: (callback = noop) ->
    modal                       = new KDModalViewWithForms
      title                     : "Share Your Session"
      content                   : ""
      overlay                   : yes
      cssClass                  : "workspace-modal"
      width                     : 500
      tabs                      :
        forms                   :
          "Share This Session"  :
            fields              :
              Label             :
                itemClass       : KDCustomHTMLView
                tagName         : "p"
                partial         : "Share the session ID with your friends to hang out together."
              SessionKey        :
                itemClass       : KDCustomHTMLView
                partial         : @sessionKey
                tagName         : "div"
                cssClass        : "key"
            buttons             :
              "Ok"              :
                title           : "Got It"
                cssClass        : "modal-clean-green"
                callback        : -> modal.destroy()

    callback modal
