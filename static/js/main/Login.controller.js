sap.ui.define([
    'sap/ui/core/mvc/Controller',
], function (Controller) {
    'use strict';

    var Login = Controller.extend("i.main.Login", {
        onInit: function () {
            this.App = this.byId('root');
        },
        buttonClick: function () {
            let input = this.byId('passwordInput');
            let pwd = input.getValue();
            $.ajax({
                url: '/issues/secret',
                method: 'POST',
                contentType: 'application/json; charset=utf-8',
                data: JSON.stringify({ secret: pwd }),
                success: (msg, status, xhr) => {
                    console.log(msg);
                    let token = xhr.getResponseHeader('Authorization');
                    $.ajaxSetup({
                        headers: {
                            "Authorization": token
                        }
                    });
                    sap.ui.xmlview({
                        viewName: 'i.main.Catalog'
                    }).loaded().then(page => {
                        if (window.Shell) {
                            window.Shell.setApp(page);
                        }
                    });

                }
            });
        }

    });
    return Login;
});