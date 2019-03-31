# GoUI

[GoUI](https://fipress.org/project/goui) is a lightweight cross-platform Go "GUI" framework. It is not really a *GUI* library, instead, it displays things through web view. 
That is, you can write cross-platform applications using JavaScript, HTML and CSS, and Go of cause. 

The basic idea is:
1. Write your app with HTML, JavaScript, and CSS like it is a single page web application.
    You may adopt whatever javascript framework you like, angular, backbone, etc. 
2. Wrap the web application to desktop or mobile application.
3. Plus, write some backend services with Go, then your Javascript code can access them like make an AJAX request
4. Plus plus, you may also write some frontend services with Javascript for your Go code to access.   

That's how GoUI came from. And now, it has some unique merits compares to other library.

1. It provides two way bindings between Go and Javascript. Both sides can register services for the other side to access.

2. Your application won't get too large because you don't need to include Chromium or any other browser into your package, since it uses Cocoa/WebKit for macOS, MSHTML for Windows and Gtk/WebKit for Linux.  
    The macOS package "hello.app" of example "hello" only takes **2.6MB**. 

3. It is extremely easy To Use. The api is super simple, and intentionally designed adapts to web developers. Register a service is like registering a REST service and request one is like making a jquery AJAX call.

4. It is powerful, not because itself but because it brings two powerful tools together.
 
    - **The UI** You can use any frontend technologies to make your web page shinning. And conveniently, you can open the Web Inspector to help you. 
    
    - **The Logic** With Javascript + Go, you can do anything.  


## Basic Usage
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
goui.Create(goui.Settings{Title: "Hello",
		Left:      200,
		Top:       50,
		Width:     400,
		Height:    510,
		Resizable: true,
		Debug:     true})
```

### The frontend - Javascript
1. add `goui.js` in your web page

2. request backend services
```
goui.request({url: "hello",
            success: function(data) {
                        document.getElementById("result").innerText = data;
             }});
```

### The backend can access frontend services

1. The frontend register services with `goui.service`
   ```
   goui.service("chat/:msg",function() {
        console.log(msg);
   })
   ```  

2. The backend invoke frontend service by `goui.RequestJSService` 
```
goui.RequestJSService(goui.JSServiceOptions{
		Url: "chat/"+msg,
	})
```

### Menu
It is extremely easy to create menu with GoUI: just define your menuDefs, and create window with them,

```
menuDefs = []goui.MenuDef{{Title: "File", Type: goui.Container, Children: []goui.MenuDef{
			{Title: "New", Type: goui.Custom, Action: "new", Handler: func() {
				println("new file")
			}},
			{Title: "Open", Type: goui.Custom, Action: "open"},
			{Type: goui.Separator},
			{Title: "Quit", Type: goui.Standard, Action: "quit"},
			...
			}
		}
		
goui.CreateWithMenu(settings,menuDefs)
``` 

For a complete demonstration, please check out the example "texteditor".

## Debugging and packaging
### Debugging
I personally recommend [GoLand](https://www.jetbrains.com/go) to debug Go code, or IntelliJ IDEA which contains GoLand. You may need to set the output directory to your working folder so that your debugging binary can read the web folder. Or other solution to make sure the binary can read the page you provided.

To debug javascript code or check web page elements, you should set `Debug` settings to `true`, then you can open the inspector from the context menu.

![Inspector](https://github.com/FIPress/GoUI/raw/master/screenshots/debug-web.png)

### Packaging
The easiest way to package GoUI applications would be through [GoUI-CLI](https://github.com/FIPress/GoUI-CLI), which will packaging native applications for all the supported platforms, macOS, Ubuntu, Windows, iOS and Android.

## Examples
Under the `example` directory, there are some examples for you to get started with.

### hello
This is the GoUI version of "Hello world". It's pretty much a starting point for you to go, so let's take a look at it in detail. 

The project only contains 3 files.

```
  ├─web
  | ├─css
  | |  └─index.css
  | ├─img
  | |  └─goui.png
  | ├─js
  | |  └─goui.js 
  | |  └─index.js
  | └─index.html
  └─main.go
```

`goui.js` - The frontend library GoUI provides for you.

`main.go` - provides a service named `hello`, and then create a window.
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

`index.html` - has a button on it, if a user clicks the button, it will request the above `hello` service.
`index.js` - frontend logic of `index.html`.
```
goui.request({url: "hello",
              success: function(data) {
                  document.getElementById("result").innerText = data;
             }});
```

Here are some screenshots.

macOS
![hello-macOS](https://github.com/FIPress/GoUI/raw/master/screenshots/hello-mac.jpeg)

Ubuntu
![hello-ubuntu](https://github.com/FIPress/GoUI/raw/master/screenshots/hello-ubuntu.jpeg)

iOS
![hello-ios](https://github.com/FIPress/GoUI/raw/master/screenshots/hello-ios.jpeg) 

### chat
Demonstrate how backend requests frontend service.

### texteditor
Demonstrate how to setup menu for desktop applications, on macOS, ubuntu and windows.

macOS
![editor-macOS](https://github.com/FIPress/GoUI/raw/master/screenshots/editor-mac.png)

Ubuntu
![editor-ubuntu](https://github.com/FIPress/GoUI/raw/master/screenshots/editor-ubuntu.jpeg)

## Progress
Currently, the basic functions are working on macOS, Ubuntu and iOS. We are working on Windows and Android, and it will be done soon. Then we will start adding features. If you have any suggestion, please don't hesitate to tell us.  