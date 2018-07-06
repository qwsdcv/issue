sap.ui.define([
    'jquery.sap.global',
    'sap/m/MessageToast',
    'sap/ui/core/mvc/Controller',
    'sap/ui/model/json/JSONModel'
], function (jQuery, MessageToast, Controller, JSONModel) {
    "use strict";

    var Catalog = Controller.extend("i.main.Catalog", {

        onInit: function () {
            this.Split = this.byId("Split");
            this.Tree = this.byId("catalog");
            this.PopInput = null;

            this.CurrentPath = null;

            this.initPopup();
        },

        getIcon(type) {
            switch (type) {
                case 'folder':
                    return 'sap-icon://folder-blank';
                case 'document':
                    return 'sap-icon://document-text';
                default:
                    return 'sap-icon://document-text';
            }
        },

        newDocument: function (event) {
            jQuery.sap.log.info("newDocument");
        },

        newFolder: function (event) {
            jQuery.sap.log.info("newFolder");



            this.PopInput.openBy(this.Tree);
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

                that.postNew({
                    title:text,
                    parent_id:null
                });
            };
        },

        postNew: function (data) {
            $.ajax({
                url: '/issues/menu',
                method: 'POST',
                dataType: 'json',
                error: (jqXHR, textStatus, errorThrown) => {
                    MessageToast.toast(textStatus);
                },
                data: data,
                success: () => { }
            });
        }

    });


    return Catalog;

});
