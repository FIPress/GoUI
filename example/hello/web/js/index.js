document.onreadystatechange = function () {
    if(document.readyState == "complete") {
        var result = "";
        document.getElementById("btn").onclick = function() {
            goui.request({url: "hello",
                success: function(data) {
                    result = result + data + "\n"
                    document.getElementById("result").innerText = result;
                }});
        }
    }
}