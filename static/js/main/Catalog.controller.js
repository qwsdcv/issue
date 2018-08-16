sap.ui.define([
    'jquery.sap.global',
    'sap/m/MessageToast',
    'sap/ui/core/mvc/Controller',
    'sap/ui/model/json/JSONModel',
    'i/pkg/Formatter'
], function (jQuery, MessageToast, Controller, JSONModel, Formatter) {
    "use strict";


    var Catalog = Controller.extend("i.main.Catalog", {

        Formatter: Formatter,

        onInit: function () {

            this.Split = this.byId("Split");
            this.Tree = this.byId("catalog");
            this.Tree.setMode("SingleSelectMaster");
            this.Spacer = this.byId("Spacer");
            this.PopInput = null;
            this.TextArea = this.byId("TypeHere");
            this.HTML = this.byId("PreviewHere");

            this.JsonModel = new JSONModel();
            this.getView().setModel(this.JsonModel);

            this.NewType = Formatter.FOLDER;

            this.initPopup();

            this.loadMenu();


            this.Converter = new showdown.Converter();
            this.Converter.setOption('tables', true);
            this.Converter.setOption('tasklists', true);
            this.Converter.setOption('emoji', true);
            this.Converter.setOption('underline', true);
        },
        getParameterByName: function (name) {
            var match = RegExp('[?&]' + name + '=([^&]*)').exec(window.location.href);
            return match && decodeURIComponent(match[1].replace(/\+/g, ' '));
        },

        isEditMode() {
            try {
                let token = $.ajaxSetup()['headers']['Authorization'];
                return token != null;
            } catch (error) {
                return false;
            }
        },

        newDocument: function (event) {
            jQuery.sap.log.info("newDocument");
            let openTarget = this.Spacer;
            if (this.CurrentSelected) {
                openTarget = this.CurrentSelected;
            }
            this.NewType = Formatter.DOCUMENT;
            this.PopInput.openBy(openTarget);
        },

        newFolder: function (event) {
            jQuery.sap.log.info("newFolder");
            let openTarget = this.Spacer;
            if (this.CurrentSelected) {
                openTarget = this.CurrentSelected;
            }
            this.NewType = Formatter.FOLDER;
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
                    if (obj.type == Formatter.FOLDER) {
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
                    this.Tree.setBusy(false);
                },
                success: (json) => {
                    that.getView().getModel().setData({ catalogs: json }, true)
                    if (json.length > 0) {
                        let initId = json[0].id;
                        that.getContent(initId);
                    }
                    this.Tree.setBusy(false);
                }
            });
        },

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
                    this.Tree.setBusy(false);
                },
                data: JSON.stringify(data),
                success: (json) => {
                    that.getView().getModel().setData({ catalogs: json }, true);
                    this.Tree.setBusy(false);
                }
            });
        },

        getComment: function () {
            let that = this;
            let commentStuff = this.byId('commentStuff');

            $.ajax({
                url: '/issues/comment/' + this.CurrentContentBinding.id,
                method: 'GET',
                dataType: 'json',
                contentType: 'application/json; charset=utf-8',
                error: (jqXHR, textStatus, errorThrown) => {
                    MessageToast.show(textStatus);
                    commentStuff.setBusy(false);
                },
                success: (json) => {
                    that.getView().getModel().setProperty('/comments', json);
                    commentStuff.setBusy(false);
                }
            });
        },

        addComment: function () {
            let commentStuff = this.byId('commentStuff');
            let commentText = this.byId('TypeComment').getValue();
            this.byId('TypeComment').setValue('');
            if (commentText && commentText != '' && commentText.trim() != '') {
                commentStuff.setBusy(true);
                let data = {
                    parent_id: this.CurrentContentBinding.id,
                    content: commentText
                };
                let that = this;
                $.ajax({
                    url: '/issues/comment' /*+ this.CurrentContentBinding.id*/,
                    method: 'POST',
                    dataType: 'json',
                    contentType: 'application/json; charset=utf-8',
                    error: (jqXHR, textStatus, errorThrown) => {
                        MessageToast.show(textStatus);
                        commentStuff.setBusy(false);
                    },
                    data: JSON.stringify(data),
                    success: (json) => {
                        that.getView().getModel().setProperty('/comments', json);
                        commentStuff.setBusy(false);
                    }
                });
            } else {
                MessageToast.show('不要评论空');
            }
        },

        go2Detail: function () {
            if (this.Split) {
                this.Split.toDetail(this.createId("detail"));
            }
        },
        selectionChange: function (oEvent) {
            this.go2Detail();
            var iItem = oEvent.getParameter("listItem");
            this.CurrentSelected = iItem/*.getBindingContextPath()*/;
            let currentPath = this.CurrentSelected.getBindingContextPath();
            let obj = this.JsonModel.getObject(currentPath);


            this.getContent(obj.id);
        },

        getContent: function (id) {
            let detailPage = this.byId("detail");
            detailPage.setBusy(true);
            let that = this;
            $.ajax({
                url: 'issues/content/' + id,
                method: 'GET',
                dataType: 'json',
                contentType: 'application/json; charset=utf-8',
                error: (jqXHR, textStatus, errorThrown) => {
                    MessageToast.show(textStatus);
                    detailPage.setBusy(false);
                },
                success: (json) => {
                    that.CurrentContentBinding = json;
                    this.getComment(id);
                    let text = json.content;
                    that.TextArea.setValue(text);
                    that.TextArea.fireLiveChange({ value: text });
                    detailPage.setBusy(false);
                }
            });
        },

        setContent: function () {

            if (this.TextAreaChange && this.CurrentContentBinding && this.CurrentContentBinding.id) {
                let text = this.TextArea.getValue();
                this.CurrentContentBinding.content = text;
                $.ajax({
                    url: 'issues/content/' + this.CurrentContentBinding.id,
                    method: 'POST',
                    dataType: 'json',
                    data: JSON.stringify(this.CurrentContentBinding),
                    contentType: 'application/json; charset=utf-8',
                    error: (jqXHR, textStatus, errorThrown) => {
                        MessageToast.show(errorThrown);
                    },
                    success: (json) => {
                        MessageToast.show("Submited.");
                    }
                });
            }
        },

        go2Master: function () {
            if (this.Split) {
                this.Split.to(this.createId("master"));
            }
        },

        getNickName: function (nick, ip) {
            if (nick && nick.trim()) {
                return nick;
            } else {
                return ip;
            }
        },
        getMarkDown: function (content) {
            let converter = this.Converter;
            let html = converter.makeHtml(content);
            let finalHtml = `<div>${html}</div>`;
            return finalHtml;
        },

        handleLiveChange: function (oEvent) {
            let sValue = oEvent.getParameter("value");

            let converter = this.Converter;
            let html = converter.makeHtml(sValue);
            let finalHtml = `<div>${html}</div>`;

            this.HTML.setContent(finalHtml);

            this.TextAreaChange = true;
        }

    });


    return Catalog;

});
