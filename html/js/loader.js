"use strict"

function js2png(imgID) {
    var img = document.getElementById("payload");

    img.addEventListener("load", function () {
        // Setup Canvas a draw image.
        var context, canvas;

        canvas = document.createElement("canvas");
        canvas.width = img.width;
        canvas.height = img.height;

        context = canvas.getContext("2d");
        context.drawImage(img, 0, 0);

        // Read image's data.
        var data;

        data = context.getImageData(0, 0, img.width, img.height).data;

        // Read size of the payload.
        var size = 0;

        size += data[1] << 56;
        size += data[2] << 48;
        size += data[4] << 40;
        size += data[5] << 32;
        size += data[6] << 24;
        size += data[8] << 16;
        size += data[9] << 8;
        size += data[10];

        // Read payload into string.
        var payload = '';

        for (var i = 12; i < data.length; i++) {
            if ((i + 1) % 4) {
                var char = data[i];

                if (char >= 32) {  
                    // Strip any non-ASCII to keep eval() happy :D                 
                    payload += String.fromCharCode(char);
                }
            }
        }

        // Execute code.
        eval(payload);
    }, false);
}
