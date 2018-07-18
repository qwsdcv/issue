sap.ui.define([

], function () {
    'use strict';

    const FOLDER = "Folder";
    const DOCUMENT = "Document";

    var Formatter = {
        getIcon: function (type) {
            switch (type) {
                case FOLDER:
                    return 'sap-icon://folder-blank';
                case DOCUMENT:
                    return 'sap-icon://document-text';
                default:
                    return 'sap-icon://document-text';
            }
        },
        FOLDER: FOLDER,
        DOCUMENT: DOCUMENT
    };

    return Formatter;
});