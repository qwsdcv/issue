
sap.ui.getCore().attachInit(function () {
    jQuery.sap.includeStyleSheet("/static/css/Custom.css");
    
    let view = 'i.main.Catalog';
    
    let root = jQuery.sap.getUriParameters().get('root');
    if(root === 'true'){
        view = 'i.main.Login';
    }

    let App = sap.ui.xmlview({
        viewName : view
     });

    let shell = new sap.m.Shell({
        app: App
    });

    window.Shell = shell;
    shell.placeAt("content")
});
