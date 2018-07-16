sap.ui.define([
    'jquery.sap.global',
    'sap/m/MessageToast',
    'sap/ui/core/mvc/Controller',
    'sap/ui/model/json/JSONModel'
], function (jQuery, MessageToast, Controller, JSONModel) {
    "use strict";


    const FOLDER = "Folder";
    const DOCUMENT = "Document";

    var Catalog = Controller.extend("i.main.Catalog", {

        onInit: function () {
            jQuery.sap.includeStyleSheet("/static/css/Custom.css");

            this.Split = this.byId("Split");
            this.Tree = this.byId("catalog");
            this.Tree.setMode("SingleSelectMaster");
            this.Spacer = this.byId("Spacer");
            this.PopInput = null;

            this.JsonModel = new JSONModel();
            this.getView().setModel(this.JsonModel);

            this.CurrentSelected = null;
            this.NewType = FOLDER;

            this.initPopup();

            this.loadMenu();
        },

        getIcon(type) {
            switch (type) {
                case FOLDER:
                    return 'sap-icon://folder-blank';
                case DOCUMENT:
                    return 'sap-icon://document-text';
                default:
                    return 'sap-icon://document-text';
            }
        },

        newDocument: function (event) {
            jQuery.sap.log.info("newDocument");
            let openTarget = this.Spacer;
            if (this.CurrentSelected) {
                openTarget = this.CurrentSelected;
            }
            this.NewType = DOCUMENT;
            this.PopInput.openBy(openTarget);
        },

        newFolder: function (event) {
            jQuery.sap.log.info("newFolder");
            let openTarget = this.Spacer;
            if (this.CurrentSelected) {
                openTarget = this.CurrentSelected;
            }
            this.NewType = FOLDER;
            this.PopInput.openBy(openTarget);
        },

        initPopup: function () {
            if (!this.PopInput) {
                this.PopInput = sap.ui.xmlfragment("i.main.NewMenu", this);
            }
            this.getView().addDependent(this.PopInput);
            let that = this;

            let input = sap.ui.getCore().byId("NewMenuItem");
            input.onsapenter = (e) => {
                let text = input.getValue();
                that.PopInput.close();
                input.setValue('');
                jQuery.sap.log.info(text);

                let pId = null;
                if (that.CurrentSelected) {
                    let currentPath = that.CurrentSelected.getBindingContextPath();
                    let obj = that.JsonModel.getObject(currentPath);
                    if (obj.type == FOLDER) {
                        pId = obj.id;
                    } else {
                        pId = obj.parent_id;
                    }

                }

                that.postNew({
                    title: text,
                    parent_id: pId,
                    type: that.NewType
                });
            };
        },
        loadMenu: function () {
            this.Tree.setBusy(true);
            let that = this
            $.ajax({
                url: '/issues/menu',
                method: 'GET',
                dataType: 'json',
                contentType: 'application/json; charset=utf-8',
                error: (jqXHR, textStatus, errorThrown) => {
                    MessageToast.show(textStatus);
                },
                success: (json) => {
                    that.recursiveSetIcon(json);
                    that.getView().getModel().setData(json)
                    this.Tree.setBusy(false);
                }
            });
        },
        recursiveSetIcon: function (data) {
            data.forEach((one) => {
                one.icon = this.getIcon(one.type);
                if (one.nodes) {
                    this.recursiveSetIcon(one.nodes);
                }
            });
        },

        /*addItems:function (jsonArray) {
            let that =this;
            jsonArray.forEach((one,index)=>{
                let item = new sap.m.StandardTreeItem({
                    icon:that.getIcon(one.type),
                    title:one.title
                });
                that.Tree.addItem(item);
            });
        },*/

        postNew: function (data) {
            this.Tree.setBusy(true);
            let that = this
            $.ajax({
                url: '/issues/menu',
                method: 'POST',
                dataType: 'json',
                contentType: 'application/json; charset=utf-8',
                error: (jqXHR, textStatus, errorThrown) => {
                    MessageToast.show(textStatus);
                },
                data: JSON.stringify(data),
                success: (json) => {
                    that.recursiveSetIcon(json);
                    that.getView().getModel().setData(json)
                    this.Tree.setBusy(false);
                }
            });
        },
        selectionChange: function (oEvent) {
            var iItem = oEvent.getParameter("listItem");
            this.CurrentSelected = iItem/*.getBindingContextPath()*/;
        }

    });


    return Catalog;

});
