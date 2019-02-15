# GoUI

[GoUI]() is a lightweight cross-platform Go "GUI" framework. It is not really a *GUI* library, instead, it displays through web view. 
That is, you can write cross-platform applications using JavaScript, HTML and CSS, and Go of cause. 

It uses Cocoa/WebKit for macOS, MSHTML for Windows and Gtk/WebKit for Linux. So you don't need to include Chromium or any other browser into your package.

GoUI provides two way bindings between Go and Javascript. Both sides can register services for the other side to access. 

The basic idea is:
1. Write your app with JavaScript, HTML and CSS like it is a single page web application
    You may whatever framework you like, angular, backbone, etc. 
2. Wrap the web application to desktop or mobile application with GoUI
3. Plus, write some backend services with Go, then your Javascript code can access them like make an AJAX request
4. Plus plus, you may also write some frontend services with Javascript for your Go code to access   

## Usage
### The backend - Go
1. Install the package by `go get`
```
go get github.com/fipress/GoUI
``` 

2. import the package
```
import (
   	"github.com/fipress/GoUI"
   )
```

3. Register services with `goui.Service`, for Javascript to access
```
goui.Service("hello", func(context *goui.Context) {
	context.Success("Hello world!")
})
```

4. Create window with `goui.Create`
```
goui.Create(goui.Settings{"Hello","./ui/test.html",20,30,800,400,true,true})
```

5. Invoke frontend service by `goui.RequestJSService`
```
goui.RequestJSService(goui.JSServiceOptions{
		Url: "chat/"+msg,
	})
```

### The frontend - Javascript
1. add `goui.js` in your web page

2. Register services with `goui.service`
```
goui.service("chat/:msg",function(msg) {
    $box.append('<div class="bot"><div><b class="avatar">B</b></div><div><div class="msg">' + msg + '</div></div</div>')
 })
```

3. request backend services
```
goui.request({url: "hello",
            success: function(data) {
                        document.getElementById("result").innerText = data;
             }});
```
