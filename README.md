# javascript-js2png

Hide JavaScript code into PNG image. For more details and a step-by-step
description of the code follow this [post][1].

## Prerequisite

* Go 1.3+ `sudo apt-get install golang-go`
* Make (optional)
* `git clone https://github.com/expobrain/javascript-js2png.git`

## Usage

To compile the JS-to-PNG encoder run:

    make js2png

To convert any payload into an image:

    bin/js2png <js_file> <png_file>

## Example

To run the embedded web server:

    make js2png
    make serve

To test the payload run the embedded web server and open the browser on
`http://localhost:8080`.

**Note**: remember that for security reasons the browser will not execute any
JavaScript code if the HTML file is loaded with the `file://` protocol.

[1]: https://www.expobrain.net/2014/09/26/hide-javascript-code-into-images/
