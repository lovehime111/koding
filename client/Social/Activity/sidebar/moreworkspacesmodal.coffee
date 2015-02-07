class MoreWorkspacesModal extends SidebarSearchModal

  constructor: (options = {}, data) ->

    options.cssClass         = KD.utils.curry 'more-modal more-workspaces', options.cssClass
    options.width            = 462
    options.title          or= 'Workspaces'
    options.disableSearch    = yes
    options.itemClass      or= ModalWorkspaceItem
    options.bindModalDestroy = no

    super options, data

  viewAppended: ->

    @addButton = new KDButtonView
      title    : "Add Workspace"
      style    : 'add-big-btn'
      loader   : yes
      callback : =>
        @destroy()
        KD.singletons.mainView.activitySidebar.addNewWorkspace @getData()

    @addSubView @addButton, '.kdmodal-content'

    super


  populate: ->

    { workspaces } = @getData()

    @listController.addItem workspace for workspace in workspaces


