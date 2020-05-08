window.goui = (function() {
    let obj = {};
    let invokeBackend;

    let osConsts = {
        linux:{pathSeparator:'/',pathListSeparator:':'},
        windows:{pathSeparator:'\\',pathListSeparator: ';'}
    };

    let Context = {
        create: function (options) {
            let obj = {};

            obj.error = function (msg) {
                if(options.error) {
                    let data = {url:options.error,data:msg};
                    invokeBackend(data);
                }
            };

            obj.success = function(data) {
                if(options.success) {
                    data = {url:options.success,data:data};
                    invokeBackend(data);
                }
            };

            return obj;
        }
    };

    let Router = {
        create: function () {
            let obj = {};

            let parsedRoutes = [];

            let optionalParam = /\((.*?)\)/g;
            let namedParam = /(\(\?)?:\w+/g;
            let splatParam = /\*\w+/g;
            let escapeRegExp = /[\-{}\[\]+?.,\\\^$|#\s]/g;

            let routeStripper = /^[#\/]|\s+$/g;

            let pathToRegExp = function (path) {
                path = path.replace(escapeRegExp, '\\$&')
                    .replace(optionalParam, '(?:$1)?')
                    .replace(namedParam, function (match, optional) {
                        return optional ? match : '([^/?]+)';
                    })
                    .replace(splatParam, '([^?]*?)');
                return new RegExp('^' + path + '(?:\\?([\\s\\S]*))?$');
            };

            obj.parse = function (path, handler) {
                var route = {};
                route.regexp = pathToRegExp(path);
                route.handler = handler;
                parsedRoutes.push(route);
            };

            obj.dispatch = function (url) {
                var matched;
                parsedRoutes.some(function (route) {
                    var args = url.match(route.regexp);
                    if (args) {
                        route.args = args.slice(1);
                        matched = route;
                        return true;
                    }
                });
                return matched;
            };

            return obj;
        }
    };

    if (window.webkit) {
        obj.os = osConsts.linux ;
        invokeBackend = function(data) {
            window.webkit.messageHandlers.goui.postMessage(JSON.stringify(data));
        };
    } else if (window.gouiAndroid){
        obj.os = osConsts.linux ;
        invokeBackend = function(data) {
            window.gouiAndroid.handleMessage(JSON.stringify(data));
        };
    } else if(window.external) {
        obj.os = osConsts.windows ;
        invokeBackend = function(data) {
            window.external.notify(JSON.stringify(data));
        };
    }

    let router = Router.create();

    let seq = 0;

    let getName = function () {
        let name = "f" + seq;
        seq++;
        return name;
    };

    //Request is to send request to the backend
    //
    //options: {
    //      url:"service function",
    //      data: data,
    //      dataType: "json, text, html or xml" ?
    //      context: callback context
    //      success: callback function on success,
    //      error: callback function on erroe
    // }

    obj.request = function (options) {
        var successName, errorName;

        if (options.success) {
            successName = getName();
            obj[successName] = function (data) {
                options.success.call(options.context, data);
                delete obj[successName];
                if (errorName) {
                    delete obj[errorName];
                }
            };
        }

        if (options.error) {
            errorName = getName();
            obj[errorName] = function (err) {
                options.error.call(options.context, err);
                delete obj[errorName];
                if (successName) {
                    delete obj[successName];
                }
            };
        }

        var req = { url: options.url};
        if(options.data) {
            req.data = JSON.stringify(options.data);
        }
        if(successName) {
            req.success = "goui." + successName;
        }
        if(errorName) {
            req.error = "goui." + errorName;
        }

        invokeBackend(req);
    };

    // service is to register a frontend service the backend can request
    obj.service = function (path,handler) {
        router.parse(path,handler);
    };

    //options: {
    //      url,
    //      data,
    //      success,
    //      error
    // }
    obj.handleRequest = function (options) {
        var ctx = Context.create(options);

        if (!options || !options.url) {
            ctx.error("Invalid request.");
            return;
        }

        var route = router.dispatch(options.url);
        if (route) {
            route.args.push(ctx);
            route.handler.apply(null, route.args);
        } else {
            ctx.error("Service not found: ", options.url)
        }
    };

    obj.escapeRegExp = function(text) {
        return text.replace(/[-[\]{}()*+?.,\\^$|#\s]/g, '\\$&');
    };


    return obj;
})();
