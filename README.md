# GoUI

[GoUI](https://fipress.org/project/goui) is a lightweight cross-platform Go "GUI" framework. It is not really a *GUI* library, instead, it displays things through web view. 
That is, you can write cross-platform applications using JavaScript, HTML and CSS, and Go of cause. 

The basic idea is:
1. Write your app with HTML, JavaScript, and CSS like it is a single page web application.
    You may adopt whatever javascript framework you like, angular, backbone, etc. 
2. Wrap the web application to desktop (or mobile in the future) application with GoUI
3. Plus, write some backend services with Go, then your Javascript code can access them like make an AJAX request
4. Plus plus, you may also write some frontend services with Javascript for your Go code to access   

GoUI has some unique merits compares to other library.

1. It provides two way bindings between Go and Javascript. Both sides can register services for the other side to access.

2. Your application won't get too large because you don't need to include Chromium or any other browser into your package, since it uses Cocoa/WebKit for macOS, MSHTML for Windows and Gtk/WebKit for Linux.

3. It is extremely easy To Use. The api is super simple, and intentionally designed adapts to web developers. Register a service is like registering a REST service and request one is like making a jquery AJAX call.

4. It is powerful, not because itself but because it brings two powerful tools together.
 
    - **The UI** You can use any frontend technologies to make your web page shinning. And conveniently, you can open the Web Inspector to help you. 
    
    - **The Logic** With Javascript + Go, you can do anything.  


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

## Examples
Under the `example` directory, there are two examples now for you to get started with.

### hello
This is the GoUI version of "Hello world". It's pretty much a starting point for you to go, so let's take a look at it in detail. 

The project only contains 3 files.

```
  ├─ui
  | ├─js
  | |  └─goui.js 
  | └─hello.html
  └─hello.go
```

`goui.js` - The library GoUI provides for you.

`hello.go` - provides a service named `hello`, and then create a window.
```
goui.Service("hello", func(context *goui.Context) {
		context.Success("Hello world!")
	})

	//create and open a window
	goui.Create(goui.Settings{Title: "Hello",
		Url:       "./ui/hello.html",
		Left:      20,
		Top:       30,
		Width:     300,
		Height:    200,
		Resizable: true,
		Debug:     true})
```

`hello.html` - has a button on it, if a user clicks the button, it will request the above `hello` service.
```
goui.request({url: "hello",
              success: function(data) {
                  document.getElementById("result").innerText = data;
             }});
```

### chat
This is an example of a fake chatbot. Please check the code out and run it if you are interested.

