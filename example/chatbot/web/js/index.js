$(document).ready(function() {
    var $input = $("#input");
    var $box = $("#chat");

     $input.on("keydown", function (e) {
        if(e.which == 13 || e.keyCode == 13) {
            var msg = $input.val();
            if(msg.length != 0) {
                sendMessage(msg)
                $input.val("");
            }
        }
    });

    var sendMessage = function(msg) {
        $box.append('<div class="user"><div></div><div><div class="msg">' + msg + '</div></div><div><b class="avatar">U</b></div></div>')
        send(msg);
    };

    var send = function(msg) {
        goui.request({url: "chat/"+msg,success: receiveMessage});
    };

    var receiveMessage = function(msg) {
        $box.append('<div class="bot"><div><b class="avatar">B</b></div><div><div class="msg">' + msg + '</div></div</div>')
    };

    goui.service("chat/:msg",receiveMessage)

    send("ready");

});
